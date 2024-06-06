package main

import "time"

var count = 1

func main() {

	go SayHello()
	go SayHi()
	time.Sleep(time.Second)
}

func SayHello() {
	time.Sleep(time.Second * 3)
	count = count + 1
}

func SayHi() {
	count++
}

// SayHello:
// 1. Picks count
// 2. Increments count
// 3. Reinitialize count

// SayHi:
// 1. Picks count
// 2. Increments count
// 3. Reinitialize count
