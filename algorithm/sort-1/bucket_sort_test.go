package sort

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	fmt.Println("array:", array)
	BucketSort(array)
	fmt.Println("array:", array)
}
