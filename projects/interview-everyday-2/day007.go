package main

import "fmt"

// recover() 必须在 defer() 函数中调用才有效
// 在调用 defer() 时 便会计算函数的参数并压入栈中
func day007Test1() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer fmt.Print(recover())
		panic(1)
	}()
	defer recover()
	panic(2)
}

func day007Test2() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer func() {
			fmt.Print(recover())
		}()
		panic(1)
	}()
	defer recover()
	panic(2)
}

func day007() {
	// day007Test1() // 21
	day007Test2() // 12
}
