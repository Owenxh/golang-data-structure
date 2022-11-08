package graph

import (
	"fmt"
	"io.vava.datastructure/util"
	"testing"
)

func TestAdjSet(t *testing.T) {
	graph := NewGraph(util.GetFileAbsolutePath("/graph/data/g.txt"))
	fmt.Println(graph)
}
