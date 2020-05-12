package tree

import (
	"fmt"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := Constructor()
	trie.Insert("abp")
	trie.Insert("app")
	// fmt.Println(trie.Search("apple"))
	// fmt.Println(trie.Search("app"))
	// fmt.Println(trie.Search("app"))
	// fmt.Println(trie.StartsWith("app"))
	// fmt.Println(trie.StartsWith("app"))
	trie.Insert("abe")
	trie.Insert("aef")
	trie.Insert("ae")
	fmt.Println(trie.SearchNode("ae"))
	fmt.Println(trie.SearchNode("ae"))
	fmt.Println(trie.SearchNode("ab"))
	fmt.Println(trie.SearchNode("aa"))
	fmt.Println(trie.SearchNode("ap"))
}
