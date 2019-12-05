package main

import (
	"fmt"
	"strings"
)

// TrimRight() 会将第二个参数字符串里面所有的字符拿出来处理
// 只要与其中任何一个字符相等 便会将其删除 想正确地截取字符串
// 可以参考 TrimSuffix() 函数
func day008Test1() {
	fmt.Println(strings.TrimRight("ABBA", "BA"))
	fmt.Println(strings.TrimSuffix("ABBA", "BA"))
}

// copy(dst, src) 函数返回 len(dst)、len(src) 之间的最小值
// 如果想要将 src 完全拷贝至 dst 必须给 dst 分配足够的内存空间
func day008Test2() {
	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)
	fmt.Println(dst)
}

func day008Test3() {
	var src, dst []int
	src = []int{1, 2, 3}
	// 分配足够的空间
	// dst = make([]int, len(src))
	// 或者直接 append
	dst = append(dst, src...)
	copy(dst, src)
	fmt.Println(dst)
}

func day008() {
	// day008Test1()
	day008Test2()
	day008Test3()
}
