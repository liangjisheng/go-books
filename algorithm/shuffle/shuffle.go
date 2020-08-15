package shuffle

import (
	"math/rand"
	"time"
)

var array = make([]int, 0)

func init() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < rand.Intn(20)+5; i++ {
		array = append(array, rand.Intn(1000))
	}
}

// Shuffle 每次随机抽出两张牌交换, 交换一定次数后结束
// 假设交换了m此, 则某张牌始终没有被交换的概率为
// (n-2)/n * (n-2)/n, … …* (n-2)/n = ((n-2)/n)^m
// 我们希望其概率小于摸个值,求出m的解.假设概率小于1/1000,对于n=52,
// m大概为176,实际上远远大于数组的长度
func Shuffle(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	// 交换n次
	for i := 0; i < n; i++ {
		idx1 := rand.Intn(n)
		idx2 := rand.Intn(n)

		array[idx1], array[idx2] = array[idx2], array[idx1]
	}
}

// Shuffle1 Fisher–Yates shuffle算法
// 每次随机选取一个数, 然后将该数与数组中最后(或最前)的元素相交换
// (如果随机选中的是最后/最前的元素, 则相当于没有发生交换);
// 然后缩小选取数组的范围, 去掉最后的元素,即之前随机抽取出的数.
// 重复上面的过程, 直到剩余数组的大小为1, 即只有一个元素时结束
func Shuffle1(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	i := n
	for i > 1 {
		j := rand.Intn(i)
		array[i-1], array[j] = array[j], array[i-1]
		i--
	}
}
