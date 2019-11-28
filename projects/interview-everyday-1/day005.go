package main

import "fmt"

const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func day005Test1() {

}

func day005() {
	const c1 = 123
	const c2 = 1.23
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)

	fmt.Printf("%T %v\n", y, y) // uint16 120
	fmt.Printf("%T %v\n", z, z) // string abc
}
