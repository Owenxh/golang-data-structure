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
	val string
	// The left child
	left *BSTNode
	// The right child
	right *BSTNode
}

// Returns true if the tree contains element.
func (t *BST) Contains(e string) bool {
	return t.contains(t.root, e)
}

func (t *BST) contains(node *BSTNode, e string) bool {
	if node == nil {
		return false
	}
	if node.val == e {
		return true
	} else if node.val > e {
		return t.contains(node.left, e)
	}
	return t.contains(node.right, e)
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
	if node.val == e {
		return node
	} else if node.val > e {
		node.left = t.add(node.left, e)
		return node
	} else {
		node.right = t.add(node.right, e)
		return node
	}
}

// Remove the element from the tree.
func (t *BST) Remove(e string) {
	t.root = t.remove(t.root, e)
}

func (t *BST) remove(node *BSTNode, e string) *BSTNode {
	if node == nil {
		return node
	}
	if node.val > e {
		node.left = t.remove(node.left, e)
		return node.left
	} else if node.val < e {
		node.right = t.remove(node.right, e)
		return node.right
	} else {
		if node.left == nil && node.right == nil {
			t.size--
			return nil
		}
		// TODO
		return nil
	}
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
	for min.left != nil {
		min = min.left
	}
	return min.val
}

// Returns max value in the tree.
func (t *BST) maximum() string {
	max := t.root
	for max.right != nil {
		max = max.right
	}
	return max.val
}
