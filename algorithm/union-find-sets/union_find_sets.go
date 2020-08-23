package unionfindsets

import "fmt"

var father = make([]int, 0)

// 我们用一个数组rank[]记录每个根节点对应的树的深度 (如果不是根节点, 其rank相当于以它作为根节点的子树的深度)
// 一开始, 把所有元素的rank(秩)设为1. 合并时比较两个根节点, 把rank较小者往较大者上合并
// 路径压缩和按秩合并如果一起使用, 时间复杂度接近 O(n), 值得注意的是, 按秩合并会带来额外的空间复杂度
var rank = make([]int, 0)

// 假如有编号为1, 2, 3, ..., n的n个元素, 我们用一个数组father[]来存储每个元素的父节点
// (因为每个元素有且只有一个父节点, 所以这是可行的). 一开始, 我们先将它们的父节点设为自己
func initUnionFindSet(n int) {
	for i := 0; i < n; i++ {
		father = append(father, i)
		rank = append(rank, 1)
	}
}

// 我们用递归的写法实现对代表元素的查询: 一层一层访问父节点, 直至根节点
// (根节点的标志就是父节点是本身). 要判断两个元素是否属于同一个集合
// 只需要看它们的根节点是否相同即可
func findRecursion(elem int) int {
	if father[elem] == elem {
		return elem
	}

	return findRecursion(father[elem])
}

// 可以使用路径压缩的方法. 既然只关心一个元素对应的根节点, 那希望每个元素到根节点的路径尽可能短, 最好只需要一步
func findRecursionPathCompress(elem int) int {
	if father[elem] == elem {
		return elem
	}

	// 将路径上所有经过的元素都直接指向代表元素
	father[elem] = findRecursionPathCompress(father[elem])
	return father[elem]
}

// 非递归查找代表元素
func findIteration(elem int) int {
	for father[elem] != elem {
		elem = father[elem]
	}

	return elem
}

// 非递归路径压缩
func findIterationPathCompress(elem int) int {
	// 查找 elem 元素所在集合的代表元素 fa
	fa := elem
	for father[fa] != fa {
		fa = father[fa]
	}

	// 此时 fa 为这个集合的代表元素
	// 将路径上所有经过的元素都直接指向代表元素
	compressPoint := elem
	for father[compressPoint] != fa {
		tmp := father[compressPoint]
		father[compressPoint] = fa
		compressPoint = tmp
	}

	return elem
}

// 合并操作也是很简单的, 先找到两个集合的代表元素, 然后将前者的父节点设为后者即可
// 当然也可以将后者的父节点设为前者
func merge(x, y int) {
	father[findRecursion(x)] = findRecursion(y)
}

func mergeWithPathCompress(x, y int) {
	father[findRecursionPathCompress(x)] = findRecursionPathCompress(y)
}

func mergeWithRank(x, y int) {
	// 先找到两个根节点
	faX := findRecursionPathCompress(x)
	faY := findRecursionPathCompress(y)
	// x, y 已经在同一个集合中, 不需要合并
	if faX == faY {
		return
	}

	// 比较2个集合的秩
	if rank[faX] < rank[faY] {
		father[faX] = faY
	} else if rank[faX] == rank[faY] {
		// faX 合并到 faY 上
		father[faX] = faY
		// 深度相同, 合并后的根节点深度加1
		rank[faY]++

		// 或者也可以 faY 合并到 faX 上
		// father[faY] = faX
		// rank[faX]++
	} else if rank[faX] > rank[faY] {
		father[faY] = faX
	}
}

// 计算一共有多少个集合
func count() int {
	n := len(father)
	res := 0
	for i := 0; i < n; i++ {
		// 代表元素父节点就是其本身
		if i == father[i] {
			res++
		}
	}
	return res
}

// 打印并查集
func printUnionFindSet() {
	n := len(father)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", father[i])
	}
	fmt.Println()
}
