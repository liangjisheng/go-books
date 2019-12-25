package main

// get(获取一个新buffer)和give(回收一个buffer到pool)这两个channel被所有goroutines使用
// 回收器对收回的buffer保持连接,并定期的丢弃那些过于陈旧可能不会再使用的buffer
// (在示例代码中这个周期是一分钟).这让程序可以自动应对爆发性的buffers需求

import (
	"container/list"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var makes int
var frees int

func makeBuffer() []byte {
	makes++
	return make([]byte, rand.Intn(5000000)+5000000)
}

type queued struct {
	when  time.Time
	slice []byte
}

func makeRecycler() (get, give chan []byte) {
	get = make(chan []byte)
	give = make(chan []byte)

	go func() {
		q := new(list.List)
		for {
			if q.Len() == 0 {
				q.PushFront(queued{when: time.Now(), slice: makeBuffer()})
			}

			e := q.Front()
			timeout := time.NewTimer(time.Minute)
			select {
			case b := <-give:
				timeout.Stop()
				q.PushFront(queued{when: time.Now(), slice: b})

			case get <- e.Value.(queued).slice:
				timeout.Stop()
				q.Remove(e)

			case <-timeout.C:
				e := q.Front()
				for e != nil {
					n := e.Next()
					if time.Since(e.Value.(queued).when) > time.Minute {
						q.Remove(e)
						e.Value = nil
					}
					e = n
				}
			}
		}
	}()

	return
}

func main() {
	pool := make([][]byte, 20)
	get, give := makeRecycler()

	var m runtime.MemStats
	for {
		b := <-get
		i := rand.Intn(len(pool))
		if pool[i] != nil {
			give <- pool[i]
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
