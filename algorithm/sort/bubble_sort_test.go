package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println("array:", array)
	BubbleSort(array)
	fmt.Println("array:", array)
}

func TestBubbleSort1(t *testing.T) {
	fmt.Println("array:", array)
	BubbleSort1(array)
	fmt.Println("array:", array)
}

func TestBubbleSort2(t *testing.T) {
	fmt.Println("array:", array)
	BubbleSort2(array)
	fmt.Println("array:", array)
}
