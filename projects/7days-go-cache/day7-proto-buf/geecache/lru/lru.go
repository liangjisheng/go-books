package lru

import (
	"container/list"
)

// Cache is a LRU cache, It is not safe for concurrent access.
type Cache struct {
	// maxBytes 是允许使用的最大内存
	maxBytes int64
	// nbytes 是当前已使用的内存
	nbytes int64
	ll     *list.List
	// 值是双向链表中对应节点的指针
	cache map[string]*list.Element
	// OnEvicted 是某条记录被移除时的回调函数 可以为 nil
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

// 键值对 entry 是双向链表节点的数据类型 在链表中仍保存每个值对应的 key
// 的好处在于淘汰队首节点时，需要用 key 从字典中删除对应的映射
type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes 用于返回值所占用的内存大小
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value 找到节点并移动到队首,队首附近的缓存就是热点缓存
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item 从队尾移除最近最少访问的节点
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}
