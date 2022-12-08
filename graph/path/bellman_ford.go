package path

import (
	"io.vava.datastructure/graph"
	"io.vava.datastructure/util"
	"math"
)

// BellmanFord 算法实现求解带负权边的最短路径问题
type BellmanFord struct {
	// 图
	graph.WeightedGraph
	// S 起点
	S int
	// 起点 S 至各顶点的距离
	Dis []int
	// 标记顶点是否已经求得最短距离
	Visited []bool
	// 从起点到顶点的最短路径上的前一个顶点
	Pre []int
}

func NewBellmanFord(g graph.WeightedGraph, s int) *BellmanFord {
	g.ValidateVertex(s)

	dis := make([]int, g.V())

	// 初始化：令各顶点的距离为正无穷，这里用 MaxInt32 表示正无穷
	for v := 0; v < g.V(); v++ {
		dis[v] = math.MaxInt32
	}
	// 起点的距离到自己为 0
	dis[s] = 0
	visited := make([]bool, g.V())
	pre := make([]int, g.V())

	// 使用最小堆存储各顶点与起点间的距离
	pq := util.NewPriorityQueue[*graph.Node](graph.LessNode)
	pq.Push(&graph.Node{V: s, Dis: dis[s]})

	for !pq.IsEmpty() {
		cur := pq.Pop().V

		// 跳过已经求得出短距离的顶点
		if visited[cur] {
			continue
		}

		// 堆中取的第一个顶点，当前计算出的该顶点与起点间的距离即为起点到它的最短距离；
		// 因为如果从其他顶点回到该顶点求得的距离肯定大于当前计算结果，不可能得到更小的值了
		visited[cur] = true

		for _, w := range g.Adj(cur) {
			if dis[cur]+g.GetWeight(cur, w) < dis[w] {
				dis[w] = dis[cur] + g.GetWeight(cur, w)
				pq.Push(&graph.Node{V: w, Dis: dis[w]})
				pre[w] = cur
			}
		}
	}

	return &BellmanFord{
		WeightedGraph: g,
		S:             s,
		Dis:           dis,
		Visited:       visited,
		Pre:           pre,
	}
}
