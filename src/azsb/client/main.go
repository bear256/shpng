package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"../../azsb"

	"github.com/xtaci/kcp-go"
)

func writeBlob(client *rpc.Client, container, blob, data string) {
	req := &azsb.Req{
		Params: map[string]interface{}{
			"container": container,
			"blob":      blob,
		},
		Body: []byte(data),
	}
	resp := &azsb.Resp{}
	err := client.Call("azsb.Write", req, resp)
	if err != nil {
		log.Fatalln(err)
	}
}

func readBlob(client *rpc.Client, container, blob string) {
	req := &azsb.Req{
		Params: map[string]interface{}{
			"container": container,
			"blob":      blob,
		}}
	resp := &azsb.Resp{}
	err := client.Call("azsb.Read", req, resp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(resp.Body.([]byte)))
}

func main() {
	conn, err := kcp.Dial("localhost:10000")
	if err != nil {
		log.Fatalln(err)
	}
	client := rpc.NewClient(conn)
	writeBlob(client, os.Args[1], os.Args[2], os.Args[3])
	readBlob(client, os.Args[1], os.Args[2])
}
