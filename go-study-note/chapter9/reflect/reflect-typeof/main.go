package main

// User ...
type User struct {
	Username string
}

// ToString ...
func (*User) ToString() {}

// Admin ...
type Admin struct {
	User
	title string
}

// Test ...
func (Admin) Test() {}
func (Admin) test() {}

func main() {
	// demo1()
	// demo2()
	// demo3()
	demo4()
}
