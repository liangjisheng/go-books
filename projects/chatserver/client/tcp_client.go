package client

import (
	"go-demos/go-chat/chatserver/protocol"
	"log"
	"net"
)

// TCPChatClient 里面设置一些私有变量用于跟踪连接的conn
type TCPChatClient struct {
	conn      net.Conn
	cmdReader *protocol.CommandReader
	cmdWriter *protocol.CommandWriter
	name      string
	error     chan error
	incoming  chan protocol.MessageCommand
}

// NewClient ...
func NewClient() *TCPChatClient {
	return &TCPChatClient{
		incoming: make(chan protocol.MessageCommand),
	}
}

// Dial 建立连接并且创建通讯协议的reader和writer
func (c *TCPChatClient) Dial(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println("dial connect err:", err)
		return err
	}

	c.conn = conn
	c.cmdReader = protocol.NewCommandReader(conn)
	c.cmdWriter = protocol.NewCommandWriter(conn)
	return nil
}

// Start ...
func (c *TCPChatClient) Start() {
	for {
		cmd, err := c.cmdReader.Read()
		if err != nil {
			c.error <- err
			break
		}

		if cmd != nil {
			switch v := cmd.(type) {
			case protocol.MessageCommand:
				c.incoming <- v
			default:
				log.Printf("Unknown command: %v", v)
			}
		}
	}
}

// Close ...
func (c *TCPChatClient) Close() {
	c.conn.Close()
}

// Incoming ...
func (c *TCPChatClient) Incoming() chan protocol.MessageCommand {
	return c.incoming
}

// Error ...
func (c *TCPChatClient) Error() chan error {
	return c.error
}

// Send ...
func (c *TCPChatClient) Send(command interface{}) error {
	return c.cmdWriter.Write(command)
}

// SetName ...
func (c *TCPChatClient) SetName(name string) error {
	return c.Send(protocol.NameCommand{
		Name: name,
	})
}

// SendMessage ...
func (c *TCPChatClient) SendMessage(message string) error {
	return c.Send(protocol.SendCommand{
		Message: message,
	})
}
