package main

import (
	"io.vava.datastructure/helper"
	"io.vava.datastructure/sort/selectsort"
)

func main() {
	arr := []int{2, 4, 6, 5, 3, 1}
	selectsort.Sort(arr)
	helper.VerifyIsSorted(arr)

	arr2 := []int{2, 4, 6, 5, 3, 1}
	selectsort.Sort(arr2)
	helper.VerifyIsSorted(arr)
}
