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

//AddVertex adds a Vertex to a Graph.
func (g *Graph) AddVertex(k int) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}

//AddEdge adds an edge to a Graph.
//Directed graph - From, To
func (g *Graph) AddEdge(from, to int) {
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v---->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.Adjacent, to) {
		err := fmt.Errorf("existing edge (%v---->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	}
}

//Contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

//GetVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == k {
			return g.Vertices[i]
		}
	}
	return nil
}

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
