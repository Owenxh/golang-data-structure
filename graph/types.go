package graph

import (
	"fmt"

	"io.vava.datastructure/tree"
)

type Cloneable[E any] interface {
	Clone() E
}

type Graph interface {
	// V the vertex count
	V() int
	// E the edge count
	E() int
	// Adj adjacency of the vertex v
	Adj(v int) []int
	// AddEdge connect vertex v & w
	AddEdge(v, w int)
	// RemoveEdge remove edge between vertex v & w
	RemoveEdge(v, w int)
	// ValidateVertex validate whether v is valid vertex
	ValidateVertex(v int)
	// Degree returns the degree of vertex v
	Degree(v int) int
	Cloneable[Graph]
}

type Edge struct {
	V int
	W int
}

func (e Edge) String() string {
	return fmt.Sprintf("%d-%d", e.V, e.W)
}

type WeightedGraph interface {
	Graph
	GetWeight(v, w int) int
}

type WeightedEdge struct {
	V      int
	W      int
	Weight int
}

func (e WeightedEdge) String() string {
	return fmt.Sprintf("(%d-%d: %d)", e.V, e.W, e.Weight)
}

type SortedWeightEdges []WeightedEdge

func (s SortedWeightEdges) Len() int {
	return len(s)
}

func (s SortedWeightEdges) Less(i, j int) bool {
	return s[i].Weight-s[j].Weight < 0
}

func (s SortedWeightEdges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type WeightEdgesHeap struct {
	SortedWeightEdges
}

func (s *WeightEdgesHeap) Push(x any) {
	s.SortedWeightEdges = append(s.SortedWeightEdges, x.(WeightedEdge))
}

func (s *WeightEdgesHeap) Pop() any {
	n := len(s.SortedWeightEdges)
	x := s.SortedWeightEdges[n-1]
	s.SortedWeightEdges = s.SortedWeightEdges[0 : n-1]
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
