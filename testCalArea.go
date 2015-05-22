package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{12, 2}
	c1 := Circle{10}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of c1 is: ", c1.area())
}
