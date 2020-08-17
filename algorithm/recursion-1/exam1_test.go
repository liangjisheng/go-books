package recursion

import (
	"fmt"
	"testing"
)

func TestGetUpMedian(t *testing.T) {
	fmt.Println("array1:", array1)
	fmt.Println("array2:", array2)
	fmt.Println(GetUpMedian(array1, array2))
}

func TestGetUpMedianIteration(t *testing.T) {
	fmt.Println("array1:", array1)
	fmt.Println("array2:", array2)
	fmt.Println(GetUpMedianIteration(array1, array2))
}
