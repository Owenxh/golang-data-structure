package graph

import (
	"fmt"
	"io.vava.datastructure/tree"
)

type Graph interface {
	// V the vertex count
	V() int
	// E the edge count
	E() int
	// Adj adjacency of the vertex v
	Adj(v int) []int
	// AddEdge connect vertex v & w
	AddEdge(v, w int)
	// ValidateVertex validate whether v is valid vertex
	ValidateVertex(v int)
	// Degree returns the degree of vertex v
	Degree(v int) int
}

type Edge struct {
	V, W int
}

func (e Edge) String() string {
	return fmt.Sprintf("%d-%d", e.V, e.W)
}

type TreeSet interface {
	Size() int
	Keys() []int
	Put(v int)
	Remove(v int)
	Contains(v int) bool
}

type treeSet struct {
	delegate *tree.AVLTree[int, bool]
}

func (t *treeSet) Size() int {
	return t.delegate.Size()
}

func (t *treeSet) Keys() []int {
	var keys []int
	t.delegate.PreOrderFunc(func(k int, v bool) {
		keys = append(keys, k)
	})
	return keys
}

func (t *treeSet) Put(v int) {
	t.delegate.Add(v, true)
}

func (t *treeSet) Remove(v int) {
	t.delegate.Remove(v)
}

func (t *treeSet) Contains(v int) bool {
	return t.delegate.Contains(v)
}

func NewTreeSet() TreeSet {
	return &treeSet{
		delegate: &tree.AVLTree[int, bool]{},
	}
}

func NewTreeSets(c int) []TreeSet {
	sets := make([]TreeSet, c)
	for i := 0; i < c; i++ {
		sets[i] = NewTreeSet()
	}
	return sets
}
