package Graphics

import "math"

type GraphicVertex struct {
	Number int
	X      int
	Y      int
}

func (v1 GraphicVertex) isEqual(v2 GraphicVertex) bool {
	return v1.Number == v2.Number
}

func generateVertices(n, sizeFrameX, sizeFrameY int) []GraphicVertex {
	var vertices []GraphicVertex
	centerX := sizeFrameX / 2
	centerY := sizeFrameY / 2
	radius := int(float64(centerY) * 0.9)

	for i := 1; i <= n; i++ {
		x := int(float64(centerX) + float64(radius)*math.Cos(float64(i)*math.Pi*2/float64(n)))
		y := int(float64(centerY) + float64(radius)*math.Sin(float64(i)*math.Pi*2/float64(n)))
		vertices = append(vertices, GraphicVertex{x, y, i})
	}
	return vertices
}
