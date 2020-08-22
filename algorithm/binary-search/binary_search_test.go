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
	fmt.Println(BinarySearchFirstLargeOrEqual(arr, 8))
}

func TestBinarySearch2(t *testing.T) {
	arr := []int{1, 2, 5, 6, 9}
	// arr := []int{0, 1, 1, 2, 2, 3}
	// arr := []int{2}
	fmt.Println("array:", arr)
	fmt.Println(BinarySearchLastLargeOrEqual(arr, 1))
}

func TestBinarySearch3(t *testing.T) {
	arr := []int{1, 2, 5, 5, 6, 9}
	// arr := []int{0, 1, 1, 2, 2, 3}
	// arr := []int{2}
	fmt.Println("array:", arr)
	fmt.Println(BinarySearchFirstLarge(arr, 5))
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

func TestLeftBound(t *testing.T) {
	arr := []int{1, 2, 4, 5, 6, 9}
	// arr := []int{0, 1, 1, 2, 2, 3}
	// arr := []int{2}
	fmt.Println("array:", arr)
	fmt.Println(leftBound(arr, 5))
}

func TestRightBound(t *testing.T) {
	arr := []int{1, 2, 5, 5, 6, 9}
	// arr := []int{0, 1, 1, 2, 2, 3}
	// arr := []int{2}
	fmt.Println("array:", arr)
	fmt.Println(rightBound(arr, 10))
}

func TestMinArray(t *testing.T) {
	arr := []int{1, 2, 5, 5, 6, 9}
	fmt.Println("array:", arr)
	fmt.Println(minArray(arr))
}

func TestMaxArray(t *testing.T) {
	arr := []int{1, 2, 5, 5, 6, 9}
	fmt.Println("array:", arr)
	fmt.Println(maxArray(arr))
}

func TestMinEatSpeed(t *testing.T) {
	piles := []int{3, 6, 7, 11}
	H := 8
	fmt.Println(minEatSpeed(piles, H))
	fmt.Println(minEatSpeedOptimize(piles, H))
}

func TestMinLoadSpeed(t *testing.T) {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	D := 4
	fmt.Println(minLoadCap(weights, D))
	fmt.Println(minLoadCapOptisize(weights, D))
}
