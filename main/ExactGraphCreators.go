package main

import "math/rand"

func createRandomGraph(n int) *Graph {
	g := NewGraph()
	if n < 4 {
		return nil
	}
	for i := 0; i < n; i++ {
		g.AddVertex(Vertex{i, 0, 0})
	}

	quantityOfActedVertices := 2
	firstNumber, secondNumber := 0, 0

	g.AddEdgeInVertexFormat(g.Vertices[1], g.Vertices[2], float64(rand.Intn(100)+1))

	for quantityOfActedVertices < n {
		firstNumber = rand.Intn(quantityOfActedVertices) + 1
		quantityOfActedVertices++
		secondNumber = quantityOfActedVertices
		g.AddEdgeInVertexFormat(g.Vertices[firstNumber], g.Vertices[secondNumber], float64(rand.Intn(100)+1))
	}

	randomQuantity := rand.Intn(2*n) + 1
	for i := 1; i < randomQuantity; i++ {
		firstNumber = rand.Intn(n) + 1
		secondNumber = rand.Intn(n) + 1
		if firstNumber != secondNumber {
			g.AddEdgeInVertexFormat(g.Vertices[firstNumber], g.Vertices[secondNumber], float64(rand.Intn(100)+1))
		}
	}

	return g
}
