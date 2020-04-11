package main

import (
	"errors"
	"fmt"
	"log"
)

// ArrayQueue ...
type ArrayQueue struct {
	elems      []int
	n          int
	head, tail int
}

// NewArrayDeque ...
func NewArrayDeque(n int) *ArrayQueue {
	return &ArrayQueue{
		elems: make([]int, n, n),
		n:     n,
	}
}

func (q *ArrayQueue) enqueue(elem int) *ArrayQueue {
	if q.tail == q.n {
		log.Println("queue is full.")
		return q // full
	}

	q.elems[q.tail] = elem
	q.tail++
	return q
}

func (q *ArrayQueue) dequeue() (int, error) {
	if q.head == q.tail {
		return -1, errors.New("queue is empty")
	}

	val := q.elems[q.head]
	q.head++
	return val, nil
}

func (q *ArrayQueue) print() {
	for i := q.head; i < q.tail; i++ {
		fmt.Printf("%d ", q.elems[i])
	}
	fmt.Println()
}
