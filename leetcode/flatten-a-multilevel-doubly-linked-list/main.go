package main

// Definition for a Node.
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	var tmp []*Node

	curr := root
	var prev *Node
	for curr != nil {
		next := &Node{Val: curr.Val}
		if prev != nil {
			prev.Next = next
		}
		next.Prev = prev
		prev = next

		if curr.Child != nil {
			tmp = append(tmp, next)
			curr = curr.Child
		} else {
			curr = curr.Next
		}
		if curr == nil && len(tmp) > 0 {
			// TODO
		}
	}
	return nil
}

func main() {

}
