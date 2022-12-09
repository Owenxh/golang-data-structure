package graph

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io.vava.datastructure/util"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// AdjMap Adjacency Tree Map
type AdjMap struct {
	v   int
	e   int
	adj []util.TreeMap
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

func (g *AdjMap) AddEdge(v int, w int) {
	g.ValidateVertex(v)
	g.ValidateVertex(w)

	if !g.adj[v].Contains(w) {
		g.adj[v].Put(w, 1)
		g.adj[w].Put(v, 1)
		g.e++
	}
}

func (g *AdjMap) RemoveEdge(v int, w int) {
	g.ValidateVertex(v)
	g.ValidateVertex(w)
	g.adj[v].Remove(w)
	g.adj[w].Remove(v)
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
	dstAdj := make([]util.TreeMap, g.V())
	for v := 0; v < g.V(); v++ {
		dstAdj[v] = util.NewTreeMap()
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

func FileAsGraph(file string) Graph {
	return FileAsWeightedGraph(file)
}

func FileAsWeightedGraph(file string) WeightedGraph {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Print(err)
		}
	}()
	return ReadWeightedGraph(f)
}

func StringAsGraph(src string) Graph {
	return StringAsWeightedGraph(src)
}

func StringAsWeightedGraph(src string) WeightedGraph {
	return ReadWeightedGraph(bytes.NewReader([]byte(src)))
}

func ReadGraph(src io.Reader) Graph {
	return readAdjMap(src)
}

func ReadWeightedGraph(src io.Reader) WeightedGraph {
	return readAdjMap(src)
}

func readAdjMap(src io.Reader) *AdjMap {
	if src == nil {
		panic("invalid io.Reader")
	}

	reg := regexp.MustCompile(`-?\d+`)
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanLines)

	var V, E int
	var adj []util.TreeMap
	if scanner.Scan() {
		tokens := reg.FindAllString(scanner.Text(), -1)
		if len(tokens) != 2 {
			panic(fmt.Sprintf("read graph V & E faild at first row; required 2 tokens but got %d count", len(tokens)))
		}
		V, _ = strconv.Atoi(tokens[0])
		E, _ = strconv.Atoi(tokens[1])
		adj = make([]util.TreeMap, V)
		for i := 0; i < len(adj); i++ {
			adj[i] = util.NewTreeMap()
		}
	}
	for row := 1; scanner.Scan(); row++ {
		tokens := reg.FindAllString(scanner.Text(), -1)
		if len(tokens) < 2 || len(tokens) > 3 {
			panic(fmt.Sprintf("read edge failed at row: %v; got %d tokens", row, len(tokens)))
		}
		v, _ := strconv.Atoi(tokens[0])
		validateVertex(v, V)
		w, _ := strconv.Atoi(tokens[1])
		validateVertex(w, V)
		if v == w {
			panic("Self path is detected!")
		}
		if adj[v].Contains(w) {
			panic("Parallel edges detected!")
		}
		weight := 1
		if len(tokens) == 3 {
			weight, _ = strconv.Atoi(tokens[2])
		}
		adj[v].Put(w, weight)
		adj[w].Put(v, weight)
	}
	return &AdjMap{
		v:   V,
		e:   E,
		adj: adj,
	}
}
