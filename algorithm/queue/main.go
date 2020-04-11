package main

import (
	"fmt"
	"log"
)

func main() {
	q := NewArrayDeque(3)
	q.enqueue(0)
	q.enqueue(1)
	q.enqueue(2)
	// q.enqueue(3)
	q.print()

	val, err := q.dequeue()
	if err != nil {
		log.Println("dequeue err:", err)
		return
	}
	fmt.Println(val)
	q.print()

	val, err = q.dequeue()
	if err != nil {
		log.Println("dequeue err:", err)
		return
	}
	fmt.Println(val)

	val, err = q.dequeue()
	if err != nil {
		log.Println("dequeue err:", err)
		return
	}
	fmt.Println(val)
}
