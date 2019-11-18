package main

import "fmt"

type person struct {
	name string
}

func day005Test1() {
	var m map[person]int
	if m == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	// Output
	// nil

	p := person{"mike"}
	fmt.Println(m[p])
	// Output
	// 0
}

func day005() {
	day005Test1()
	day005Test2()
}

func day005f(num ...int) {
	num[0] = 18
}

func day005Test2() {
	i := []int{5, 6, 7}
	day005f(i...)
	fmt.Println(i[0])
	// Output
	// 18
}
