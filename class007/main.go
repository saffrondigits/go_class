package main

import (
	"fmt"
)

func main() {
	circle1 := &Circle{}
	fmt.Println("Radius:", circle1.radius)

	circle1.SetRadius(5)
	fmt.Println("Radius:", circle1.radius)

	area := circle1.Area()
	fmt.Printf("area: %v\n", area)

	perimeter := circle1.Perimeter()
	fmt.Printf("perimeter: %v\n", perimeter)
}

// Method is function that is associated with a type
// If a method is implemented by a type, it will be called by the same type

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c *Circle) SetRadius(r float64) {
	c.radius = r
}
