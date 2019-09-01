package main

import "fmt"

// 匿名接⼝可⽤作变量类型，或结构成员
type tester struct {
	s interface {
		String() string
	}
}

type user struct {
	id   int
	name string
}

func (u *user) String() string {
	return fmt.Sprintf("user: %d, %s", u.id, u.name)
}

func anonymousInterface() {
	t := tester{&user{1, "ljs"}}
	fmt.Println(t.s.String())
}
