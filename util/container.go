package util

import "io.vava.datastructure/tree"

type Heap[T any] struct {
	Slice  []T
	IsLess func(a, b T) bool
}

func NewHeap[T any](less func(a, b T) bool) *Heap[T] {
	return &Heap[T]{IsLess: less}
}

func (h *Heap[T]) Len() int {
	return len(h.Slice)
}

func (h *Heap[T]) Less(i, j int) bool {
	return h.IsLess(h.Slice[i], h.Slice[j])
}

func (h *Heap[T]) Swap(i, j int) {
	h.Slice[i], h.Slice[j] = h.Slice[j], h.Slice[i]
}

func (h *Heap[T]) Push(x any) {
	h.Slice = append(h.Slice, x.(T))
}

func (h *Heap[T]) Pop() any {
	n := len(h.Slice)
	x := h.Slice[n-1]
	h.Slice = h.Slice[0 : n-1]
	return &x
}

type TreeMap interface {
	Size() int
	Put(k, v int)
	Get(k int) int
	Remove(v int)
	Contains(v int) bool
	Keys() []int
	EntrySet() []Entry
}

type Entry struct {
	K int
	V int
}

type treeMap struct {
	delegate *tree.AVLTree[int, int]
}

func (t *treeMap) Size() int {
	return t.delegate.Size()
}

func (t *treeMap) Put(k, v int) {
	t.delegate.Add(k, v)
}

func (t *treeMap) Get(k int) int {
	return t.delegate.Get(k)
}

func (t *treeMap) Remove(v int) {
	t.delegate.Remove(v)
}

func (t *treeMap) Contains(v int) bool {
	return t.delegate.Contains(v)
}

func (t *treeMap) Keys() []int {
	var keys []int
	t.delegate.InOrderFunc(func(k int, v int) {
		keys = append(keys, k)
	})
	return keys
}

func (t *treeMap) EntrySet() []Entry {
	var kvs []Entry
	t.delegate.InOrderFunc(func(k int, v int) {
		kvs = append(kvs, Entry{k, v})
	})
	return kvs
}

type TreeSet interface {
	Size() int
	Add(k int)
	Remove(v int)
	Contains(v int) bool
	Keys() []int
}

type treeSet struct {
	TreeMap
}

func (t *treeSet) Add(k int) {
	t.TreeMap.Put(k, 1)
}

func NewTreeMap() TreeMap {
	return &treeMap{
		delegate: &tree.AVLTree[int, int]{},
	}
}

func NewTreeMaps(c int) []TreeMap {
	maps := make([]TreeMap, c)
	for i := 0; i < c; i++ {
		maps[i] = NewTreeMap()
	}
	return maps
}

func NewTreeSet() TreeSet {
	return &treeSet{
		NewTreeMap(),
	}
}
