package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

var graph *dijkstra.Graph

func testGroup(nodes []dijkstra.Vertex, deep int) int {
	var newnodes []dijkstra.Vertex
	zero := nodes[0]
	for _, n := range nodes {
		if _, err := graph.Shortest(zero.ID, n.ID); err != nil {
			newnodes = append(newnodes, n)
		}
	}

	if len(newnodes) > 0 {
		deep = testGroup(newnodes, deep+1)
	}
	return deep
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	graph = dijkstra.NewGraph()

	for scanner.Scan() {
		value := scanner.Text()
		data := strings.Split(value, " <-> ")
		node := data[0]
		if _, err := graph.GetMapping(data[0]); err != nil {
			graph.AddMappedVertex(data[0])
		}
		data = strings.Split(data[1], ", ")
		for _, t := range data {
			if _, err := graph.GetMapping(t); err != nil {
				graph.AddMappedVertex(t)
			}
			graph.AddMappedArc(node, t, 1)
		}
	}

	fmt.Println(testGroup(graph.Verticies, 1))
}
