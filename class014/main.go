package main

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	sq := &Square{5}
	sq.Area()

	c1 := &Circle{3}
	c1.Area()

	calculateArea(c1)
}

func calculateArea(shape Shape) float64 {
	return shape.Area()
}

// defer, panic, recover errors
// Closures
// Packages
