package main

import "fmt"

func test(fn func() int) int {
	return fn()
}

// FormatFunc ...
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func funcSign() {
	// 直接将匿名函数当参数
	s1 := test(func() int { return 100 })

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(s1, s2)
}
