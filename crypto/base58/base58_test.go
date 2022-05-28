package base58

import (
	"testing"
)

func TestBase58(t *testing.T) {
	v1 := "test"
	vb := Encode([]byte(v1))
	t.Log(string(vb))

	sb := Decode(vb)
	t.Log(string(sb))

	//github.com/mr-tron/base58
	//t.Log(base58.Encode([]byte("test"))) //3yZe7d
}
