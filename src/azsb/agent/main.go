package main

import (
	"log"
	"net/rpc"
	"os"

	"../../azsb"
	"../../config"

	kcp "github.com/xtaci/kcp-go"
)

func main() {
	accountName, accountKey := config.Toml().StorageAccountInfo()
	blobs := azsb.NewBlobs(accountName, accountKey)
	agent := azsb.NewRPC(blobs)
	server := rpc.NewServer()
	server.RegisterName("azsb", agent)
	listener, err := kcp.Listen(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	server.Accept(listener)
}
