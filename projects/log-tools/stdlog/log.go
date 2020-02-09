package main

import "log"

func demo() {
	u := user{
		Name: "ljs",
		Age:  18,
	}

	log.SetPrefix("Login: ")
	log.Printf("%s login, age:%d", u.Name, u.Age)

	// 设置选项可在每条输出的文本前增加一些额外信息，如日期时间、文件名等
	// log库提供了 6 个选项
	// src/log/log.go
	// const (
	// 	Ldate = 1 << iota
	// 	Ltime
	// 	Lmicroseconds
	// 	Llongfile
	// 	Lshortfile
	// 	LUTC
	// )

	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	log.Printf("%s login, age:%d", u.Name, u.Age)
}
