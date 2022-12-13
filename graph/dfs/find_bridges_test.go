package dfs

import (
	"testing"

	"io.vava.datastructure/graph"
)

func TestFindBridges(t *testing.T) {
	graphTexts := []string{Graph1Text, Graph2Text, Graph3Text, Graph4Text}
	for _, text := range graphTexts {
		g := graph.TextAsGraph(text)
		t.Log("Graph's bridges?", NewFindBridges(g).Result())
	}
}
