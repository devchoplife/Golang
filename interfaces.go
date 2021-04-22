package main

import (
	"fmt"
	"math"
)

//Interface allows us to look at different types a sif they are the same
//eg the following struct both have area method
//the interface below menas any type that has the method area() ad returns float64 is of type shape
//You can only use use things defined in the interface when accessing the interface
type Shape interface {
	area() float64
}

//You can create more than one interface
type Shape2 interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) area() float64 {
	return r.width * r.height
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{4.5}
	r1 := Rect{5, 7}

	shapes := []Shape{&c1, &r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}

	for _, shape := range shapes {
		fmt.Println(getArea(shape))
	}
}
