package recursion

import (
	"fmt"
	"math/rand"
)

// 用递归方式反转链表

// Node ...
type Node struct {
	data int
	next *Node
}

// ReverseList ...
func ReverseList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	newHead := ReverseList(head.next)
	tmp := head.next
	tmp.next = head
	head.next = nil
	return newHead
}

// GetList ...
func GetList() *Node {
	head := &Node{
		data: rand.Intn(20) + 1,
		next: nil,
	}
	last := head
	for i := 0; i < 4; i++ {
		last.next = &Node{
			data: rand.Intn(20) + 1,
			next: nil,
		}
		last = last.next
	}
	return head
}

// PrintList ...
func PrintList(head *Node) {
	last := head
	for last != nil {
		fmt.Printf("%d ", last.data)
		last = last.next
	}
	fmt.Println()
}
