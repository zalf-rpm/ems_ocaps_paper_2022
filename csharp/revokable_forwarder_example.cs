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

        // do @0 (msg :Text);
        public async Task Do(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Alice::do | msg: " + msg);
            Console.WriteLine("@Alice::do | sending foo(carol) to Bob");
            if(_bob != null && _carol != null) {
                await _bob.Foo(Proxy.Share(_carol), cancellationToken_);
            }
        }

        // set @0 (bob :Bob, carol :Carol);
        public Task Set(Schema.IBob bob, Schema.ICarol carol, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Alice::set");
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
        readonly object _lock = new object();
        Schema.ICarol _carol;

        // do @0 (msg :Text);
        public async Task Do(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Bob::do | msg: " + msg);
            if(_carol != null) {
                Console.WriteLine("@Bob::do | sending do(msg) to Carol");
                await _carol.Do("<Bobs DO message to Carol>", cancellationToken_);
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
        // do @0 (msg :Text);
        public Task Do(string msg, CancellationToken cancellationToken_ = default) {
            Console.WriteLine("@Carol::do | msg: " + msg);
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

                Console.WriteLine("@main | sending set(bob,carol) to Alice");
                await alice.Set(Proxy.Share(bob), carol);
                Console.WriteLine("@main | sending do(msg) to Alice");
                await alice.Do("<mains DO message to Alice>");

                Console.WriteLine("main | sending do(msg) to Bob");
                await bob.Do("<mains DO message to Bob>");

                //System.Threading.Thread.Sleep(4000);
                Console.WriteLine("@main | finished");

                //while (true) System.Threading.Thread.Sleep(1000); 
            }
        }
    }
}
