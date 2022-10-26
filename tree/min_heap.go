package tree

import (
	"errors"
	"io.vava.datastructure/array"
)

type MinHeap struct {
	arr *array.Array
}

func NewMinHeap() *MinHeap {
	return &MinHeap{arr: array.New()}
}

func (h *MinHeap) get(i int) int {
	return h.arr.Get(i)
}

func (h *MinHeap) set(i, e int) {
	h.arr.Set(i, e)
}

func (h *MinHeap) IsEmpty() bool {
	return h.arr.IsEmpty()
}

func (h *MinHeap) Size() int {
	return h.arr.Size()
}

func (h *MinHeap) swap(i, j int) {
	t := h.get(i)
	h.set(i, h.get(j))
	h.set(j, t)
}

func (h *MinHeap) FindMin() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("can't extract min value from empty array")
	}
	return h.arr.GetFirst(), nil
}

func (h *MinHeap) Add(e int) {
	h.arr.AddLast(e)
	h.siftUp(h.arr.Size() - 1)
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 && h.get(i) < h.get(parent(i)) {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MinHeap) siftDown(i int) {
	for leftChild(i) < h.arr.Size() {
		j := leftChild(i)

		if r := rightChild(i); r < h.Size() && h.get(r) < h.get(j) {
			j = r
		}

		if h.get(j) > h.get(i) {
			break
		}

		h.swap(i, j)
		i = j
	}
}

func (h *MinHeap) ExtractMin() (int, error) {
	min, err := h.FindMin()
	if err != nil {
		return 0, err
	}

	h.swap(0, h.Size()-1)
	h.arr.RemoveLast()
	h.siftDown(0)

	return min, nil
}
