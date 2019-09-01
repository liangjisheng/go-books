package main

func length(s string) int {
	println("call length.")
	return len(s)
}

func for1() {
	s := "abcd"

	for i, n := 0, length(s); i < n; i++ {
		println(i, s[i])
	}
}
