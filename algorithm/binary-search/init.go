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
