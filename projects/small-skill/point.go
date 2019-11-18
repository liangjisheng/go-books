package main

import (
	"fmt"
)

// People ...
type People interface {
	Show()
}

// Student ...
type Student struct{}

// Show ...
func (stu *Student) Show() {
}

func live() People {
	var stu *Student
	return stu
}

func interface1() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB") // result
	}
}
