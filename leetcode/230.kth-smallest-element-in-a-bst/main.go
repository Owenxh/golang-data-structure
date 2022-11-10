// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	list := inOrder(root, []int{})
	return list[k-1]
}

func inOrder(root *TreeNode, list []int) []int {
	if root == nil {
		return list
	}
	list = inOrder(root.Left, list)
	return inOrder(root.Right, append(list, root.Val))
}

func main() {

}
