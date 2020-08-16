package recursion

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	head := GetList()
	PrintList(head)
	head = ReverseList(head)
	PrintList(head)
}
