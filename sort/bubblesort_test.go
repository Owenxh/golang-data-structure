package sort

import (
	"testing"

	"io.vava.datastructure/util"
)

func TestBubbleSort(t *testing.T) {
	util.DoSortTests(BubbleSort, t, [][]int{util.RandomIntSlice(10000)})
}
