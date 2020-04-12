package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 海量数据M中topN查找,使用最小堆,时间复杂度为(MlogN)

func topN(nums []int, n int) (resp []int) {
	// 构造最小堆
	minHeap := NewMinHeap()
	for i := 0; i < n; i++ {
		minHeap.insert(nums[i])
	}

	// 遍历剩余数据,如果比跟节点大,则替换跟节点,遍历完后堆中留下的数据便是topN
	minVal, _ := minHeap.GetMinVal()
	for _, v := range nums[n:] {
		if v > minVal {
			minHeap.Replace(v)
			minVal, _ = minHeap.GetMinVal()
		}
	}

	// 依次删除堆中数据,从小到大获取topN
	for i := 0; i < n; i++ {
		v, _ := minHeap.Sift()
		resp = append(resp, v)
	}
	return
}

func topNTest() {
	nums := []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < 1000000; i++ {
		nums = append(nums, 50+rand.Intn(2000000))
	}

	res := topN(nums, 10)
	fmt.Println("最小堆排序输出topN:")
	fmt.Println(res)

	fmt.Println("go自带排序接口排序结果:")
	sort.Ints(nums)
	fmt.Println(nums[len(nums)-10 : len(nums)])
}
