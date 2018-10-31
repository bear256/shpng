package azsb

import (
	"github.com/golang/protobuf/proto"
)

// Fetch a packet from blob
func (blob *Blob) Fetch(id string) []byte {
	return blob.Packets[id]
}

// Add a packet into blob
func (blob *Blob) Add(id string, packet []byte) {
	blob.Packets[id] = packet
}

// Bytes serialize blob
func (blob *Blob) Bytes() ([]byte, error) {
	return proto.Marshal(blob)
}
