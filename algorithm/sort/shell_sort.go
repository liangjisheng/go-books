package sort

// 先将整个待排元素序列分割成若干个子序列(由相隔某个“增量”的元素组成的)分别进行直接插入排序
// 然后依次缩减增量再进行排序, 待整个序列中的元素基本有序(增量足够小)时, 再对全体元素进行一
// 次直接插入排序

// ShellSort ...
func ShellSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	for gap := n / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < n; i++ {
			cur := array[i]
			j := i - gap
			for ; j >= 0; j -= gap {
				if array[j] > cur {
					array[j+gap] = array[j]
				} else {
					break
				}
			}

			array[j+gap] = cur
		}
	}
}
