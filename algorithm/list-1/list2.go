package main

// 主要讲解快慢指针的应用: 链表有环的情况

// 判断链表是否有环, 如果有, 找到环的入口位置, 要求空间复杂度为O(1)

// 检测一个链表是否有环, 并返回相遇的节点, 判断有环为啥要返回相遇的结点
// 而不是返回 true 或 false 呢. 因为题目中还有一个要求, 判断环的入口位置
func (list *List) detectCrossNode() *Node {
	fast := list.head
	slow := list.head

	for fast.next != nil && fast != nil {
		fast = fast.next.next
		slow = slow.next

		// fast 现在指向尾节点的 next
		if fast == nil {
			return nil
		}

		// 快慢指针相遇
		if slow.data == fast.data {
			return slow
		}
	}

	// fast 指向尾节点
	return nil
}

// 要找到入口结点, 只需定义两个指针, 一个指针指向 head, 一个指针指向快慢指针的相遇点
// 然后这两个指针不断遍历(每次走一步), 当它们指向同一个结点时即是环的入口结点
func (list *List) getRingEntryNode() *Node {
	// 获取快慢指针相遇结点
	crossNode := list.detectCrossNode()
	// 如果没有相遇点, 则没有环
	if crossNode == nil {
		return nil
	}

	// 分别定义两个指针, 一个指向头结点, 一个指向快慢指针相交结点
	tmp1 := list.head
	tmp2 := crossNode

	// 两者相遇点即为环的入口结点
	for tmp1.data != tmp2.data {
		tmp1 = tmp1.next
		tmp2 = tmp2.next
	}
	return tmp1
}
