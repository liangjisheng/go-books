package server

import (
	"fmt"
	"net"
)

// Server ...
type Server struct {
	listener net.Listener
}

// NewServer ...
func NewServer(address ...string) []*Server {
	s := make([]*Server, 0)
	for _, a := range address {
		lst, _ := net.Listen("tcp", a)
		s = append(s, &Server{
			listener: lst,
		})
	}
	return s
}

// Run ...
func (s *Server) Run() {
	for {
		conn, _ := s.listener.Accept()
		go s.handler(conn)
	}
}

func (s *Server) handler(conn net.Conn) {
	for {
		msg := make([]byte, 1024)
		_, _ = conn.Read(msg)
		fmt.Println(string(msg))
		_, _ = conn.Write([]byte(fmt.Sprintf("message from Server: %s", s.listener.Addr().String())))
	}
}
