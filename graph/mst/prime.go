package mst

import (
	"container/heap"
	"io.vava.datastructure/graph"
	"io.vava.datastructure/graph/dfs"
)

// Prime 算法实现最小生成树 - Minimum Tree Spanning
func Prime(g graph.WeightedGraph) ([]graph.WeightedEdge, bool) {
	// 图中所有顶点必须是连通的
	cc := dfs.NewCC(g)
	if cc.Count() > 1 {
		return nil, false
	}

	var mst []graph.WeightedEdge
	visited := make([]bool, g.V())
	visited[0] = true

	var hp graph.WeightEdgesHeap
	for _, w := range g.Adj(0) {
		hp.Push(graph.WeightedEdge{V: 0, W: w, Weight: g.GetWeight(0, w)})
	}
	heap.Init(&hp)
	for hp.Len() > 0 {
		edge := heap.Pop(&hp).(*graph.WeightedEdge)
		// 两个顶点 w & w 属于同一切分
		if visited[edge.W] && visited[edge.V] {
			continue
		}
		var nv = edge.W
		if visited[edge.W] {
			nv = edge.V
		}
		for _, w := range g.Adj(nv) {
			heap.Push(&hp, graph.WeightedEdge{V: nv, W: w, Weight: g.GetWeight(nv, w)})
		}

		visited[edge.W] = true
		visited[edge.V] = true
		mst = append(mst, *edge)
	}
	return mst, true
}
