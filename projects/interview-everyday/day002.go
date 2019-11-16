package main

import "fmt"

var (
	// := 仅适用于局部变量, 不适用于全局变量和全局常量
	// size := 1024
	size    = 1024
	maxSize = size * 2
)

func day002() {
	list := new([]int)
	// first argument to append must be slice; have *[]int
	// list = append(list, 1)
	*list = append(*list, 1)
	fmt.Println(list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	// cannot use s2 (type []int) as type int in append
	// s1 = append(s1, s2)
	s1 = append(s1, s2...)
	fmt.Println(s1)

	fmt.Println(size, maxSize)
}
