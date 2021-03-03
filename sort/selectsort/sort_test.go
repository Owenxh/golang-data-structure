package selectsort

import (
	"testing"
	"time"

	"io.vava.datastructure/sort/helper"
)

func testData() [][]int {
	arrays := [][]int{
		helper.GenerateRandomIntSlice(1000, 1000),
		helper.GenerateRandomIntSlice(10000, 10000),
		helper.GenerateRandomIntSlice(100000, 100000),
	}
	return arrays
}

func internalTestSort(sort func([]int), t *testing.T) {
	arrays := testData()

	for _, arr := range arrays {
		start := time.Now()
		sort(arr)
		t.Logf("sort slice(length %d) cost %v", len(arr), time.Now().Sub(start))
		helper.VerifyIsSorted(arr)
	}
}

//func TestSort(t *testing.T) {
//	arr := []int{2, 4, 6, 5, 3, 1}
//	Sort(arr)
//
//	if !reflect.DeepEqual(arr, []int{1, 2, 3, 4, 5, 6}) {
//		t.Fatalf("Sort failed")
//	}
//}

func TestSort(t *testing.T) {
	internalTestSort(Sort, t)
}

func TestSort2(t *testing.T) {
	internalTestSort(Sort2, t)
}

func BenchmarkSort(b *testing.B) {
	arr := helper.GenerateRandomIntSlice(10000, 10000)
	b.ResetTimer()
	Sort(arr)
}
