package stack

// 双向循环链表实现栈

type node struct {
	prev, next *node
	value string
}

type Stack struct {
	root *node
	len int
}

func NewStack() *Stack {
	n := &node{}
	n.prev = n
	n.next = n
	return &Stack{root:n}
}

// Push 放到最后
func (s *Stack) Push(item string) {
	n := &node{value:item}

	s.root.prev.next = n
	n.prev = s.root.prev

	n.next = s.root
	s.root.prev = n
	s.len++
}

// Pop 获取最后一个元素
func (s *Stack) Pop() string {
	item := s.root.prev
	if item == s.root {
		return ""
	}

	s.root.prev = item.prev
	item.prev.next = s.root

	// 避免内存泄露
	item.prev = nil
	item.next = nil
	s.len--
	return item.value
}
