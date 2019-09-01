package main

import "fmt"

func typeDeduce() {
	var o interface{} = &user{1, "ljs"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}

	u := o.(*user)
	// u := o.(user)  // panic: interface is *main.User, not main.User
	fmt.Println(u)
}

func typeDeduce1() {
	var o interface{} = &user{1, "ljs"}

	switch v := o.(type) {
	case nil: // o == nil
		fmt.Println("nil")
	case fmt.Stringer: // interface
		fmt.Println("fmt.Stringer:", v)
	case func() string: // func
		fmt.Println("func:", v())
	case *user: // struct
		fmt.Printf("user: %d, %s\n", v.id, v.name)
	default:
		fmt.Println("unknown")
	}
}
