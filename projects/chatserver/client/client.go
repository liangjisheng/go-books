package client

import (
	"go-demos/go-chat/chatserver/protocol"
)

// ChatClient 客户端通过Dial()连接到服务端，Start() Close()负责停止和关闭服务
// Send()用于发送指令。SetName()和SendMessage()负责设置用户名以及发送消息的逻辑封装
// 最后Incoming()返回一个channel,作为和服务端建立起来作为通讯的连接通道
type ChatClient interface {
	Dial(address string) error
	Send(command interface{}) error
	SendMessage(message string) error
	SetName(name string) error
	Start()
	Close()
	Error() chan error
	Incoming() chan protocol.MessageCommand
}
