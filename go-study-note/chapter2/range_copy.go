package main

import "fmt"

func rangeCopy() {
	a := [3]int{0, 1, 2}

	// range 会复制对象
	for i, v := range a {
		if i == 0 {
			a[1], a[2] = 999, 999
			fmt.Println(a)
		}

		a[i] = v + 100
	}

	fmt.Println(a)

	// slice 不会被复制
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		if i == 0 {
			s = s[:3]
			s[2] = 100
			v++
		}

		println(i, v)
	}

	for i := 0; i < len(s); i++ {
		print(s[i])
	}
	println()
}

func rangeCopy1() {
	array := []int{10, 20, 30}
	for _, v := range array {
		v++
	}
	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		array[i]++
	}
	fmt.Println(array)
}
