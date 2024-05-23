package main

import "fmt"

func main() {
	y := 10

	increment(&y)

	fmt.Printf("y: %v\n", y)
}

func increment(x *int) {
	*x = *x + 1
}

func decrement(x int) {
	x--
}
