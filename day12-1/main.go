package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RyanCarrier/dijkstra"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	graph := dijkstra.NewGraph()
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
	sum := 0
	zero, _ := graph.GetMapping("0")
	for _, n := range graph.Verticies {
		if _, err := graph.Shortest(zero, n.ID); err == nil {
			sum++
		}
	}
	fmt.Println(sum)
}
