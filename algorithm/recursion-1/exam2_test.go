package recursion

import (
	"fmt"
	"testing"
)

func TestFindKth(t *testing.T) {
	// fmt.Println("array1:", array1)
	// fmt.Println("array2:", array2)
	// fmt.Println(FindKth(array1, array2, 5))

	arr1 := []int{8, 10, 23, 42, 54, 86, 90}
	arr2 := []int{3, 12, 50, 56, 65, 69, 75}
	fmt.Println("array1:", arr1)
	fmt.Println("array2:", arr2)
	fmt.Println(FindKth(arr1, arr2, 6))
}
