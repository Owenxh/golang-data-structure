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
		return node
	} else if node.val < e {
		node.right = t.remove(node.right, e)
		return node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			t.size--
			return right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			t.size--
			return left
		} else {
			successor := t.minNode(node.right)
			//successor.right = t.remove(node.right, successor.val)
			successor.right, successor.left = t.removeMin(node.right), node.left
			node.left, node.right = nil, nil
			return successor
		}
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
	if t.root == nil {
		panic("tree is empty")
	}
	return t.minNode(t.root).val
}

func (t *BST) minNode(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	min := node
	for min.left != nil {
		min = min.left
	}
	return min
}

// Returns max value in the tree.
func (t *BST) maximum() string {
	if t.root == nil {
		panic("tree is empty")
	}
	return t.maxNode(t.root).val
}

func (t *BST) maxNode(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	max := node
	for max.right != nil {
		max = max.right
	}
	return max
}

func (t *BST) RemoveMin() (val string) {
	val = t.minimum()
	t.root = t.removeMin(t.root)
	return
}

func (t *BST) removeMin(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	if node.left == nil {
		r := node.right
		node.right = nil
		t.size--
		return r
	}
	node.left = t.removeMin(node.left)
	return node
}

func (t *BST) RemoveMax() (val string) {
	val = t.maximum()
	t.root = t.removeMax(t.root)
	return
}

func (t *BST) removeMax(node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	if node.right == nil {
		l := node.left
		node.left = nil
		t.size--
		return l
	}
	node.right = t.removeMax(node.right)
	return node
}
