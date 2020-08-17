package recursion

import (
	"fmt"
	"testing"
)

func TestCombination(t *testing.T) {
	arr := make([]int, 0)
	selectArray := make([]byte, 0)
	for i := 0; i < 5; i++ {
		arr = append(arr, i)
		selectArray = append(selectArray, 0)
	}
	fmt.Println("array:", arr)
	combination(arr, 0, selectArray)
}
