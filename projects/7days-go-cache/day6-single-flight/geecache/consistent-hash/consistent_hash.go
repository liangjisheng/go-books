package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map constains all hashed keys
type Map struct {
	// hash 函数可以自定义, 默认为crc32.ChecksumIEEE
	hash Hash
	// 虚拟节点倍数
	replicas int
	// hash环
	keys []int // Sorted
	// 虚拟节点与真实节点的映射表 hashMap, 键是虚拟节点的哈希值, 值是真实节点的名称
	hashMap map[int]string
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}

	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key.
// 获取给定 key 所在的节点
// 顺时针找到第一个匹配的虚拟节点的下标 idx 从 m.keys 中获取到对应的哈希值
// 如果 idx == len(m.keys) 说明应选择 m.keys[0] 因为 m.keys 是一个环状结构
// 所以用取余数的方式来处理这种情况
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	// Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})

	// 这里idx对 len(m.keys) 求余是因为有可能key在m.keys中没有找到
	// 此时需要将key缓存在第一台节点上,而sort.Search()未找到时会返回
	// n 也就是 len(m.keys), 求余后正好放在第一台
	// 通过 hashMap 映射得到真实的节点
	return m.hashMap[m.keys[idx%(len(m.keys))]]
}
