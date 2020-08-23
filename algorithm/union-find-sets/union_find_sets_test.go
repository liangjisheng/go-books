package unionfindsets

import (
	"fmt"
	"testing"
)

func TestUnionFindSet(t *testing.T) {
	initUnionFindSet(10)
	printUnionFindSet()

	merge(0, 1)
	merge(0, 2)
	merge(3, 4)
	merge(3, 1)
	merge(5, 7)
	merge(7, 8)
	merge(7, 8)
	printUnionFindSet()

	fmt.Println("count:", count())
}

func TestUnionFindSet1(t *testing.T) {
	initUnionFindSet(10)
	printUnionFindSet()

	mergeWithRank(0, 1)
	mergeWithRank(1, 2)
	mergeWithRank(2, 3)
	mergeWithRank(3, 4)
	mergeWithRank(4, 5)
	mergeWithRank(5, 6)
	mergeWithRank(6, 7)
	mergeWithRank(7, 8)
	mergeWithRank(8, 9)
	printUnionFindSet()

	fmt.Println("count:", count())
}

// 题目背景
// 若某个家族人员过于庞大，要判断两个是否是亲戚，确实还很不容易，现在给出某个亲戚关系图，
// 求任意给出的两个人是否具有亲戚关系。
// 题目描述
// 规定：x和y是亲戚，y和z是亲戚，那么x和z也是亲戚。如果x,y是亲戚，那么x的亲戚都是y的亲戚，y的亲戚也都是x的亲戚。
// 输入格式
// 第一行：三个整数n,m,p，（n<=5000,m<=5000,p<=5000），分别表示有n个人，m个亲戚关系，询问p对亲戚关系。
// 以下m行：每行两个数Mi，Mj，1<=Mi，Mj<=N，表示Mi和Mj具有亲戚关系。
// 接下来p行：每行两个数Pi，Pj，询问Pi和Pj是否具有亲戚关系。
// 输出格式
// P行，每行一个’Yes’或’No’。表示第i个询问的答案为“具有”或“不具有”亲戚关系。

// 这其实是一个很有现实意义的问题。我们可以建立模型，把所有人划分到若干个不相交的集合中
// 每个集合里的人彼此是亲戚。为了判断两个人是否为亲戚，只需看它们是否属于同一个集合即可
// 因此，这里就可以考虑用并查集进行维护了

func TestRelativeQuestion(t *testing.T) {
	var n, m, p, x, y int
	fmt.Scanf("%d%d%d", &n, &m, &p)
	initUnionFindSet(n)

	for i := 0; i < m; i++ {
		fmt.Scanf("%d%d", &x, &y)
		mergeWithRank(x, y)
	}
	for i := 0; i < p; i++ {
		fmt.Scanf("%d%d", &x, &y)
		faX := findRecursionPathCompress(x)
		faY := findRecursionPathCompress(y)
		if faX == faY {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
