package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Head     *Node
	Capacity int
	Length   int
}

func NewStackUsingList(size int) *Stack {
	return &Stack{
		Head:     nil,
		Capacity: size,
	}
}

type Node struct {
	Data int
	Next *Node
}

func (s *Stack) Push(data int) {
	if s.Capacity == s.Length {
		fmt.Println("stack overflow")
		return
	}
	newNode := &Node{Data: data}
	if s.Head == nil {
		s.Head = newNode
		s.Length++
		return
	}
	newNode.Next = s.Head
	s.Head = newNode
	s.Length++
}

func (s *Stack) Pop(data int) (int, error) {
	if s.Length == 0 {
		return 0, errors.New("stack underflow")
	}
	popped := s.Head.Data
	s.Head = s.Head.Next
	s.Length--
	return popped, nil
}

func (s *Stack) Peek(data int) {
	fmt.Println("the value popped to be next:", s.Head.Data)
}
