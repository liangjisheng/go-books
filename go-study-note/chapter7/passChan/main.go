package main

import "fmt"

type request struct {
	data []int
	ret  chan int
}

func newRequest(data ...int) *request {
	return &request{data, make(chan int, 1)}
}

func process(req *request) {
	x := 0
	for _, i := range req.data {
		x += i
	}

	req.ret <- x
}

func main() {
	req := newRequest(10, 20, 30)
	process(req)
	fmt.Println(<-req.ret)
}
