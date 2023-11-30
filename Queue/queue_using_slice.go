package queue

import (
	"errors"
	"fmt"
)

type queue struct {
	items []int
	Size  int
}

func NewQueueUsingSlice(size int) *queue {
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

func (q *queue) Peek() {
	fmt.Println("the next value to be removed:", q.items[0])
}
