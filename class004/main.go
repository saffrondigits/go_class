package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name string
	age  int
}

func main() {
	students := make(map[string]Student)

	names := make([]string, 0)

	var command string // Shrey

	for {
		fmt.Printf("Enter command: ")
		fmt.Scanf("%s", &command)

		if command == "exit" {
			break
		} else if command == "list" {
			for _, name := range names { // ["dave", "shrey", "piyush"]
				s := students[name]
				fmt.Printf("%s is %d years old\n", s.name, s.age)
			}
		} else {
			var age string
			fmt.Printf("Enter age: ")
			fmt.Scanf("%s", &age)

			ageInt := convertToInt(age)

			if ageInt == -1 {
				fmt.Printf("Invalid age\n")
				continue
			}
			// stud :=

			students[command] = Student{
				name: command,
				age:  ageInt,
			}
			names = append(names, command)
		}
	}
}

func convertToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	if i < 0 {
		return -1
	}
	return i
}
