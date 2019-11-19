package main

import "fmt"

func day007() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
	// Output
	// nil

	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"])
	// Output
	// 0
}

// cap() 函数适用于 array, slice, channel, 不适用于 map
