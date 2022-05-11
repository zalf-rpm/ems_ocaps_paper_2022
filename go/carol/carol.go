package main

import (
	"context"
	"fmt"
	"log"

	"capnproto.org/go/capnp/v3/rpc"
	"github.com/zalf-rpm/ems_ocaps_paper_2022/go/capnp"
	"github.com/zalf-rpm/ems_ocaps_paper_2022/go/helper"
)

const defaultPort = 9993

func main() {
	fmt.Println("May I introduce myself. I'm Carol.")

	// read args, listen to incomming connections
	l, err := helper.ListenerFromCmdlArgs(defaultPort)
	if err != nil {
		log.Fatal(err)
	}

	// error reporter for brocken connections
	errChan := helper.ErrorReporter()

	var instanceOfCarol int
	for {
		// accept incomming connections
		c, err := l.Accept()
		fmt.Printf("Carol: request from %v\n", c.RemoteAddr())
		if err != nil {
			errChan <- err
			continue
		}

		a := carol{
			name:   "Bob",
			number: instanceOfCarol,
		}
		instanceOfCarol++

		// Listen for calls, using  bootstrap interface (in a new go thread)
		main := capnp.Carol_ServerToClient(&a, nil)
		rpc.NewConn(rpc.NewPackedStreamTransport(c), &rpc.Options{BootstrapClient: main.Client, ErrorReporter: &helper.ConnError{Out: errChan}})
		// this connection will be close when the client closes the connection
	}
}

type carol struct {
	name   string
	number int
}

// Carol_Server interface
func (b *carol) Do(c context.Context, call capnp.Actor_do) error {
	msg, err := call.Args().Msg()
	if err != nil {
		return err
	}
	fmt.Println("@Carol::do | msg:", msg)
	return nil
}
