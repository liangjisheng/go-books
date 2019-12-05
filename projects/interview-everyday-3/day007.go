package main

import "fmt"

// 当 i 的值为 0、128 是会发生相等情况，注意 byte 是 uint8 的别名
func day007Test1() {
	count := 0
	for i := range [256]struct{}{} {
		m, n := byte(i), int8(i)
		if n == -n {
			fmt.Println("i:", i, "n:", n)
			count++
		}
		if m == -m {
			fmt.Println("i:", i, "m:", m)
			count++
		}
	}
}

func day007Test2() {
	a := [3]int{0, 1, 2}
	s := a[1:2]

	s[0] = 11
	s = append(s, 12)
	s = append(s, 13) // cap 不足造成重新分配内存
	s[0] = 21

	fmt.Println(a) // 0 11 12
	fmt.Println(s) // 21 12 13
}

func day007() {
	// day007Test1()
	day007Test2()
}
