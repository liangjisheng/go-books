package proxy

import (
	"io"
	"net"
)

// Proxy ...
type Proxy struct {
	listener       net.Listener
	backendAddress string
}

// NewProxy ..
func NewProxy(agentAddress, serverAddress string) *Proxy {
	lst, _ := net.Listen("tcp", agentAddress)
	return &Proxy{
		listener:       lst,
		backendAddress: serverAddress,
	}
}

// Run ...
func (p *Proxy) Run() {
	for {
		frontConn, _ := p.listener.Accept()
		backendConn, _ := net.Dial("tcp", p.backendAddress)
		go p.Agent(frontConn, backendConn)
	}
}

// Agent ...
func (p *Proxy) Agent(frontConn, backendConn net.Conn) {
	go io.Copy(backendConn, frontConn)
	go io.Copy(frontConn, backendConn)
}
