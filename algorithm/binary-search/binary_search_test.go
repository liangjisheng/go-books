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

func TestBinarySearch1(t *testing.T) {
	arr := []int{1, 2, 5, 6, 9}
	// arr := []int{0, 1, 2, 2, 2, 3}
	// arr := []int{2}
	fmt.Println("array:", arr)
	fmt.Println(BinarySearch1(arr, 8))
}

func TestBinarySearch2(t *testing.T) {
	arr := []int{1, 2, 5, 6, 9}
	// arr := []int{0, 1, 1, 2, 2, 3}
	// arr := []int{2}
	fmt.Println("array:", arr)
	fmt.Println(BinarySearch2(arr, 1))
}

func TestBinarySearch3(t *testing.T) {
	arr := []int{1, 2, 5, 5, 6, 9}
	// arr := []int{0, 1, 1, 2, 2, 3}
	// arr := []int{2}
	fmt.Println("array:", arr)
	fmt.Println(BinarySearch3(arr, 5))
}

func TestPow(t *testing.T) {
	for i := 1; i <= 10; i++ {
		fmt.Println(pow(2, i))
	}
}

func TestRotatedBinarySearch(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("array:", arr)
	fmt.Println(RotatedBinarySearch(arr, 3))
}
