package graph

import (
	"fmt"
	"testing"

	"io.vava.datastructure/util"
)

func TestAdjSet(t *testing.T) {
	graph := FileAsGraph(util.GetFileAbsolutePath("/data/g.txt"))
	fmt.Println(graph)
}
