package main

import "fmt"

type NumArray struct {
	tree   []int
	data   []int
	merger func(int, int) int
}

func Constructor(nums []int) NumArray {
	st := NumArray{
		data: make([]int, len(nums), len(nums)),
		tree: make([]int, 4*len(nums), 4*len(nums)),
		merger: func(i int, i2 int) int {
			return i + i2
		},
	}
	copy(st.data, nums)
	st.tree = make([]int, 4*len(nums), 4*len(nums))
	st.buildNumArray(0, 0, len(st.data)-1)
	return st
}

func (st *NumArray) buildNumArray(treeIndex int, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}

	leftTreeIndex, rightTreeIndex := st.leftChild(treeIndex), st.rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildNumArray(leftTreeIndex, l, mid)
	st.buildNumArray(rightTreeIndex, mid+1, r)
	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *NumArray) leftChild(index int) int {
	return index*2 + 1
}

func (st *NumArray) rightChild(index int) int {
	return index*2 + 2
}

func (st *NumArray) SumRange(left int, right int) int {
	return st.Query(left, right)
}

func (st *NumArray) Query(queryL, queryR int) int {
	if queryL > queryR || queryL < 0 || queryR >= len(st.data) {
		panic(fmt.Sprintf("illegal query range [%v, %v]", queryL, queryR))
	}
	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

func (st *NumArray) query(treeIndex, l, r, queryL, queryR int) int {
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
