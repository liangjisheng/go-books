package main

// 给定一个单向链表的头结点head,节点的值类型是整型, 再给定一个整数privot
// 实现一个调整链表的函数, 将链表调整为左部分都是值小于privot的节点
// 中间部分都是值等于privot的节点, 右部分都是大于privot的节点
// 且对某部分内部节点的顺序不做要求
// 例如: 链表 9-0-4-5-1, pivot=3
// 调整后是 1-0-4-9-5, 也可以是0-1-9-5-4

func (list *List) toArray() []int {
	array := make([]int, 0)
	tmp := list.head.next
	for tmp != nil {
		array = append(array, tmp.data)
		tmp = tmp.next
	}
	return array
}

// 解法一: 一种很简单的方法就是用一个数组来存链表的节点, 按照某个值把他们进行划分
//  时间和空间复杂度都为 O(N)
func (list *List) split1(privot int) *List {
	array := list.toArray()
	n := len(array)

	arrayLess := make([]int, 0)
	arrayEqual := make([]int, 0)
	arrayGreat := make([]int, 0)

	for i := 0; i < n; i++ {
		if array[i] < privot {
			arrayLess = append(arrayLess, array[i])
		} else if array[i] == privot {
			arrayEqual = append(arrayEqual, array[i])
		} else if array[i] > privot {
			arrayGreat = append(arrayGreat, array[i])
		}
	}

	list1 := NewList()
	lenLess := len(arrayLess)
	for i := 0; i < lenLess; i++ {
		list1.TailInsert(arrayLess[i])
	}

	lenEqual := len(arrayEqual)
	for i := 0; i < lenEqual; i++ {
		list1.TailInsert(arrayEqual[i])
	}

	lenGreat := len(arrayGreat)
	for i := 0; i < lenGreat; i++ {
		list1.TailInsert(arrayGreat[i])
	}
	return list1
}

// 也可以采取使用3个指针, 把原链表依次划分成三个部分的链表, 然后再把他们合并起来
// 这种做法不但空间复杂度为 O(1), 而且内部节点的顺序也是和原链表一样的
func (list *List) split2(privot int) *List {
	var lessBegin, lessEnd *Node
	var equalBegin, equalEnd *Node
	var greatBegin, greatEnd *Node
	list1 := NewList()

	// 划分
	tmp := list.head.next
	for tmp != nil {
		next := tmp.next
		tmp.next = nil
		if tmp.data < privot {
			// 第一个小于 privot 的值
			if lessBegin == nil {
				lessBegin = tmp
				lessEnd = tmp
			} else {
				lessEnd.next = tmp
				lessEnd = lessEnd.next
			}
		} else if tmp.data == privot {
			// 第一个等于 privot 的值
			if equalBegin == nil {
				equalBegin = tmp
				equalEnd = tmp
			} else {
				equalEnd.next = tmp
				equalEnd = equalEnd.next
			}
		} else {
			if greatBegin == nil {
				greatBegin = tmp
				greatEnd = tmp
			} else {
				greatEnd.next = tmp
				greatEnd = greatEnd.next
			}
		}

		tmp = next
	}

	// 把三部分串连起来

	// 第一部分为空
	if lessEnd == nil {
		if equalEnd == nil {
			list1.head.next = greatBegin
			return list1
		}
		equalEnd.next = greatBegin
		list1.head.next = equalBegin
		return list1
	}

	// 第二部分为空
	if equalEnd == nil {
		lessEnd.next = greatBegin
		list1.head.next = lessBegin
		return list1
	}

	lessEnd.next = equalBegin
	equalEnd.next = greatBegin
	list1.head.next = lessBegin
	return list1
}
