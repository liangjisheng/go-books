package server

import (
	"errors"
	"go-demos/go-chat/chatserver/protocol"
	"io"
	"log"
	"net"
	"sync"
)

type client struct {
	conn   net.Conn
	name   string
	writer *protocol.CommandWriter
}

// TCPChatServer ...
type TCPChatServer struct {
	listener net.Listener
	clients  []*client
	mutex    *sync.Mutex
}

var (
	UnknownClient = errors.New("Unknown client")
)

// NewServer ...
func NewServer() *TCPChatServer {
	return &TCPChatServer{
		mutex: &sync.Mutex{},
	}
}

// Listen ...
func (s *TCPChatServer) Listen(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("server listen err:", err)
		return err
	}

	s.listener = l
	log.Printf("Listening on %v\n", address)
	return nil
}

// Close ...
func (s *TCPChatServer) Close() {
	s.listener.Close()
}

// Start ...
func (s *TCPChatServer) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("connect err:", err)
			continue
		}

		client := s.accept(conn)
		go s.serve(client)
	}
}

// Broadcast ...
func (s *TCPChatServer) Broadcast(command interface{}) error {
	for _, client := range s.clients {
		err := client.writer.Write(command)
		if err != nil {
			log.Printf("write command %+v err %+v", command, err)
			continue
		}
	}
	return nil
}

func (s *TCPChatServer) accept(conn net.Conn) *client {
	log.Printf("Accepting connection from %v, total clients: %v", conn.RemoteAddr().String(), len(s.clients)+1)
	s.mutex.Lock()
	defer s.mutex.Unlock()
	client := &client{
		conn:   conn,
		writer: protocol.NewCommandWriter(conn),
	}
	s.clients = append(s.clients, client)
	return client
}

func (s *TCPChatServer) remove(client *client) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// remove the connections from clients array
	for i, check := range s.clients {
		if check == client {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
		}
	}

	log.Printf("Closing connection from %v", client.conn.RemoteAddr().String())
	client.conn.Close()
}

func (s *TCPChatServer) serve(client *client) {
	cmdReader := protocol.NewCommandReader(client.conn)
	defer s.remove(client)

	for {
		cmd, err := cmdReader.Read()
		if err != nil && err != io.EOF {
			log.Printf("Read error: %v", err)
			break
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.SendCommand:
				go s.Broadcast(protocol.MessageCommand{
					Message: v.Message,
					Name:    client.name,
				})

			case protocol.NameCommand:
				client.name = v.Name
			}
		}

		if err == io.EOF {
			break
		}
	}
}
