package bfs

import (
	"fmt"
	"io.vava.datastructure/graph"
	"io.vava.datastructure/queue"
	"testing"
)

const TestGraphString = `7 7
						 0 1
						 0 2
						 1 3
						 1 4
						 2 3
						 2 6
						 5 6`

type BFSVisitor struct {
	graph.Graph
	visited []bool
	fn      func(v int)
}

func BFS(g graph.Graph, fn func(v int)) {
	visitor := &BFSVisitor{
		Graph:   g,
		visited: make([]bool, g.V()),
		fn:      fn,
	}
	for v := 0; v < g.V(); v++ {
		if !visitor.visited[v] {
			visitor.bfs(v)
		}
	}
}

func (b *BFSVisitor) bfs(s int) {
	q := queue.NewLoopQueue[int]()
	q.Enqueue(s)
	b.visited[s] = true

	for !q.IsEmpty() {
		v := q.Dequeue()
		b.fn(v)
		for _, w := range b.Graph.Adj(v) {
			if !b.visited[w] {
				q.Enqueue(w)
				b.visited[w] = true
			}
		}
	}
}

func TestBFSVisitor(t *testing.T) {
	g := graph.StringAsGraph(TestGraphString)
	BFS(g, func(v int) {
		fmt.Print(v, " ")
	})
	fmt.Println()
}
