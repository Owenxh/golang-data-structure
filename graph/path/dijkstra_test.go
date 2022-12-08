package path

import (
	"fmt"
	"io.vava.datastructure/graph"
	"testing"
)

const DijkstraGText = `5 8
					0 1 4
					0 2 2
					1 2 1
					1 3 2
					1 4 3
					2 3 4
					2 4 5
					3 4 1`

func TestDijkstra(t *testing.T) {
	graphTexts := []string{DijkstraGText}
	for _, text := range graphTexts {
		g := graph.StringAsWeightedGraph(text)
		dij := NewDijkstra(g, 0)
		for v := 0; v < g.V(); v++ {
			fmt.Printf("Distance %d -> %d is %d, ", dij.S, v, dij.DistTo(v))
			fmt.Printf("paths are %d\n", dij.Path(v))
		}
	}
}
