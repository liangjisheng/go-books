package recursion

import (
	"math/rand"
	"time"
)

var array1 = make([]int, 0)
var array2 = make([]int, 0)

func init() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(10) + 1
	for i := 0; i < n; i++ {
		array1 = append(array1, rand.Intn(100))
		array2 = append(array2, rand.Intn(100))
	}

	BubbleSort(array1)
	BubbleSort(array2)
}

// BubbleSort 如果序列已经完全有序,可以提前终止冒泡排序,相当于提前进行终止
func BubbleSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := n - 1; i > 0; i-- {
		sorted := true
		for j := 1; j <= i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
				sorted = false
			}
		}

		if sorted {
			break
		}
	}
}
