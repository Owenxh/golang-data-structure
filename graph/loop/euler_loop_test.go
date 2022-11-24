package loop

import (
	"io.vava.datastructure/graph"
	"testing"
)

const EulerG1Text = `5 6
					0 1
					0 2
					1 2
					2 3
					2 4
					3 4`

const EulerG2Text = `11 15
					0 1
					0 3
					1 2
					1 4
					1 5
					2 5
					3 4
					4 5
					4 6
					5 7
					6 7
					7 8
					7 9
					8 10
					9 10`

func TestEulerLoop(t *testing.T) {
	graphTexts := []string{EulerG1Text, EulerG2Text}
	for _, text := range graphTexts {
		g := graph.StringAsGraph(text)
		hl := NewEulerLoop(g)
		t.Log("Graph has euler loop?", hl.HasEulerLoop())
		if hl.HasEulerLoop() {
			t.Log("Graph's euler loop: ", hl.Result())
		}
	}
}
