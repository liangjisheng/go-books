package main

func main() {
	s := NewStack(3)
	s.push(1)
	s.push(2)
	s.push(3)
	s.print()

	s.pop()
	s.print()
}
