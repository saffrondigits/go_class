package main

import "fmt"

type User struct {
	Name string
}

func main() {
	u1 := User{Name: "A User"}

	calculateArea(u1)
	calculateArea(4)
	calculateArea(22.1)
	calculateArea("string")

	fmt.Println(true)
}

func calculateArea(rand interface{}) {
	var a int
	a = rand.(int)
	fmt.Println("Interface", a)
}

// Strings
// Strconv
// os
// io
// time
// fmt
// errors
// math/rand
// net/http
