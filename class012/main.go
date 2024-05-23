package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) UpdateAge() {
	u.Age++
}

func main() {
	usr := User{
		Name: "John",
		Age:  30,
	}

	fmt.Println("Age before update", usr.Age)
	usr.UpdateAge()
	fmt.Println("Age after update", usr.Age)

}
