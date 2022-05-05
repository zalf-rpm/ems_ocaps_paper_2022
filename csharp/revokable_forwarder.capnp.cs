using Capnp;
using Capnp.Rpc;
using System;
using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;

namespace Ems_ocaps_paper.Schema
{
    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xa27a66204e1f5179UL), Proxy(typeof(Actor_Proxy)), Skeleton(typeof(Actor_Skeleton))]
    public interface IActor : IDisposable
    {
        Task Do(string msg, CancellationToken cancellationToken_ = default);
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xa27a66204e1f5179UL)]
    public class Actor_Proxy : Proxy, IActor
    {
        public async Task Do(string msg, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Actor.Params_Do.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Actor.Params_Do()
            {Msg = msg};
            arg_?.serialize(in_);
            using (var d_ = await Call(11707782470238687609UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Actor.Result_Do>(d_);
                return;
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xa27a66204e1f5179UL)]
    public class Actor_Skeleton : Skeleton<IActor>
    {
        public Actor_Skeleton()
        {
            SetMethodTable(Do);
        }

        public override ulong InterfaceId => 11707782470238687609UL;
        async Task<AnswerOrCounterquestion> Do(DeserializerState d_, CancellationToken cancellationToken_)
        {
            using (d_)
            {
                var in_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Actor.Params_Do>(d_);
                await Impl.Do(in_.Msg, cancellationToken_);
                var s_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Actor.Result_Do.WRITER>();
                return s_;
            }
        }
    }

    public static class Actor
    {
        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xfb6aefd15395d9a0UL)]
        public class Params_Do : ICapnpSerializable
        {
            public const UInt64 typeId = 0xfb6aefd15395d9a0UL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                Msg = reader.Msg;
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
                writer.Msg = Msg;
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public string Msg
            {
                get;
                set;
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
                public string Msg => ctx.ReadText(0, null);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 1);
                }

                public string Msg
                {
                    get => this.ReadText(0, null);
                    set => this.WriteText(0, value, null);
                }
            }
        }

        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xd9e5a1fdbab984dbUL)]
        public class Result_Do : ICapnpSerializable
        {
            public const UInt64 typeId = 0xd9e5a1fdbab984dbUL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 0);
                }
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xfd124cf35abe03caUL), Proxy(typeof(Alice_Proxy)), Skeleton(typeof(Alice_Skeleton))]
    public interface IAlice : Ems_ocaps_paper.Schema.IActor
    {
        Task Set(Ems_ocaps_paper.Schema.IBob bob, Ems_ocaps_paper.Schema.ICarol carol, CancellationToken cancellationToken_ = default);
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xfd124cf35abe03caUL)]
    public class Alice_Proxy : Proxy, IAlice
    {
        public async Task Set(Ems_ocaps_paper.Schema.IBob bob, Ems_ocaps_paper.Schema.ICarol carol, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Alice.Params_Set.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Alice.Params_Set()
            {Bob = bob, Carol = carol};
            arg_?.serialize(in_);
            using (var d_ = await Call(18235722449259725770UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Alice.Result_Set>(d_);
                return;
            }
        }

        public async Task Do(string msg, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Actor.Params_Do.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Actor.Params_Do()
            {Msg = msg};
            arg_?.serialize(in_);
            using (var d_ = await Call(11707782470238687609UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Actor.Result_Do>(d_);
                return;
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xfd124cf35abe03caUL)]
    public class Alice_Skeleton : Skeleton<IAlice>
    {
        public Alice_Skeleton()
        {
            SetMethodTable(Set);
        }

        public override ulong InterfaceId => 18235722449259725770UL;
        async Task<AnswerOrCounterquestion> Set(DeserializerState d_, CancellationToken cancellationToken_)
        {
            using (d_)
            {
                var in_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Alice.Params_Set>(d_);
                await Impl.Set(in_.Bob, in_.Carol, cancellationToken_);
                var s_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Alice.Result_Set.WRITER>();
                return s_;
            }
        }
    }

    public static class Alice
    {
        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xc14b91e8a10701f2UL)]
        public class Params_Set : ICapnpSerializable
        {
            public const UInt64 typeId = 0xc14b91e8a10701f2UL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                Bob = reader.Bob;
                Carol = reader.Carol;
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
                writer.Bob = Bob;
                writer.Carol = Carol;
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public Ems_ocaps_paper.Schema.IBob Bob
            {
                get;
                set;
            }

            public Ems_ocaps_paper.Schema.ICarol Carol
            {
                get;
                set;
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
                public Ems_ocaps_paper.Schema.IBob Bob => ctx.ReadCap<Ems_ocaps_paper.Schema.IBob>(0);
                public Ems_ocaps_paper.Schema.ICarol Carol => ctx.ReadCap<Ems_ocaps_paper.Schema.ICarol>(1);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 2);
                }

                public Ems_ocaps_paper.Schema.IBob Bob
                {
                    get => ReadCap<Ems_ocaps_paper.Schema.IBob>(0);
                    set => LinkObject(0, value);
                }

                public Ems_ocaps_paper.Schema.ICarol Carol
                {
                    get => ReadCap<Ems_ocaps_paper.Schema.ICarol>(1);
                    set => LinkObject(1, value);
                }
            }
        }

        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xfa5948440e852aa3UL)]
        public class Result_Set : ICapnpSerializable
        {
            public const UInt64 typeId = 0xfa5948440e852aa3UL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 0);
                }
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xd52ad152ab041ea2UL), Proxy(typeof(Bob_Proxy)), Skeleton(typeof(Bob_Skeleton))]
    public interface IBob : Ems_ocaps_paper.Schema.IActor
    {
        Task Foo(Ems_ocaps_paper.Schema.ICarol carol, CancellationToken cancellationToken_ = default);
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xd52ad152ab041ea2UL)]
    public class Bob_Proxy : Proxy, IBob
    {
        public async Task Foo(Ems_ocaps_paper.Schema.ICarol carol, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Bob.Params_Foo.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Bob.Params_Foo()
            {Carol = carol};
            arg_?.serialize(in_);
            using (var d_ = await Call(15360319632087195298UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Bob.Result_Foo>(d_);
                return;
            }
        }

        public async Task Do(string msg, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Actor.Params_Do.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Actor.Params_Do()
            {Msg = msg};
            arg_?.serialize(in_);
            using (var d_ = await Call(11707782470238687609UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Actor.Result_Do>(d_);
                return;
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xd52ad152ab041ea2UL)]
    public class Bob_Skeleton : Skeleton<IBob>
    {
        public Bob_Skeleton()
        {
            SetMethodTable(Foo);
        }

        public override ulong InterfaceId => 15360319632087195298UL;
        async Task<AnswerOrCounterquestion> Foo(DeserializerState d_, CancellationToken cancellationToken_)
        {
            using (d_)
            {
                var in_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Bob.Params_Foo>(d_);
                await Impl.Foo(in_.Carol, cancellationToken_);
                var s_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Bob.Result_Foo.WRITER>();
                return s_;
            }
        }
    }

    public static class Bob
    {
        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xcc80d0c2361b67faUL)]
        public class Params_Foo : ICapnpSerializable
        {
            public const UInt64 typeId = 0xcc80d0c2361b67faUL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                Carol = reader.Carol;
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
                writer.Carol = Carol;
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public Ems_ocaps_paper.Schema.ICarol Carol
            {
                get;
                set;
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
                public Ems_ocaps_paper.Schema.ICarol Carol => ctx.ReadCap<Ems_ocaps_paper.Schema.ICarol>(0);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 1);
                }

                public Ems_ocaps_paper.Schema.ICarol Carol
                {
                    get => ReadCap<Ems_ocaps_paper.Schema.ICarol>(0);
                    set => LinkObject(0, value);
                }
            }
        }

        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0x9a19936b9f4b0f30UL)]
        public class Result_Foo : ICapnpSerializable
        {
            public const UInt64 typeId = 0x9a19936b9f4b0f30UL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 0);
                }
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xc37310a51edca8c5UL), Proxy(typeof(Carol_Proxy)), Skeleton(typeof(Carol_Skeleton))]
    public interface ICarol : Ems_ocaps_paper.Schema.IActor
    {
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xc37310a51edca8c5UL)]
    public class Carol_Proxy : Proxy, ICarol
    {
        public async Task Do(string msg, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Actor.Params_Do.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Actor.Params_Do()
            {Msg = msg};
            arg_?.serialize(in_);
            using (var d_ = await Call(11707782470238687609UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Actor.Result_Do>(d_);
                return;
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xc37310a51edca8c5UL)]
    public class Carol_Skeleton : Skeleton<ICarol>
    {
        public Carol_Skeleton()
        {
            SetMethodTable();
        }

        public override ulong InterfaceId => 14083618761091098821UL;
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xcde6a707e4219348UL), Proxy(typeof(Forwarder_Proxy)), Skeleton(typeof(Forwarder_Skeleton))]
    public interface IForwarder : Ems_ocaps_paper.Schema.ICarol
    {
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xcde6a707e4219348UL)]
    public class Forwarder_Proxy : Proxy, IForwarder
    {
        public async Task Do(string msg, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Actor.Params_Do.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Actor.Params_Do()
            {Msg = msg};
            arg_?.serialize(in_);
            using (var d_ = await Call(11707782470238687609UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Actor.Result_Do>(d_);
                return;
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xcde6a707e4219348UL)]
    public class Forwarder_Skeleton : Skeleton<IForwarder>
    {
        public Forwarder_Skeleton()
        {
            SetMethodTable();
        }

        public override ulong InterfaceId => 14836729674752693064UL;
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xbd0ca7ed013308baUL), Proxy(typeof(Revoker_Proxy)), Skeleton(typeof(Revoker_Skeleton))]
    public interface IRevoker : Ems_ocaps_paper.Schema.IForwarder
    {
        Task Revoke(CancellationToken cancellationToken_ = default);
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xbd0ca7ed013308baUL)]
    public class Revoker_Proxy : Proxy, IRevoker
    {
        public async Task Revoke(CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Revoker.Params_Revoke.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Revoker.Params_Revoke()
            {};
            arg_?.serialize(in_);
            using (var d_ = await Call(13622447609258117306UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Revoker.Result_Revoke>(d_);
                return;
            }
        }

        public async Task Do(string msg, CancellationToken cancellationToken_ = default)
        {
            var in_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Actor.Params_Do.WRITER>();
            var arg_ = new Ems_ocaps_paper.Schema.Actor.Params_Do()
            {Msg = msg};
            arg_?.serialize(in_);
            using (var d_ = await Call(11707782470238687609UL, 0, in_.Rewrap<DynamicSerializerState>(), false, cancellationToken_).WhenReturned)
            {
                var r_ = CapnpSerializable.Create<Ems_ocaps_paper.Schema.Actor.Result_Do>(d_);
                return;
            }
        }
    }

    [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xbd0ca7ed013308baUL)]
    public class Revoker_Skeleton : Skeleton<IRevoker>
    {
        public Revoker_Skeleton()
        {
            SetMethodTable(Revoke);
        }

        public override ulong InterfaceId => 13622447609258117306UL;
        async Task<AnswerOrCounterquestion> Revoke(DeserializerState d_, CancellationToken cancellationToken_)
        {
            using (d_)
            {
                await Impl.Revoke(cancellationToken_);
                var s_ = SerializerState.CreateForRpc<Ems_ocaps_paper.Schema.Revoker.Result_Revoke.WRITER>();
                return s_;
            }
        }
    }

    public static class Revoker
    {
        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xb457bead0d4bb698UL)]
        public class Params_Revoke : ICapnpSerializable
        {
            public const UInt64 typeId = 0xb457bead0d4bb698UL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 0);
                }
            }
        }

        [System.CodeDom.Compiler.GeneratedCode("capnpc-csharp", "1.3.0.0"), TypeId(0xbc1b102199ee0e6bUL)]
        public class Result_Revoke : ICapnpSerializable
        {
            public const UInt64 typeId = 0xbc1b102199ee0e6bUL;
            void ICapnpSerializable.Deserialize(DeserializerState arg_)
            {
                var reader = READER.create(arg_);
                applyDefaults();
            }

            public void serialize(WRITER writer)
            {
            }

            void ICapnpSerializable.Serialize(SerializerState arg_)
            {
                serialize(arg_.Rewrap<WRITER>());
            }

            public void applyDefaults()
            {
            }

            public struct READER
            {
                readonly DeserializerState ctx;
                public READER(DeserializerState ctx)
                {
                    this.ctx = ctx;
                }

                public static READER create(DeserializerState ctx) => new READER(ctx);
                public static implicit operator DeserializerState(READER reader) => reader.ctx;
                public static implicit operator READER(DeserializerState ctx) => new READER(ctx);
            }

            public class WRITER : SerializerState
            {
                public WRITER()
                {
                    this.SetStruct(0, 0);
                }
            }
        }
    }
}