package graph

import "fmt"

//Graph
type Graph struct {
	Vertices []*Vertex
}

//Vertex
type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

//AddVertex adds a Vertex to a Graph
func (g *Graph) AddVertex(k int) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}

//AddEdge

//Contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

//GetVertex

//Print will print the adjency list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex %v: ", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf(" %v", v.Key)
		}
	}
	fmt.Println()
}
