package main

// 返回区间 [2, n) 中有几个素数
// int countPrimes(int n)

// 比如 countPrimes(10) 返回 4
// 因为 2,3,5,7 是素数

// 首先从 2 开始, 我们知道 2 是一个素数, 那么 2 × 2 = 4, 3 × 2 = 6, 4 × 2 = 8...
// 都不可能是素数了
// 然后我们发现 3 也是素数, 那么 3 × 2 = 6, 3 × 3 = 9, 3 × 4 = 12也都不可能是素数了
// 看到这里, 你是否有点明白这个排除法的逻辑了呢? 先看第一版代码

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	// 初始化为 true
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			// i 的倍数不可能是素数了
			for j := i * 2; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 统计 [2-n) 中素数的个数
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

// 优化
func countPrimes1(n int) int {
	isPrime := make([]bool, n)
	// 初始化为 true
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}

	// 由于因子的对称性, 其中的 for 循环只需要遍历 [2,sqrt(n)] 就够了
	// 外层的 for 循环只需要遍历到 sqrt(n)
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			// i 的倍数不可能是素数了
			// 比如 i = 4 时算法会标记 4 × 2 = 8, 4 × 3 = 12 等等数字,
			// 但是 8 和 12 已经被 i = 2 和 i = 3 的 2 × 4 和 3 × 4 标记过了
			// 优化: 让j从i的平方开始遍历, 而不是从 2 * i 开始
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	// 统计 [2-n) 中素数的个数
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
}

// 优化后的算法有一个名字, 叫做 Sieve of Eratosthenes
