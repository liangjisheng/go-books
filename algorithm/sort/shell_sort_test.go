package sort

import (
	"fmt"
	"testing"
)

func TestShellSort(t *testing.T) {
	fmt.Println("array:", array)
	ShellSort(array)
	fmt.Println("array:", array)
}
