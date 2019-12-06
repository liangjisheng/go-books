package main

import "fmt"

const (
	// Century ...
	Century = 100
	// Decade ...
	Decade = 010
	// Year ...
	Year = 001
)

func day009Test1() {
	// 八进制数以 0 开头 十六进制数以 0x 开头 所以 Decade 表示十进制的 8
	fmt.Println(Century + 2*Decade + 2*Year) // 118

	// ^ 作为二元运算符时表示按位异或
	fmt.Println(3^2+4^2 == 5^2)
	fmt.Println(3*3+4*4 == 5*5)

	fmt.Println(6^2+8^2 == 10^2)
	fmt.Println(6*6+8*8 == 10*10)
}

func day009() {
	// day009Test1()

	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}
