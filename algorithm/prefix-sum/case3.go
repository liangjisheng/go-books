package main

// 题目描述: 给定一个数组A[0,1,…,n-1],请构建一个数组B[0,1,…,n-1],
// 其中B中的元素B[i]=A[0]A[1]…A[i-1]A[i+1]…A[n-1]. 不能使用除法
// 可以使用前缀和类似的思想
// 1、先算出 B[i] = A[0] * …A[i-1]
// 2、接着算 B[i] = A[i+1]…A[n-1] * B[i](B[i]是步骤1 算出来的值)

func multiply(arrA []int) []int {
	n := len(arrA)
	arrB := make([]int, n)
	arrB[0] = 1

	// 先求 B[i] = A[0]*...*A[i-1]
	for i := 1; i < n; i++ {
		arrB[i] = arrB[i-1] * arrA[i-1]
	}

	// 再求 B[i] = A[i+1]*...*A[n-1]
	tmp := arrA[n-1]
	for i := n - 2; i >= 0; i-- {
		arrB[i] *= tmp
		tmp *= arrA[i]
	}

	return arrB
}
