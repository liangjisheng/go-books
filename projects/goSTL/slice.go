package main

import (
	"fmt"

	"github.com/liyue201/gostl/ds/slice"
)

func sliceTest() {
	a := slice.IntSlice(make([]int, 0))
	a = append(a, 2)
	a = append(a, 1)
	a = append(a, 3)
	fmt.Printf("%v\n", a)
}
