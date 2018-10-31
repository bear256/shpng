package rpck

import (
	"fmt"
	"net/rpc"

	kcp "github.com/xtaci/kcp-go"
)

// Server is defined for serving asdb clients
type Server struct {
	*rpc.Server
}

// NewServer returns an adbs server instance
func NewServer() *Server {
	server := rpc.NewServer()
	return &Server{server}
}

// Serve method to make server listening on a port with kcp listener
func (s *Server) Serve(laddr string) error {
	listener, err := kcp.Listen(laddr)
	if err != nil {
		return err
	}
	fmt.Println("Listening on", listener.Addr())
	s.Accept(listener)
	return nil
}

func (s *Server) RegisterMap(rcvrs map[string]interface{}) error {
	for name, rcvr := range rcvrs {
		err := s.RegisterName(name, rcvr)
		if err != nil {
			return err
		}
	}
	return nil
}

func DefaultPort() string {
	sum := 0
	for _, c := range "rpck" {
		idx := int(c - 'a')
		sum = sum*10 + (idx+1)%4
	}
	return fmt.Sprintf(":%d", sum%(65536-1024)+1024)
}
