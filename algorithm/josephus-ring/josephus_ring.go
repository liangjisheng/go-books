package main

import (
	"container/ring"
	"fmt"
)

// 使用Ring模拟约瑟夫环问题

type Player struct {
	position int  // 位置
	alive    bool // 是否存活
}

func josephus() {
	const (
		playerCount = 41 // 玩家人数
		startPos    = 1  // 开始报数位置
	)

	deadline := 3

	r := ring.New(playerCount)

	// 设置所有玩家的初始化
	for i := 1; i <= playerCount; i++ {
		r.Value = &Player{i, true}
		r = r.Next()
	}

	// 如果开始报数的位置不为1，则设置开始位置
	if startPos > 1 {
		r = r.Move(startPos - 1)
	}

	counter := 1   // 报数从1开始，因为下面的循环从第二个开始计算
	deadCount := 0 // 死亡人数，初始值为0

	for deadCount < playerCount { // 直到所有人都死亡，否则循环一直执行
		r = r.Next()

		if r.Value.(*Player).alive {
			counter++
		}

		// 如果报数为deadline,则此人淘汰出局
		if counter == deadline {
			r.Value.(*Player).alive = false
			fmt.Printf("Player %d died.\n", r.Value.(*Player).position)
			deadCount++

			counter = 0 // 报数置为0
		}
	}
}
