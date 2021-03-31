package tree

// Binary search tree implementation.
type BST struct {
	// The root of the tree
	root *BSTNode
	// The tree size
	size int
}

// The node element of the BST.
type BSTNode struct {
	// The value
	Val string
	// The left child
	Left *BSTNode
	// The Right child
	Right *BSTNode
}

// Returns true if the tree contains element.
func (t *BST) Contains(e string) bool {
	return t.contains(t.root, e)
}

func (t *BST) contains(node *BSTNode, e string) bool {
	if node == nil {
		return false
	}
	if node.Val == e {
		return true
	} else if node.Val > e {
		return t.contains(node.Left, e)
	}
	return t.contains(node.Right, e)
}

// Add the element to the tree.
func (t *BST) Add(e string) {
	t.root = t.add(t.root, e)
}

func (t *BST) add(node *BSTNode, e string) *BSTNode {
	if node == nil {
		t.size++
		return &BSTNode{e, nil, nil}
	}
	if node.Val == e {
		return node
	} else if node.Val > e {
		node.Left = t.add(node.Left, e)
		return node
	} else {
		node.Right = t.add(node.Right, e)
		return node
	}
}

// Remove the element from the tree.
func (t *BST) Remove(e string) {
	// TODO
}

// Returns the if the tree is empty.
func (t *BST) IsEmpty(e string) bool {
	return t.size == 0
}

func (t *BST) Size() int {
	return t.size
}

// Returns min value in the tree.
func (t *BST) minimum() string {
	min := t.root
	for min.Left != nil {
		min = min.Left
	}
	return min.Val
}

// Returns max value in the tree.
func (t *BST) maximum() string {
	max := t.root
	for max.Right != nil {
		max = max.Right
	}
	return max.Val
}
