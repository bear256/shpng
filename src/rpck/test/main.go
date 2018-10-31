package main

import (
	"../../rpck"
)

func main() {
	server := rpck.NewServer()
	server.Serve(rpck.DefaultPort())
}
