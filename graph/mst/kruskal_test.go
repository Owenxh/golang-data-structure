package mst

import (
	"fmt"
	"io.vava.datastructure/graph"
	"testing"
)

func TestKruskal_Result(t *testing.T) {
	gText := `7 12
			0 1 2
			0 3 7
			0 5 2
			1 2 1
			1 3 4
			1 4 3
			1 5 5
			2 4 4
			2 5 4
			3 4 1
			3 6 5
			4 6 7`
	g := graph.StringAsWeightedGraph(gText)
	fmt.Println(g)
	k := NewKruskal(g)
	fmt.Println("Minimum Tree Spanning:", k.Result())
}
