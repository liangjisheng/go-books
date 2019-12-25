package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	pool := make([][]byte, 20)
	buffer := make(chan []byte, 5)
	var m runtime.MemStats
	makes := 0

	for {
		var b []byte
		select {
		case b = <-buffer:
		default:
			makes++
			b = makeBuffer()
		}

		i := rand.Intn(len(pool))
		if pool[i] != nil {
			select {
			case buffer <- pool[i]:
				pool[i] = nil
			default:
			}
		}

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
