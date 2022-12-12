package graph

import (
	"fmt"
	"strings"

	"io.vava.datastructure/tree"
)

// AdjMap Adjacency Tree Map
type AdjMap struct {
	v        int
	e        int
	adj      []tree.TreeMap
	directed bool
}

func (g *AdjMap) ValidateVertex(v int) {
	if v < 0 || v > g.v {
		panic(fmt.Sprintf("vertex %v is invalid", v))
	}
}

func (g *AdjMap) E() int {
	return g.e
}

func (g *AdjMap) V() int {
	return g.v
}

func (g *AdjMap) IsDirected() bool {
	return g.directed
}

func (g *AdjMap) AddEdge(v int, w int) {
	g.ValidateVertex(v)
	g.ValidateVertex(w)

	if g.adj[v].Contains(w) {
		return
	}

	g.adj[v].Put(w, 1)
	if !g.directed {
		g.adj[w].Put(v, 1)
	}
	g.e++
}

func (g *AdjMap) RemoveEdge(v int, w int) {
	g.ValidateVertex(v)
	g.ValidateVertex(w)
	g.adj[v].Remove(w)
	if !g.directed {
		g.adj[w].Remove(v)
	}
}

func (g *AdjMap) Adj(v int) []int {
	g.ValidateVertex(v)
	return g.adj[v].Keys()
}

func (g *AdjMap) Degree(v int) int {
	g.ValidateVertex(v)
	return g.adj[v].Size()
}

func (g *AdjMap) GetWeight(v, w int) int {
	g.ValidateVertex(v)
	g.ValidateVertex(w)
	return g.adj[v].Get(w)
}

func (g *AdjMap) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("V = %v, E = %v\n", g.v, g.e))
	for vertex := 0; vertex < g.v; vertex++ {
		sb.WriteString(fmt.Sprintf("%v: ", vertex))
		for _, entry := range g.adj[vertex].EntrySet() {
			sb.WriteString(fmt.Sprintf("(%v: %v) ", entry.K, entry.V))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (g *AdjMap) Clone() Graph {
	dstAdj := make([]tree.TreeMap, g.V())
	for v := 0; v < g.V(); v++ {
		dstAdj[v] = tree.NewTreeMap()
		for _, entry := range g.adj[v].EntrySet() {
			dstAdj[v].Put(entry.K, entry.V)
		}
	}

	return &AdjMap{
		v:   g.V(),
		e:   g.E(),
		adj: dstAdj,
	}
}

func validateVertex(v int, max int) {
	if v < 0 || v > max {
		panic(fmt.Sprintf("vertex %v is invalid", v))
	}
}
