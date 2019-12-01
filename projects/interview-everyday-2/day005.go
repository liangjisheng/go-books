package main

import "fmt"

func day005Test1() {
	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k) // 9

	for k = 0; k < 3; k++ {
	}
	fmt.Println(k) // 3

	for k = range (*[3]int)(nil) {
	}
	fmt.Println(k) // 2
}

func day005Test2() {
	nil := 123
	fmt.Println(nil)
	// 前作用域中 预定义的 nil 被覆盖 此时 nil 是 int 类型值 不能赋值给 map 类型
	// cannot use nil (type int) as type map[string]int in assignment
	// var _ map[string]int = nil
}

func day005() {
	day005Test1()
}
