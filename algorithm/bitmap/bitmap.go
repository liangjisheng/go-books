package main

// Bitmap ...
type Bitmap struct {
	elems []byte
	size  uint64 // 存多少位bits
}

// NewBitMap ...
func NewBitMap(size uint64) *Bitmap {
	return &Bitmap{
		elems: make([]byte, size/8+1, size/8+1),
		size:  size,
	}
}

// Add 将第n位设置为1
func (b *Bitmap) Add(n uint64) {
	index := n >> 3      // 等价于 n/8
	position := n & 0x07 // 等价于 n%8
	b.elems[index] |= 1 << position
}

// Delete 将第n位设置为0
func (b *Bitmap) Delete(n uint64) {
	index := n >> 3
	position := n & 0x07
	b.elems[index] &= ^(1 << position)
}

// Contain 判断第n位是否为1
func (b *Bitmap) Contain(n uint64) bool {
	index := n >> 3
	position := n & 0x07
	return (b.elems[index] & (1 << position)) != 0
}
