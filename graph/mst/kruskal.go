package mst

import (
	"io.vava.datastructure/graph"
	"io.vava.datastructure/graph/dfs"
	"io.vava.datastructure/tree"
	"sort"
)

// Kruskal 算法实现最小生成树 - Minimum Tree Spanning
type Kruskal struct {
	graph.WeightedGraph
	mst []graph.WeightedEdge
}

func NewKruskal(g graph.WeightedGraph) *Kruskal {
	// 图中所有顶点必须是连通的
	cc := dfs.NewCC(g)
	if cc.Count() > 1 {
		return &Kruskal{
			WeightedGraph: g,
		}
	}

	var edges graph.SortedWeightEdges
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if v < w {
				edges = append(edges, graph.WeightedEdge{V: v, W: w, Weight: g.GetWeight(v, w)})
			}
		}
	}
	sort.Sort(edges)

	var mst []graph.WeightedEdge
	// 并查集
	uf := tree.NewUnionFind5(g.V())
	for _, edge := range edges {
		if !uf.IsConnected(edge.V, edge.W) {
			mst = append(mst, edge)
			uf.UnionElements(edge.V, edge.W)
		}
	}
	return &Kruskal{
		WeightedGraph: g,
		mst:           mst,
	}
}

func (k *Kruskal) Result() []graph.WeightedEdge {
	return k.mst
}
