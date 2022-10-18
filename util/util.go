package util

import (
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func RandomIntSlice(length int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(1<<31 - 1)
	}
	return res
}

func RandomIntSliceWithBound(length, bound int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(bound)
	}
	return res
}

func ReversedIntSlice(length int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = length - i - 1
	}
	return res
}

func OrderedIntSlice(length, bound int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = i
	}
	return res
}

func RandomIntSlices(sizes ...int) [][]int {
	var data [][]int
	for _, size := range sizes {
		data = append(data, RandomIntSliceWithBound(size, size))
	}
	return data
}

func VerifyIsSorted(arr []int) {
	var prev int
	for i, v := range arr {
		if i > 0 && v < prev {
			panic("Not a sorted slice")
		}
		prev = v
	}
}

func DoSortTest(sort func([]int), t *testing.T, data []int) {
	DoSortTests(sort, t, [][]int{data})
}

func DoSortTests(sort func([]int), t *testing.T, data [][]int) {
	for _, arr := range data {
		sortFnName := sortFnName(sort)

		t.Logf("sort slice(length %9d) using %16v", len(arr), sortFnName)
		start := time.Now()
		sort(arr)
		cost := time.Now().Sub(start)
		
		t.Logf("sort slice(length %9d) using %16v cost %v", len(arr), sortFnName, cost)
		VerifyIsSorted(arr)
	}
}

func sortFnName(sort func([]int)) string {
	res := runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name()
	index := strings.LastIndex(res, ".") + 1
	res = string([]rune(res)[index:])
	return res
}

func Swap(arr []int, i, j int) {
	//t := arr[i]
	//arr[i] = arr[j]
	//arr[j] = t
	arr[i], arr[j] = arr[j], arr[i]
}

func Copy(src []int) []int {
	return CopyArrayFrom(src, 0, len(src)-1)
}

func CopyArrayFrom(src []int, l, r int) []int {
	ret := make([]int, r-l+1, r-l+1)
	CopyArray(src, l, r, ret)
	return ret
}

func CopyArray(src []int, l, r int, dest []int) {
	internalCopyArray(src, l, dest, l, r-l+1)
}

func internalCopyArray(src []int, srcPos int, dest []int, destPos, length int) {
	for i, j := srcPos, 0; i < srcPos+length; i++ {
		dest[destPos+j] = src[i]
		j++
	}
}
