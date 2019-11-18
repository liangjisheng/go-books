package main

import "fmt"

func hello() []string {
	return nil
}

func day004() {
	// 将 hello() 赋值给变量 h，而不是函数的返回值
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	// Output
	// not nil

	// 将函数的返回值赋值给变量 h1
	h1 := hello()
	if h1 == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	// Output
	// nil
}
