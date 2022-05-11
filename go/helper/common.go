package helper

import (
	"flag"
	"fmt"
	"net"
)

// ErrorReporter interface
type ConnError struct {
	Out chan<- error
}

func (cerr *ConnError) ReportError(err error) {
	cerr.Out <- err
}

// Init - read args, open connection
func ListenerFromCmdlArgs(defaultPort int) (net.Listener, error) {
	// read cmd line args
	addr := flag.String("addr", "localhost", "address of this process")
	port := flag.Int("port", defaultPort, "port")
	flag.Parse()

	// listen to incomming connections
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *addr, *port))
	return l, err
}

// ErrorReporter
func ErrorReporter() chan error {
	// error reporter for brocken connections
	errChan := make(chan error)
	go func(errC <-chan error) {
		for {
			err := <-errC
			fmt.Println(err)
		}
	}(errChan)
	return errChan
}
