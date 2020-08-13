package sort

import (
	"fmt"
	"testing"
)

func TestCountSort(t *testing.T) {
	fmt.Println("array:", array)
	CountSort(array)
	fmt.Println("array:", array)
}

func TestCountSort1(t *testing.T) {
	fmt.Println("array:", array)
	CountSort1(array)
	fmt.Println("array:", array)
}
