package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct{
	width int
	height int
}

func (r *rect) area() float64 {
	return float64(r.height) * float64(r.width)
}

func (r *rect) perim() float64 {
	return float64(2 * (r.width + r.height))
}

type circle struct {
	radius float64
}

func(c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func(c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	var g geometry
	g = &c
	fmt.Println(g.area())
	measure(&r)
	measure(&c)
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
