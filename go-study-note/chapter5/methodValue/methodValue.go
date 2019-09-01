package main

import "fmt"

// 根据调⽤者不同，⽅法分为两种表现形式
// instance.method(args...) ---> <type>.func(instance, args...)
// 前者称为 method value，后者 method expression
// 两者都可像普通函数那样赋值和传参，区别在于 method value 绑定实例，⽽ method
// expression 则须显式传参

type user struct {
	id   int
	name string
}

func (u *user) test() {
	fmt.Printf("%p, %v\n", u, u)
}

func main() {
	u := user{1, "ljs"}
	fmt.Printf("u: %p\n", &u)

	u.test()

	mValue := u.test // 隐式传递 receiver
	mValue()

	mExpression := (*user).test // 显式传递 receiver
	mExpression(&u)
}
