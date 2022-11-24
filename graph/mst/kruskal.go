package mst

import (
	"io.vava.datastructure/graph"
	"io.vava.datastructure/graph/dfs"
)

// Kruskal 算法实现最小生成树 - Minimum Tree Spanning
type Kruskal struct {
	graph.WeightedGraph
	mst []graph.WeightedEdge
}

func NewKruskal(g graph.WeightedGraph) *Kruskal {
	kruskal := &Kruskal{
		WeightedGraph: g,
	}

	// 图中所有顶点必须是连通的
	cc := dfs.NewCC(g)
	if cc.Count() > 1 {
		return kruskal
	}

	var edges []graph.WeightedEdge
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if v < w {
				edges = append(edges, graph.WeightedEdge{V: v, W: w, Weight: g.GetWeight(v, w)})
			}
		}
	}

	return kruskal
}
