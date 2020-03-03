package tree

import (
	"fmt"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := Constructor()
	trie.Insert("abp")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("abe")
	trie.Insert("aef")
	fmt.Println(trie.SearchNode("ae"))
}
