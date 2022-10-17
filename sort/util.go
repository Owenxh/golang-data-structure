package sort

import (
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func randomIntSlice(length int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(1<<31 - 1)
	}
	return res
}

func randomIntSliceWithBound(length, bound int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = rand.Intn(bound)
	}
	return res
}

func reversedIntSlice(length int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = length - i - 1
	}
	return res
}

func orderedIntSlice(length, bound int) []int {
	res := make([]int, length, length)
	for i := 0; i < length; i++ {
		res[i] = i
	}
	return res
}

func randomIntSlices(sizes ...int) [][]int {
	var data [][]int
	for _, size := range sizes {
		data = append(data, randomIntSliceWithBound(size, size))
	}
	return data
}

func verifyIsSorted(arr []int) {
	var prev int
	for i, v := range arr {
		if i > 0 && v < prev {
			panic("Not a sorted slice")
		}
		prev = v
	}
}

func applySortTest(sort func([]int), t *testing.T, data [][]int) {
	for _, arr := range data {
		start := time.Now()
		sort(arr)
		cost := time.Now().Sub(start)
		sortFnName := sortFnName(sort)
		t.Logf("sort slice(length %9d) using %16v cost %v", len(arr), sortFnName, cost)
		verifyIsSorted(arr)
	}
}

func sortFnName(sort func([]int)) string {
	res := runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name()
	index := strings.LastIndex(res, ".") + 1
	res = string([]rune(res)[index:])
	return res
}

func swap(arr []int, i, j int) {
	//t := arr[i]
	//arr[i] = arr[j]
	//arr[j] = t
	arr[i], arr[j] = arr[j], arr[i]
}
