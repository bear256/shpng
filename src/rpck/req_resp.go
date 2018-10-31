package rpck

// Req is the input argument of RPC method
type Req struct {
	Params map[string]interface{}
	Body   []byte
}

// Resp is the output argument of RPC method
type Resp struct {
	Body []byte
}
