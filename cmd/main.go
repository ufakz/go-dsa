package main

import (
	"fmt"

	"github.com/farouqu/go-dsa/pkg/bstree"
	"github.com/farouqu/go-dsa/pkg/graph"
)

func main() {
	//Test Graph
	g := &graph.Graph{}

	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 1)

	//invalid case
	g.AddEdge(5, 5)

	g.Print()

	//Test Binary Search Tree
	tree := bstree.Node{Key: 100}
	tree.Insert(50)
	tree.Insert(200)
	tree.Insert(300)

	fmt.Println(tree.Search(30))
	fmt.Println(tree)
}
