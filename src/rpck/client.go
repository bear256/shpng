package rpck

import (
	"net/rpc"

	"github.com/xtaci/kcp-go"
)

// Client is defined for communicating with server
type Client struct {
	*rpc.Client
}

// NewClient returns an asdb client instance
func NewClient(raddr string) (*Client, error) {
	conn, err := kcp.Dial(raddr)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClient(conn)
	return &Client{client}, nil
}
