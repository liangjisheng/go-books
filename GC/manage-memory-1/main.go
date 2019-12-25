package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// 这里有一个简单的程序制造了大量的垃圾(garbage),每秒钟创建一个 5,000,000
// 到 10,000,000 bytes 的数组. 程序维持了20个这样的数组,其他的则被丢弃
// 程序这样设计是为了模拟一种非常常见的情况:随着时间的推移,程序中的不同部分
// 申请了内存,有一些被保留,但大部分不再重复使用.在Go语言网络编程中,用goroutines
// 来处理网络连接和网络请求时(network connections or requests),通常goroutines
// 都会申请一块内存(比如slice来存储收到的数据)然后就不再使用它们了.随着时间的推移
// 会有大量的内存被网络连接(network connections)使用,连接累积的垃圾come and gone

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 20)
	var m runtime.MemStats
	makes := 0

	for {
		b := makeBuffer()
		makes++
		i := rand.Intn(len(pool))
		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m) // 获取堆的使用信息
		fmt.Printf("%d,%d,%d,%d,%d,%d\n",
			m.HeapSys,      // 程序向应用程序申请的内存
			bytes,          // 当前pool里总共申请的字节大小
			m.HeapAlloc,    // 堆上目前分配的内存
			m.HeapIdle,     // 堆上目前没有使用的内存
			m.HeapReleased, // 回收到操作系统的内存
			makes)          // 总共申请了多少次内存
	}
}
