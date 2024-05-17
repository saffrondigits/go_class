package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Email   string
	Hobbies []string
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
			Age:   45,
			Email: "mail@alice.com",
		},
		{
			Name:  "Bob",
			Age:   25,
			Email: "mail@bob.com",
		},
		{
			Name:  "Carol",
			Age:   18,
			Email: "mail@carol.com",
		},
	}

	for _, v := range persons {
		fmt.Printf("Name: %s, Age: %d, Email : %s\n", v.Name, v.Age, v.Email)
	}

	fmt.Println("\nPeople who are 18 or less")
	// Print the name of people who are 18 or less
	for _, v := range persons {
		if v.Age <= 18 {
			fmt.Printf("Name: %s, Age: %d, Email : %s\n", v.Name, v.Age, v.Email)
		}
	}

	fmt.Println("\nPeople who are above 18")
	// Print the name of people who are above 18
	for _, v := range persons {
		if v.Age > 18 {
			fmt.Printf("Name: %s, Age: %d, Email : %s\n", v.Name, v.Age, v.Email)
		}
	}
}

// TODO:

// Populate hobbies
// Print people havign more than 2 hobbies
// Print people wiho plays guitar
// Print people who can play guitat and teach SAP
