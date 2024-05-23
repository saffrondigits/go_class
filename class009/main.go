package main

import (
	"fmt"
	"time"

	"github.com/good-binary/utility/random"
)

func main() {

	for i := 0; i < 10; i++ {
		name := random.RandomFullName()
		fmt.Println(name)
		time.Sleep(time.Second)
	}
}
