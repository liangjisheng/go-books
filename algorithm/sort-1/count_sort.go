package sort

// 平均时间复杂度: O(n + k)
// 最佳时间复杂度: O(n + k)
// 最差时间复杂度: O(n + k)
// 空间复杂度: O(n + k)

// 当输入的元素是n 个0到k之间的整数时, 它的运行时间是 O(n + k).
// 在实际工作中, 当k=O(n)时, 我们一般会采用计数排序, 这时的运行时间为O(n)
// 计数排序需要两个额外的数组用来对元素进行计数和保存排序的输出结果, 所以空间复杂度为O(k+n)

// 计数排序使用一个额外的数组C, 其中第i个元素是待排序数组A中值等于i的元素的个数
// 计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中
// 作为一种线性时间复杂度的排序, 计数排序要求输入的数据必须是有确定范围的整数
// 用来计数的数组C的长度取决于待排序数组中数据的范围(等于待排序数组的最大值与最小值的差加上1)
// 然后进行分配、收集处理:
// 分配: 扫描一遍原始数组, 以当前值-minValue作为下标, 将该下标的计数器增1
// 收集: 扫描一遍计数器数组, 按顺序把值收集起来

// 改进工作: CountSort 和 CountSort1 这2个函数都只能处理非负数的数组排序
// 负数排序、负数和非负数混合排序需要修改排序算法

// 计数算法只能使用在已知序列中的元素在0-k之间, 且要求排序的复杂度在线性效率上
// 计数排序和基数排序很类似, 都是非比较型排序算法. 但是, 它们的核心思想是不同的
// 基数排序主要是按照进制位对整数进行依次排序, 而计数排序主要侧重于对有限范围内对象的统计
// 基数排序可以采用计数排序来实现

func maxNum(array []int) int {
	maxI := 0
	n := len(array)
	for i := 1; i < n; i++ {
		if array[maxI] < array[i] {
			maxI = i
		}
	}
	return array[maxI]
}

// CountSort ...
func CountSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	maxNum := maxNum(array)
	countArray := make([]int, maxNum+1)
	resultArray := make([]int, n)

	// 分配: 统计每个元素出现的频率
	for i := 0; i < n; i++ {
		countArray[array[i]]++
	}

	// 收集: 反向填充目标数组
	k := n - 1
	for i := maxNum; i >= 0; i-- {
		for j := countArray[i]; j > 0; j-- {
			resultArray[k] = i
			k--
		}
	}

	for i := 0; i < n; i++ {
		array[i] = resultArray[i]
	}
}

func minNum(array []int) int {
	minI := 0
	n := len(array)
	for i := 1; i < n; i++ {
		if array[minI] > array[i] {
			minI = i
		}
	}
	return array[minI]
}

// 优化改进, 场景分析: 举个极端的例子, 如果排序的数组有200W个元素, 但是这200W个数的值
// 都在1000000-1000100, 也就说有100个数, 总共重复了200W次, 现在要排序, 怎么办?
// 这种情况排序, 计数排序应该是首选. 但是这时候n的值为200W, 如果按原来的算法,
// k的值10001000, 但是此时c中真正用到的地方只有100个, 这样对空间造成了极大的浪费
// 改进思路: 针对c数组的大小, 优化计数排序

// CountSort1 ...
func CountSort1(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	maxNum := maxNum(array)
	minNum := minNum(array)
	// 改变计数数组的大小为要排序数组的极差值加1
	countArray := make([]int, maxNum-minNum+1)
	resultArray := make([]int, n)

	// 分配: 统计每个元素出现的频率
	for i := 0; i < n; i++ {
		countArray[array[i]-minNum]++
	}

	// 收集: 反向填充目标数组
	k := n - 1
	for i := maxNum - minNum; i >= 0; i-- {
		for j := countArray[i]; j > 0; j-- {
			resultArray[k] = i + minNum
			k--
		}
	}

	for i := 0; i < n; i++ {
		array[i] = resultArray[i]
	}
}
