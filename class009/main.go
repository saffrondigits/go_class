package main

import "fmt"

func main() {
	a := 5

	fmt.Printf("Value of a: %d\n", a)
	fmt.Printf("Memory location of a: %v\n", &a)

	b := &a
	fmt.Printf("Value of b: %v\n", b)
	fmt.Printf("Memory location of b: %v\n", &b)
	fmt.Printf("Value of a: %v\n", *b)
}
