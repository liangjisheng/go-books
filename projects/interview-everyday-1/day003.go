package main

import "fmt"

func day003Test1() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
	// Output
	// 2 or 3
	// r range map 是无序的 如果第一次循环到 A 则输出 3 否则输出 2
}

func day003() {
	day003Test1()
}
