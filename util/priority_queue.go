package util

import (
	"container/heap"
)

type PriorityQueue[E interface{}] interface {
	Size() int
	IsEmpty() bool
	Push(E)
	Pop() E
}

type priorityQueue[E interface{}] struct {
	delegate *internalHeap[E]
}

func (pq *priorityQueue[E]) Size() int {
	return pq.delegate.Len()
}

func (pq *priorityQueue[E]) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *priorityQueue[E]) Push(e E) {
	heap.Push(pq.delegate, e)
}

func (pq *priorityQueue[E]) Pop() E {
	return heap.Pop(pq.delegate).(E)
}

func NewPriorityQueue[T interface{}](less func(a, b T) bool) PriorityQueue[T] {
	return &priorityQueue[T]{delegate: &internalHeap[T]{IsLess: less}}
}

func InitPriorityQueue[T interface{}](slice []T, less func(a, b T) bool) PriorityQueue[T] {
	delegate := &internalHeap[T]{Slice: slice, IsLess: less}
	heap.Init(delegate)
	return &priorityQueue[T]{delegate: delegate}
}

type internalHeap[T any] struct {
	Slice  []T
	IsLess func(a, b T) bool
}

func (h *internalHeap[T]) Len() int {
	return len(h.Slice)
}

func (h *internalHeap[T]) Less(i, j int) bool {
	return h.IsLess(h.Slice[i], h.Slice[j])
}

func (h *internalHeap[T]) Swap(i, j int) {
	h.Slice[i], h.Slice[j] = h.Slice[j], h.Slice[i]
}

func (h *internalHeap[T]) Push(x any) {
	h.Slice = append(h.Slice, x.(T))
}

func (h *internalHeap[T]) Pop() any {
	n := len(h.Slice)
	x := h.Slice[n-1]
	h.Slice = h.Slice[0 : n-1]
	return x
}
