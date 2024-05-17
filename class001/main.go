package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	p1 := Person{
		Name:  "John",
		Age:   30,
		Email: "mail@john.com",
	}

	p2 := Person{
		Name:  "Bob",
		Age:   20,
		Email: "mail@bob.com",
	}

	var persons = []Person{p1, p2}

	fmt.Printf("p1: %+v\n", p1)
	fmt.Printf("p2: %+v\n", p2)
	fmt.Printf("persons: %v\n", persons)
}
