package main

import (
	"fmt"
	"time"
)

func chanParam() {
	p := make(chan bool)
	fmt.Printf("原始chan的内存地址是: %p\n", &p)

	// 使用make函数返回的是一个hchan类型的指针*hchan, 这不是与map一个道理嘛
	// 实际我们的 fun (p chan bool)与fun (p *hchan)是一样的, 实际上在作为
	// 传递参数时还是使用了指针的副本进行传递, 属于值传递
	go func(p chan bool) {
		fmt.Printf("函数里接收到chan的内存地址是: %p\n", &p)
		// 模拟耗时
		time.Sleep(2 * time.Second)
		p <- true
	}(p)

	select {
	case l := <-p:
		fmt.Println(l)
	}
}
