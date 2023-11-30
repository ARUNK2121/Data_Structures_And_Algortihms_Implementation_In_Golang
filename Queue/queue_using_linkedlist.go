package queue

import (
	"errors"
	"fmt"
)

// initializer function which returns a new instance of queue
func NewQueueUsingList(size int) *Queue {
	return &Queue{
		Head:     nil,
		Capacity: size,
	}
}

// structure for holding the queue
type Queue struct {
	Head     *Node
	Tail     *Node
	Capacity int
	Length   int
}

// structure for the node
type Node struct {
	Data int
	Next *Node
	Prev *Node
}

// enque adds the given data into queue
func (q *Queue) Enqueue(data int) {
	if q.Capacity == q.Length {
		fmt.Println("queue is full")
		return
	}
	newNode := &Node{Data: data}
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
		q.Length++
		return
	}

	newNode.Next = q.Head
	q.Head.Prev = newNode
	q.Head = newNode

	q.Length++
}

// dequeue returns the data in a FIFO order
func (q *Queue) DeQueue() (int, error) {
	fmt.Println("head:", q.Head)
	fmt.Println("tail:", q.Tail)
	if q.Length == 0 {
		return 0, errors.New("queue is empty")
	}

	if q.Tail == q.Head {
		fmt.Println("reached")
		data := q.Tail.Data
		fmt.Println("data:", data)
		q.Head = nil
		fmt.Println("67")
		q.Tail = nil
		fmt.Println("56")
		return data, nil
	}
	//take the data from tail
	popped := q.Tail.Data
	q.Tail = q.Tail.Prev

	q.Tail.Next = nil
	q.Length--
	return popped, nil
}

// peek displays data to be dequeued next
func (q *Queue) Peek() {
	if q.Tail == nil {
		fmt.Println("queue is empty")
		return
	}
	fmt.Println("the value popped to be next:", q.Tail.Data)
}
