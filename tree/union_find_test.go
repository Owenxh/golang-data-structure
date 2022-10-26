package tree

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestUnionFind(t *testing.T) {
	count := 10_000_000

	// uf2 := NewUnionFind2(count)
	// testUF(t, uf2, count)

	uf3 := NewUnionFind3(count)

	testUF(t, uf3, count)

	uf4 := NewUnionFind4(count)
	testUF(t, uf4, count)

	uf5 := NewUnionFind5(count)
	testUF(t, uf5, count)

	uf6 := NewUnionFind6(count)
	testUF(t, uf6, count)
}

func testUF(t *testing.T, uf UF, m int) {
	size := uf.GetSize()

	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.UnionElements(a, b)
	}

	for i := 0; i < m; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.IsConnected(a, b)
	}

	fmt.Printf("[%v] cost time:%v\n", reflect.TypeOf(uf), time.Now().Sub(start))
}
