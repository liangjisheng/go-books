package stack

import (
	"errors"
	"fmt"
	"log"
)

// Stack ...
type Stack struct {
	elems []int
	n     int
	top   int
}

// NewStack ...
func NewStack(n int) *Stack {
	return &Stack{
		elems: make([]int, n, n),
		n:     n,
	}
}

func (s *Stack) push(elem int) *Stack {
	if s.top == s.n {
		log.Println("stack is full.")
		return s
	}

	s.elems[s.top] = elem
	s.top++
	return s
}

func (s *Stack) pop() (int, error) {
	if s.top == 0 {
		log.Println("stack is empty.")
		return -1, errors.New("stack is empty")
	}

	s.top--
	return s.elems[s.top], nil
}

func (s *Stack) print() {
	for top := s.top - 1; top >= 0; top-- {
		fmt.Printf("%d ", s.elems[top])
	}
	fmt.Println()
}
