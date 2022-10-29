package main

import "fmt"

type Node struct {
	l     int
	r     int
	left  *Node
	right *Node
	count int
}

func (node *Node) Print() {
	fmt.Printf("[l:%v, r:%v, count:%v]\n", node.l, node.r, node.count)
	if node.left != nil {
		node.left.Print()
	}
	if node.right != nil {
		node.right.Print()
	}
}

func calculateRange(nums []int) (int, int) {
	l, r := nums[0], nums[0]
	for _, v := range nums {
		if v < l {
			l = v
		}
		if v > r {
			r = v
		}
	}
	return l, r
}

func buildST(l, r int) *Node {
	root := &Node{
		l: l,
		r: r,
	}

	if l == r {
		return root
	}

	m := l + (r-l)/2

	root.left = buildST(l, m)
	root.right = buildST(m+1, r)

	return root
}

func (node *Node) Insert(insertL, insertR, v int) {
	if node.l == insertL && node.r == insertR {
		node.count += v
		return
	}

	m := node.l + (node.r-node.l)/2
	if m >= insertR {
		node.left.Insert(insertL, insertR, v)
	} else if m < insertL {
		node.right.Insert(insertL, insertR, v)
	} else {
		node.left.Insert(insertL, m, v)
		node.right.Insert(m+1, insertR, v)
	}
	node.count = node.left.count + node.right.count
}

func (node *Node) Count(queryL, queryR int) int {
	if queryL > queryR {
		return 0
	}

	if node.l == queryL && node.r == queryR {
		return node.count
	}

	m := node.l + (node.r-node.l)/2
	if m >= queryR {
		return node.left.Count(queryL, queryR)
	} else if m < queryL {
		return node.right.Count(queryL, queryR)
	} else {
		lRet := node.left.Count(queryL, m)
		rRet := node.right.Count(m+1, queryR)
		return lRet + rRet
	}
}

func countSmaller(nums []int) []int {
	l, r := calculateRange(nums)
	st := buildST(l, r)

	ret := make([]int, len(nums), len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = st.Count(l, nums[i]-1)
		st.Insert(nums[i], nums[i], 1)
	}

	return ret
}

func main() {
	nums := []int{2, 0, 1}
	fmt.Println(countSmaller(nums))

}
