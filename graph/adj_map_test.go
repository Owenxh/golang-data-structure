package graph

import (
	"fmt"
	"io.vava.datastructure/util"
	"testing"
)

func TestAdjSet(t *testing.T) {
	graph := FileAsGraph(util.GetFileAbsolutePath("/data/g.txt"))
	fmt.Println(graph)
}
