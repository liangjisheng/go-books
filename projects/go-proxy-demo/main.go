package main

import (
	"goproxydemo/client"
	"goproxydemo/proxy"
	"goproxydemo/server"
)

func main() {
	servers := server.NewServer("127.0.0.1:9998", "127.0.0.1:9999")
	for _, s := range servers {
		go s.Run()
	}

	p := proxy.NewProxy("127.0.0.1:9000", "127.0.0.1:9999")
	go p.Run()

	for _, c := range client.NewClient("127.0.0.1:9000", 3) {
		go c.Run()
	}
	select {}
}
