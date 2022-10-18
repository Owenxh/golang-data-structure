package sort

import (
	"io.vava.datastructure/util"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	util.DoSortTests(BubbleSort, t, util.RandomIntSlices(10000, 100000, 200000))
}
