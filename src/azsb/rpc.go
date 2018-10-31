package azsb

import "fmt"

// Req is the input argument of RPC method
type Req struct {
	Params map[string]interface{}
	Body   []byte
}

// Resp is the output argument of RPC method
type Resp struct {
	Body interface{}
}

// RPC is defined for net/rpc
type RPC struct {
	blobs *Blobs
}

// NewRPC returns a RPC instance
func NewRPC(blobs *Blobs) *RPC {
	return &RPC{blobs}
}

func (rpc *RPC) Read(req *Req, resp *Resp) error {
	container, blob, err := checkoutBlobParams(req)
	if err != nil {
		return err
	}
	data, err := rpc.blobs.Read(container, blob)
	resp.Body = data
	return err
}

func (rpc *RPC) Write(req *Req, resp *Resp) error {
	container, blob, err := checkoutBlobParams(req)
	if err != nil {
		return err
	}
	err = rpc.blobs.Write(container, blob, req.Body)
	return err
}

func checkoutBlobParams(req *Req) (string, string, error) {
	var err error
	container, has := req.Params["container"].(string)
	if !has {
		err = fmt.Errorf("missing container in request")
	}
	blob, has := req.Params["blob"].(string)
	if !has {
		err = fmt.Errorf("missing blob in request")
	}
	return container, blob, err
}
