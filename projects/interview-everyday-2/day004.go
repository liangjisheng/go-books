package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	sync.Mutex
}

// 方法使用值接收器将导致锁失效 将 Mutex 作为匿名字段时
// 相关的方法必须使用指针接收者 否则会导致锁机制失效
func (d data) test(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func (d *data) test1(s string) {
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}

func day004() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data

	go func() {
		defer wg.Done()
		// d.test("read") // 锁失效
		d.test1("read")
	}()

	go func() {
		defer wg.Done()
		// d.test("write") // 锁失效
		d.test1("write")
	}()

	wg.Wait()
}
