package main

import "fmt"

type Stack struct {
	items []int
	count int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
	s.count++
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
	s.count--
}

func (s *Stack) Display() {
	for _, v := range s.items {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}
func main() {
	stack := &Stack{}
	stack.Push(5)
	stack.Push(3)
	stack.Push(9)
	stack.Display()
	stack.Push(7)
	stack.Display()

	stack.Push(1)
	stack.Display()
	stack.Push(4)
	stack.Display()
	stack.Pop()
	stack.Display()
	stack.Pop()
	stack.Display()
}
