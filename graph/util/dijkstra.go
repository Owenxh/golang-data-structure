package util

import (
	"io.vava.datastructure/graph"
	"sort"
)

// Dijkstra 算法实现求解有权图单源最小路径
type Dijkstra struct {
	graph.WeightedGraph
	S       int
	Dis     []int
	Visited []bool
}

type Node struct {
	V   int // V 顶点
	Dis int // Dis 距离
}

type Nodes []Node

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i].Dis < n[j].Dis
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type NodesHeap struct {
	Nodes
}

func (h *NodesHeap) Push(x any) {
	h.Nodes = append(h.Nodes, x.(Node))
}

func (h *NodesHeap) Pop() any {
	n := len(h.Nodes)
	x := h.Nodes[n-1]
	h.Nodes = h.Nodes[0 : n-1]
	return &x
}

type Heap[T any] struct {
	Slice  []T
	sorter sort.Interface
}

func (h *Heap[T]) Len() int {
	return h.sorter.Len()
}

func (h *Heap[T]) Less(i, j int) bool {
	return h.sorter.Less(i, j)
}

func (h *Heap[T]) Swap(i, j int) {
	h.sorter.Swap(i, j)
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
