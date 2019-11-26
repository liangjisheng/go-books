package server

// ChatServer interface 能让这个server的行为更加清晰明了
type ChatServer interface {
	Listen(address string) error
	Broadcast(command interface{}) error
	Start()
	Close()
}
