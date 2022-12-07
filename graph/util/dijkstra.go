package util

import (
	"io.vava.datastructure/graph"
	"io.vava.datastructure/util"
	"math"
)

// Dijkstra 算法实现求解有权图单源最小路径
type Dijkstra struct {
	// 图
	graph.WeightedGraph
	// S 起点
	S int
	// 起点 S 至各顶点的距离
	Dis []int
	// 标记顶点是否已经求得最短距离
	Visited []bool
}

func NewDijkstra(g graph.WeightedGraph, s int) *Dijkstra {
	g.ValidateVertex(s)

	dis := make([]int, g.V())

	// 初始化：令各顶点的距离为正无穷，这里用 MaxInt32 表示正无穷
	for v := 0; v < g.V(); v++ {
		dis[v] = math.MaxInt32
	}
	// 起点的距离到自己为 0
	dis[s] = 0
	visited := make([]bool, g.V())

	hp := util.NewHeap[*graph.Node](graph.LessNode)
	hp.PushElement(&graph.Node{V: s, Dis: dis[s]})

	for hp.Len() > 0 {
		//node := hp.PopElement()
		//if
	}

	return &Dijkstra{
		WeightedGraph: g,
		S:             s,
		Dis:           dis,
		Visited:       visited,
	}
}

func (d *Dijkstra) IsConnectedTo(v int) bool {
	d.ValidateVertex(v)
	return d.Visited[v]
}

func (d *Dijkstra) DistTo(v int) int {
	d.ValidateVertex(v)
	return d.Dis[v]
}
