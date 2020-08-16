package recursion

import (
	"fmt"
	"testing"
)

func TestFibonacciRecursion(t *testing.T) {
	var i uint64 = 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", FibonacciRecursion(i))
	}
	fmt.Println()
}

func TestFibonacciIteration(t *testing.T) {
	var i uint64 = 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", FibonacciIteration(i))
	}
	fmt.Println()
}
