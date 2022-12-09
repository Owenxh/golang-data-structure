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
)

func FileAsGraph(file string) Graph {
	return FileAsWeightedGraph(file, false)
}

func FileAsDirectedGraph(file string) Graph {
	return FileAsWeightedGraph(file, true)
}

func FileAsWeightedGraph(file string, directed bool) WeightedGraph {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Print(err)
		}
	}()
	return ReadWeightedGraph(f, directed)
}

func StringAsGraph(src string) Graph {
	return StringAsWeightedGraph(src, false)
}

func StringAsDirectedGraph(src string) Graph {
	return StringAsWeightedGraph(src, true)
}

func StringAsWeightedGraph(src string, directed bool) WeightedGraph {
	return ReadWeightedGraph(bytes.NewReader([]byte(src)), directed)
}

func ReadGraph(src io.Reader) Graph {
	return readAdjMap(src, false)
}

func ReadWeightedGraph(src io.Reader, directed bool) WeightedGraph {
	return readAdjMap(src, directed)
}

func readAdjMap(src io.Reader, directed bool) *AdjMap {
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
		if !directed {
			adj[w].Put(v, weight)
		}
	}
	return &AdjMap{
		v:        V,
		e:        E,
		adj:      adj,
		directed: directed,
	}
}
