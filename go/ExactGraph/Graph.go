package ExactGraph

import "math/rand"

type Graph struct {
	edges    []Edge
	vertices []Vertex
}

func NewGraph() *Graph {
	return &Graph{
		edges:    []Edge{},
		vertices: []Vertex{},
	}
}

func GenerateGraph(n int) *Graph {
	g := NewGraph()
	if n < 4 {
		return nil
	}
	for i := 0; i < n; i++ {
		g.AddVertex(Vertex{i})
	}

	quantityOfActedVertices := 2
	firstNumber, secondNumber := 0, 0

	g.AddEdgeInVertexFormat(g.vertices[1], g.vertices[2], rand.Intn(100)+1)

	for quantityOfActedVertices < n {
		firstNumber = rand.Intn(quantityOfActedVertices) + 1
		quantityOfActedVertices++
		secondNumber = quantityOfActedVertices
		g.AddEdgeInVertexFormat(g.vertices[firstNumber], g.vertices[secondNumber], rand.Intn(100)+1)
	}

	randomQuantity := rand.Intn(2*n) + 1
	for i := 1; i < randomQuantity; i++ {
		firstNumber = rand.Intn(n) + 1
		secondNumber = rand.Intn(n) + 1
		if firstNumber != secondNumber {
			g.AddEdgeInVertexFormat(g.vertices[firstNumber], g.vertices[secondNumber], rand.Intn(100)+1)
		}
	}

	return g
}
