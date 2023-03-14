package main

type Graph struct {
	Edges    []Edge
	Vertices []Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Edges:    []Edge{},
		Vertices: []Vertex{},
	}
}
