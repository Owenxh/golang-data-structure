package dfs

import (
	"io.vava.datastructure/graph"
)

type FindCutPoints struct {
	graph.Graph
	visited []bool
	count   int
	order   []int
	low     []int
	res     map[int]struct{}
}

func NewFindCutPoints(g graph.Graph) *FindCutPoints {
	f := &FindCutPoints{
		Graph:   g,
		visited: make([]bool, g.V()),
		order:   make([]int, g.V()),
		low:     make([]int, g.V()),
		res:     make(map[int]struct{}),
	}
	for v := 0; v < g.V(); v++ {
		if !f.visited[v] {
			f.dfs(v, v)
		}
	}
	return f
}

// dfs Depth-First-Search
func (f *FindCutPoints) dfs(v int, parent int) {
	f.visited[v] = true
	f.order[v] = f.count
	f.low[v] = f.order[v]
	f.count++
	var child int
	for _, w := range f.Graph.Adj(v) {
		if !f.visited[w] {
			// fmt.Printf("<start> parent:%d -> [v:%d -> w:%d] \n", parent, v, w)
			f.dfs(w, v)
			// fmt.Printf("<end> parent:%d -> [v:%d -> w:%d] \n", parent, v, w)
			f.low[v] = min(f.low[v], f.low[w])
			if v != parent && f.low[w] >= f.order[v] {
				f.res[v] = struct{}{}
				// fmt.Printf("v:%d is cut point\n", v)
			}
			child++
			// 根结点有多个孩子，则它是一个割点
			if v == parent && child > 1 {
				f.res[v] = struct{}{}
			}
		} else if w != parent {
			f.low[v] = min(f.low[v], f.low[w])
		}
	}
}

func (f *FindCutPoints) Result() []int {
	if len(f.res) == 0 {
		return nil
	}
	points := make([]int, len(f.res))
	var i int
	for k := range f.res {
		points[i] = k
		i++
	}
	return points
}
