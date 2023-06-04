package main

import (
	"fmt"
	"sort"
	"sync"
)

type Set struct {
	sync.RWMutex
	m map[int]bool
}

func New(items ...int) *Set {
	s := &Set{
		m: make(map[int]bool, len(items)),
	}
	s.Add(items...)
	return s
}

func (s *Set) Duplicate() *Set {
	s.Lock()
	defer s.Unlock()
	r := &Set{
		m: make(map[int]bool, len(s.m)),
	}
	for e := range s.m {
		r.m[e] = true
	}
	return r
}

func (s *Set) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		s.m[v] = true
	}
}

func (s *Set) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		delete(s.m, v)
	}
}

func (s *Set) Has(items ...int) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range items {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}
	return true
}

func (s *Set) Count() int {
	s.Lock()
	defer s.Unlock()
	return len(s.m)
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

func (s *Set) Empty() bool {
	s.Lock()
	defer s.Unlock()
	return len(s.m) == 0
}

func (s *Set) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

func (s *Set) SortedList() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// Union 并集
// 获取 s 与参数的并集，结果存入 s
func (s *Set) Union(sets ...*Set) {
	// 为了防止多例程死锁，不能同时锁定两个集合
	// 所以这里没有锁定 s，而是创建了一个临时集合
	r := s.Duplicate()
	// 获取并集
	for _, set := range sets {
		set.Lock()
		for e := range set.m {
			r.m[e] = true
		}
		set.Unlock()
	}
	// 将结果转入 s
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
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
		set.Lock()
		for e := range set.m {
			r.m[e] = true
		}
		set.Unlock()
	}
	return r
}

// Minus 差集
// 获取 s 与所有参数的差集，结果存入 s
func (s *Set) Minus(sets ...*Set) {
	// 为了防止多例程死锁，不能同时锁定两个集合
	// 所以这里没有锁定 s，而是创建了一个临时集合
	r := s.Duplicate()
	// 获取差集
	for _, set := range sets {
		set.Lock()
		for e := range set.m {
			delete(r.m, e)
		}
		set.Unlock()
	}
	// 将结果转入 s
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
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
		for e := range set.m {
			delete(r.m, e)
		}
	}
	return r
}

// Intersect 交集
// 获取 s 与其它参数的交集，结果存入 s
func (s *Set) Intersect(sets ...*Set) {
	// 为了防止多例程死锁，不能同时锁定两个集合
	// 所以这里没有锁定 s，而是创建了一个临时集合
	r := s.Duplicate()
	// 获取交集
	for _, set := range sets {
		set.Lock()
		for e := range r.m {
			if _, ok := set.m[e]; !ok {
				delete(r.m, e)
			}
		}
		set.Unlock()
	}
	// 将结果转入 s
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
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
		for e := range r.m {
			if _, ok := set.m[e]; !ok {
				delete(r.m, e)
			}
		}
	}
	return r
}

// Complement 补集
// 获取 s 相对于 full 的补集，结果存入 s
func (s *Set) Complement(full *Set) {
	r := full.Duplicate()
	s.Lock()
	defer s.Unlock()
	// 获取补集
	for e := range s.m {
		delete(r.m, e)
	}
	// 将结果转入 s
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
	}
}

// Complement 补集（函数）
// 获取 sub 相对于 full 的补集，并返回
func Complement(sub, full *Set) *Set {
	r := full.Duplicate()
	sub.Lock()
	defer sub.Unlock()
	for e := range sub.m {
		delete(r.m, e)
	}
	return r
}

func setRoutineSafe() {
	s1 := New(1, 2, 3, 4, 5, 6, 7, 8)
	s2 := New(3, 4, 5, 6)
	s3 := New(1, 2, 5, 6, 8, 9)

	// 创建 10 个 goroutine 同步操作 s2，看会不会死锁
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			for i := 0; i < 100; i++ {
				s2.Union(s1) // 获取并集
				fmt.Printf("%2v：s2 + %v = %v\n", n, s1.SortedList(), s2.SortedList())

				s2.Minus(s3) // 获取差集
				fmt.Printf("%2v：s2 - %v = %v\n", n, s3.SortedList(), s2.SortedList())

				s2.Intersect(s1) // 获取交集
				fmt.Printf("%2v：s2 * %v = %v\n", n, s1.SortedList(), s2.SortedList())

				s2.Complement(s1) // 获取 s2 相对于 s1 的补集
				fmt.Printf("%2v：%v / s2 = %v\n", n, s1.SortedList(), s2.SortedList())
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
