package recursion

import "testing"

func TestPermutation(t *testing.T) {
	permutation(3)
}

func TestPermutationRecursion(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 4; i++ {
		arr = append(arr, i)
	}
	permutationRecursion(arr, 0)
}

func TestPermutationDict(t *testing.T) {
	arr := make([]int, 0)
	for i := 0; i < 4; i++ {
		arr = append(arr, i)
	}
	permutationDict(arr)
}
