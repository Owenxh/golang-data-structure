// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/solutions/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	return buildTree(list, 0, len(list)-1)

}

func buildTree(list []int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{
			Val: list[l],
		}
	}
	mid := l + (r-l)/2
	root := &TreeNode{
		Val: list[mid],
	}
	root.Left = buildTree(list, l, mid-1)
	root.Right = buildTree(list, mid+1, r)
	return root
}

func main() {

}
