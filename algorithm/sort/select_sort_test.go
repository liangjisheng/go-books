package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	fmt.Println("array:", array)
	selectSort(array)
	fmt.Println("array:", array)
}

func TestSelectSort1(t *testing.T) {
	fmt.Println("array:", array)
	selectSort1(array)
	fmt.Println("array:", array)
}
