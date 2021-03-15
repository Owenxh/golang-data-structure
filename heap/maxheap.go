package heap

import (
	"errors"

	"io.vava.datastructure/array"
)

type MaxHeap struct {
	array *array.Array
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{array: array.New()}
}

func (h MaxHeap) IsEmpty() bool {
	return h.array.IsEmpty()
}

func (h MaxHeap) Size() int {
	return h.array.Size()
}

func (h MaxHeap) ExtractMax() (int, error) {
	if h.array.IsEmpty() {
		return 0, errors.New("can't extract max value from empty h")
	}

	res, err := h.FindMax()
	if err != nil {
		return res, err
	}

	h.swap(0, h.Size()-1)
	h.array.RemoveLast()
	h.siftDown(0)

	return res, nil
}

func (h MaxHeap) FindMax() (int, error) {
	if h.array.IsEmpty() {
		return 0, errors.New("can't extract max value from empty h")
	}

	return h.array.GetFirst(), nil
}

func (h MaxHeap) Add(e int) {
	h.array.AddLast(e)
	h.siftUp(h.array.Size() - 1)
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

func (h MaxHeap) siftUp(i int) {
	for i > 0 && h.array.Get(parent(i)) < h.array.Get(i) {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

func (h MaxHeap) siftDown(i int) {
	for leftChild(i) < h.array.Size() {
		j := leftChild(i)
		if j+1 < h.array.Size() && h.array.Get(j+1) > h.array.Get(j) {
			j += 1
		}
		if h.array.Get(i) > h.array.Get(j) {
			break
		}
		h.swap(i, j)
		i = j
	}
}

func (h MaxHeap) swap(i, j int) {
	t := h.array.Get(i)
	h.array.Set(i, h.array.Get(j))
	h.array.Set(j, t)
}
