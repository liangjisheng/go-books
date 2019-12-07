package main

import (
	"fmt"
	"time"
)

// A ...
func A(string string) string {
	return string + string
}

// B ...
func B(len int) int {
	return len + len
}

// 	func C(val, default string) string {
//        if val == "" {
//            return default
//        }
//        return val
//    }

// C() 函数的 default 属于关键字 string 和 len 是预定义标识符，可以在局部使用
// nil 也可以当做变量使用，不过不建议写这样的代码，可读性不好

func f1() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func f2() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}

func day002() {
	ch := make(chan int, 1)
	go func() {
		// select 会随机选择一个 case, 选完 case 之后才会进入函数
		select {
		case ch <- f1():
		case ch <- f2():
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch) // 1、2随机输出
}
