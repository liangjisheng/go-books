package main

import "fmt"

type user struct {
	id   int
	name string
}

type manager struct {
	user
}

func (u *user) toString() string {
	return fmt.Sprintf("user: %p, %v", u, u)
}

// 通过匿名字段，可获得和继承类似的复⽤能⼒。依据编译器查找次序，只需在外层定义同
// 名⽅法，就可以实现 "override"

func anonymousField() {
	m := manager{user{1, "ljs"}}
	fmt.Printf("manager: %p\n", &m)
	fmt.Println(m.toString())
}
