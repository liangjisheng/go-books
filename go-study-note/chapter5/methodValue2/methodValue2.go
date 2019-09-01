package main

import "fmt"

type user struct {
	id   int
	name string
}

func (u *user) testPointer() {
	fmt.Printf("%p, %v\n", u, u)
}

func (u user) testValue() {
	fmt.Printf("%p, %v\n", &u, u)
}

func main() {
	u := user{1, "ljs"}
	fmt.Printf("user: %p, %v\n", &u, u)

	mv := user.testValue
	mv(u)

	mp := (*user).testPointer
	mp(&u)

	mp2 := (*user).testValue // *User ⽅法集包含 TestValue
	mp2(&u)                  // 签名变为 func TestValue(self *User)
	// 实际依然是 receiver value copy
}
