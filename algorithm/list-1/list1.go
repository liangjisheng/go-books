package main

// 主要讲解快慢指针的应用

// 题目1
// 给定一个带有头结点 head 的非空单链表, 返回链表的中间结点.
// 如果有两个中间结点, 则返回第二个中间结点

// 解法一 要知道链表的中间结点, 首先我们需要知道链表的长度, 说到链表长度大家想到了啥
// 还记得我们在上文中说过哨兵结点可以保存链表的长度吗, 这样直接 从 head 的后继结点
// 开始遍历 链表长度 / 2 次即可找到中间结点
func (list *List) findMiddleNodeMethod1() *Node {
	tmp := list.head.next
	midLen := list.length / 2 // 链表长度奇偶通用
	for midLen > 0 {
		tmp = tmp.next
		midLen--
	}
	return tmp
}

// 解法二 如果哨兵结点里没有定义长度呢, 那就要遍历一遍链表拿到链表长度(定义为 length)了
// 然后再从头结点开始遍历 length / 2 次即为中间结点

// 解法三
// 解法二由于要遍历两次链表, 显得不是那么高效, 那能否只遍历一次链表就能拿到中间结点呢
// 这里就引入我们的快慢指针了, 主要有三步
// 1.快慢指针同时指向 head 的后继结点
// 2.慢指针走一步, 快指针走两步
// 3.不断地重复步骤2, 什么时候停下来呢, 这取决于链表的长度是奇数还是偶数
// 如果链表长度为奇数, 当 fast.next = null 时,slow 为中间结点
// 如果链表长度为偶数, 当 fast = null 时,slow 为中间结点
// 由以上分析可知: 当 fast = null 或者 fast.next = null 时,
// 此时的 slow 结点即为我们要求的中间结点, 否则不断地重复步骤 2
func (list *List) findMiddleNodeWithSlowFastPointer() *Node {
	slow := list.head.next
	fast := list.head.next

	for fast != nil && fast.next != nil {
		fast = fast.next.next // 快指针走两步
		slow = slow.next      // 慢指针走一步
	}

	// 此时的 slow 结点即为中间结点
	return slow
}

// 题目2
// 输入一个链表, 输出该链表中的倒数第 k 个结点. 比如链表为 head-->1-->2-->3-->4-->5
// 求倒数第三个结点(即值为 3 的节点)

// 解法一 如果要求逆序的第 k 个结点, 常规的做法是先顺序遍历一遍链表, 拿到链表长度
// 然后再遍历 链表长度-k 次即可, 这样要遍历两次链表, 不是那么高效
func (list *List) findKthToTailMethod1(k int) *Node {
	if k > list.length {
		return nil
	}

	tmp := list.head.next
	step := list.length - k
	for step > 0 {
		tmp = tmp.next
		step--
	}
	return tmp
}

// 解法二 快慢指针解法
// 1.首先让快慢指针同时指向 head 的后继结点
// 2.快指针往前走 k- 1 步,先走到第 k 个结点
// 3.快慢指针同时往后走一步, 不断重复此步骤, 直到快指针走到尾结点
// 此时的 slow 结点即为我们要找的倒序第 k 个结点
func (list *List) findKthToTailMethod2(k int) *Node {
	slow := list.head.next
	fast := list.head.next

	// 快指针先移到第k个结点
	tmpK := k - 1
	for tmpK > 0 {
		// 链表为空或者 k 大于链表长度
		if fast == nil {
			return nil
		}
		fast = fast.next
		tmpK--
	}

	// slow 和 fast 同时往后移, 直到 fast 走到尾结点
	for fast.next != nil {
		fast = fast.next
		slow = slow.next
	}
	return slow
}

// 给定一个单链表, 设计一个算法实现链表向右旋转 K 个位置. 举例：给定
// head->1->2->3->4->5->NULL, K=3,右旋后即为 head->3->4->5-->1->2->NULL
// 分析: 这道题其实是对求倒序第 K 个位置的的一个变形, 主要思路如下
// 先找到倒数第 K+1 个结点, 此结点的后继结点即为倒数第 K 个结点
// 将倒数第 K+1 结点的的后继结点设置为 null
// 将 head 的后继结点设置为以上所得的倒数第 K 个结点, 将原尾结点的后继结点设置为原 head 的后继结点
func (list *List) reversedKthToTail(k int) {
	// 倒数第 k+1 个节点
	kPreNode := list.findKthToTailMethod2(k + 1)
	// 倒数第 k 个节点
	kNode := kPreNode.next
	kPreNode.next = nil

	headNext := list.head.next
	list.head.next = kNode

	// 寻找尾节点
	tmp := kNode
	for tmp.next != nil {
		tmp = tmp.next
	}
	// 尾结点的后继结点设置为原 head 的后继结点
	tmp.next = headNext
}

// 输入一个链表, 删除该链表中的倒数第 k 个结点
func (list *List) deleteKthToTail(k int) {
	// 倒数第 k+1 个节点
	kPreNode := list.findKthToTailMethod2(k + 1)
	if kPreNode == nil || kPreNode.next == nil {
		return
	}
	kPreNode.next = kPreNode.next.next
}
