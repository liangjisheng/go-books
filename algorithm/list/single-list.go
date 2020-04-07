package main

import (
	"fmt"
)

// ListNode ...
type ListNode struct {
	key  int
	next *ListNode
}

// List ...
type List struct {
	head *ListNode
}

// NewList ...
func NewList() *List {
	return &List{}
}

// Insert ...
func (list *List) Insert(key int) *List {
	node := &ListNode{
		key:  key,
		next: nil,
	}
	if list.head == nil {
		list.head = node
		return list
	}

	tailNode := list.GetTailNode()
	tailNode.next = node
	return list
}

// GetTailNode ...
func (list *List) GetTailNode() *ListNode {
	node := list.head
	for node != nil {
		if node.next != nil {
			node = node.next
		} else {
			return node
		}
	}
	return node
}

// Find ...
func (list *List) Find(key int) *ListNode {
	node := list.head
	for node != nil {
		if node.key == key {
			return node
		}
		node = node.next
	}
	return node
}

// Print ...
func (list *List) Print() {
	node := list.head
	for node != nil {
		fmt.Printf("%d ", node.key)
		node = node.next
	}
	fmt.Println()
}

// Delete ...
func (list *List) Delete(key int) *List {
	if list.head == nil {
		return list
	}
	if list.head.key == key {
		list.head = list.head.next
		return list
	}

	preNode := list.head
	node := list.head.next
	for node != nil {
		if node.key == key {
			preNode.next = node.next
			return list
		}
		preNode = node
		node = node.next
	}
	return list
}
