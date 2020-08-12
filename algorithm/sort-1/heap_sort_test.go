package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	fmt.Println("array:", array)
	HeapSort(array)
	fmt.Println("array:", array)
}
