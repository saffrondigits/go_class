package main

import "fmt"

type Car struct {
	Color string
	Speed int
}

func (c *Car) Display() {
	fmt.Printf("The car is %s and it's speed is %d\n", c.Color, c.Speed)
}

func (c *Car) Paint(newColor string) {
	c.Color = newColor
}

func (c *Car) Accelerate() {
	c.Speed = c.Speed + 1
}

func main() {
	myCar := &Car{
		Color: "Black",
		Speed: 45,
	}

	myCar.Display()

	myCar.Accelerate()

	myCar.Display()

	myCar.Paint("yellow")

	myCar.Display()
}
