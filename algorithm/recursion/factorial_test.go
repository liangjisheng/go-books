package recursion

import (
	"fmt"
	"testing"
)

func TestFactorialRecursion(t *testing.T) {
	var i uint64 = 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", FactorialRecursion(i))
	}
	fmt.Println()
}

func TestFactorialIteration(t *testing.T) {
	var i uint64 = 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", FactorialIteration(i))
	}
	fmt.Println()
}
