package sort

import (
	"math/rand"
	"time"
)

// 平均时间复杂度: O(n + k)
// 最佳时间复杂度: O(n + k)
// 最差时间复杂度: O(n ^ 2)
// 空间复杂度: O(n * k)
// 稳定性: 稳定

// 桶排序最好情况下使用线性时间O(n), 桶排序的时间复杂度, 取决与对各个桶之间数据
// 进行排序的时间复杂度, 因为其它部分的时间复杂度都为O(n). 很显然, 桶划分的越小,
// 各个桶之间的数据越少, 排序所用的时间也会越少. 但相应的空间消耗就会增大

// 桶排序(Bucket sort)或所谓的箱排序, 是一个排序算法, 工作的原理是将数组分到有限数量的桶里
// 每个桶再个别排序(有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排序),
// 最后依次把各个桶中的记录列出来记得到有序序列. 桶排序是鸽巢排序的一种归纳结果
// 当要被排序的数组内的数值是均匀分配的时候, 桶排序使用线性时间(Θ(n))
// 但桶排序并不是比较排序, 他不受到O(nlogn)下限的影响

// 假设数据分布在[0, 100)之间, 每个桶内部用链表表示, 在数据入桶的同时插入排序
// 然后把各个桶中的数据合并

// Node ...
type Node struct {
	data int
	next *Node
}

// 将val有序插入到链表中
func insert(head *Node, val int) *Node {
	if head == nil {
		return &Node{
			data: val,
		}
	}

	list := &Node{
		data: 0,
		next: head,
	}
	pre := list
	cur := list.next

	for cur != nil && cur.data <= val {
		pre = cur
		cur = cur.next
	}

	node := &Node{
		data: val,
		next: cur,
	}
	pre.next = node
	return list.next
}

// 有序合并2个链表, 返回合并后的head
func merge(head1, head2 *Node) *Node {
	i := head1
	j := head2
	var head Node
	k := &head

	for i != nil && j != nil {
		if i.data <= j.data {
			k.next = i
			i = i.next
		} else {
			k.next = j
			j = j.next
		}
		k = k.next
	}

	if i != nil {
		k.next = i
	}

	if j != nil {
		k.next = j
	}

	return head.next
}

// PreBucketNum 每个桶最多多少个元素
var PreBucketNum = 2

// BucketSort ...
func BucketSort(array []int) {
	n := len(array)
	if n <= 1 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	bucketNum := n/PreBucketNum + 1
	bucket := make([]*Node, bucketNum)

	for i := 0; i < n; i++ {
		bucketI := rand.Intn(bucketNum)
		bucket[bucketI] = insert(bucket[bucketI], array[i])
	}

	head := bucket[0]
	for i := 1; i < bucketNum; i++ {
		head = merge(head, bucket[i])
	}

	for i := 0; i < n; i++ {
		array[i] = head.data
		head = head.next
	}
}

// 桶排序是计数排序的变种, 它利用了函数的映射关系, 高效与否的关键就在于这个映射函数的确定
// 把计数排序中相邻的m个”小桶”放到一个”大桶”中, 在分完桶后, 对每个桶进行排序(一般用快排),
// 然后合并成最后的结果
// 算法思想和散列中的开散列法差不多, 当冲突时放入同一个桶中; 可应用于数据量分布比较均匀
// 或比较侧重于区间数量时
// 桶排序最关键的建桶, 如果桶设计得不好的话桶排序是几乎没有作用的. 通常情况下
// 上下界有两种取法, 第一种是取一个10^n或者是2^n的数, 方便实现
// 另一种是取数列的最大值和最小值然后均分作桶.
