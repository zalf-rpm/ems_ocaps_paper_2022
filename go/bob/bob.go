package main

import (
	"context"
	"fmt"
	"log"

	"capnproto.org/go/capnp/v3/rpc"
	"github.com/zalf-rpm/ems_ocaps_paper_2022/go/capnp"
	"github.com/zalf-rpm/ems_ocaps_paper_2022/go/helper"
)

const defaultPort = 9992

func main() {
	fmt.Println("Nice to meet you. I'm Bob.")

	// read args, listen to incomming connections
	l, err := helper.ListenerFromCmdlArgs(defaultPort)
	if err != nil {
		log.Fatal(err)
	}

	// error reporter for brocken connections
	errChan := helper.ErrorReporter()

	var instanceOfBob int
	for {
		// accept incomming connections
		c, err := l.Accept()
		fmt.Printf("Bob: request from %v\n", c.RemoteAddr())
		if err != nil {
			errChan <- err
			continue
		}

		a := bob{
			name:   "Bob",
			number: instanceOfBob,
			carol:  nil,
		}
		instanceOfBob++
		// Listen for calls, using  bootstrap interface (in a new go thread)
		main := capnp.Bob_ServerToClient(&a, nil)
		rpc.NewConn(rpc.NewPackedStreamTransport(c), &rpc.Options{BootstrapClient: main.Client, ErrorReporter: &helper.ConnError{Out: errChan}})
		// this connection will be close when the client closes the connection
	}
}

type bob struct {
	name   string
	number int
	carol  *capnp.Carol
}

// Bob_Server interface
func (self *bob) Foo(c context.Context, call capnp.Bob_foo) error {
	// give Bob a new Interface of Carol
	fmt.Println("@Bob::foo")
	carol := call.Args().Carol().AddRef()
	self.carol = &carol
	return nil
}

func (self *bob) Do(c context.Context, call capnp.Actor_do) error {

	msg, err := call.Args().Msg()
	if err != nil {
		return err
	}
	fmt.Println("@Bob::do | msg:", msg)

	// forward the message send to Bob to Carol
	if self.carol != nil {
		resultFuture, release := self.carol.Do(context.Background(), func(a capnp.Actor_do_Params) error {
			return a.SetMsg("<Bobs DO message to Carol>") // ??? I think this should be a msg
		})
		defer release()
		_, err := resultFuture.Struct()
		if err != nil {
			return err
		}
	}
	return nil
}
