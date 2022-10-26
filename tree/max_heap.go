package tree

import (
	"errors"

	"io.vava.datastructure/array"
)

type MaxHeap struct {
	arr *array.Array
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{arr: array.New()}
}

func (h *MaxHeap) get(i int) int {
	return h.arr.Get(i)
}

func (h *MaxHeap) set(i, e int) {
	h.arr.Set(i, e)
}

func (h *MaxHeap) IsEmpty() bool {
	return h.arr.IsEmpty()
}

func (h *MaxHeap) Size() int {
	return h.arr.Size()
}

func (h *MaxHeap) ExtractMax() (int, error) {
	if h.arr.IsEmpty() {
		return 0, errors.New("can't extract max value from empty h")
	}

	res, err := h.FindMax()
	if err != nil {
		return res, err
	}

	h.swap(0, h.Size()-1)
	h.arr.RemoveLast()
	h.siftDown(0)

	return res, nil
}

func (h *MaxHeap) FindMax() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("can't extract max value from empty h")
	}

	return h.arr.GetFirst(), nil
}

func (h *MaxHeap) Add(e int) {
	h.arr.AddLast(e)
	h.siftUp(h.arr.Size() - 1)
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父节点的索引
func parent(i int) int {
	return (i - 1) / 2
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(i int) int {
	return i*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(i int) int {
	return i*2 + 2
}

func (h *MaxHeap) siftUp(i int) {
	for i > 0 && h.get(parent(i)) < h.get(i) {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

func (h *MaxHeap) siftDown(i int) {
	for leftChild(i) < h.arr.Size() {
		j := leftChild(i)
		if j+1 < h.Size() && h.get(j+1) > h.get(j) {
			j += 1
		}
		if h.get(i) > h.get(j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h *MaxHeap) swap(i, j int) {
	t := h.get(i)
	h.set(i, h.get(j))
	h.set(j, t)
}
