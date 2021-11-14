package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

type Rectangle struct {
	a, b float64
}

func (c Circle) getArea() float64 {
	return c.r * c.r * math.Pi
}

func (r Rectangle) getArea() float64 {
	return r.a * r.b
}

type Shape interface {
	getArea() float64
}

func getArea(s Shape) float64 {
	return s.getArea()
}

func main() {

	circle := Circle{10}
	rect := Rectangle{5, 20}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rect))
}
