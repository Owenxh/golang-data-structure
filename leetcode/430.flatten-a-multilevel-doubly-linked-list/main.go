// link https://leetcode-cn.com/problems/flatten-a-multilevel-doubly-linked-list/
package main

// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

//func flatten(root *Node) *Node {
//	if root == nil {
//		return root
//	}
//	var res, prev *Node
//	var tmp []*Node
//	curr := root
//	for curr != nil {
//		var next *Node
//		if curr.Child != nil {
//			tmp = append(tmp, curr.Next)
//			next = curr.Child
//		} else {
//			next = curr.Next
//			for next == nil && len(tmp) > 0 {
//				next = tmp[len(tmp)-1]
//				tmp = tmp[0 : len(tmp)-1]
//			}
//		}
//		if prev != nil {
//			prev.Next = curr
//		}
//		if res == nil {
//			res = curr
//		}
//		curr.Prev, curr.Next, curr.Child = prev, nil, nil
//		prev, curr = curr, next
//	}
//	return res
//}

// 给定一个链表，返回其 head, tail 结点
func flatten(root *Node) *Node {
	flatten0(root, nil)
	return root
}

func flatten0(node, prev *Node) *Node {
	if node == nil {
		return prev
	}
	child, next := node.Child, node.Next
	node.Prev, node.Next, node.Child = prev, nil, nil
	if prev != nil {
		prev.Next = node
	}
	prev = flatten0(child, node)
	return flatten0(next, prev)
}

func main() {

}
