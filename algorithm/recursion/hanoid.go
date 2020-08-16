package recursion

import "fmt"

// 将一个圆盘从 a 移动到 b 上
func move(a, b string) {
	fmt.Println(a, "->", b)
}

// Hanoid 将 n 个圆盘从 a 经由 b 移动到 c 上
func Hanoid(n int, a, b, c string) {
	if n <= 0 {
		return
	}

	// 将 n-1 个圆盘从 a 经由 c 移动到 b 上
	Hanoid(n-1, a, c, b)
	// 将最后一个圆盘从 a 移动到 c 上
	move(a, c)
	// 将 n-1 个圆盘从 b 经由 a 移动到 c 上
	Hanoid(n-1, b, a, c)
}
