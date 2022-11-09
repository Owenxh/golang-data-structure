package dfs

import (
	"io.vava.datastructure/graph"
	"io.vava.datastructure/util"
	"testing"
)

func TestCC(t *testing.T) {
	g := graph.NewGraph(util.GetFileAbsolutePath("/graph/data/g.txt"))
	cc := NewCC(g)
	t.Log("cc count:", cc.Count())
}

func TestCCDistributed(t *testing.T) {
	g := graph.NewGraph(util.GetFileAbsolutePath("/graph/data/g.txt"))
	cc := NewCCDistributed(g)
	t.Log("connected components count:", cc.Count())
	t.Logf("%v and %v is connected? %v", 1, 6, cc.IsConnected(1, 6))
	t.Logf("%v and %v is connected? %v", 2, 5, cc.IsConnected(2, 5))
	t.Log("connected components:", cc.Components())
}

func TestSingleSourcePath(t *testing.T) {
	g := graph.NewGraph(util.GetFileAbsolutePath("/graph/data/g.txt"))
	v := 0
	p := NewSingleSourcePath(g, v)

	vertices := []int{
		0, 1, 2, 3, 4, 5, 6,
	}
	for _, w := range vertices {
		t.Logf("Path %v to %v: %v", v, w, p.Path(w))
	}
}
