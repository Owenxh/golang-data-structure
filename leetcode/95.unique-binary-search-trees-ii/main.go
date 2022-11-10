package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTrees0(1, n)
}

func generateTrees0(l int, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	var res []*TreeNode
	for i := l; i <= r; i++ {
		leftTrees := generateTrees0(l, i-1)
		rightTrees := generateTrees0(i+1, r)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				root := &TreeNode{
					Val:   i,
					Left:  leftTree,
					Right: rightTree,
				}
				res = append(res, root)
			}
		}
	}
	return res
}

func preOrder(node *TreeNode, fn func(*TreeNode)) {
	if node == nil {
		return
	}
	fn(node)
	preOrder(node.Left, fn)
	preOrder(node.Right, fn)
}

func main() {
	for _, t := range generateTrees(3) {
		preOrder(t, func(node *TreeNode) {
			fmt.Print(node.Val, " ")
		})
		fmt.Println()
	}
}
