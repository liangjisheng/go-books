package sort

import (
	"fmt"
	"testing"
)

func TestMergeSortRecursive(t *testing.T) {
	fmt.Println("array:", array)
	MergeSortRecursive(array)
	fmt.Println("array:", array)
}

func TestMergeSortIteration(t *testing.T) {
	fmt.Println("array:", array)
	MergeSortIteration(array)
	fmt.Println("array:", array)
}
