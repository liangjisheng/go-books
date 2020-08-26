package main

import "fmt"

// 静态区间求和(一维前缀和)
func prefixSum1() {
	fmt.Print("input array number: ")
	n := 0
	fmt.Scanf("%d", &n)
	fmt.Scanln()

	a := make([]int, n+1)
	b := make([]int, n+1)

	fmt.Print("input array: ")
	elem := 0
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &elem)
		a[i] = elem
		b[i] = b[i-1] + a[i]
	}
	fmt.Println("a:", a[1:])
	fmt.Println("b:", b[1:])
	fmt.Scanln()

	fmt.Print("input select times: ")
	times := 0
	fmt.Scanf("%d", &times)
	fmt.Scanln()

	left := 0
	right := 0
	for times > 0 {
		fmt.Print("input left and right: ")
		fmt.Scanf("%d%d", &left, &right)
		fmt.Scanln()
		fmt.Printf("sum a[%d]+...+a[%d]: %d\n", left, right, b[right]-b[left-1])
		times--
	}
}
