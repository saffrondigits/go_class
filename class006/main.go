package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	student1 := Student{
		name: "John",
	}

	SayHello(student1)

	student2 := Student{
		name: "Alice",
	}

	student2.SayHello()

}

func SayHello(s Student) {
	fmt.Printf("Hello, %s\n", s.name)
}

// A method is nothing but a function that has a caller type
// Methods can have paremeters and return values
// But if they are implemented by a type then the same type can call the function

func (s Student) SayHello() {
	fmt.Printf("Hello, %s\n", s.name)
}
