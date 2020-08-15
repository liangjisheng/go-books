package shuffle

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	fmt.Println("array:", array)
	Shuffle(array)
	fmt.Println("array:", array)
}

func TestShuffle1(t *testing.T) {
	fmt.Println("array:", array)
	Shuffle1(array)
	fmt.Println("array:", array)
}
