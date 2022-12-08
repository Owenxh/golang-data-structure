package path

import (
	"fmt"
	"io.vava.datastructure/graph"
	"testing"
)

func TestBellmanFord_DistTo(t *testing.T) {
	graphTexts := []string{DijkstraGText}
	for _, text := range graphTexts {
		g := graph.StringAsWeightedGraph(text)
		bf := NewBellmanFord(g, 0)
		for v := 0; v < g.V(); v++ {
			fmt.Printf("Distance %d -> %d is %d, ", bf.S, v, bf.DistTo(v))
			fmt.Printf("paths are %d\n", bf.Path(v))
		}
		fmt.Println()
	}
}
