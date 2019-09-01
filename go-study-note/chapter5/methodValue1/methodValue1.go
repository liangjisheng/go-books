package main

import "fmt"

// 需要注意， method value 会复制 receiver
// 在汇编层⾯， method value 和闭包的实现⽅式相同，实际返回 FuncVal 类型对象。
// FuncVal { method_address, receiver_copy }

type user struct {
	id   int
	name string
}

func (u user) test() {
	fmt.Println(u)
}

func main() {
	u := user{1, "Tom"}
	mValue := u.test // ⽴即复制 receiver，因为不是指针类型，不受后续修改影响

	u.id, u.name = 2, "Jack"
	u.test() // {2 Jack}

	mValue() // {1 Tom}
}
