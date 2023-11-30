package main

import (
	"errors"
	"fmt"
)

func main() {

	q := queue{}
	fmt.Println(q)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)

	fmt.Println(q)

	a, err := q.DeQueue()
	if err != nil {
		fmt.Println(err.Error())
	}
	b, err := q.DeQueue()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(a)
	fmt.Println(b)
}

type queue struct {
	items []int
	Size  int
}

func NewQueue(size int) *queue {
	return &queue{
		items: []int{},
		Size:  size,
	}
}

func (q *queue) Enqueue(value int) {
	if q.Size == len(q.items) {
		fmt.Println("queue is full")
		return
	}
	q.items = append(q.items, value)
	fmt.Println("added:", value)
}

func (q *queue) DeQueue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}

	toRemove := q.items[0]
	q.items = q.items[1:]

	return toRemove, nil
}
