package stack

import "math"

// 双向循环链表实现栈

type nodeInt struct {
	prev, next *nodeInt
	value int
}

type StackInt struct {
	root *nodeInt
	len int
}

func NewStackInt() *StackInt {
	n := &nodeInt{}
	n.prev = n
	n.next = n
	return &StackInt{root:n}
}

// Push 放到最后
func (s *StackInt) Push(item int) {
	n := &nodeInt{value:item}

	s.root.prev.next = n
	n.prev = s.root.prev

	n.next = s.root
	s.root.prev = n
	s.len++
}

// Pop 获取最后一个元素
func (s *StackInt) Pop() int {
	item := s.root.prev
	if item == s.root {
		return math.MaxInt32
	}

	s.root.prev = item.prev
	item.prev.next = s.root

	// 避免内存泄露
	item.prev = nil
	item.next = nil
	s.len--
	return item.value
}
