package main

// 主要讲解链表反转

import (
	"fmt"
)

// 使用一个虚拟的节点作为头结点, 即我们常说的哨兵

// Node ...
type Node struct {
	data int
	next *Node
}

// List ...
type List struct {
	head   *Node
	length int
}

// NewNode ...
func NewNode(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

// NewList ...
func NewList() *List {
	return &List{
		head:   NewNode(0), // 哨兵结点
		length: 0,
	}
}

// TailInsert 在尾插入新的节点
func (list *List) TailInsert(data int) {
	tmp := list.head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = NewNode(data)
	list.length++
}

// HeadInsert 在头插入新的节点
func (list *List) HeadInsert(data int) {
	node := NewNode(data)
	node.next = list.head.next
	list.head.next = node
	list.length++
}

// Print ...
func (list *List) Print() {
	fmt.Print("list: ")
	tmp := list.head.next
	for tmp != nil {
		fmt.Printf("%d ", tmp.data)
		tmp = tmp.next
	}
	fmt.Println()
}

// 给定单向链表的头指针和一个节点指针, 定义一个函数在 O(1) 内删除这个节点
// 注意题目没有规定说不能改变结点中的值, 所以有一种很巧妙的方法, 狸猫换太子!
// 我们先通过结点 2 找到结点 3, 再把节点 3 的值赋给结点 2, 此时结点 2 的值变成了 3
// 这时候问题就转化成了上图这种比较简单的需求, 即根据结点 2 把结点 3 移除即可
// 不过需要注意的是这种解题技巧只适用于被删除的指定结点是中间结点的情况
// 如果指定结点是尾结点, 还是要老老实实地找到尾结点的前继结点, 再把尾结点删除
func (list *List) removeSelectedNode(deleteNode *Node) {
	// deleteNode 是尾节点
	if deleteNode.next == nil {
		tmp := list.head
		for tmp.next != deleteNode {
			tmp = tmp.next
		}
		// 找到尾结点的前继结点, 把尾结点删除
		tmp.next = nil
	} else {
		nextNode := deleteNode.next
		// 将删除结点的后继结点的值赋给被删除结点
		deleteNode.data = nextNode.data
		// 将 nextNode 结点删除
		deleteNode.next = nextNode.next
		nextNode.next = nil
	}
}

// Invert 时间/空间复杂度 由于递归调用了 n 次 Invert 函数
// 所以时间复杂度显然是 O(n), 空间复杂度呢, 没有用到额外的空间, 但是由于递归调用了
// n 次 Invert 函数, 压了 n 次栈, 所以空间复杂度也是 O(n)
func (list *List) Invert(node *Node) *Node {
	// 递归结束条件,只有一个节点结束递归
	if node.next == nil {
		return node
	}

	// 步骤 1: 先翻转 node 之后的链表
	newHead := list.Invert(node.next)

	// 步骤 2: 再把原 node 节点后继结点的后继结点指向 node, node 的后继节点设置为空(防止形成环)
	node.next.next = node
	node.next = nil

	// 步骤 3: 返回翻转后的头结点
	return newHead
}

// IterationInvert 用迭代的思路来做由于循环了 n 次, 显然时间复杂度为 O(n),
// 另外由于没有额外的空间使用, 也未像递归那样调用递归函数不断压栈
// 所以空间复杂度是 O(1),对比递归, 显然应该使用迭代的方式来处理
func (list *List) IterationInvert(node *Node) *Node {
	pre := node
	cur := node.next
	// pre 是第一个结点, 避免翻转链表后形成环
	pre.next = nil

	for cur != nil {
		// 务必注意!!!: 在 cur 指向 pre 之前一定要先保留 cur 的后继结点
		// 不然如果 cur 先指向 pre, 之后就再也找不到后继结点了
		nextNode := cur.next
		cur.next = pre
		pre = cur
		cur = nextNode
	}

	// 此时 pre 指向的是原链表的尾结点, 翻转后即为链表 head 的后继结点
	return pre
}

// IterationInvertPartList 变形题 1: 给定一个链表的头结点 head,以及两个整数 from 和 to, 在链表上把
// 第 from 个节点和第 to 个节点这一部分进行翻转. 例如: 给定如下链表,
// from = 2, to = 4 head-->5-->4-->3-->2-->1 将其翻转后
// 链表变成 head-->5--->2-->3-->4-->1
func (list *List) IterationInvertPartList(fromIdx, toIdx int) {
	var fromPre, from *Node
	var to, toNext *Node

	// 步骤 1: 找到 from-1, from, to, to+1 这四个结点
	tmp := list.head.next
	curIdx := 1 // 头结点的index为1
	for tmp != nil {
		if curIdx == fromIdx-1 {
			fromPre = tmp
		} else if curIdx == fromIdx {
			from = tmp
		} else if curIdx == toIdx {
			to = tmp
		} else if curIdx == toIdx+1 {
			toNext = tmp
		}

		tmp = tmp.next
		curIdx++
	}

	// from 或 to 都超过尾结点不翻转
	if from == nil || to == nil {
		return
	}

	// 步骤2: 使用循环迭代法翻转从 from 到 to 的结点
	pre := from
	cur := from.next
	pre.next = nil
	for cur != toNext {
		nextNode := cur.next
		cur.next = pre
		pre = cur
		cur = nextNode
	}

	// 步骤3: 将 from-1 节点指向 to 结点(如果从 head 的后继结点开始翻转,
	// 则需要重新设置 head 的后继结点), 将 from 结点指向 to + 1 结点
	if fromPre != nil {
		fromPre.next = to
	} else {
		list.head.next = to
	}
	from.next = toNext
}

// IterationInvertLinkedListEveryK 变形题 2: 给出一个链表, 每 k 个节点一组进行翻转,
// 并返回翻转后的链表, k 是一个正整数, 它的值小于或等于链表的长度. 如果节点总数不是
// k 的整数倍,那么将最后剩余节点保持原有顺序, 示例: 给定这个链表:
// head-->1->2->3->4->5 当 k = 2 时, 应当返回: head-->2->1->4->3->5
// 当 k = 3 时, 应当返回: head-->3->2->1->4->5 说明:
// 你的算法只能使用常数的额外空间
// 你不能只是单纯的改变节点内部的值, 而是需要实际的进行节点交换
func (list *List) IterationInvertLinkedListEveryK(node *Node, k int) *Node {
	tmp := node.next
	step := 0 // 计数用来找出首节点和尾节点

	var startK, endK *Node // k 个一组链表中的头节点和尾节点
	startKPre := node      // k 个一组链表中的头节点的前置节点

	for tmp != nil {
		// tmp 的下一个节点, 因为由于翻转, tmp 的后继结点会变,要提前保存
		tmpNext := tmp.next
		if step == 0 {
			startK = tmp
			step++
		} else if step == k-1 {
			// 此时找到了 k 个一组链表区间的尾结点(endK), 对这段链表用迭代进行翻转
			endK = tmp

			pre := startK
			cur := startK.next
			// 已经到尾节点
			if cur == nil {
				break
			}
			endKNext := endK.next
			pre.next = nil
			for cur != endKNext {
				next := cur.next
				cur.next = pre
				pre = cur
				cur = next
			}

			// 翻转后此时 endK 和 startK 分别是是 k 个一组链表中的首尾结点
			startKPre.next = endK
			startK.next = endKNext

			// 当前的 k 个一组翻转完了, 开始下一个 k 个一组的翻转
			startKPre = startK
			step = 0
		} else {
			step++
		}
		tmp = tmpNext
	}

	return node
}

// InvertLinkedListEveryK ...
func (list *List) InvertLinkedListEveryK(node *Node, k int) *Node {
	// 结束条件
	if node == nil || node.next == nil {
		return node
	}

	// 前进 k-1 步
	cur := node
	for i := 1; i < k && cur != nil; i++ {
		cur = cur.next
	}
	// 判断是否能组成一组, 不能的话就结束
	if cur == nil {
		return node
	}

	// tmp 指向剩余的链表
	tmp := cur.next
	cur.next = nil

	// 把 k 个节点进行反转
	newHead := list.Invert(node)

	// 把之后的部分链表进行每K个节点逆转
	newTmpHead := list.InvertLinkedListEveryK(tmp, k)

	// 把两部分节点连接起来
	node.next = newTmpHead

	return newHead
}

// ReverseIterationInvertLinkedListEveryK 变形题 3: 变形 2 针对的是顺序的 k 个一组翻转
// 那如何逆序 k 个一组进行翻转呢
// 例如: 给定如下链表, head-->1-->2-->3-->4-->5 逆序 k 个一组翻转后, 链表变成
// head-->1--->3-->2-->5-->4 (k = 2 时)
func (list *List) ReverseIterationInvertLinkedListEveryK(k int) {
	// 先翻转链表
	list.head.next = list.IterationInvert(list.head.next)
	// k 个一组翻转链表
	list.head = list.IterationInvertLinkedListEveryK(list.head, k)
	// 再次翻转链表
	list.head.next = list.IterationInvert(list.head.next)
}
