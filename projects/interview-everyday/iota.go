package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func iotatest() {
	fmt.Println(x, y, z, k, p)
	// Output
	// 0 2 zz zz 5
}
