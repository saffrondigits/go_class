package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	var students = map[string]Student{
		"27th main": Student{
			name: "John",
			age:  30,
		},
		"10th cross": Student{
			name: "Bob",
			age:  20,
		},
		"1st aventu": Student{
			name: "Dave",
			age:  30,
		},
	}

	fmt.Printf("students: %v\n", students["1st aventu"])
}
