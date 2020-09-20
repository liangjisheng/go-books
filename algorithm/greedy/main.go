package main

import "fmt"

// 比如说一个算法问题使用暴力解法需要指数级时间, 如果能使用动态规划消除重叠子问题,
// 就可以降到多项式级别的时间, 如果满足贪心选择性质, 那么可以进一步降低时间复杂度,
// 达到线性级别的

func main() {
	// fmt.Println(smallestSubsequence("cabcc"))
	// fmt.Println(smallestSubsequence("cdadabcc"))
	// fmt.Println(smallestSubsequence("ecbacba"))

	// array := []int{2, 3, 1, 1, 4, 2, 1}
	// fmt.Printf("step: %d\n", jump1(array))
	// fmt.Printf("step: %d\n", jump2(array))

	intvs := twoDimSlice{
		oneDimSlice{10, 16},
		oneDimSlice{2, 8},
		oneDimSlice{1, 6},
		oneDimSlice{7, 12},
	}

	fmt.Println(intervalSchedule(intvs))
}
