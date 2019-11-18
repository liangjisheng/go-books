package main

import (
	"fmt"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

// Iter ...
func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{}) // 解除注释看看！
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()

		for elem, value := range set.s {
			ch <- elem
			println("Iter:", elem, value)
		}

		close(ch)
		set.RUnlock()
	}()
	return ch
}

func itertest() {
	th := threadSafeSet{
		s: []interface{}{"1", "2"},
	}

	v := <-th.Iter()
	fmt.Printf("%s %v\n", "ch", v)
}
