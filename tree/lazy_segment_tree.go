package tree

import (
	"fmt"
	"math"
)

type LazySegmentTree struct {
	tree   []int
	lazy   []int
	size   int
	merger func(int, int) int
}

func (st *LazySegmentTree) String() string {
	return fmt.Sprint(st.tree)
}

// Java method
// Height of segment tree
// int x = (int) (Math.ceil(Math.log(n) / Math.log(2)));
// Maximum size of segment tree
// int max_size = 2 * (int) Math.pow(2, x) - 1;

// NewLazySegmentTree creates a lazy updatable segment tree
func NewLazySegmentTree(data []int) *LazySegmentTree {
	h := math.Ceil(math.Log2(float64(len(data))))
	sz := 2*int(math.Pow(2, h)) - 1

	st := LazySegmentTree{
		tree: make([]int, sz, sz),
		lazy: make([]int, sz, sz),
		size: len(data),
		merger: func(a int, b int) int {
			return a + b
		},
	}
	st.buildLazySegmentTree(0, data, 0, len(data)-1)
	return &st
}

func (st *LazySegmentTree) buildLazySegmentTree(treeIndex int, data []int, l, r int) {
	if l == r {
		st.tree[treeIndex] = data[l]
		return
	}

	leftTree, rightTree := st.leftChild(treeIndex), st.rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildLazySegmentTree(leftTree, data, l, mid)
	st.buildLazySegmentTree(rightTree, data, mid+1, r)

	st.tree[treeIndex] = st.merger(st.tree[leftTree], st.tree[rightTree])
}

// 给定指定节点下标，计算其左子树节点下标
func (st *LazySegmentTree) leftChild(index int) int {
	return index*2 + 1
}

// 给定指定节点下标，计算其右子树节点下标
func (st *LazySegmentTree) rightChild(index int) int {
	return index*2 + 2
}

// Query 查询区间范围 [queryL, queryR] 的值
func (st *LazySegmentTree) Query(queryL, queryR int) int {
	if queryL > queryR || queryL < 0 || queryR >= st.size {
		panic(fmt.Sprintf("illegal query range [%v, %v]", queryL, queryR))
	}
	return st.query(0, 0, st.size-1, queryL, queryR)
}

func (st *LazySegmentTree) query(root, l, r, queryL, queryR int) int {
	st.lazyUpdate(root, l, r)

	// 区间 [queryL, queryR] 正好相等
	if l == queryL && r == queryR {
		return st.tree[root]
	}

	m := l + (r-l)/2

	left := st.leftChild(root)
	right := st.rightChild(root)

	if queryR <= m {
		return st.query(left, l, m, queryL, queryR)
	}

	if queryL > m {
		return st.query(right, m+1, r, queryL, queryR)
	}

	leftRet := st.query(left, l, m, queryL, m)
	rightRet := st.query(right, m+1, r, m+1, queryR)

	return st.merger(leftRet, rightRet)
}

// UpdateRange update segment tree from range [updateL, updateR], every element need to add diff
func (st *LazySegmentTree) UpdateRange(updateL, updateR int, diff int) {
	if updateL < 0 || updateR >= st.size || updateR < updateL {
		panic(fmt.Sprintf("update range [%v, %v] is illegal", updateL, updateR))
	}
	st.updateRange(0, 0, st.size-1, updateL, updateR, diff)
}

func (st *LazySegmentTree) lazyUpdate(root int, l int, r int) {
	if st.lazy[root] != 0 {
		fmt.Println(fmt.Sprintf("==== update range [%v, %v] = %v", l, r, st.lazy[root]))
		// update self
		st.tree[root] += (r - l + 1) * st.lazy[root]

		// not a leaf node
		if l != r {
			st.lazy[st.leftChild(root)] += st.lazy[root]
			st.lazy[st.rightChild(root)] += st.lazy[root]
		}

		st.lazy[root] = 0
	}
}

func (st *LazySegmentTree) updateRange(root int, l int, r int, updateL int, updateR int, diff int) {
	st.lazyUpdate(root, l, r)

	// current segment is fully in range
	if l == updateL && r == updateR {
		// update self
		st.tree[root] += (r - l + 1) * diff

		// not a leaf node
		if l != r {
			st.lazy[st.leftChild(root)] += diff
			st.lazy[st.rightChild(root)] += diff
		}
		return
	}

	left := st.leftChild(root)
	right := st.rightChild(root)

	mid := l + (r-l)/2

	if mid >= updateR {
		st.updateRange(left, l, mid, updateL, updateR, diff)
	} else if mid+1 <= updateL {
		st.updateRange(right, mid+1, r, updateL, updateR, diff)
	} else {
		st.updateRange(left, l, mid, updateL, mid, diff)
		st.updateRange(right, mid+1, r, mid+1, updateR, diff)
	}
	st.tree[root] = st.merger(st.tree[left], st.tree[right])
}

//func (st *LazySegmentTree) updateRange2(root int, l int, r int, updateL int, updateR int, diff int) {
//	st.lazyUpdate(root, l, r)
//
//	// out of range
//	if l > r || l > updateR || r < updateL {
//		return
//	}
//
//	// current segment is fully in range
//	if l >= updateL && r <= updateR {
//		// update self
//		st.tree[root] += (r - l + 1) * diff
//
//		// not a leaf node
//		if l != r {
//			leftTreeIndex, rightTreeIndex := st.leftChild(root), st.rightChild(root)
//			st.lazy[leftTreeIndex] = diff
//			st.lazy[rightTreeIndex] = diff
//		}
//		return
//	}
//
//	leftChild := st.leftChild(root)
//	rightChild := st.rightChild(root)
//
//	mid := l + (r-l)/2
//
//	st.updateRange(leftChild, l, mid, updateL, updateR, diff)
//	st.updateRange(rightChild, mid+1, r, updateL, updateR, diff)
//
//	st.tree[root] = st.merger(st.tree[leftChild], st.tree[rightChild])
//}
