package dfs

import (
	"testing"

	"io.vava.datastructure/graph"
	"io.vava.datastructure/util"
)

const TestGraphPath = "/data/g.txt"

func TestCC(t *testing.T) {
	g := graph.FileAsGraph(util.GetFileAbsolutePath(TestGraphPath))
	cc := NewCC(g)
	t.Log("cc count:", cc.Count())
}

func TestCCDistributed(t *testing.T) {
	g := graph.FileAsGraph(util.GetFileAbsolutePath(TestGraphPath))
	cc := NewCCDistributed(g)
	t.Log("connected components count:", cc.Count())
	t.Logf("%v and %v is connected? %v", 1, 6, cc.IsConnected(1, 6))
	t.Logf("%v and %v is connected? %v", 2, 5, cc.IsConnected(2, 5))
	t.Log("connected components:", cc.Components())
}

func TestSingleSourcePath(t *testing.T) {
	g := graph.FileAsGraph(util.GetFileAbsolutePath(TestGraphPath))
	v := 0
	p := NewSingleSourcePath(g, v)

	vertices := []int{
		0, 1, 2, 3, 4, 5, 6,
	}
	for _, w := range vertices {
		t.Logf("Path %v → %v: %v", v, w, p.Path(w))
	}
}

func TestCycleDetection(t *testing.T) {
	g1 := graph.FileAsGraph(util.GetFileAbsolutePath(TestGraphPath))
	t.Log("Graph has cycle?", NewCycleDetection(g1).HasCycle())

	text := `7 5
			 0 1
			 0 2
			 1 3
			 1 4
			 2 6`
	g2 := graph.StringAsGraph(text)
	t.Log("Graph has cycle?", NewCycleDetection(g2).HasCycle())

	text = `4 4
			0 1
			0 3
			1 2
			2 3`
	g3 := graph.StringAsGraph(text)
	t.Log("Graph has cycle?", NewCycleDetection(g3).HasCycle())
}

func TestBipartitionDetection(t *testing.T) {
	text := `7 6
			0 1
			0 2
			1 3
			1 4
			2 3
			2 6`
	g := graph.StringAsGraph(text)
	t.Log("Graph is bipartite?", NewBipartitionDetection(g).IsBipartite())

	text = `4 6
			0 1
			0 2
			0 3
			1 2
			1 3
			2 3`
	g = graph.StringAsGraph(text)
	t.Log("Graph is bipartite?", NewBipartitionDetection(g).IsBipartite())

	text = `4 4
			0 1
			0 3
			1 2
			2 3`
	g = graph.StringAsGraph(text)
	t.Log("Graph is bipartite?", NewBipartitionDetection(g).IsBipartite())
}

const (
	g1Text = `7 8
	0 1
	0 2
	1 3
	2 3
	3 5
	4 5
	4 6
	5 6`

	g2Text = `12 16
	0 1
	0 2
	1 3
	2 3
	3 5
	4 5
	4 6
	4 7
	5 6
	6 8
	8 9
	8 10
	8 11
	9 10
	9 11
	10 11`

	g3Text = `5 6
	0 1
	0 2
	1 2
	2 3
	2 4
	3 4`

	treeText = `7 6
	0 1
	0 3
	1 6
	2 3
	2 5
	3 4`
)

func TestFindBridges(t *testing.T) {
	graphTexts := []string{g1Text, g2Text, g3Text, treeText}
	for _, text := range graphTexts {
		g := graph.StringAsGraph(text)
		t.Log("Graph's bridges?", NewFindBridges(g).Result())
	}
}

func TestFindCutPoints(t *testing.T) {
	graphTexts := []string{g1Text, g2Text, g3Text, treeText}
	for _, text := range graphTexts {
		g := graph.StringAsGraph(text)
		t.Log("Graph's cut points?", NewFindCutPoints(g).Result())
	}
}

func TestFindCutPointsVerbose(t *testing.T) {
	graphTexts := []string{g3Text}
	for _, text := range graphTexts {
		g := graph.StringAsGraph(text)
		t.Log("Graph's cut points?", NewFindCutPointsWithLog(g, true).Result())
	}
}
