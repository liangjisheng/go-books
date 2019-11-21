package main

import "fmt"

func day000Test1() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println("a:", len(a), cap(a)) // 0 3
	fmt.Println("b:", len(b), cap(b)) // 2 3
	fmt.Println("c:", len(c), cap(c)) // 1 2
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func day000Test2() {
	fmt.Println(f1()) // 1
	fmt.Println(f2()) // 5
	fmt.Println(f3()) // 1
}

// Person ...
type Person struct {
	age int
}

func day000Test3() {
	person := &Person{28}

	// 1. person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中
	// 等到最后执行该 defer 语句的时候取出，即输出 28
	defer fmt.Println(person.age)

	// 2. defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29
	// 所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3. 闭包引用，输出 29
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

func day000Test4() {
	person := &Person{28}

	// 1. person.age 这一行代码跟之前含义是一样的，此时是将 28 当做 defer 函数的参数
	// 会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28
	defer fmt.Println(person.age)

	// 2. defer 缓存的是结构体 Person{28} 的地址，这个地址指向的结构体没有被改变
	// 最后 defer 语句后面的函数执行的时候取出仍是 28
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3. 闭包引用 person 的值已经被改变，指向结构体 Person{29} 所以输出 29
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}

func day000() {
	// day000Test1()
	// day000Test2()
	// day000Test3() // 29 29 28
	day000Test4() // 29 28 28
}
