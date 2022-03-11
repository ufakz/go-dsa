package main

import (
	"github.com/farouqu/go-dsa/pkg/graph"
)

func main() {
	g := &graph.Graph{}

	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}

	g.Print()
}
