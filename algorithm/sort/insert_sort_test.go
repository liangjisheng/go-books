package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	fmt.Println("array:", array)
	InsertSort(array)
	fmt.Println("array:", array)
}
