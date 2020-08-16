package recursion

import (
	"fmt"
	"testing"
)

func TestForgJumpStep(t *testing.T) {
	var i uint64 = 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", ForgJumpStep(i))
	}
	fmt.Println()
}

func TestForgJumpStepIteration(t *testing.T) {
	var i uint64 = 1
	for ; i <= 10; i++ {
		fmt.Printf("%d ", ForgJumpStepIteration(i))
	}
	fmt.Println()
}

func TestForgJumpStepIterationDecimal(t *testing.T) {
	var i uint64 = 1000
	for ; i <= 1001; i++ {
		fmt.Printf("%s ", ForgJumpStepIterationDecimal(i).String())
	}
	fmt.Println()
}
