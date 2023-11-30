package main

import (
	"errors"
	"fmt"
)

func main() {
	s := NewStack(10)

	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)

	popped, err := s.pop()
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	fmt.Println("popped:", popped)
	s.pop()

	s.peek()
}

type stack struct {
	array []int
	Size  int
}

func NewStack(stacksize int) *stack {
	return &stack{
		array: []int{},
		Size:  stacksize,
	}
}

func (s *stack) push(value int) {
	if s.Size == len(s.array) {
		fmt.Println("stack overflow")
		return
	}
	s.array = append(s.array, value)
	s.Size++
	fmt.Println("added:", value)
}

func (s *stack) pop() (int, error) {
	if len(s.array) == 0 {
		return 0, errors.New("stack underflow")
	}
	picked := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	fmt.Println("removed")
	return picked, nil
}

func (s *stack) peek() {
	if len(s.array) == 0 {
		fmt.Println("stack is empty")
	}

	fmt.Println("next value to be popped :", s.array[0])
}
