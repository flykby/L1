package point

import (
	"math"
)

type Point struct {
	x float32
	y float32
}

func NewPoint(x, y float32) *Point {
	return &Point{
		x : x,
		y : y,
	}
}

func (p *Point) GetDistance(point *Point) float32 {
	distance := math.Sqrt(math.Pow(float64(p.x - point.x), 2) + math.Pow(float64(p.y - point.y), 2))
	return float32(distance)
}
