package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Point
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * math.Pow(c.Radius, 2)
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Point
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Area:%.3f Perimeter:%.3f\n", s.Area(), s.Perimeter())
}

type Movable interface {
	Move(dx, dy float64)
	PrintMovableInfo()
}

func (p *Point) Move(dx, dy float64) {
	p.x += dx
	p.y += dy
}

func (p Point) PrintMovableInfo() {
	fmt.Printf("Coords: x-> %.2f y-> %.2f\n", p.x, p.y)
}

func main() {
	rec := Rectangle{Width: 3, Height: 4, Point: Point{x: 1, y: 2}}
	cir := Circle{Radius: 3, Point: Point{x: 2, y: 1}}
	rec.PrintMovableInfo()
	cir.PrintMovableInfo()
	rec.Move(3, 4)
	cir.Move(3, 4)
	rec.PrintMovableInfo()
	cir.PrintMovableInfo()
	PrintShapeInfo(rec)
	PrintShapeInfo(cir)
}
