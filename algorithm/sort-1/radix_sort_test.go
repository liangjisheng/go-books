package sort

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	fmt.Println("array:", array)
	RadixSort(array)
	fmt.Println("array:", array)
}
