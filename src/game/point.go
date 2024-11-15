package game

import "math"

type Point struct {
	X float64
	Y float64
}

func (v Point) Normalize() Point {
	magnitude := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Point{v.X / magnitude, v.Y / magnitude}
}
