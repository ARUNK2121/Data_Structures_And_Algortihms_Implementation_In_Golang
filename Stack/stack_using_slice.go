package stack

import (
	"errors"
	"fmt"
)

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

func (s *stack) Peekush(value int) {
	if s.Size == len(s.array) {
		fmt.Println("stack overflow")
		return
	}
	s.array = append(s.array, value)
	fmt.Println("added:", value)
}

func (s *stack) Pop() (int, error) {
	if len(s.array) == 0 {
		return 0, errors.New("stack underflow")
	}
	picked := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	fmt.Println("removed")
	return picked, nil
}

func (s *stack) Peek() {
	if len(s.array) == 0 {
		fmt.Println("stack is empty")
	}

	fmt.Println("next value to be popped :", s.array[0])
}
