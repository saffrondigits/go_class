package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

type Student struct {
	Name  string
	Age   int
	Email string
}

func main() {
	persons := []Person{
		{
			Name:  "John",
			Age:   30,
			Email: "mail@john.com",
		},
		{
			Name:  "Alice",
			Age:   30,
			Email: "mail@alice.com",
		},
		{
			Name:  "Bob",
			Age:   25,
			Email: "mail@bob.com",
		},
		{
			Name:  "Carol",
			Age:   25,
			Email: "mail@carol.com",
		},
	}

	for _, v := range persons {
		fmt.Printf("Name: %s, Age: %d, Email : %s\n", v.Name, v.Age, v.Email)
	}

}
