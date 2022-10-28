package array

import (
	"reflect"
	"testing"
)

const N int = 10

func generateSequentialArray(size int) *Array[int] {
	if size <= 0 {
		panic("MyArray size must > 0")
	}

	arr := NewWithCapacity[int](size)
	for i := 0; i < size; i++ {
		arr.AddLast(i)
	}
	return arr
}

func TestArray_Size(t *testing.T) {
	size := 21
	arr := generateSequentialArray(size)

	if arr.Size() != size {
		t.Fatal("Size method test failed.")
	}
}

func TestArray_Contains(t *testing.T) {
	n := 10
	arr := generateSequentialArray(n)

	for i := 0; i < n; i++ {
		if !arr.Contains(i) {
			t.Fatal("Get method test failed.")
		}
	}
}

func TestArray_Find(t *testing.T) {
	n := 10
	arr := generateSequentialArray(n)

	for i := 0; i < n; i++ {
		if arr.Find(i) != i {
			t.Fatal("Find method test failed.")
		}
	}
}

func TestArray_Get(t *testing.T) {
	n := 10
	arr := generateSequentialArray(n)

	for i := 0; i < n; i++ {
		if arr.Get(i) != i {
			t.Fatal("Get method test failed.")
		}
	}
}

func TestArray_GetAll(t *testing.T) {
	arr := generateSequentialArray(10)

	if !reflect.DeepEqual(arr.GetAll(), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Fatal("GetAll method test failed.")
	}
}
