package tree

import "fmt"

type UnionFind2 struct {
	parent []int
}

func NewUnionFind2(size int) *UnionFind2 {
	parent := make([]int, size, size)

	// 初始化时让每个节点的根结点指向自己
	for i := 0; i < size; i++ {
		parent[i] = i
	}

	return &UnionFind2{
		parent: parent,
	}
}

func (u *UnionFind2) GetSize() int {
	return len(u.parent)
}

func (u *UnionFind2) IsConnected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UnionFind2) UnionElements(p int, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot != qRoot {
		u.parent[pRoot] = qRoot
	}
}

// 查找节点 p 所在树的根结点
func (u *UnionFind2) find(p int) int {
	if p < 0 || p >= len(u.parent) {
		panic(fmt.Sprintf("illegal parameter q:%v", p))
	}

	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

type UnionFind3 struct {
	parent []int
	sz     []int
}

func NewUnionFind3(size int) *UnionFind3 {
	parent := make([]int, size, size)
	sz := make([]int, size, size)

	// 初始化时让每个节点的根结点指向自己
	for i := 0; i < size; i++ {
		parent[i] = i
		sz[i] = 1
	}

	return &UnionFind3{
		parent: parent,
		sz:     sz,
	}
}

func (u *UnionFind3) GetSize() int {
	return len(u.parent)
}

func (u *UnionFind3) IsConnected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UnionFind3) UnionElements(p int, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}

	// 将节点数少的树合并到节点数多的树上
	if u.sz[pRoot] < u.sz[qRoot] {
		u.parent[pRoot] = qRoot
		u.sz[qRoot] += u.sz[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.sz[pRoot] += u.sz[qRoot]
	}
}

// 查找节点 p 所在树的根结点
func (u *UnionFind3) find(p int) int {
	if p < 0 || p >= len(u.parent) {
		panic(fmt.Sprintf("illegal parameter q:%v", p))
	}

	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

type UnionFind4 struct {
	parent []int
	rank   []int
}

func NewUnionFind4(size int) *UnionFind4 {
	parent := make([]int, size, size)
	rank := make([]int, size, size)

	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &UnionFind4{
		parent: parent,
		rank:   rank,
	}
}

func (u *UnionFind4) GetSize() int {
	return len(u.parent)
}

func (u *UnionFind4) IsConnected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UnionFind4) UnionElements(p int, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}

	// 将 rank 数低的集合合并到 rank 较大的集合上
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		u.parent[qRoot] = pRoot
		u.rank[pRoot] += 1
	}
}

func (u *UnionFind4) find(p int) int {
	if p < 0 || p >= len(u.parent) {
		panic(fmt.Sprintf("illegal parameter q:%v", p))
	}

	for p != u.parent[p] {
		p = u.parent[p]
	}
	return p
}

// Path Compression
type UnionFind5 struct {
	parent []int
	rank   []int
}

func NewUnionFind5(size int) *UnionFind5 {
	parent := make([]int, size, size)
	rank := make([]int, size, size)

	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &UnionFind5{
		parent: parent,
		rank:   rank,
	}
}

func (u *UnionFind5) GetSize() int {
	return len(u.parent)
}

func (u *UnionFind5) IsConnected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UnionFind5) UnionElements(p int, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}

	// 将 rank 数低的集合合并到 rank 较大的集合上
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		u.parent[qRoot] = pRoot
		u.rank[pRoot] += 1
	}
}

// 查找过程，查找元素 p 对应的集合编号
// O(h) 复杂度，h 为树的高度
func (u *UnionFind5) find(p int) int {
	if p < 0 || p >= len(u.parent) {
		panic(fmt.Sprintf("illegal parameter q:%v", p))
	}

	for p != u.parent[p] {
		u.parent[p] = u.parent[u.parent[p]]
		p = u.parent[p]
	}
	return p
}

// Path Compression
type UnionFind6 struct {
	parent []int
	rank   []int
}

func NewUnionFind6(size int) *UnionFind6 {
	parent := make([]int, size, size)
	rank := make([]int, size, size)

	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &UnionFind6{
		parent: parent,
		rank:   rank,
	}
}

func (u *UnionFind6) GetSize() int {
	return len(u.parent)
}

func (u *UnionFind6) IsConnected(p int, q int) bool {
	return u.find(p) == u.find(q)
}

func (u *UnionFind6) UnionElements(p int, q int) {
	pRoot := u.find(p)
	qRoot := u.find(q)

	if pRoot == qRoot {
		return
	}

	// 将 rank 数低的集合合并到 rank 较大的集合上
	if u.rank[pRoot] < u.rank[qRoot] {
		u.parent[pRoot] = qRoot
	} else if u.rank[pRoot] > u.rank[qRoot] {
		u.parent[qRoot] = pRoot
	} else {
		u.parent[qRoot] = pRoot
		u.rank[pRoot] += 1
	}
}

// 查找过程，查找元素 p 对应的集合编号
// O(h) 复杂度，h 为树的高度
func (u *UnionFind6) find(p int) int {
	if p < 0 || p >= len(u.parent) {
		panic(fmt.Sprintf("illegal parameter q:%v", p))
	}

	if p != u.parent[p] {
		u.parent[p] = u.find(u.parent[p])
	}
	return u.parent[p]
}
