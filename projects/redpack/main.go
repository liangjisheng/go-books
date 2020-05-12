package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// RedPack ...
type RedPack struct {
	ID              string
	Num             int // 红包个数
	NumDelivered    int // 已拆包个数
	Amount          int // 红包金额
	AmountDelivered int // 已发放金额
}

// NewRedPack ...
func NewRedPack(num int, amount int) *RedPack {
	return &RedPack{
		Num:             num,
		NumDelivered:    0,
		Amount:          amount,
		AmountDelivered: 0,
	}
}

func init() {
	// 如果放到get函数里, 会导致耗时很长
	rand.Seed(time.Now().Unix())
}

func (pack *RedPack) get() int {
	if pack.NumDelivered == pack.Num {
		return 0
	}

	// 最后一个红包直接返回
	if pack.Num == pack.NumDelivered+1 {
		amount := pack.Amount - pack.AmountDelivered
		pack.NumDelivered++
		pack.AmountDelivered = pack.Amount
		return amount
	}

	// 动态计算剩余红包的平均值
	avg := (pack.Amount - pack.AmountDelivered) / (pack.Num - pack.NumDelivered)
	// 随机计算红包金额(剩余红包的平均值的0到2倍之间)
	// rand.Seed(time.Now().Unix())
	calAmount := rand.Intn(2 * avg)
	// 红包金额最少1分
	if 0 == calAmount {
		calAmount = 1
	}
	// 保证后续每个人至少有1分
	if (pack.Amount - pack.AmountDelivered - calAmount) < (pack.Num - pack.NumDelivered - 1) {
		calAmount = pack.Amount - pack.AmountDelivered - (pack.Num - pack.NumDelivered - 1)
	}

	pack.NumDelivered++
	pack.AmountDelivered += calAmount
	return calAmount
}

func main() {
	start := time.Now().UnixNano() / 1e6
	wg := &sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			redPack := NewRedPack(10, 10000)
			total := 0
			for i := 0; i < 10; i++ {
				amount := redPack.get()
				total += amount
			}
			wg.Done()
		}(wg)
	}

	wg.Wait()
	fmt.Printf("100000个红发发放完毕,耗时：%d ms", (time.Now().UnixNano()/1e6 - start))
}
