package client

import (
	"fmt"
	"net"
)

// Client ...
type Client struct {
	conn net.Conn
}

// NewClient ...
func NewClient(proxyAddress string, num int) []*Client {
	c := make([]*Client, 0)
	for i := 0; i < num; i++ {
		conn, _ := net.Dial("tcp", proxyAddress)
		c = append(c, &Client{conn: conn})
	}
	return c
}

// Run ...
func (c *Client) Run() {
	_, _ = c.conn.Write([]byte(fmt.Sprintf("message from client: %s", c.conn.LocalAddr().String())))
	msg := make([]byte, 1024)
	c.conn.Read(msg)
	fmt.Println(string(msg))
}
