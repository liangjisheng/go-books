package main

import "fmt"

// T ...
type T int

// IsClosed check channel close
func IsClosed(ch <-chan T) bool {
	select {
	case <-ch: // 通道关闭后会返回对应类型的nil值
		return true
	default:
	}
	return false
}

func case1() {
	c := make(chan T)
	fmt.Println(IsClosed(c)) // false
	close(c)
	fmt.Println(IsClosed(c)) // true
}
