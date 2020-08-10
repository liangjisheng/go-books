package sort

// 通过构建有序序列, 对于未排序数据, 在已排序序列中从后向前扫描, 找到相应位置并插入

// InsertSort ...
func InsertSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		cur := array[i]
		j := i - 1
		for ; j >= 0 && array[j] >= cur; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = cur
	}
}
