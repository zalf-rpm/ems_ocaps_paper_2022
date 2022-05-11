package main

import (
	"context"
	"fmt"
	"log"

	"capnproto.org/go/capnp/v3/rpc"
	"github.com/zalf-rpm/ems_ocaps_paper_2022/go/capnp"
	"github.com/zalf-rpm/ems_ocaps_paper_2022/go/helper"
)

const defaultPort = 9991

func main() {
	fmt.Println("Hello I'm Alice")

	// read args, listen to incomming connections
	l, err := helper.ListenerFromCmdlArgs(defaultPort)
	if err != nil {
		log.Fatal(err)
	}

	// error reporter for brocken connections
	errChan := helper.ErrorReporter()

	var instanceOfAlice int
	for {
		// accept incomming connections
		c, err := l.Accept()
		fmt.Printf("Alice: request from %v\n", c.RemoteAddr())
		if err != nil {
			errChan <- err
			continue
		}

		a := alice{
			name:   "Alice",
			number: instanceOfAlice,
			bob:    nil,
			carol:  nil,
		}
		instanceOfAlice++

		// Listen for calls, using  bootstrap interface (in a new go thread)
		main := capnp.Alice_ServerToClient(&a, nil)
		rpc.NewConn(rpc.NewPackedStreamTransport(c), &rpc.Options{BootstrapClient: main.Client, ErrorReporter: &helper.ConnError{Out: errChan}})
		// this connection will be close when the client closes the connection
	}
}

type alice struct {
	name   string
	number int
	bob    *capnp.Bob
	carol  *capnp.Carol
}

// Alice_Server interface
func (self *alice) Set(c context.Context, call capnp.Alice_set) error {
	fmt.Println("@Alice::set")
	bob := call.Args().Bob().AddRef()
	carol := call.Args().Carol().AddRef()
	self.bob = &bob
	self.carol = &carol

	return nil
}

func (self *alice) Do(c context.Context, call capnp.Actor_do) error {

	msg, err := call.Args().Msg()
	if err != nil {
		return err
	}
	fmt.Println("@Alice::do | msg:", msg)
	fmt.Println("@Alice::do | sending foo(carol) to Bob")
	if self.bob != nil && self.carol != nil {

		resultFuture, release := self.bob.Foo(context.Background(), func(b capnp.Bob_foo_Params) error {
			err := b.SetCarol(*self.carol)
			return err
		})
		defer release()
		_, err := resultFuture.Struct()
		if err != nil {
			return err
		}
	}

	return nil
}
