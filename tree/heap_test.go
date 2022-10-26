package tree

import (
	"math/rand"
	"testing"
)

func TestMaxHeap_ExtractMax(t *testing.T) {
	n, h := 500, NewMaxHeap()
	for i := 0; i < n; i++ {
		h.Add(rand.Intn(10000))
	}
	res := make([]int, n, n)
	for i := 0; i < n; i++ {
		if v, err := h.ExtractMax(); err == nil {
			res[i] = v
		} else {
			t.Fatal(err)
		}
	}

	for i := 1; i < n; i++ {
		if res[i-1] < res[i] {
			t.Fatal("Not a valid MaxHeap.")
		}
	}

	t.Log(res)
}

func TestMinHeap_ExtractMin(t *testing.T) {
	n, h := 500, NewMinHeap()
	for i := 0; i < n; i++ {
		h.Add(rand.Intn(10000))
	}
	res := make([]int, n, n)
	for i := 0; i < n; i++ {
		if v, err := h.ExtractMin(); err == nil {
			res[i] = v
		} else {
			t.Fatal(err)
		}
	}

	for i := 1; i < n; i++ {
		if res[i-1] > res[i] {
			t.Fatal("Not a valid MaxHeap.")
		}
	}

	t.Log(res)
}
