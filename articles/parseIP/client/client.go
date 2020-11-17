package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

// Response ...
type Response struct {
	Country   string
	Province  string
	City      string
	ISP       string
	Latitude  float64
	Longitude float64
	TimeZone  string
}

// Agrs  ...
type Agrs struct {
	IPString string
}

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:3344")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	var res Response
	err = client.Call("IP2addr.Address", Agrs{"222.130.243.164"}, &res)
	if err != nil {
		log.Fatal("ip2addr error:", err)
	}
	fmt.Println(res)
}
