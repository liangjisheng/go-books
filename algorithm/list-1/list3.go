package main

// 主要讲解三种方法优雅判断回文链表

// 方法1: 可以利用另一个链表来做辅助, 把原链表的节点使用头插法, 插入到另一个链表中

// 时间复杂度为 O(n), 空间复杂度为 O(n)
func (list *List) isPalindromicList() bool {
	list1 := NewList()

	// 遍历原链表, 构造新的链表, 新的链表是原链表的反转
	tmp := list.head.next
	for tmp != nil {
		list1.HeadInsert(tmp.data)
		tmp = tmp.next
	}

	// 遍历原链表和新链表, 判断是否相等
	tmp = list.head.next
	tmp1 := list1.head.next
	for tmp != nil && tmp1 != nil {
		if tmp.data != tmp1.data {
			return false
		}
		tmp = tmp.next
		tmp1 = tmp1.next
	}
	return true
}

// 方法2: 真的需要全部插入新链表吗? 其实也可以让原链表的后半部分插入新的链表就可以了
// 然后把新链表中的元素与原链表的前半部分对比, 这样做的话空间复杂度会减少一半
func (list *List) isPalindromicList2() bool {
	// 首先使用快慢指针法找到链表的中间节点
	midNode := list.findMiddleNodeWithSlowFastPointer()

	// 从中间节点遍历原链表, 构造新的链表, 新的链表是原链表后半部分的反转
	list1 := NewList()
	for midNode != nil {
		list1.HeadInsert(midNode.data)
		midNode = midNode.next
	}

	// 遍历原链表和新链表, 判断是否相等
	tmp := list.head.next
	tmp1 := list1.head.next
	for tmp != nil && tmp1 != nil {
		if tmp.data != tmp1.data {
			return false
		}
		tmp = tmp.next
		tmp1 = tmp1.next
	}
	return true
}

// 方法3: 可以把链表的后半部分进行反转, 然后再用后半部分与前半部分进行比较就可以了
// 这种做法额外空间复杂度只需要 O(1), 时间复杂度为 O(n)
func (list *List) isPalindromicList3() bool {
	// 首先使用快慢指针法找到链表的中间节点
	midNode := list.findMiddleNodeWithSlowFastPointer()

	// 原链表后半部分反转后的 new head
	revHead := list.Invert(midNode)

	// 遍历原链表和新链表, 判断是否相等
	tmp := list.head.next
	tmp1 := revHead
	for tmp != nil && tmp1 != nil {
		if tmp.data != tmp1.data {
			return false
		}
		tmp = tmp.next
		tmp1 = tmp1.next
	}
	return true
}
