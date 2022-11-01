package tree

import (
	"fmt"
)

// 只保留最低位的 1
func lowBit(x int) int {
	return x & (-x)
}

type BinaryIndexedTree struct {
	data []int
	tree []int
	size int
}

func NewBIT(nums []int) *BinaryIndexedTree {
	sz := len(nums)
	data := make([]int, sz)
	copy(data, nums)

	tree := make([]int, sz+1)

	ret := &BinaryIndexedTree{
		data: data,
		tree: tree,
		size: sz,
	}

	for i, v := range data {
		ret.update0(i, v)
	}

	//fmt.Println(ret.tree)

	return ret
}

func (b *BinaryIndexedTree) update0(i int, v int) {
	for j := i + 1; j <= b.size; j += lowBit(j) {
		b.tree[j] += v
		//log.Printf("Add %v to tree index %v", v, j)
	}
}

func (b *BinaryIndexedTree) Update(i int, v int) {
	delta := v - b.data[i]
	b.data[i] = v
	b.update0(i, delta)
}

func (b *BinaryIndexedTree) findSum(i int) int {
	sum := 0
	for i > 0 {
		sum += b.tree[i]
		i -= lowBit(i)
		//log.Printf("Sum tree index %v", i)
	}
	return sum
}

func (b *BinaryIndexedTree) SumRange(l int, r int) int {
	if l < 0 || r >= b.size || l > r {
		panic(fmt.Sprintf("invalid range [%d, %d]", l, r))
	}

	// [0, r] - [0, l-1]
	return b.findSum(r+1) - b.findSum(l)
}
