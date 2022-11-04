package sort

import (
	"testing"

	"io.vava.datastructure/util"
)

func TestBubbleSort(t *testing.T) {
	util.TestSort(t, BubbleSort, util.RandomIntSlice(10000))
}
