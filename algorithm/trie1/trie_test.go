package trie1

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))    // 返回 true
	fmt.Println(trie.Search("app"))      // 返回 false
	fmt.Println(trie.PrefixMatch("app")) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // 返回 true
	trie.Insert("abc")

	fmt.Println(trie.getAllWords())
	trie.DeleteWrod("abc")
	fmt.Println(trie.getAllWords())
	trie.DeleteWrod("app")
	fmt.Println(trie.getAllWords())
}
