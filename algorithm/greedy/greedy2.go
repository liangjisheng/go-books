package main

import (
	"sort"
)

// https://mp.weixin.qq.com/s?__biz=Mzg2NzA4MTkxNQ==&mid=2247485839&idx=2&sn=3a4d94fb57a38cada8cd538e3f06a1e4&scene=21#wechat_redirect

// 本文解决一个很经典的贪心算法问题 Interval Scheduling(区间调度问题).
// 给你很多形如 [start,end] 的闭区间, 请你设计一个算法, 算出这些区间中最多有几个
// 互不相交的区间, (注意边界相同并不算相交)

// 正确的思路其实很简单, 可以分为以下三步:
// 从区间集合 intvs 中选择一个区间 x, 这个 x 是在当前所有区间中结束最早的(end 最小)
// 把所有与 x 区间相交的区间从区间集合 intvs 中删除
// 重复步骤 1 和 2, 直到 intvs 为空为止.  之前选出的那些 x 就是最大不相交子集
// 可以按每个区间的end数值升序排序, 因为这样处理之后实现步骤 1 和步骤 2 都方便很多

type oneDimSlice []int
type twoDimSlice []oneDimSlice

func (s twoDimSlice) Len() int           { return len(s) }
func (s twoDimSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s twoDimSlice) Less(i, j int) bool { return s[i][1] < s[j][1] }

// 返回互不相交的区间个数
func intervalSchedule(intvs twoDimSlice) int {
	// 按 end 升序排序, twoDimSlice 实现了排序接口, 可以直接调用 sort.Sort() 排序
	sort.Sort(intvs)

	// 至少应该有一个区间不相交
	count := 1
	// 排序后, 第一个区间就是 x
	xEnd := intvs[0][1]

	n := len(intvs)
	for i := 1; i < n; i++ {
		start := intvs[i][0]
		if start >= xEnd {
			// 更新 xEnd 为下一个不相交区间的 end
			count++
			xEnd = intvs[i][1]
		}
	}

	return count
}
