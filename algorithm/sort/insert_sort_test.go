package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	fmt.Println("array:", array)
	InsertSort(array)
	fmt.Println("array:", array)
}

func TestInsertSort1(t *testing.T) {
	fmt.Println("array:", array)
	InsertSort1(array)
	fmt.Println("array:", array)
}

func TestBinaryInsertSort(t *testing.T) {
	fmt.Println("array:", array)
	BinaryInsertSort(array)
	fmt.Println("array:", array)
}
