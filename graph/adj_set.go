package graph

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// AdjSet Adjacency Set
type AdjSet struct {
	v   int
	e   int
	adj []TreeSet
}

func (g *AdjSet) validate(v int) {
	if v < 0 || v > g.v {
		panic(fmt.Sprintf("vertex %v is invalid", v))
	}
}

func (g *AdjSet) E() int {
	return g.e
}

func (g *AdjSet) V() int {
	return g.v
}

func (g *AdjSet) AddEdge(v int, w int) {
	g.validate(v)
	g.validate(w)

	if !g.adj[v].Contains(w) {
		g.adj[v].Put(w)
		g.adj[w].Put(v)
		g.e++
	}
}

func (g *AdjSet) Adj(v int) []int {
	g.validate(v)
	return g.adj[v].Keys()
}

func (g *AdjSet) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("V = %v, E = %v\n", g.v, g.e))
	for vertex := 0; vertex < g.v; vertex++ {
		sb.WriteString(fmt.Sprintf("%v: ", vertex))
		vertices := g.adj[vertex].Keys()
		for _, w := range vertices {
			sb.WriteString(fmt.Sprintf("%v ", w))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func validateVertex(v int, max int) {
	if v < 0 || v > max {
		panic(fmt.Sprintf("vertex %v is invalid", v))
	}
}

func NewGraph(file string) Graph {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Print(err)
		}
	}()

	reg := regexp.MustCompile(`\d`)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var V, E int
	var adj []TreeSet
	for row := 0; scanner.Scan(); row++ {
		tokens := reg.FindAllString(scanner.Text(), -1)
		if len(tokens) != 2 {
			panic(fmt.Sprintf("invalid data of row: %v", row))
		}
		if row == 0 {
			V, _ = strconv.Atoi(tokens[0])
			E, _ = strconv.Atoi(tokens[1])
			adj = make([]TreeSet, V)
			for i := 0; i < len(adj); i++ {
				adj[i] = NewTreeSet()
			}
		} else {
			v, _ := strconv.Atoi(tokens[0])
			validateVertex(v, V)
			w, _ := strconv.Atoi(tokens[1])
			validateVertex(w, V)
			if v == w {
				panic("Self loop is detected!")
			}
			if adj[v].Contains(w) {
				panic("Parallel edges detected!")
			}
			adj[v].Put(w)
			adj[w].Put(v)
		}
	}
	return &AdjSet{
		v:   V,
		e:   E,
		adj: adj,
	}
}
