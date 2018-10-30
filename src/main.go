package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang/protobuf/proto"

	"./azsb"

	"github.com/golang/protobuf/ptypes"
)

func main() {
	// timestamp.TimestampProto(time.Now())
	blob := &azsb.Blob{
		Data: map[string]*azsb.Packet{
			"One": &azsb.Packet{
				Raw: []byte("A new raw message"),
				Trend: []*azsb.PV{
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 1000},
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 200},
				},
			},
			"Two": &azsb.Packet{
				Raw: []byte("latest message"),
				Trend: []*azsb.PV{
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 300},
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 400},
				},
			},
			"Three": &azsb.Packet{
				Raw: []byte("newest message"),
				Trend: []*azsb.PV{
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 300},
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 400},
				},
			},
			"Four": &azsb.Packet{
				Raw: []byte("message from ch4"),
				Trend: []*azsb.PV{
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 300},
					&azsb.PV{Ts: ptypes.TimestampNow(), Count: 4000},
				},
			},
		},
	}

	fmt.Println(blob)
	fname := "test.dat"
	// out, _ := proto.Marshal(blob)
	out, _ := blob.Bytes()
	fmt.Println(string(out), len(out))
	ioutil.WriteFile(fname, out, 0644)
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	blob2 := &azsb.Blob{}
	proto.Unmarshal(in, blob2)
	fmt.Println(blob2)
	for id, packet := range blob2.Data {
		fmt.Println(id, packet.Trend[0].Count)
		fmt.Println(ptypes.Timestamp(packet.Trend[0].Ts))
		fmt.Println(ptypes.TimestampString(packet.Trend[0].Ts))
		fmt.Println(time.Unix(packet.Trend[0].Ts.Seconds, int64(packet.Trend[0].Ts.Nanos)).UTC())
		for i, pv := range packet.Trend {
			fmt.Println("\t", i, pv)
		}
	}
	fmt.Println(blob2.Data["One"])
}
