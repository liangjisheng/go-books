package main

import (
	"go-demos/go-chat/chatserver/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen("127.0.0.1:8080")
	s.Start()
}
