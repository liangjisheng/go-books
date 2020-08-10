package sort

import (
	"math/rand"
	"time"
)

var array = make([]int, 0)

func init() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(10) + 5
	for i := 0; i < n; i++ {
		array = append(array, rand.Intn(1000))
	}
}
