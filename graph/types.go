package graph

import (
	"io.vava.datastructure/tree"
)

type Graph interface {
	V() int
	E() int
	Adj(v int) []int
	AddEdge(v, w int)
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
