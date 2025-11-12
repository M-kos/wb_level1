package point

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p *Point) DistanceTo(other Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}
