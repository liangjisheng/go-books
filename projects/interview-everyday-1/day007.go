package main

import "fmt"

func day007Test1() {
	var fn1 = func() {}
	var fn2 = func() {}

	// invalid operation: fn1 != fn2 (func can only be compared to nil)
	// if fn1 != fn2 {
	// 	println("fn1 not equal fn2")
	// }

	if fn1 == nil {
		println("fn1 is nil")
	} else {
		println("fn1 is not nil") // not nil
	}
	if fn2 == nil {
		println("fn2 is nil")
	} else {
		println("fn2 is not nil") // not nil
	}
}

type mapT1 struct {
	n int
}

func day007Test2() {
	m1 := make(map[int]mapT1)
	// cannot assign to struct field m[0].n in map
	// map[key]struct 中 struct 是不可寻址的 所以无法直接赋值
	// m1[0].n = 1
	fmt.Println(m1)

	m2 := make(map[int]int)
	m2[0] = 0
	fmt.Println(m2)
}

func day007() {
	// day007Test1()
	day007Test2()
}
