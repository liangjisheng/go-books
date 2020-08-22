package search

import (
	"math/rand"
	"sort"
	"time"
)

// 二分查找在实际问题中的应用, 首先思考使用 for 循环暴力解决问题, 观察代码是否如下形式
// for i := 0; i < n; i++ {
//     if isOK(i) {
// 	    return answer
// 	}
// }
// 如果是, 那么就可以使用二分搜索优化搜索空间: 如果要求最小值就是搜索左侧边界的二分,
// 如果要求最大值就用搜索右侧边界的二分

var array = make([]int, 0)

func init() {
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(10) + 5
	for i := 0; i < n; i++ {
		array = append(array, rand.Intn(1000))
	}

	sort.Ints(array)
}

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

func minArray(array []int) int {
	n := len(array)
	if n == 0 {
		panic("array len is 0")
	}

	minI := 0
	for i := 1; i < n; i++ {
		if array[i] < array[minI] {
			minI = i
		}
	}
	return array[minI]
}

func maxArray(array []int) int {
	n := len(array)
	if n == 0 {
		panic("array len is 0")
	}

	maxI := 0
	for i := 1; i < n; i++ {
		if array[i] > array[maxI] {
			maxI = i
		}
	}
	return array[maxI]
}

func sumArray(array []int) int {
	sum := 0
	n := len(array)
	for i := 0; i < n; i++ {
		sum += array[i]
	}
	return sum
}
