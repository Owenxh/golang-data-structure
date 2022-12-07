package util

import "io.vava.datastructure/graph"

// Dijkstra 算法实现求解有权图单源最小路径
type Dijkstra struct {
	graph.WeightedGraph
	S       int
	Dis     []int
	Visited []bool
}
