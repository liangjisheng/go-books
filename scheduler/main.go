package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取主机的逻辑 CPU 个数
	num := runtime.NumCPU()
	fmt.Println(num)
}
