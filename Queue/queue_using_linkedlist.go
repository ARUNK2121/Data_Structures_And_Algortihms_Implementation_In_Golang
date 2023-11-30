package main

import (
	"errors"
	"fmt"
)

func NewQueueUsingList(size int) *Queue {
	return &Queue{
		Head:     nil,
		Capacity: size,
	}
}

type Queue struct {
	Head     *Node
	Tail     *Node
	Capacity int
	Length   int
}

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func (q *Queue) Enqueue(data int) {
	if q.Capacity == q.Length {
		fmt.Println("queue is full")
		return
	}
	newNode := &Node{Data: data}
	if q.Head == nil {
		q.Head = newNode
		q.Tail = q.Head
		q.Length++
		return
	}
	newNode.Next = q.Head
	q.Head.Prev = newNode
	q.Head = newNode
	q.Length++
}

func (q *Queue) DeQueue(data int) (int, error) {
	if q.Length == 0 {
		return 0, errors.New("queue is empty")
	}
	popped := q.Tail.Data
	q.Tail = q.Tail.Prev
	q.Tail.Next = nil
	q.Length--
	return popped, nil
}

func (q *Queue) Peek(data int) {
	fmt.Println("the value popped to be next:", q.Tail.Data)
}
