package main

func add(x, y int) (z int) {
	defer func() {
		println(z)
	}()

	z = x + y
	return z + 200 // 执行顺序: z = z + 200 -> (call defer) -> (ret)
}
