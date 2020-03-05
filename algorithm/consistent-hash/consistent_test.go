package consistenthash

import (
	"fmt"
	"testing"
)

func TestNewConsistent(t *testing.T) {
	hash := NewConsistent()
	hash.Add("1000")
	hash.Add("2000")
	hash.Add("3000")
	fmt.Println(hash.hashSortNodes)
	fmt.Println(hash.GetNode())
	fmt.Println(hash.Get("1"))
	fmt.Println(hash.Get("2"))
	fmt.Println(hash.Get("3"))
	fmt.Println(hash.Get("4"))
	fmt.Println(hash.Get("5"))
	fmt.Println(hash.Get("6"))
	fmt.Println(hash.Get("7"))
	fmt.Println(hash.Get("8"))
	fmt.Println(hash.Get("9"))
	fmt.Println(hash.Get("10"))
	fmt.Println("remove", hash.Remove("3000"))
	fmt.Println(hash.GetNode())
	fmt.Println(hash.Get("1"))
	fmt.Println(hash.Get("2"))
	fmt.Println(hash.Get("3"))
	fmt.Println(hash.Get("4"))
	fmt.Println(hash.Get("5"))
	fmt.Println(hash.Get("6"))
	fmt.Println(hash.Get("7"))
	fmt.Println(hash.Get("8"))
	fmt.Println(hash.Get("9"))
	fmt.Println(hash.Get("10"))
}
