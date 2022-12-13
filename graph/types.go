package graph

import (
	"fmt"
)

type Cloneable[E any] interface {
	Clone() E
}

type Graph interface {
	// V the vertex count
	V() int
	// E the edge count
	E() int
	// Adj adjacency of the vertex v
	Adj(v int) []int
	// AddEdge connect vertex v & w
	AddEdge(v, w int)
	// RemoveEdge remove edge between vertex v & w
	RemoveEdge(v, w int)
	// ValidateVertex validate whether v is valid vertex
	ValidateVertex(v int)
	// Degree returns the degree of vertex v
	Degree(v int) int
	// IsDirected returns true if directed graph
	IsDirected() bool
	// Indegree
	Indegree(int) int
	// Outdegree
	Outdegree(int) int
	// Cloneable graph
	Cloneable[Graph]
}

type Edge struct {
	V int
	W int
}

func (e Edge) String() string {
	return fmt.Sprintf("%d-%d", e.V, e.W)
}

type WeightedGraph interface {
	Graph
	GetWeight(v, w int) int
}

type WeightedEdge struct {
	V      int
	W      int
	Weight int
}

func LessWeightedEdge(a, b *WeightedEdge) bool {
	return a.Weight-b.Weight < 0
}

func (e WeightedEdge) String() string {
	return fmt.Sprintf("(%d-%d: %d)", e.V, e.W, e.Weight)
}

type Node struct {
	V   int // V 顶点
	Dis int // Dis 距离
}

func LessNode(a, b *Node) bool {
	return a.Dis-b.Dis < 0
}
