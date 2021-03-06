package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename := "../my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tails %+v file close reopen,filenam:%s\n", tails, filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg.Text)
	}
}
