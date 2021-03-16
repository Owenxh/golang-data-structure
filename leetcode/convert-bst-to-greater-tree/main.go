// https://leetcode-cn.com/problems/convert-bst-to-greater-tree/

package main

var i int = 0
var sum int = 0

func convertBST(root *TreeNode) *TreeNode {
	sum, i = 0, 0
	m := make(map[int]int)
	inOrder(root, m)
	inOrderSetNewValue(root, m)
	return root
}

func inOrder(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, m)

	m[len(m)] = root.Val
	sum += root.Val

	inOrder(root.Right, m)
}

func inOrderSetNewValue(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}
	inOrderSetNewValue(root.Left, m)

	if i > 0 {
		sum = sum - m[i-1]
	}
	root.Val = sum
	i += 1

	inOrderSetNewValue(root.Right, m)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
