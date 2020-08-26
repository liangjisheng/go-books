package main

import "fmt"

// 输入一个n行m列的整数矩阵，再输入q个询问，每个询问包含四个整数x1, y1, x2, y2
// 表示一个子矩阵的左上角坐标和右下角坐标
// 对于每个询问输出子矩阵中所有数的和
// 输入格式

// 第一行包含三个整数n，m，q
// 接下来n行，每行包含m个整数，表示整数矩阵
// 接下来q行，每行包含四个整数x1, y1, x2, y2，表示一组询问
// 输出格式
// 共q行, 每行输出一个询问的结果
// 数据范围

// 1≤n,m≤1000,
// 1≤q≤200000,
// 1≤x1≤x2≤n,
// 1≤y1≤y2≤m,
// −1000≤矩阵内元素的值≤1000
// 输入样例：

// 3 4 3
// 1 7 2 4
// 3 6 2 8
// 2 1 2 3
// 1 1 2 2
// 2 1 3 4
// 1 3 3 4

// 输出样例:

// 17
// 27
// 21

// 二维前缀和公式
// s[i][j]=s[i-1][j]+s[i][j-1]-s[i-1][j-1]+a[i][j]

// 2点之间(x1, y1), (x2, y2) 的子矩阵和公式
// s[x2][y2]-s[x1-1][y2]-s[x2][y1-1]+s[x1-1][y1-1]

// 子矩阵的和问题(二维前缀和)
func prefixSum2() {
	var n, m, q int
	fmt.Print("input n, m, q: ")
	fmt.Scanln(&n, &m, &q)

	// 构造二维数组
	a := make([][]int, n+1)
	s := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		a[i] = make([]int, m+1)
		s[i] = make([]int, m+1)
	}
	fmt.Println("a:")
	printArray(a)
	fmt.Println("s:")
	printArray(s)

	fmt.Println("input two dimension array:")
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&a[i][j])
		}
		fmt.Scanln()
	}
	fmt.Println("a:")
	printArray(a)

	// 求前缀和
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + a[i][j]
		}
	}
	fmt.Println("s:")
	printArray(s)

	var x1, y1, x2, y2 int
	for q > 0 {
		fmt.Print("input two point (x1, y1, x2, y2): ")
		fmt.Scanln(&x1, &y1, &x2, &y2)
		sum := s[x2][y2] - s[x1-1][y2] - s[x2][y1-1] + s[x1-1][y1-1]
		fmt.Println(sum)
		q--
	}
}

func printArray(a [][]int) {
	n := len(a)
	if n == 0 {
		return
	}
	m := len(a[0])
	if m == 0 {
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%3d", a[i][j])
		}
		fmt.Println()
	}
}
