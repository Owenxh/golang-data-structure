// link: https://leetcode-cn.com/problems/binary-search-tree-iterator

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	array []int
}

func (it *BSTIterator) inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	it.inOrder(root.Left)
	it.array = append(it.array, root.Val)
	it.inOrder(root.Right)
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.inOrder(root)
	return
}

func (it *BSTIterator) Next() int {
	if len(it.array) == 0 {
		return 0
	}
	ret := it.array[0]
	it.array = it.array[1:]
	return ret
}

func (it *BSTIterator) HasNext() bool {
	return len(it.array) > 0
}

type BSTIterator2 struct {
	stack []*TreeNode
	curr  *TreeNode
}

func Constructor2(root *TreeNode) BSTIterator2 {
	return BSTIterator2{curr: root}
}

func (it *BSTIterator2) Next() int {
	for c := it.curr; c != nil; c = c.Left {
		it.stack = append(it.stack, c)
	}
	it.curr, it.stack = it.stack[len(it.stack)-1], it.stack[:len(it.stack)-1]
	ret := it.curr.Val
	it.curr = it.curr.Right
	return ret
}

func (it *BSTIterator2) HasNext() bool {
	return it.curr != nil || len(it.stack) > 0
}

func main() {

}
