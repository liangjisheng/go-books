package main

func main() {
	list := NewList()
	for i := 0; i < 10; i++ {
		list.Insert(i)
	}
	list.Print()

	// node := list.Find(2)
	// fmt.Println("find:", node.key)

	// list.Delete(0)
	// list.Print()

	list.Delete(9)
	list.Print()

	// list.Delete(11)
	// list.Print()
}
