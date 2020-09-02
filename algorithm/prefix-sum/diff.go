package main

import "fmt"

// 一维差分
// B[i] = a[i] - a[i - 1]
// 给区间[l, r]中的每个数加上c：B[l] += c, B[r + 1] -= c
// n 为原数组个数, m 为对区间[l, r]进行多少次操作

func diff1() {
	fmt.Print("input n, m: ")
	var n, m int
	fmt.Scanf("%d%d", &n, &m)
	fmt.Scanln()

	fmt.Print("input array: ")
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Scanln()

	// 构建一维差分
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			diff[i] = array[i]
		} else {
			diff[i] = array[i] - array[i-1]
		}
	}

	fmt.Println("diff array:")
	printArr(diff)

	// 对区间 [l, r] 每个数加上 c,
	var l, r, c int
	for m > 0 {
		fmt.Print("input l, r, c: ")
		fmt.Scanf("%d%d%d", &l, &r, &c)
		fmt.Scanln()

		diff[l] += c
		diff[r+1] -= c
		m--
	}

	// 对差分数组求前缀和得到操作后的数组
	prefixSum := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			prefixSum[i] = diff[i]
		} else {
			prefixSum[i] = prefixSum[i-1] + diff[i]
		}
	}

	fmt.Println("result array:")
	printArr(prefixSum)
}

func printArr(array []int) {
	n := len(array)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()
}
