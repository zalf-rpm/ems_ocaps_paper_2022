// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
)

type Actor struct{ Client *capnp.Client }

// Actor_TypeID is the unique identifier for the type Actor.
const Actor_TypeID = 0xa27a66204e1f5179

func (c Actor) Do(ctx context.Context, params func(Actor_do_Params) error) (Actor_do_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Actor_do_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Actor_do_Results_Future{Future: ans.Future()}, release
}

func (c Actor) AddRef() Actor {
	return Actor{
		Client: c.Client.AddRef(),
	}
}

func (c Actor) Release() {
	c.Client.Release()
}

// A Actor_Server is a Actor with a local implementation.
type Actor_Server interface {
	Do(context.Context, Actor_do) error
}

// Actor_NewServer creates a new Server from an implementation of Actor_Server.
func Actor_NewServer(s Actor_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Actor_Methods(nil, s), s, c, policy)
}

// Actor_ServerToClient creates a new Client from an implementation of Actor_Server.
// The caller is responsible for calling Release on the returned Client.
func Actor_ServerToClient(s Actor_Server, policy *server.Policy) Actor {
	return Actor{Client: capnp.NewClient(Actor_NewServer(s, policy))}
}

// Actor_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Actor_Methods(methods []server.Method, s Actor_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Do(ctx, Actor_do{call})
		},
	})

	return methods
}

// Actor_do holds the state for a server call to Actor.do.
// See server.Call for documentation.
type Actor_do struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Actor_do) Args() Actor_do_Params {
	return Actor_do_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Actor_do) AllocResults() (Actor_do_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Actor_do_Results{Struct: r}, err
}

type Actor_do_Params struct{ capnp.Struct }

// Actor_do_Params_TypeID is the unique identifier for the type Actor_do_Params.
const Actor_do_Params_TypeID = 0xfb6aefd15395d9a0

func NewActor_do_Params(s *capnp.Segment) (Actor_do_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Actor_do_Params{st}, err
}

func NewRootActor_do_Params(s *capnp.Segment) (Actor_do_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Actor_do_Params{st}, err
}

func ReadRootActor_do_Params(msg *capnp.Message) (Actor_do_Params, error) {
	root, err := msg.Root()
	return Actor_do_Params{root.Struct()}, err
}

func (s Actor_do_Params) String() string {
	str, _ := text.Marshal(0xfb6aefd15395d9a0, s.Struct)
	return str
}

func (s Actor_do_Params) Msg() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Actor_do_Params) HasMsg() bool {
	return s.Struct.HasPtr(0)
}

func (s Actor_do_Params) MsgBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Actor_do_Params) SetMsg(v string) error {
	return s.Struct.SetText(0, v)
}

// Actor_do_Params_List is a list of Actor_do_Params.
type Actor_do_Params_List struct{ capnp.List }

// NewActor_do_Params creates a new list of Actor_do_Params.
func NewActor_do_Params_List(s *capnp.Segment, sz int32) (Actor_do_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Actor_do_Params_List{l}, err
}

func (s Actor_do_Params_List) At(i int) Actor_do_Params { return Actor_do_Params{s.List.Struct(i)} }

func (s Actor_do_Params_List) Set(i int, v Actor_do_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Actor_do_Params_List) String() string {
	str, _ := text.MarshalList(0xfb6aefd15395d9a0, s.List)
	return str
}

// Actor_do_Params_Future is a wrapper for a Actor_do_Params promised by a client call.
type Actor_do_Params_Future struct{ *capnp.Future }

func (p Actor_do_Params_Future) Struct() (Actor_do_Params, error) {
	s, err := p.Future.Struct()
	return Actor_do_Params{s}, err
}

type Actor_do_Results struct{ capnp.Struct }

// Actor_do_Results_TypeID is the unique identifier for the type Actor_do_Results.
const Actor_do_Results_TypeID = 0xd9e5a1fdbab984db

func NewActor_do_Results(s *capnp.Segment) (Actor_do_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Actor_do_Results{st}, err
}

func NewRootActor_do_Results(s *capnp.Segment) (Actor_do_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Actor_do_Results{st}, err
}

func ReadRootActor_do_Results(msg *capnp.Message) (Actor_do_Results, error) {
	root, err := msg.Root()
	return Actor_do_Results{root.Struct()}, err
}

func (s Actor_do_Results) String() string {
	str, _ := text.Marshal(0xd9e5a1fdbab984db, s.Struct)
	return str
}

// Actor_do_Results_List is a list of Actor_do_Results.
type Actor_do_Results_List struct{ capnp.List }

// NewActor_do_Results creates a new list of Actor_do_Results.
func NewActor_do_Results_List(s *capnp.Segment, sz int32) (Actor_do_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Actor_do_Results_List{l}, err
}

func (s Actor_do_Results_List) At(i int) Actor_do_Results { return Actor_do_Results{s.List.Struct(i)} }

func (s Actor_do_Results_List) Set(i int, v Actor_do_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Actor_do_Results_List) String() string {
	str, _ := text.MarshalList(0xd9e5a1fdbab984db, s.List)
	return str
}

// Actor_do_Results_Future is a wrapper for a Actor_do_Results promised by a client call.
type Actor_do_Results_Future struct{ *capnp.Future }

func (p Actor_do_Results_Future) Struct() (Actor_do_Results, error) {
	s, err := p.Future.Struct()
	return Actor_do_Results{s}, err
}

type Alice struct{ Client *capnp.Client }

// Alice_TypeID is the unique identifier for the type Alice.
const Alice_TypeID = 0xfd124cf35abe03ca

func (c Alice) Set(ctx context.Context, params func(Alice_set_Params) error) (Alice_set_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xfd124cf35abe03ca,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Alice",
			MethodName:    "set",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Alice_set_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Alice_set_Results_Future{Future: ans.Future()}, release
}
func (c Alice) Do(ctx context.Context, params func(Actor_do_Params) error) (Actor_do_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Actor_do_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Actor_do_Results_Future{Future: ans.Future()}, release
}

func (c Alice) AddRef() Alice {
	return Alice{
		Client: c.Client.AddRef(),
	}
}

func (c Alice) Release() {
	c.Client.Release()
}

// A Alice_Server is a Alice with a local implementation.
type Alice_Server interface {
	Set(context.Context, Alice_set) error

	Do(context.Context, Actor_do) error
}

// Alice_NewServer creates a new Server from an implementation of Alice_Server.
func Alice_NewServer(s Alice_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Alice_Methods(nil, s), s, c, policy)
}

// Alice_ServerToClient creates a new Client from an implementation of Alice_Server.
// The caller is responsible for calling Release on the returned Client.
func Alice_ServerToClient(s Alice_Server, policy *server.Policy) Alice {
	return Alice{Client: capnp.NewClient(Alice_NewServer(s, policy))}
}

// Alice_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Alice_Methods(methods []server.Method, s Alice_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfd124cf35abe03ca,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Alice",
			MethodName:    "set",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Set(ctx, Alice_set{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Do(ctx, Actor_do{call})
		},
	})

	return methods
}

// Alice_set holds the state for a server call to Alice.set.
// See server.Call for documentation.
type Alice_set struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Alice_set) Args() Alice_set_Params {
	return Alice_set_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Alice_set) AllocResults() (Alice_set_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Alice_set_Results{Struct: r}, err
}

type Alice_set_Params struct{ capnp.Struct }

// Alice_set_Params_TypeID is the unique identifier for the type Alice_set_Params.
const Alice_set_Params_TypeID = 0xc14b91e8a10701f2

func NewAlice_set_Params(s *capnp.Segment) (Alice_set_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Alice_set_Params{st}, err
}

func NewRootAlice_set_Params(s *capnp.Segment) (Alice_set_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Alice_set_Params{st}, err
}

func ReadRootAlice_set_Params(msg *capnp.Message) (Alice_set_Params, error) {
	root, err := msg.Root()
	return Alice_set_Params{root.Struct()}, err
}

func (s Alice_set_Params) String() string {
	str, _ := text.Marshal(0xc14b91e8a10701f2, s.Struct)
	return str
}

func (s Alice_set_Params) Bob() Bob {
	p, _ := s.Struct.Ptr(0)
	return Bob{Client: p.Interface().Client()}
}

func (s Alice_set_Params) HasBob() bool {
	return s.Struct.HasPtr(0)
}

func (s Alice_set_Params) SetBob(v Bob) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

func (s Alice_set_Params) Carol() Carol {
	p, _ := s.Struct.Ptr(1)
	return Carol{Client: p.Interface().Client()}
}

func (s Alice_set_Params) HasCarol() bool {
	return s.Struct.HasPtr(1)
}

func (s Alice_set_Params) SetCarol(v Carol) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(1, in.ToPtr())
}

// Alice_set_Params_List is a list of Alice_set_Params.
type Alice_set_Params_List struct{ capnp.List }

// NewAlice_set_Params creates a new list of Alice_set_Params.
func NewAlice_set_Params_List(s *capnp.Segment, sz int32) (Alice_set_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Alice_set_Params_List{l}, err
}

func (s Alice_set_Params_List) At(i int) Alice_set_Params { return Alice_set_Params{s.List.Struct(i)} }

func (s Alice_set_Params_List) Set(i int, v Alice_set_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Alice_set_Params_List) String() string {
	str, _ := text.MarshalList(0xc14b91e8a10701f2, s.List)
	return str
}

// Alice_set_Params_Future is a wrapper for a Alice_set_Params promised by a client call.
type Alice_set_Params_Future struct{ *capnp.Future }

func (p Alice_set_Params_Future) Struct() (Alice_set_Params, error) {
	s, err := p.Future.Struct()
	return Alice_set_Params{s}, err
}

func (p Alice_set_Params_Future) Bob() Bob {
	return Bob{Client: p.Future.Field(0, nil).Client()}
}

func (p Alice_set_Params_Future) Carol() Carol {
	return Carol{Client: p.Future.Field(1, nil).Client()}
}

type Alice_set_Results struct{ capnp.Struct }

// Alice_set_Results_TypeID is the unique identifier for the type Alice_set_Results.
const Alice_set_Results_TypeID = 0xfa5948440e852aa3

func NewAlice_set_Results(s *capnp.Segment) (Alice_set_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Alice_set_Results{st}, err
}

func NewRootAlice_set_Results(s *capnp.Segment) (Alice_set_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Alice_set_Results{st}, err
}

func ReadRootAlice_set_Results(msg *capnp.Message) (Alice_set_Results, error) {
	root, err := msg.Root()
	return Alice_set_Results{root.Struct()}, err
}

func (s Alice_set_Results) String() string {
	str, _ := text.Marshal(0xfa5948440e852aa3, s.Struct)
	return str
}

// Alice_set_Results_List is a list of Alice_set_Results.
type Alice_set_Results_List struct{ capnp.List }

// NewAlice_set_Results creates a new list of Alice_set_Results.
func NewAlice_set_Results_List(s *capnp.Segment, sz int32) (Alice_set_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Alice_set_Results_List{l}, err
}

func (s Alice_set_Results_List) At(i int) Alice_set_Results {
	return Alice_set_Results{s.List.Struct(i)}
}

func (s Alice_set_Results_List) Set(i int, v Alice_set_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Alice_set_Results_List) String() string {
	str, _ := text.MarshalList(0xfa5948440e852aa3, s.List)
	return str
}

// Alice_set_Results_Future is a wrapper for a Alice_set_Results promised by a client call.
type Alice_set_Results_Future struct{ *capnp.Future }

func (p Alice_set_Results_Future) Struct() (Alice_set_Results, error) {
	s, err := p.Future.Struct()
	return Alice_set_Results{s}, err
}

type Bob struct{ Client *capnp.Client }

// Bob_TypeID is the unique identifier for the type Bob.
const Bob_TypeID = 0xd52ad152ab041ea2

func (c Bob) Foo(ctx context.Context, params func(Bob_foo_Params) error) (Bob_foo_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xd52ad152ab041ea2,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Bob",
			MethodName:    "foo",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Bob_foo_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Bob_foo_Results_Future{Future: ans.Future()}, release
}
func (c Bob) Do(ctx context.Context, params func(Actor_do_Params) error) (Actor_do_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Actor_do_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Actor_do_Results_Future{Future: ans.Future()}, release
}

func (c Bob) AddRef() Bob {
	return Bob{
		Client: c.Client.AddRef(),
	}
}

func (c Bob) Release() {
	c.Client.Release()
}

// A Bob_Server is a Bob with a local implementation.
type Bob_Server interface {
	Foo(context.Context, Bob_foo) error

	Do(context.Context, Actor_do) error
}

// Bob_NewServer creates a new Server from an implementation of Bob_Server.
func Bob_NewServer(s Bob_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Bob_Methods(nil, s), s, c, policy)
}

// Bob_ServerToClient creates a new Client from an implementation of Bob_Server.
// The caller is responsible for calling Release on the returned Client.
func Bob_ServerToClient(s Bob_Server, policy *server.Policy) Bob {
	return Bob{Client: capnp.NewClient(Bob_NewServer(s, policy))}
}

// Bob_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Bob_Methods(methods []server.Method, s Bob_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xd52ad152ab041ea2,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Bob",
			MethodName:    "foo",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Foo(ctx, Bob_foo{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Do(ctx, Actor_do{call})
		},
	})

	return methods
}

// Bob_foo holds the state for a server call to Bob.foo.
// See server.Call for documentation.
type Bob_foo struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Bob_foo) Args() Bob_foo_Params {
	return Bob_foo_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Bob_foo) AllocResults() (Bob_foo_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Bob_foo_Results{Struct: r}, err
}

type Bob_foo_Params struct{ capnp.Struct }

// Bob_foo_Params_TypeID is the unique identifier for the type Bob_foo_Params.
const Bob_foo_Params_TypeID = 0xcc80d0c2361b67fa

func NewBob_foo_Params(s *capnp.Segment) (Bob_foo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Bob_foo_Params{st}, err
}

func NewRootBob_foo_Params(s *capnp.Segment) (Bob_foo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Bob_foo_Params{st}, err
}

func ReadRootBob_foo_Params(msg *capnp.Message) (Bob_foo_Params, error) {
	root, err := msg.Root()
	return Bob_foo_Params{root.Struct()}, err
}

func (s Bob_foo_Params) String() string {
	str, _ := text.Marshal(0xcc80d0c2361b67fa, s.Struct)
	return str
}

func (s Bob_foo_Params) Carol() Carol {
	p, _ := s.Struct.Ptr(0)
	return Carol{Client: p.Interface().Client()}
}

func (s Bob_foo_Params) HasCarol() bool {
	return s.Struct.HasPtr(0)
}

func (s Bob_foo_Params) SetCarol(v Carol) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Bob_foo_Params_List is a list of Bob_foo_Params.
type Bob_foo_Params_List struct{ capnp.List }

// NewBob_foo_Params creates a new list of Bob_foo_Params.
func NewBob_foo_Params_List(s *capnp.Segment, sz int32) (Bob_foo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Bob_foo_Params_List{l}, err
}

func (s Bob_foo_Params_List) At(i int) Bob_foo_Params { return Bob_foo_Params{s.List.Struct(i)} }

func (s Bob_foo_Params_List) Set(i int, v Bob_foo_Params) error { return s.List.SetStruct(i, v.Struct) }

func (s Bob_foo_Params_List) String() string {
	str, _ := text.MarshalList(0xcc80d0c2361b67fa, s.List)
	return str
}

// Bob_foo_Params_Future is a wrapper for a Bob_foo_Params promised by a client call.
type Bob_foo_Params_Future struct{ *capnp.Future }

func (p Bob_foo_Params_Future) Struct() (Bob_foo_Params, error) {
	s, err := p.Future.Struct()
	return Bob_foo_Params{s}, err
}

func (p Bob_foo_Params_Future) Carol() Carol {
	return Carol{Client: p.Future.Field(0, nil).Client()}
}

type Bob_foo_Results struct{ capnp.Struct }

// Bob_foo_Results_TypeID is the unique identifier for the type Bob_foo_Results.
const Bob_foo_Results_TypeID = 0x9a19936b9f4b0f30

func NewBob_foo_Results(s *capnp.Segment) (Bob_foo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Bob_foo_Results{st}, err
}

func NewRootBob_foo_Results(s *capnp.Segment) (Bob_foo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Bob_foo_Results{st}, err
}

func ReadRootBob_foo_Results(msg *capnp.Message) (Bob_foo_Results, error) {
	root, err := msg.Root()
	return Bob_foo_Results{root.Struct()}, err
}

func (s Bob_foo_Results) String() string {
	str, _ := text.Marshal(0x9a19936b9f4b0f30, s.Struct)
	return str
}

// Bob_foo_Results_List is a list of Bob_foo_Results.
type Bob_foo_Results_List struct{ capnp.List }

// NewBob_foo_Results creates a new list of Bob_foo_Results.
func NewBob_foo_Results_List(s *capnp.Segment, sz int32) (Bob_foo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Bob_foo_Results_List{l}, err
}

func (s Bob_foo_Results_List) At(i int) Bob_foo_Results { return Bob_foo_Results{s.List.Struct(i)} }

func (s Bob_foo_Results_List) Set(i int, v Bob_foo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Bob_foo_Results_List) String() string {
	str, _ := text.MarshalList(0x9a19936b9f4b0f30, s.List)
	return str
}

// Bob_foo_Results_Future is a wrapper for a Bob_foo_Results promised by a client call.
type Bob_foo_Results_Future struct{ *capnp.Future }

func (p Bob_foo_Results_Future) Struct() (Bob_foo_Results, error) {
	s, err := p.Future.Struct()
	return Bob_foo_Results{s}, err
}

type Carol struct{ Client *capnp.Client }

// Carol_TypeID is the unique identifier for the type Carol.
const Carol_TypeID = 0xc37310a51edca8c5

func (c Carol) Do(ctx context.Context, params func(Actor_do_Params) error) (Actor_do_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Actor_do_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Actor_do_Results_Future{Future: ans.Future()}, release
}

func (c Carol) AddRef() Carol {
	return Carol{
		Client: c.Client.AddRef(),
	}
}

func (c Carol) Release() {
	c.Client.Release()
}

// A Carol_Server is a Carol with a local implementation.
type Carol_Server interface {
	Do(context.Context, Actor_do) error
}

// Carol_NewServer creates a new Server from an implementation of Carol_Server.
func Carol_NewServer(s Carol_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Carol_Methods(nil, s), s, c, policy)
}

// Carol_ServerToClient creates a new Client from an implementation of Carol_Server.
// The caller is responsible for calling Release on the returned Client.
func Carol_ServerToClient(s Carol_Server, policy *server.Policy) Carol {
	return Carol{Client: capnp.NewClient(Carol_NewServer(s, policy))}
}

// Carol_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Carol_Methods(methods []server.Method, s Carol_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Do(ctx, Actor_do{call})
		},
	})

	return methods
}

type Forwarder struct{ Client *capnp.Client }

// Forwarder_TypeID is the unique identifier for the type Forwarder.
const Forwarder_TypeID = 0xcde6a707e4219348

func (c Forwarder) Do(ctx context.Context, params func(Actor_do_Params) error) (Actor_do_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Actor_do_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Actor_do_Results_Future{Future: ans.Future()}, release
}

func (c Forwarder) AddRef() Forwarder {
	return Forwarder{
		Client: c.Client.AddRef(),
	}
}

func (c Forwarder) Release() {
	c.Client.Release()
}

// A Forwarder_Server is a Forwarder with a local implementation.
type Forwarder_Server interface {
	Do(context.Context, Actor_do) error
}

// Forwarder_NewServer creates a new Server from an implementation of Forwarder_Server.
func Forwarder_NewServer(s Forwarder_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Forwarder_Methods(nil, s), s, c, policy)
}

// Forwarder_ServerToClient creates a new Client from an implementation of Forwarder_Server.
// The caller is responsible for calling Release on the returned Client.
func Forwarder_ServerToClient(s Forwarder_Server, policy *server.Policy) Forwarder {
	return Forwarder{Client: capnp.NewClient(Forwarder_NewServer(s, policy))}
}

// Forwarder_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Forwarder_Methods(methods []server.Method, s Forwarder_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Do(ctx, Actor_do{call})
		},
	})

	return methods
}

type Revoker struct{ Client *capnp.Client }

// Revoker_TypeID is the unique identifier for the type Revoker.
const Revoker_TypeID = 0xbd0ca7ed013308ba

func (c Revoker) Revoke(ctx context.Context, params func(Revoker_revoke_Params) error) (Revoker_revoke_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xbd0ca7ed013308ba,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Revoker",
			MethodName:    "revoke",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Revoker_revoke_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Revoker_revoke_Results_Future{Future: ans.Future()}, release
}
func (c Revoker) Do(ctx context.Context, params func(Actor_do_Params) error) (Actor_do_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Actor_do_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Actor_do_Results_Future{Future: ans.Future()}, release
}

func (c Revoker) AddRef() Revoker {
	return Revoker{
		Client: c.Client.AddRef(),
	}
}

func (c Revoker) Release() {
	c.Client.Release()
}

// A Revoker_Server is a Revoker with a local implementation.
type Revoker_Server interface {
	Revoke(context.Context, Revoker_revoke) error

	Do(context.Context, Actor_do) error
}

// Revoker_NewServer creates a new Server from an implementation of Revoker_Server.
func Revoker_NewServer(s Revoker_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Revoker_Methods(nil, s), s, c, policy)
}

// Revoker_ServerToClient creates a new Client from an implementation of Revoker_Server.
// The caller is responsible for calling Release on the returned Client.
func Revoker_ServerToClient(s Revoker_Server, policy *server.Policy) Revoker {
	return Revoker{Client: capnp.NewClient(Revoker_NewServer(s, policy))}
}

// Revoker_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Revoker_Methods(methods []server.Method, s Revoker_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xbd0ca7ed013308ba,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Revoker",
			MethodName:    "revoke",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Revoke(ctx, Revoker_revoke{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xa27a66204e1f5179,
			MethodID:      0,
			InterfaceName: "capnp/revokable_forwarder.capnp:Actor",
			MethodName:    "do",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Do(ctx, Actor_do{call})
		},
	})

	return methods
}

// Revoker_revoke holds the state for a server call to Revoker.revoke.
// See server.Call for documentation.
type Revoker_revoke struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Revoker_revoke) Args() Revoker_revoke_Params {
	return Revoker_revoke_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Revoker_revoke) AllocResults() (Revoker_revoke_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Revoker_revoke_Results{Struct: r}, err
}

type Revoker_revoke_Params struct{ capnp.Struct }

// Revoker_revoke_Params_TypeID is the unique identifier for the type Revoker_revoke_Params.
const Revoker_revoke_Params_TypeID = 0xb457bead0d4bb698

func NewRevoker_revoke_Params(s *capnp.Segment) (Revoker_revoke_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Revoker_revoke_Params{st}, err
}

func NewRootRevoker_revoke_Params(s *capnp.Segment) (Revoker_revoke_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Revoker_revoke_Params{st}, err
}

func ReadRootRevoker_revoke_Params(msg *capnp.Message) (Revoker_revoke_Params, error) {
	root, err := msg.Root()
	return Revoker_revoke_Params{root.Struct()}, err
}

func (s Revoker_revoke_Params) String() string {
	str, _ := text.Marshal(0xb457bead0d4bb698, s.Struct)
	return str
}

// Revoker_revoke_Params_List is a list of Revoker_revoke_Params.
type Revoker_revoke_Params_List struct{ capnp.List }

// NewRevoker_revoke_Params creates a new list of Revoker_revoke_Params.
func NewRevoker_revoke_Params_List(s *capnp.Segment, sz int32) (Revoker_revoke_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Revoker_revoke_Params_List{l}, err
}

func (s Revoker_revoke_Params_List) At(i int) Revoker_revoke_Params {
	return Revoker_revoke_Params{s.List.Struct(i)}
}

func (s Revoker_revoke_Params_List) Set(i int, v Revoker_revoke_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Revoker_revoke_Params_List) String() string {
	str, _ := text.MarshalList(0xb457bead0d4bb698, s.List)
	return str
}

// Revoker_revoke_Params_Future is a wrapper for a Revoker_revoke_Params promised by a client call.
type Revoker_revoke_Params_Future struct{ *capnp.Future }

func (p Revoker_revoke_Params_Future) Struct() (Revoker_revoke_Params, error) {
	s, err := p.Future.Struct()
	return Revoker_revoke_Params{s}, err
}

type Revoker_revoke_Results struct{ capnp.Struct }

// Revoker_revoke_Results_TypeID is the unique identifier for the type Revoker_revoke_Results.
const Revoker_revoke_Results_TypeID = 0xbc1b102199ee0e6b

func NewRevoker_revoke_Results(s *capnp.Segment) (Revoker_revoke_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Revoker_revoke_Results{st}, err
}

func NewRootRevoker_revoke_Results(s *capnp.Segment) (Revoker_revoke_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Revoker_revoke_Results{st}, err
}

func ReadRootRevoker_revoke_Results(msg *capnp.Message) (Revoker_revoke_Results, error) {
	root, err := msg.Root()
	return Revoker_revoke_Results{root.Struct()}, err
}

func (s Revoker_revoke_Results) String() string {
	str, _ := text.Marshal(0xbc1b102199ee0e6b, s.Struct)
	return str
}

// Revoker_revoke_Results_List is a list of Revoker_revoke_Results.
type Revoker_revoke_Results_List struct{ capnp.List }

// NewRevoker_revoke_Results creates a new list of Revoker_revoke_Results.
func NewRevoker_revoke_Results_List(s *capnp.Segment, sz int32) (Revoker_revoke_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return Revoker_revoke_Results_List{l}, err
}

func (s Revoker_revoke_Results_List) At(i int) Revoker_revoke_Results {
	return Revoker_revoke_Results{s.List.Struct(i)}
}

func (s Revoker_revoke_Results_List) Set(i int, v Revoker_revoke_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s Revoker_revoke_Results_List) String() string {
	str, _ := text.MarshalList(0xbc1b102199ee0e6b, s.List)
	return str
}

// Revoker_revoke_Results_Future is a wrapper for a Revoker_revoke_Results promised by a client call.
type Revoker_revoke_Results_Future struct{ *capnp.Future }

func (p Revoker_revoke_Results_Future) Struct() (Revoker_revoke_Results, error) {
	s, err := p.Future.Struct()
	return Revoker_revoke_Results{s}, err
}

const schema_ff44938a6ce76f11 = "x\xda\x94TQHSQ\x18\xfe\xff\x9ds\xbd\x0a\xce" +
	"<]AJr:F\xa8\xe0\xb2YA\xbeL\x975" +
	"\xc1\x8a\xddE\x84A\xc46\xef\xa4\x9c]\xb9w&\xfa" +
	"\x14Q/\xbe\x95\x04\xd5Ki\x84A\xf4\x12Q$%" +
	"T\xf4\x12\x09\xf9\xd0\x83\x94\x0fAe/\x15\xd4S\x86" +
	"\xdc8g\xbb\xf7\x0e\xe6X\xbem\xf7?\xff\xf7\x7f\xdf" +
	"\xf7\x7f\xe7t<\xc1n\xcfni\xb5\x02@\x8dI\x15" +
	"V\xc7\x96\xfe[\xc3\xd3\xdbn\x02\x0b \x00\x95\x01:" +
	"\x07\xc8E\x04jM\xa8\xbe\xa3M\xe9\xc9Y`M\xc4" +
	"b\xfajfj\xba\xd7\x02\xc0\xce\x83$\x84\xcaq\"" +
	"\x03(*\x89*\x13\xfc\x97u\xfdq\xbf\xf7\xc1\xc2\x89" +
	"G\xc0Zl\x9c\x04y\xc8q\x86k~\xdch\xaem" +
	"xVPQ\xc9<\xaf\xccWv\xe2\xf7\xb9\xea\xe7E" +
	"\x13zH\x04\x15UL8B\xa2\xca\x18\xa9\x07\xb0~" +
	"\xa1<\xf3\xedJ\xff\x0b`;\x11@\xf2p \x8dL" +
	"!\xa02F\xc6\x01\xad\xd7\xf7>6\xde\xad5_\x15" +
	"\xc1-q\xc2\x9f\x04\xdc\x0a\x91\x95\x15\x01\xb76\xd4\xb0" +
	"\xef\xe5\xbb\x0bos\xc2%\xe4p\x8bd\x92\xc3-\x93" +
	"0\xa0\xd57\xdd\xfcY\x9e\xfb\xbaX\x04\xf7\x87\xc4Q" +
	"\xf1r!J\x15\x95\x95*\xca\xe1f\x1b\xe9\xfd\xf8R" +
	"\xdb\xfb\x0dN\xfb\x91\x9f\x03P$\x1aU\xda\xc5\xe9\x0f" +
	"\x97\x9e\xce\xaf\xcf|Y\xcei\x11\x9e\xec\xa0S\xdc\x93" +
	";m\x97kz\xfb\x06\xd6\x0a*^z\x95Wn/" +
	"_;\xb6\xf4\xf3\xec\xdf\xbc~Ax\x9d\xaf\x8a\xc3s" +
	"\xc2o\xc8\xc2\xc9\xdf\x87\xb7\xae\x17Qh\xa5!T\xf6" +
	"\x0b\x0a{iT9E\xeb!b\xa5\x12\xa3\xe7Fw" +
	"\x19Z\xd5y}8\x91\xcch\xa7\xd3\xba1\x9e0\x06" +
	"5#(J]\x11=\x19L\xebz \xae\x99c\x99" +
	"\xac\x09vC\xa9\xf3\xbe\x9eTV7b\x88*%\x12" +
	"\x80\xc3\x17m\xb1\x8cm\x07\x0f\x93d2\xa8wc\x0c" +
	"\xd1\xa1 \x95\x82\x8c\xf3a\x9a\x11\x14C\xb5@\xcc\x97" +
	"0\x12#\xe6\xa6\xfb\xe2\xe1\x9c\x04\xa7\x91\x96i\x04W" +
	"\x85\x9dk\xb4c\xccX\x97P\x11\xce\x81w\xa3J\xb1" +
	" ,\x00\xe5\xe9\xf5d\xce\xa4\xb4\xa0\xa9e\x031!" +
	"\x08@\xad$\x14\x80\"\x00k\xf5\x03\xa8\x01\x82j\x87" +
	"\x07\x19b\x1d\xf2\x8f\xed!\x00\xb5\x85\xa0\xba\xc7\x83r" +
	"RO\"s\x03\x07\x88\x0c\xd0\x97J\x18z\x06\x99{" +
	"\x0br\xdf\x1d6\xa4\xe4\xde\x0e\xf0\xce\x18b\x8cHB" +
	"\x8bs\xf1\x0b\xb4\xd0r)\x09\xe7\xa4\xa8\xd4\x11\xe2\x0d" +
	"\x09]\xa8\xd6y\xca\xb3+\x89\x7f(\xff\x1f\x8d\x02\x86" +
	".\x0a\x94\xd7'G\xf4\xa4\xbbO\xfb\xda\xa3\xfd\xf01" +
	"\xe6\x17\xfb\x94\xd3\xba\x9e_\xe6F\x06\x94^&O}" +
	"p\xd0\xbd'\x9b\x0a@\xae\x09\xcdM\xcc\xc9g\xa6\xd0" +
	"i\xbf\xeb\xb4<b\x0ea5x\xb0\xfa\xbfv/\x98" +
	"\xb8\xee\xd8o,\xda\xcf\x90\xe3\x8e\xa9e\x8b\xdd\xf9\x17" +
	"\x00\x00\xff\xff\x9b\x9d\xee\xe9"

func init() {
	schemas.Register(schema_ff44938a6ce76f11,
		0x9a19936b9f4b0f30,
		0xa27a66204e1f5179,
		0xb457bead0d4bb698,
		0xbc1b102199ee0e6b,
		0xbd0ca7ed013308ba,
		0xc14b91e8a10701f2,
		0xc37310a51edca8c5,
		0xcc80d0c2361b67fa,
		0xcde6a707e4219348,
		0xd52ad152ab041ea2,
		0xd9e5a1fdbab984db,
		0xfa5948440e852aa3,
		0xfb6aefd15395d9a0,
		0xfd124cf35abe03ca)
}
