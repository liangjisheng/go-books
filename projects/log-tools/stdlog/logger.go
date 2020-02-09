package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func selfLogger() {
	u := user{
		Name: "ljs",
		Age:  18,
	}

	buf := &bytes.Buffer{}

	// 第一个参数为io.Writer 我们可以使用io.MultiWriter实现多目的地输出
	logger := log.New(buf, "", log.Lshortfile|log.LstdFlags)
	logger.Printf("%s login, age:%d", u.Name, u.Age)
	fmt.Print(buf.String())

	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logger = log.New(io.MultiWriter(writer1, writer2, writer3), "", log.Lshortfile|log.LstdFlags)
	logger.Printf("%s login, age:%d", u.Name, u.Age)
	fmt.Print(writer1.String())
}
