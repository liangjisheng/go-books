package main

import (
	"fmt"
	"sort"
)

type Set map[int]bool

func New(items ...int) *Set {
	s := make(Set, len(items))
	s.Add(items...)
	return &s
}

// Duplicate 创建副本
func (s *Set) Duplicate() *Set {
	r := make(Set, len(*s))
	for e := range *s {
		r[e] = true
	}
	return &r
}

func (s *Set) Add(items ...int) {
	for _, v := range items {
		(*s)[v] = true
	}
}

func (s *Set) Remove(items ...int) {
	for _, v := range items {
		delete(*s, v)
	}
}

func (s *Set) Has(items ...int) bool {
	for _, v := range items {
		if _, ok := (*s)[v]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) Count() int {
	return len(*s)
}

func (s *Set) Clear() {
	*s = map[int]bool{}
}

func (s *Set) Empty() bool {
	return len(*s) == 0
}

func (s *Set) List() []int {
	list := make([]int, 0, len(*s))
	for item := range *s {
		list = append(list, item)
	}
	return list
}

func (s *Set) SortedList() []int {
	list := (*s).List()
	sort.Ints(list)
	return list
}

// Union 并集 获取 s 与参数的并集，结果存入 s
func (s *Set) Union(sets ...*Set) {
	for _, set := range sets {
		for e := range *set {
			(*s)[e] = true
		}
	}
}

// Union 并集（函数）
// 获取所有参数的并集，并返回
func Union(sets ...*Set) *Set {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取并集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range *set {
			(*r)[e] = true
		}
	}
	return r
}

// Minus 差集
// 获取 s 与所有参数的差集，结果存入 s
func (s *Set) Minus(sets ...*Set) {
	for _, set := range sets {
		for e := range *set {
			delete(*s, e)
		}
	}
}

// Minus 差集（函数）
// 获取第 1 个参数与其它参数的差集，并返回
func Minus(sets ...*Set) *Set {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取差集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range *set {
			delete(*r, e)
		}
	}
	return r
}

// Intersect 交集
// 获取 s 与其它参数的交集，结果存入 s
func (s *Set) Intersect(sets ...*Set) {
	for _, set := range sets {
		for e := range *s {
			if _, ok := (*set)[e]; !ok {
				delete(*s, e)
			}
		}
	}
}

// Intersect 交集（函数）
// 获取所有参数的交集，并返回
func Intersect(sets ...*Set) *Set {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取交集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range *r {
			if _, ok := (*set)[e]; !ok {
				delete(*r, e)
			}
		}
	}
	return r
}

// Complement 补集
// 获取 s 相对于 full 的补集，结果存入 s
func (s *Set) Complement(full *Set) {
	r := s.Duplicate()
	s.Clear()
	for e := range *full {
		if _, ok := (*r)[e]; !ok {
			(*s)[e] = true
		}
	}
}

// Complement 补集（函数）
// 获取 sub 相对于 full 的补集，并返回
func Complement(sub, full *Set) *Set {
	r := full.Duplicate()
	for e := range *sub {
		delete(*r, e)
	}
	return r
}

func setRoutineUnsafe() {
	s1 := New(1, 2, 3, 4, 5, 6, 7, 8)
	s2 := New(3, 4, 5, 6)
	s3 := New(5, 6, 8, 9)
	r1 := Union(s1, s2, s3)     // 获取并集
	r2 := Minus(s1, s2, s3)     // 获取差集
	r3 := Intersect(s1, s2, s3) // 获取交集
	r4 := Complement(s2, s1)    // 获取 s2 相对于 s1 的补集
	fmt.Println(r1.SortedList())
	fmt.Println(r2.SortedList())
	fmt.Println(r3.SortedList())
	fmt.Println(r4.SortedList())
}
