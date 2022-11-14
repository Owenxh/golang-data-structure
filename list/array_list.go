package list

import (
	"fmt"
	"reflect"
)

type ArrayList[E any] struct {
	elements []E
	size     int
}

func (array *ArrayList[E]) String() string {
	return fmt.Sprintf("MyArray:{elements:%v, size:%d}", array.elements, array.size)
}

func New[E any]() *ArrayList[E] {
	return NewWithCapacity[E](10)
}

func NewWithCapacity[E any](capacity int) *ArrayList[E] {
	if capacity <= 0 {
		panic("MyArray capacity must > 0")
	}
	return &ArrayList[E]{
		elements: make([]E, capacity),
		size:     0,
	}
}

func (array *ArrayList[E]) resize(capacity int) {
	arr := make([]E, capacity)
	for i := 0; i < array.size; i++ {
		arr[i] = array.elements[i]
	}

	array.elements = arr
}

func (array *ArrayList[E]) AddLast(e E) {
	array.Add(array.size, e)
}

func (array *ArrayList[E]) AddFirst(e E) {
	array.Add(0, e)
}

func (array *ArrayList[E]) Add(index int, e E) {
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

func (array *ArrayList[E]) RemoveLast() E {
	return array.Remove(array.size - 1)
}

func (array *ArrayList[E]) RemoveFirst() E {
	return array.Remove(0)
}

func (array *ArrayList[E]) Remove(index int) E {
	if index < 0 || index >= array.size {
		panic("Remove index must >= 0 and < MyArray's size")
	}

	res := array.elements[index]

	for i := index + 1; i < array.size; i++ {
		array.elements[i-1] = array.elements[i]
	}
	var e E
	array.elements[array.size-1] = e
	array.size--

	if array.size > 0 && array.size <= len(array.elements)/4 {
		array.resize(len(array.elements) / 2)
	}

	return res
}

func (array *ArrayList[E]) Contains(e E) bool {
	return array.Find(e) >= 0
}

func (array *ArrayList[E]) Find(e E) int {
	for i, v := range array.elements {
		if reflect.DeepEqual(v, e) {
			return i
		}
	}

	return -1
}

func (array *ArrayList[E]) Set(index int, e E) {
	if index < 0 || index >= array.size {
		panic("Set index must >= 0 and < MyArray's size")
	}
	array.elements[index] = e
}

func (array *ArrayList[E]) GetFirst() E {
	return array.Get(0)
}

func (array *ArrayList[E]) GetLast() E {
	return array.Get(array.size - 1)
}

func (array *ArrayList[E]) Get(index int) E {
	if index < 0 || index >= array.size {
		panic("Find index must >= 0 and < MyArray's size")
	}

	return array.elements[index]
}

func (array *ArrayList[E]) GetAll() []E {
	res := make([]E, array.size)
	for i := 0; i < array.size; i++ {
		res[i] = array.elements[i]
	}
	return res
}

func (array *ArrayList[E]) Size() int {
	return array.size
}

func (array *ArrayList[E]) IsEmpty() bool {
	return array.size == 0
}
