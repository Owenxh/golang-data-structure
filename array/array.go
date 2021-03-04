package array

import "fmt"

type Array struct {
	elements []int
	size     int
}

func (array *Array) String() string {
	return fmt.Sprintf("Array{elements: %v, size:%d}", array.elements, array.size)
}

func NewArray() *Array {
	return NewArrayWithCapacity(20)
}

func NewArrayWithCapacity(capacity int) *Array {
	if capacity <= 0 {
		panic("Array capacity must > 0")
	}

	return &Array{make([]int, capacity, capacity), 0}
}

func (array *Array) resize(capacity int) {
	arr := make([]int, capacity, capacity)
	for i := 0; i < array.size; i++ {
		arr[i] = array.elements[i]
	}

	array.elements = arr
}

func (array *Array) AddLast(e int) {
	array.Insert(array.size, e)
}

func (array *Array) AddFirst(e int) {
	array.Insert(0, e)
}

func (array *Array) Insert(index, e int) {
	if index < 0 || index > array.size {
		panic("Insert index mustn't less than 0 or large than Array's size")
	}

	if array.size >= len(array.elements) {
		array.resize(len(array.elements) * 2)
	}

	for i := array.size - 1; i >= index; i-- {
		array.elements[i+1] = array.elements[i]
	}

	array.elements[index] = e
	array.size++
}

func (array *Array) DeleteLast() {
	array.Delete(array.size - 1)
}

func (array *Array) DeleteFirst() {
	array.Delete(0)
}

func (array *Array) Delete(index int) {
	if index < 0 || index >= array.size {
		panic("Delete index must >= 0 and < Array's size")
	}

	for i := index + 1; i < array.size; i++ {
		array.elements[i-1] = array.elements[i]
	}
	array.elements[array.size-1] = 0
	array.size--

	if array.size > 0 && array.size <= len(array.elements)/4 {
		array.resize(len(array.elements) / 2)
	}
}

func (array *Array) Contains(e int) bool {
	return array.Find(e) >= 0
}

func (array *Array) Find(e int) int {
	for i, v := range array.elements {
		if v == e {
			return i
		}
	}

	return -1
}

func (array *Array) Size() int {
	return array.size
}

func (array *Array) Get(index int) int {
	if index < 0 || index >= array.size {
		panic("Find index must >= 0 and < Array's size")
	}

	return array.elements[index]
}

func (array *Array) GetAll() []int {
	res := make([]int, array.size, array.size)
	for i := 0; i < array.size; i++ {
		res[i] = array.elements[i]
	}
	return res
}
