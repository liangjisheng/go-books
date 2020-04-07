package main

import (
	"github.com/dynport/gossh"

	"log"
)

// MakeLogger returns a function of type gossh.Writer func(...interface{})
// MakeLogger just adds a prefix (DEBUG, INFO, ERROR)
func MakeLogger(prefix string) gossh.Writer {
	return func(args ...interface{}) {
		log.Println((append([]interface{}{prefix}, args...))...)
	}
}

func main() {
	client := gossh.New("117.51.148.112", "dc2-user")
	// my default agent authentication is used. use
	// client.SetPassword("<secret>")
	// client.SetPrivateKey("<PrivateKey path>") if not set this path will read this prv key:$HOME/.ssh/id_rsa
	// for password authentication
	client.DebugWriter = MakeLogger("DEBUG")
	client.InfoWriter = MakeLogger("INFO ")
	client.ErrorWriter = MakeLogger("ERROR")

	defer client.Close()

	rsp, e := client.Execute("uptime")
	if e != nil {
		client.ErrorWriter(e.Error())
	}
	client.InfoWriter(rsp.String())

	rsp, e = client.Execute("echo -n $(cat /proc/version); cat /does/not/exists")
	if e != nil {
		client.ErrorWriter(e.Error())
		client.ErrorWriter("STDOUT: " + rsp.Stdout())
		client.ErrorWriter("STDERR: " + rsp.Stderr())
	}
}
