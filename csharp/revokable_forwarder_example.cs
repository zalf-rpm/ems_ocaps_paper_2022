using System;
using Capnp.Rpc;
using Capnp.Rpc.Interception;
using System.Net;
using System.Text.Json;
using System.Threading.Tasks;
using System.Threading;
using System.Diagnostics;

namespace Ems_ocaps_paper
{
    class Alice : Schema.IAlice {
        readonly object _lock = new object();
        private Schema.IBob _bob;
        private Schema.ICarol _carol;
        private Schema.IForwarder _forwarder = new Forwarder();
        private Schema.IRevoker _revoker = new Revoker();

        // do @0 (msg :Text);
        public async Task Act(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Alice::act | msg: " + msg);
            Console.WriteLine("@Alice::act | sending foo(carol) to Bob");
            if(_bob != null && _carol != null) {
                await _forwarder.SetActor(Proxy.Share(_revoker), cancellationToken_);
                await _revoker.SetActor(Proxy.Share(_carol), cancellationToken_);
                //await _bob.Foo(Proxy.Share(_carol), cancellationToken_);
                await _bob.Foo(Proxy.Share(_forwarder), cancellationToken_);
            }
        }

        // set @0 (bob :Bob, carol :Carol);
        public Task SetBobAndCarol(Schema.IBob bob, Schema.ICarol carol, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Alice::setBobAndCarol");
            lock(_lock){
                if (bob != _bob) {
                    _bob?.Dispose();
                    _bob = bob;
                }
                if (carol != _carol) {
                    _carol?.Dispose();
                    _carol = carol;
                }
            }
            return Task.CompletedTask;
        }

        // revokeCarol     @1 ();
        public Task RevokeCarol(CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Alice::revokeCarol");
            return _revoker.Revoke();
        }

        public void Dispose() {
            _bob?.Dispose();
            _carol?.Dispose();
        }

        static private TcpRpcServer _server;
        static public void run(int tcpPort) {
            _server = new TcpRpcServer();
            _server.Main = new Alice();
            _server.StartAccepting(IPAddress.Any, tcpPort);
        }
    }

    class Bob : Schema.IBob {
        private readonly object _lock = new object();
        private Schema.ICarol _carol;
        private int _count = 1;

        // act @0 (msg :Text);
        public async Task Act(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Bob::act | msg: " + msg);
            if(_carol != null) {
                Console.WriteLine("@Bob::act | sending act(msg) to Carol");
                await _carol.Act("<Bobs " + _count + ". ACT message to Carol>", cancellationToken_);
                _count++;
            }   
        }

        // foo @0 (carol :Carol);
        public Task Foo(Schema.ICarol carol, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Bob::foo");
            lock (_lock) {
                if (carol != _carol) {
                    _carol?.Dispose();
                    _carol = carol;
                }
            }
            return Task.CompletedTask;
        }

        public void Dispose() {}

        static private TcpRpcServer _server;
        static public void run(int tcpPort) {
            _server = new TcpRpcServer();
            _server.Main = new Bob();
            _server.StartAccepting(IPAddress.Any, tcpPort);
        }
    }

    class Carol : Schema.ICarol {
        // act @0 (msg :Text);
        public Task Act(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Carol::act | msg: " + msg);
            return Task.CompletedTask;
        }

        public void Dispose() {}

        static private TcpRpcServer _server;
        static public void run(int tcpPort) {
            _server = new TcpRpcServer();
            _server.Main = new Carol();
            _server.StartAccepting(IPAddress.Any, tcpPort);
        }
    }

    class Forwarder : Schema.IForwarder {
        private readonly object _lock = new object();
        private Schema.IActor _actor;

        // act @0 (msg :Text);
        public async Task Act(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Forwarder::act | msg: " + msg);
            await _actor?.Act(msg, cancellationToken_);
        }

        // setActor @0 (a :Actor);
        public Task SetActor(Schema.IActor actor, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Forwarder::setActor");
            lock (_lock) {
                if (actor != _actor) {
                    _actor?.Dispose();
                    _actor = actor;
                }
            }
            return Task.CompletedTask;
        }

        public void Dispose() {}

        static private TcpRpcServer _server;
        static public void run(int tcpPort) {
            _server = new TcpRpcServer();
            _server.Main = new Forwarder();
            _server.StartAccepting(IPAddress.Any, tcpPort);
        }
    }

    class Revoker : Schema.IRevoker {
        private readonly object _lock = new object();
        private Schema.IActor _actor;

        // act @0 (msg :Text);
        public async Task Act(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Revoker::act | msg: " + msg);
            if (_actor != null) await _actor.Act(msg, cancellationToken_);
            else Console.WriteLine("Access to Actor has been revoked.");
        }

        // setActor @0 (a :Actor);
        public Task SetActor(Schema.IActor actor, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Revoker::setActor");
            lock (_lock) {
                if (actor != _actor) {
                    _actor?.Dispose();
                    _actor = actor;
                }
            }
            return Task.CompletedTask;
        }

        // revoke @0 ();
        public Task Revoke(CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Revoker::revoke");
            lock (_lock) {
                _actor?.Dispose();
                _actor = null;
            }
            return Task.CompletedTask;
        }

        public void Dispose() {}

        static private TcpRpcServer _server;
        static public void run(int tcpPort) {
            _server = new TcpRpcServer();
            _server.Main = new Revoker();
            _server.StartAccepting(IPAddress.Any, tcpPort);
        }
    }

    class Program
    {
        static async Task Main(string[] args)
        {
            string actor = "aio";
            int alicePort = 9991;
            int bobPort = 9992;
            int carolPort = 9993;

            for (var i = 0; i < args.Length; i++)
            {
                try
                {
                    if (args[i].StartsWith("actor")) actor = args[i].Split('=')[1];
                }
                catch (System.Exception) { }
            }

            if(actor == "alice") Alice.run(alicePort);
            else if(actor == "bob") Bob.run(bobPort);
            else if(actor == "carol") Carol.run(carolPort);
            else if(actor == "main")
            {
                using var aliceCon = new TcpRpcClient("127.0.0.1", alicePort);
                var alice = aliceCon.GetMain<Schema.IAlice>();
                using var bobCon = new TcpRpcClient("127.0.0.1", bobPort);
                var bob = bobCon.GetMain<Schema.IBob>();
                using var carolCon = new TcpRpcClient("127.0.0.1", carolPort);
                var carol = carolCon.GetMain<Schema.ICarol>();

                Console.WriteLine("@main | sleep for 1s");
                System.Threading.Thread.Sleep(1000);

                Console.WriteLine("@main | sending setBobAndCarol(bob,carol) to Alice");
                await alice.SetBobAndCarol(Proxy.Share(bob), carol);
                Console.WriteLine("@main | sending act(msg) to Alice");
                await alice.Act("<mains ACT message to Alice>");

                Console.WriteLine("@main | sending act(msg) to Bob");
                await bob.Act("<mains 1st ACT message to Bob>");
                await bob.Act("<mains 2nd ACT message to Bob>");

                Console.WriteLine("@main | sending revokeCarol to Alice");
                await alice.RevokeCarol();

                Console.WriteLine("@main | sending act(msg) to Bob");
                await bob.Act("<mains 3rd ACT message to Bob>");

                //System.Threading.Thread.Sleep(4000);
                Console.WriteLine("@main | finished");

                //while (true) System.Threading.Thread.Sleep(1000); 
            }
        }
    }
}
