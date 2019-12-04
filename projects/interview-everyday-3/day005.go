package main

import (
	"fmt"
	"sync"
)

func test1(arr *[3]int) {
	(*arr)[0] = 88
}

func day005Test1() {
	arr := [3]int{1, 2, 3}
	test1(&arr)
	fmt.Println(arr)
}

func day005Test2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
		// 协程里面，使用 wg.Add(1) 但是没有 wg.Done()，导致 panic()
		wg.Add(1)
	}()
	wg.Wait()
}

func day005() {
	// day005Test1()
	// day005Test2()
}
