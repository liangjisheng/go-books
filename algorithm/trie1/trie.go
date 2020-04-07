package trie1

import (
	"fmt"
)

// Trie ...
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

// NewTrie ...
func NewTrie() *Trie {
	return &Trie{}
}

// Insert ...
func (trie *Trie) Insert(word string) {
	node := trie
	for _, c := range word {
		if node.next[c-'a'] == nil {
			node.next[c-'a'] = NewTrie()
		}
		node = node.next[c-'a']
	}
	node.isEnd = true
}

// Search 精确查找
func (trie *Trie) Search(word string) bool {
	node := trie
	for _, c := range word {
		node = node.next[c-'a']
		if node == nil {
			return false
		}
	}
	return node.isEnd
}

// PrefixMatch 前缀匹配
func (trie *Trie) PrefixMatch(prefix string) bool {
	node := trie
	for _, c := range prefix {
		node = node.next[c-'a']
		if node == nil {
			return false
		}
	}
	return true
}

// 返回与前缀prefix匹配的最后一个结点的地址
func (trie *Trie) _prefix(prefix string) *Trie {
	node := trie
	for _, c := range prefix {
		node = node.next[c-'a']
		if node == nil {
			return nil
		}
	}
	return node
}

// _getWords 获取以root为起始结点的所有单词
func _getWords(root *Trie, word string) []string {
	if nil == root {
		return nil
	}
	allWords := make([]string, 0, 0)
	if root.isEnd {
		allWords = append(allWords, word)
	}
	for i := 0; i < 26; i++ {
		word = word + fmt.Sprintf("%s", string([]byte{byte('a' + i)}))
		allWords = append(allWords, _getWords(root.next[i], word)...)
		word = word[:len(word)-1]
	}
	return allWords
}

// 获取以prefix为前缀的所有单词
func (trie *Trie) getAllWordsOfPrefix(prefix string) []string {
	node := trie._prefix(prefix)
	words := _getWords(node, prefix)
	return words
}

// 获取Trie中所有单词
func (trie *Trie) getAllWords() []string {
	return trie.getAllWordsOfPrefix("")
}

func _deleteWord(node **Trie, word string, d int) {
	if d == len(word) {
		(*node).isEnd = false
	} else {
		_deleteWord(&((*node).next[word[d]-'a']), word, d+1)
	}

	if (*node).isEnd {
		return
	}
	for _, v := range (*node).next {
		if v != nil {
			return
		}
	}
	*node = nil
}

// DeleteWrod ...
func (trie *Trie) DeleteWrod(word string) {
	if !trie.Search(word) {
		return
	}
	_deleteWord(&trie, word, 0)
}
