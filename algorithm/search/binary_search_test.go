package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	fmt.Println("array:", array)
	for i := 0; i < len(array); i++ {
		fmt.Print("target: ", array[i], " index: ")
		fmt.Println(BinarySearch(array, array[i]))
	}
}
