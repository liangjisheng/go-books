package main

import (
	"fmt"
	"net"
	"time"
)

// DetecdIPAndPort 检测本机是否可以和给定的ip端口建立TCP连接
func DetecdIPAndPort(ip, port string) bool {
	address := net.JoinHostPort(ip, port)
	// 3 秒超时
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return false
	}
	if conn != nil {
		_ = conn.Close()
		return true
	}
	return false
}

func tcpGather(ip string, ports []string) map[string]string {
	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results[port] = "failed"
			// todo log handler
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}
	return results
}

func main() {
	fmt.Printf("%v\n", tcpGather("127.0.0.1", []string{"8080"}))
}
