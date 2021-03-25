package tree

// Binary search tree implementation.
type BST struct {
	// The root of the tree
	Root *BSTNode
	// The tree size
	Size int
}

// The node element of the BST.
type BSTNode struct {
	// The value
	Val int
	// The left child
	Left *BSTNode
	// The Right child
	Right *BSTNode
}

// Returns true if the tree contains element.
func (t *BST) Contains(e int) bool {
	// TODO

	return false
}

// Add the element to the tree.
func (t *BST) Add(e int) {
	// TODO
}

// Remove the element from the tree.
func (t *BST) Remove(e int) {
	// TODO
}

// Returns the if the tree is empty.
func (t *BST) IsEmpty(e int) bool {
	return t.Size == 0
}

// Returns min value in the tree.
func (t *BST) minimum() int {
	// TODO
	return 0
}

// Returns max value in the tree.
func (t *BST) maximum() int {
	// TODO
	return 0
}
