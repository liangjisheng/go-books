package main

// struct Iface
// {
// 		Itab* tab;
// 		// 数据指针持有的是目标对象的只读复制品，复制完整对象或指针
// 		void* data;
// };
// struct Itab
// {
// 		InterfaceType* inter;
// 		Type* type;
// 		void (*fun[])(void);
// };

import "fmt"

func interfaceCopy() {
	u := user{1, "ljs"}
	var i interface{} = u

	u.id = 2
	u.name = "jack"

	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(user))
}

func interfaceModify() {
	u := user{1, "ljs"}
	var vi, pi interface{} = u, &u

	// 数据指针持有的是目标对象的只读复制品，复制完整对象或指针
	// vi.(user).name = "jack" // Error: cannot assign to vi.(User).name
	pi.(*user).name = "jack"

	fmt.Printf("%v\n", vi.(user))  // {1 ljs}
	fmt.Printf("%v\n", pi.(*user)) // user: 1, jack
	fmt.Printf("%v\n", vi.(user))  // {1 ljs}
}
