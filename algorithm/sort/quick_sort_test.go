package sort

import (
	"fmt"
	"testing"
)

func TestQuickSortRecursive(t *testing.T) {
	fmt.Println("array:", array)
	QuickSortRecursive(array)
	fmt.Println("array:", array)
}

func TestQuickSortIteration(t *testing.T) {
	fmt.Println("array:", array)
	QuickSortIteration(array)
	fmt.Println("array:", array)
}
