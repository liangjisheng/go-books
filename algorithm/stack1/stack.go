package stack

// 数组实现栈

//Stack ...
type Stack struct {
	items []string
	current int
}

//NewStack ...
func NewStack() *Stack {
	return &Stack{
		items: make([]string, 10),
		current: 0,
	}
}

func (s *Stack) Push(item string) {
	s.current++
	if len(s.items) == s.current {
		s.items = append(s.items, item)
		return
	}
	s.items[s.current] = item
}

func (s *Stack) Pop() string {
	if s.current == 0 {
		return ""
	}

	item := s.items[s.current]
	s.current--
	return item
}
