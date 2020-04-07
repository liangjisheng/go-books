package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SkipListNode 节点
type SkipListNode struct {
	key  int         // key为排序使用的字段
	data interface{} // data为实际存储的数据
	next []*SkipListNode
}

// SkipList 跳表
type SkipList struct {
	head   *SkipListNode
	tail   *SkipListNode
	length int
	level  int
	mut    *sync.RWMutex
	rand   *rand.Rand
}

// 随机生成的层数要满足P=0.5的几何分布,大致可以理解为
// 掷硬币连续出现正面的次数,就是我们要的层数
func (list *SkipList) randomLevel() int {
	level := 1
	for ; level < list.level && list.rand.Uint32()&0x1 == 1; level++ {
	}
	return level
}

// NewSkipList 构造方法
func NewSkipList(level int) *SkipList {
	list := &SkipList{}
	if level <= 0 {
		level = 32
	}
	list.level = level
	list.head = &SkipListNode{next: make([]*SkipListNode, level, level)}
	list.tail = &SkipListNode{}
	list.mut = &sync.RWMutex{}
	list.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	for index := range list.head.next {
		list.head.next[index] = list.tail
	}
	return list
}

// Add insert
func (list *SkipList) Add(key int, data interface{}) {
	list.mut.Lock()
	defer list.mut.Unlock()

	level := list.randomLevel() // 确定插入的层数

	// 找到没一层需要插入的位置
	update := make([]*SkipListNode, level, level)
	node := list.head
	for index := level - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node.key > key {
				update[index] = node // 找到一个插入位置
				break
			} else if node.key == key {
				node1.data = data // 已经有这个数据,则更新实际数据
				return
			} else {
				node = node1
			}
		}
	}

	// 执行插入
	newNode := &SkipListNode{key, data, make([]*SkipListNode, level, level)}
	for index, node := range update {
		node.next[index], newNode.next[index] = newNode, node.next[index]
	}
	list.length++
}

// Remove ...
func (list *SkipList) Remove(key int) bool {
	list.mut.Lock()
	defer list.mut.Unlock()

	// 查找要删除的节点
	node := list.head
	remove := make([]*SkipListNode, list.level, list.level)
	var target *SkipListNode
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > key {
				break
			} else if node1.key == key {
				remove[index] = node
				target = node1
				break
			} else {
				node = node1
			}
		}
	}

	// 执行删除
	if target != nil {
		for index, node1 := range remove {
			if node1 != nil {
				node1.next[index] = target.next[index]
			}
		}
		list.length--
		return true
	}
	return false
}

// Find ...
func (list *SkipList) Find(key int) interface{} {
	list.mut.Lock()
	defer list.mut.Unlock()

	node := list.head
	for index := len(node.next) - 1; index >= 0; index-- {
		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > key {
				break
			} else if node1.key == key {
				return node1.data
			} else {
				node = node1
			}
		}
	}
	return nil
}

// Print ...
func (list *SkipList) Print() {
	node := list.head
	for index := len(node.next) - 1; index >= 0; index-- {
		node1 := node.next[index]
		for {
			if node1 != list.tail {
				fmt.Print(node1.key, " ")
				node1 = node1.next[index]
			} else {
				fmt.Println()
				break
			}
		}
	}
}

// Length 获取数据总量
func (list *SkipList) Length() int {
	list.mut.Lock()
	defer list.mut.Unlock()
	return list.length
}

func main() {
	skipList := NewSkipList(5)
	for i := 0; i < 20; i++ {
		skipList.Add(i, i)
	}
	skipList.Print()
}
