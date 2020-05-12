package tree

// Trie ...
type Trie struct {
	node   map[rune]*Trie
	val    string
	isLast bool
}

// Constructor ...
func Constructor() Trie {
	return Trie{node: make(map[rune]*Trie, 26), isLast: false}
}

// Insert ...
func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.node[v] == nil { // 没找到该节点
			vt := &Trie{node: make(map[rune]*Trie, 26), isLast: false}
			t.node[v] = vt
		}
		t = t.node[v] // 找到节点，跳到下一个节点
	}
	t.val = word
	t.isLast = true
}

// Search 返回 t.isLast 表示精确查找
func (t *Trie) Search(word string) bool {
	for _, v := range word {
		if t.node[v] == nil { // 没找到该节点
			return false
		}
		t = t.node[v] // 找到节点，跳到下一个节点
	}
	return t.isLast
}

// StartsWith 只要前缀匹配就返回 true
func (t *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if t.node[v] == nil { // 没找到该节点
			return false
		}
		t = t.node[v] // 找到节点，跳到下一个节点
	}
	return true
}

// SearchNode 查找以key开始的所有数据
func (t *Trie) SearchNode(key string) (res []interface{}) {
	root := t
	for _, v := range key {
		if v, ok := root.node[v]; ok {
			// root.node = v.node
			root = v
		} else {
			return
		}
	}
	var queue []*Trie
	queue = append(queue, root)
	for len(queue) > 0 {
		var childQueue []*Trie
		for _, v := range queue { // 遍历这层的数据
			if v.isLast {
				res = append(res, v.val)
			}
			for _, vi := range v.node { // 把下一次的数据进行遍历
				childQueue = append(childQueue, vi)
			}
		}
		queue = childQueue
	}
	// fmt.Println(len(res))
	return
}
