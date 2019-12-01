package main

import "fmt"

func day006Test1() {
	var x int8 = -128
	fmt.Printf("%08b, %d\n", x, x) // -10000000, -128
	var y = x / -1
	fmt.Printf("%08b, %d\n", y, y) // -10000000, -128
}

// F ...
func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

// defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中
// 到真正执行 defer() 的时候取出
func day006Test2() {
	f := F(5)
	defer func() {
		fmt.Println(f())
	}()
	defer fmt.Println(f())

	i := f()
	fmt.Println(i)
}

func day006() {
	// day006Test1()
	day006Test2() // 7 6 8
}
