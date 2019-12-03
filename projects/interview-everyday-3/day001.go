package main

import "fmt"

// Slice ...
type Slice []int

// NewSlice ...
func NewSlice() Slice {
	return make(Slice, 0)
}

// Add ...
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

func day001Test1() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}

func day001Test2() {
	s := NewSlice()
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}

func day001() {
	// day001Test1() // 132
	day001Test2() // 312
}
