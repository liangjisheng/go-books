package sort

// 平均时间复杂度: O(N^2)
// 最差时间复杂度: O(N^2)
// 空间复杂度: O(1)
// 排序方式: In-place
// 稳定性: 稳定

// (1) 最好情况: 序列已经是升序排列, 在这种情况下, 需要进行的比较操作需(n-1)次即可
// (2) 最坏情况: 序列是降序排列, 那么此时需要进行的比较共有n(n-1)/2次

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

// InsertSort1 ...
func InsertSort1(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		cur := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > cur {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = cur
	}
}

// BinaryInsertSort 直接插入排序每次往前插入时, 是按顺序依次往前查找, 数据量较大时
// 必然比较耗时, 效率低, 改进思路: 在往前找合适的插入位置时采用二分查找的方式, 即折半插入
// 在已经排序的元素序列中二分查找到第一个比它大的数的位置
func BinaryInsertSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		cur := array[i]
		left := 0
		right := i - 1

		for left <= right {
			mid := (left + right) / 2
			if array[mid] > cur {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		for j := i; j > left; j-- {
			array[j] = array[j-1]
		}

		array[left] = cur
	}
}
