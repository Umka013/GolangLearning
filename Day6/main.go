package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) newPoint(x, y float64) {
	p.x = x
	p.y = y
}

func (point Point) Distance(point2 Point) float64 {
	return math.Sqrt(math.Pow(point.x-point2.x, 2) + math.Pow(point.y-point2.y, 2))
}

func main() {
	p1 := Point{x: 10, y: 10}
	p2 := Point{x: 20, y: 20}
	p1.newPoint(100, 100)
	fmt.Println(p1.Distance(p2))
}
