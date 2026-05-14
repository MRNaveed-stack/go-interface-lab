package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func totalArea(shapes []Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func Area(shape Shape) float64 {
	area := shape.Area()
	fmt.Println("Area of the shape is: ", area)
	return area
}

func main() {
	c := Circle{Radius: 1}
	r := Rectangle{Width: 1, Height: 2}
	shapes := []Shape{
		Circle{Radius: 4},
		Circle{Radius: 9},
		Rectangle{Width: 3, Height: 7},
		Rectangle{Width: 1, Height: 5},
	}

	fmt.Println("Total area of all the shapes are: ", totalArea(shapes))

	fmt.Println("area of the circle is: ", c.Area())
	fmt.Println("area of the rectangle is: ", r.Area())

}
