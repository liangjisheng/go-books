package main

import "fmt"

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	// fmt.Println(fibonacci1(10))
	// fmt.Println(fibonacci2(10))
	// fmt.Println(fibonacci3(10))
	// fmt.Println(fibonacci4(10))

	// coins := []int{1, 2, 5}
	// fmt.Println(coinChange1(coins, 11))
	// fmt.Println(coinChange2(coins, 11))
	// fmt.Println(coinChange3(coins, 11))

	// text := "test"
	// pattern := "test"
	// fmt.Println("isMatch1:", isMatch1(text, pattern))
	// fmt.Println("isMatch2:", isMatch2(text, pattern))
	// fmt.Println("isMatch3:", isMatch3(text, pattern))

	// text = "aa"
	// pattern = "a*"
	// text = "aab"
	// pattern = "c*a*b"
	// text = "ab"
	// pattern = ".*"

	// fmt.Println("isMatch4:", isMatch4(text, pattern))
	// fmt.Println("isMatch5:", isMatch5(text, pattern))

	// nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	// fmt.Println(lengthOfLIS(nums))

	// piles := []int{3, 9, 1, 2}
	// fmt.Println(stoneGame(piles))

	// fmt.Println("maxA1:", maxA1(3))
	// fmt.Println("maxA1:", maxA1(7))
	// fmt.Println("maxA1:", maxA1(30))

	// fmt.Println("maxA1Memo:", maxA1Memo(3))
	// fmt.Println("maxA1Memo:", maxA1Memo(7))
	// fmt.Println("maxA1Memo:", maxA1Memo(20))

	fmt.Println("maxA2:", maxA2(3))
	fmt.Println("maxA2:", maxA2(7))
	fmt.Println("maxA2:", maxA2(30))
}
