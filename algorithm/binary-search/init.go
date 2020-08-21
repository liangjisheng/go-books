package search

import (
	"math/rand"
	"sort"
	"time"
)

var array = make([]int, 0)

func init() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(10) + 5
	for i := 0; i < n; i++ {
		array = append(array, rand.Intn(1000))
	}

	sort.Ints(array)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func minArray(array []int) int {
	n := len(array)
	if n == 0 {
		panic("array len is 0")
	}

	minI := 0
	for i := 1; i < n; i++ {
		if array[i] < array[minI] {
			minI = i
		}
	}
	return array[minI]
}

func maxArray(array []int) int {
	n := len(array)
	if n == 0 {
		panic("array len is 0")
	}

	maxI := 0
	for i := 1; i < n; i++ {
		if array[i] > array[maxI] {
			maxI = i
		}
	}
	return array[maxI]
}
