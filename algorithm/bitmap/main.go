package main

import (
	"fmt"
)

func main() {
	bitMap := NewBitMap(100)
	bitMap.Add(98)
	bitMap.Add(9)
	bitMap.Add(19)
	bitMap.Add(29)

	fmt.Println(bitMap.Contain(19))
	bitMap.Delete(19)
	fmt.Println(bitMap.Contain(19))
}
