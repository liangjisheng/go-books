package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"
)

// 一共 m * n 个数, 表示为 m * n 的矩阵, 其中存储的元素值为 [0, m*n-1]
// 要求随机选 [i, j] 翻过来查看是否是 0, 如果是 0, 则将这个数字翻转过来
// 接着随机选 [i, j], 寻找数字 1, 默认最开始的时候所有的数字都是看不见的
// 一直找到 m*n-1, 要求打印出翻转的过程
// 这个想法来自于 chainup 团建时候的翻 16 个公司文化字

// Point ...
type Point struct {
	x, y int
}

var (
	m = 4
	n = 4
)

var (
	originMatrix  = make([][]int, m) // 原始数据
	knowMatrix    = make([][]int, m) // 表示已经知道的数据
	reverseMatrix = make([][]int, m) // 表示现在已经翻过来的数据
)

var array = make([]int, m*n)
var position = make(map[int]Point) // 记录某个数所在的索引
var resPath = list.New()           // 记录翻过的路径

func init() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < m*n; i++ {
		array[i] = i
	}

	for i := 0; i < m; i++ {
		originMatrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			index := rand.Intn(m * n)
			for array[index] == -1 {
				index = rand.Intn(m * n)
			}
			originMatrix[i][j] = array[index]
			array[index] = -1
		}
	}

	fmt.Println("origin matrix:")
	printMatrix(originMatrix)

	for i := 0; i < m; i++ {
		knowMatrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			knowMatrix[i][j] = -1
		}
	}

	for i := 0; i < m; i++ {
		reverseMatrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			reverseMatrix[i][j] = -1
		}
	}
}

func printMatrix(mat [][]int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%3d", mat[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// 找第一个为0的数字, 从{0,0}处开始寻找, 一直到{m-1, n-1}
	currentNum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			knowMatrix[i][j] = originMatrix[i][j]
			position[originMatrix[i][j]] = Point{i, j}
			resPath.PushBack(Point{i, j})
			if originMatrix[i][j] == currentNum {
				reverseMatrix[i][j] = originMatrix[i][j]
				currentNum++
				// fmt.Println("reverse matrix:")
				// printMatrix(reverseMatrix)
			}
		}
	}

	// fmt.Println("number position:")
	// for i := 0; i < m*n; i++ {
	// 	fmt.Printf("num: %d, position: %v\n", i, position[i])
	// }

	for i := currentNum; i < m*n; i++ {
		resPath.PushBack(position[i])
		pos := position[i]
		reverseMatrix[pos.x][pos.y] = i
		// fmt.Println("reverse matrix:")
		// printMatrix(reverseMatrix)
	}

	fmt.Println("reverse path:")
	pathLen := resPath.Len()
	for i := 0; i < pathLen; i++ {
		res := resPath.Front()
		fmt.Printf("times: %d, position: %v\n", i+1, res.Value.(Point))
		resPath.Remove(res)
	}

	// fmt.Println("know matrix:")
	// printMatrix(knowMatrix)

	fmt.Println("reverse matrix:")
	printMatrix(reverseMatrix)
}
