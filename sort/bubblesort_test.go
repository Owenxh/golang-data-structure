package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	applySortTest(BubbleSort, t, randomIntSlices(10000, 100000, 200000))
}
