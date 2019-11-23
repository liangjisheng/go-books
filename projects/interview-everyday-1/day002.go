package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

func day002Test1() {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice) // 1 2 0 0 0
	change(slice[0:2]...)
	fmt.Println(slice) // 1 2 3 0 0
}

func day002Test2() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) // 1 12 13 4 5
	fmt.Println("a = ", a) // 1 12 13 4 5
}

func day002() {
	// day002Test1()
	day002Test2()
}
