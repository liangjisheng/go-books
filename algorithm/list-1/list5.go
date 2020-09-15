package main

import "fmt"

// cycleList 将一个链表转换成环形链表
func cycleList(list *List) {
	tmp := list.head.next
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = list.head.next
}

// 环形单链表约瑟夫问题, 返回最后存活的节点 时间复杂度为 O(n * m)
func josephus1(cyclelist *List, m int) *Node {
	pre := cyclelist.head
	cur := cyclelist.head.next

	num := 1
	for cur.next != cur {
		if num == m {
			fmt.Printf("%d dead\n", cur.data)
			pre.next = cur.next // 删除 cur
			cur = cur.next
			num = 1
		} else {
			pre = cur
			cur = cur.next
			num++
		}
	}

	return cur
}

func josephus2(cyclelist *List, m int) *Node {
	// 链表中没有数据
	if cyclelist.head.next == nil || m < 1 {
		return cyclelist.head.next
	}

	dest := liveNumber(cyclelist.length, m)
	tmp := cyclelist.head.next
	for i := 1; i < dest; i++ {
		tmp = tmp.next
	}
	return tmp
}

// liveNumber 表示链表中元素个数为 n 时, 返回存活的编号 [1-n]
func liveNumber(n, m int) int {
	// 只有一个元素时, 递归结束条件
	if n == 1 {
		return 1
	}

	return (liveNumber(n-1, m)+m-1)%n + 1
}
