package bitoperation

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	fmt.Println(sum(5))
}

func TestSum2(t *testing.T) {
	fmt.Println(sum2(5, 6))
}

func TestMult(t *testing.T) {
	fmt.Println(mult(5, -6))
}

func TestOdd(t *testing.T) {
	if IsOdd(4) {
		fmt.Println("4 is odd")
	} else {
		fmt.Println("4 is even")
	}

	if IsOdd(5) {
		fmt.Println("5 is odd")
	} else {
		fmt.Println("5 is even")
	}
}

func TestSwap(t *testing.T) {
	x, y := 4, 5
	fmt.Printf("x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("x=%d, y=%d\n", x, y)
}

func TestFindOnlyOne(t *testing.T) {
	array := []int{1, 1, 2, 2, 3, 3, 4, 4, 5}
	fmt.Println(findOnlyOne(array))
}

func TestPow(t *testing.T) {
	fmt.Println(pow(2, 10))
}

func TestNumberOf1(t *testing.T) {
	fmt.Println(numberOf1(13))
}
