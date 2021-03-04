package main

import (
	"fmt"

	"io.vava.datastructure/array"
)

const N int = 10

func applyFn(fn func(int)) {
	for i := 0; i < N; i++ {
		fn(i)
	}
}

func main() {
	arr := array.NewArrayWithCapacity(4)
	applyFn(func(i int) {
		arr.AddLast(i)
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.DeleteFirst()
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.AddFirst(i)
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.DeleteLast()
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.Insert(arr.Size()/2, i)
		fmt.Println(arr)
	})

	applyFn(func(i int) {
		arr.Delete(arr.Size() / 2)
		fmt.Println(arr)
	})
}
