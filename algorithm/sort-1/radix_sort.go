package sort

// 时间复杂度: O(k*N)
// 空间复杂度: O(k + N)
// 稳定性: 稳定

// 将所有待比较数值(正整数)统一为同样的数位长度, 数位较短的数前面补零
// 从最低位开始, 依次进行一次排序
// 这样从最低位排序一直到最高位排序完成以后, 数列就变成一个有序序列

// 设待排序的数组R[1..n], 数组中最大的数是d位数, 基数为r
// (如基数为10, 即10进制, 最大有10种可能, 即最多需要10个桶来映射数组元素)

// 基数排序与计数排序、桶排序这三种排序算法都利用了桶的概念, 但对桶的使用方法上有明显差异
// 基数排序: 根据键值的每位数字来分配桶; 计数排序: 每个桶只存储单一键值;
// 桶排序: 每个桶存储一定范围的数值
// 基数排序不是直接根据元素整体的大小进行元素比较, 而是将原始列表元素分成多个部分
// 对每一部分按一定的规则进行排序, 进而形成最终的有序列表

// 辅助函数, 求数据的最大位数
func maxBit(array []int) int {
	n := len(array)
	if n <= 0 {
		return -1
	}

	// 先求出最大数, 再求其位数
	maxNum := array[0]
	for i := 1; i < n; i++ {
		if maxNum < array[i] {
			maxNum = array[i]
		}
	}

	d := 1
	for maxNum >= 10 {
		maxNum /= 10
		d++
	}
	return d
}

// m^n
func pow(m, n int) int {
	res := 1
	for n > 0 {
		res *= m
		n--
	}
	return res
}

// RadixSort ...
func RadixSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	maxBit := maxBit(array)
	res := make([]int, n)

	for i := 0; i < maxBit; i++ {
		// 计数器
		count := make([]int, 10)

		// 统计每个桶中的记录数
		for j := 0; j < n; j++ {
			tmp := (array[j] / pow(10, i)) % 10
			count[tmp]++
		}

		// 将res中的位置依次分配给每个桶
		for j := 1; j < 10; j++ {
			count[j] = count[j-1] + count[j]
		}

		// 将所有桶中记录依次收集到res中
		for j := n - 1; j >= 0; j-- {
			tmp := (array[j] / pow(10, i)) % 10
			res[count[tmp]-1] = array[j]
			count[tmp]--
		}

		// 将res的内容复制到array中
		for j := 0; j < n; j++ {
			array[j] = res[j]
		}
	}
}
