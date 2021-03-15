package array

import "fmt"

type Array struct {
	elements []int
	size     int
}

func (array *Array) String() string {
	return fmt.Sprintf("MyArray:{elements:%v, size:%d}", array.elements, array.size)
}

func New() *Array {
	return NewWithCapacity(10)
}

func NewWithCapacity(capacity int) *Array {
	if capacity <= 0 {
		panic("MyArray capacity must > 0")
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
	array.Add(array.size, e)
}

func (array *Array) AddFirst(e int) {
	array.Add(0, e)
}

func (array *Array) Add(index, e int) {
	if index < 0 || index > array.size {
		panic("Add index mustn't less than 0 or large than MyArray's size")
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

func (array *Array) RemoveLast() int {
	return array.Remove(array.size - 1)
}

func (array *Array) RemoveFirst() int {
	return array.Remove(0)
}

func (array *Array) Remove(index int) int {
	if index < 0 || index >= array.size {
		panic("Remove index must >= 0 and < MyArray's size")
	}

	res := array.elements[index]

	for i := index + 1; i < array.size; i++ {
		array.elements[i-1] = array.elements[i]
	}
	array.elements[array.size-1] = 0
	array.size--

	if array.size > 0 && array.size <= len(array.elements)/4 {
		array.resize(len(array.elements) / 2)
	}

	return res
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

func (array *Array) Set(index, e int) {
	if index < 0 || index >= array.size {
		panic("Set index must >= 0 and < MyArray's size")
	}
	array.elements[index] = e
}

func (array *Array) GetFirst() int {
	return array.Get(0)
}

func (array *Array) GetLast() int {
	return array.Get(array.size - 1)
}

func (array *Array) Get(index int) int {
	if index < 0 || index >= array.size {
		panic("Find index must >= 0 and < MyArray's size")
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

func (array *Array) Size() int {
	return array.size
}

func (array *Array) IsEmpty() bool {
	return array.size == 0
}
