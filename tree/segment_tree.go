package tree

import "fmt"

type SegmentTree[E any] struct {
	tree   []E
	data   []E
	merger func(E, E) E
}

type Merger[E any] func(l, r E) E

func NewSegmentTree[E any](data []E, merger func(E, E) E) *SegmentTree[E] {
	st := SegmentTree[E]{}
	copy(st.data, data)
	st.merger = merger
	st.tree = make([]E, 4*len(data), 4*len(data))
	st.buildSegmentTree(0, 0, len(st.data)-1)
	return &st
}

func (st *SegmentTree[E]) buildSegmentTree(treeIndex int, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}

	leftTreeIndex, rightTreeIndex := st.leftChild(treeIndex), st.rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)
	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree[E]) Get(index int) E {
	if index < 0 || index >= len(st.data) {
		panic("index is illegal")
	}
	return st.data[index]
}

// 给定指定节点下标，计算其左子树节点下标
func (st *SegmentTree[E]) leftChild(index int) int {
	return index*2 + 1
}

// 给定指定节点下标，计算其右子树节点下标
func (st *SegmentTree[E]) rightChild(index int) int {
	return index*2 + 2
}

// 数组下标以 0 开始作为根结点的二叉树左、右、父节点下标计算方式
//
// 给定指定节点下标，计算其父节点下标
// func (st *SegmentTree[E]) parent(index int) int {
// 	return (index - 1) / 2
// }

// 查询区间范围 [queryL, queryR] 的值
func (st *SegmentTree[E]) Query(queryL, queryR int) E {
	if queryL > queryR || queryL < 0 || queryR >= len(st.data) {
		panic(fmt.Sprintf("illegal query range [%v, %v]", queryL, queryR))
	}
	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// 在以 treeIndex 为根节点的线段树中 [l,r] 的范围里，搜索区间 [queryL, queryR] 的值
func (st *SegmentTree[E]) query(treeIndex, l, r, queryL, queryR int) E {
	// 区间 [queryL, queryR] 正好相等
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}

	m := l + (r-l)/2

	leftChildIndex := st.leftChild(treeIndex)
	rightChildIndex := st.rightChild(treeIndex)

	// 区间 [queryL, queryR] 全位于左子树上
	if queryR <= m {
		return st.query(leftChildIndex, l, m, queryL, queryR)
	}

	// 区间 [queryL, queryR] 全位于右子树上
	if queryL > m {
		return st.query(rightChildIndex, m+1, r, queryL, queryR)
	}

	// 区间 [queryL, queryR] 部分们于左子树，部分位于右子树

	leftRet := st.query(leftChildIndex, l, m, queryL, m)
	rightRet := st.query(rightChildIndex, m+1, r, m+1, queryR)

	return st.merger(leftRet, rightRet)
}
