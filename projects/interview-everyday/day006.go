package main

import "fmt"

func day006Test1() {
	a := 5
	b := 8.1
	// invalid operation: a + b (mismatched types int and float64)
	// fmt.Println(a + b)
	fmt.Println(float64(a) + b)
}

func day006Test2() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
	// Output
	// 4
}

func day006Test3() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	// invalid operation: a == b (mismatched types [2]int and [3]int)
	// if a == b {
	// 	fmt.Println("equal")
	// } else {
	// 	fmt.Println("not equal")
	// }
}

func day006() {
	// day006Test1()
	// day006Test2()
}
