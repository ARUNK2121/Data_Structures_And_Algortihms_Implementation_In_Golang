package linkedlist

import (
	"errors"
	"fmt"
)

const (
	ListEmptyError = "list is empty"
)

// structure represents the linked list
type List struct {
	Head *Node //head of linked list
	Tail *Node //tail of linked list
}

// structure represents the nodes in list
type Node struct {
	Data int   //actual data inside the node
	Next *Node //reference to the next node
	Prev *Node //reference to the previous node
}

// Insert takes a value and insert into the list
func (l *List) Insert(value int) {
	new := &Node{Data: value}
	if l.Head == nil {
		l.Head = new
		l.Tail = new
		return
	}

	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	l.Head.Prev = new
	new.Next = l.Head
	l.Head = new
}

// TraverseFroward will print the data inside
// list in the order of head to tail
func (l *List) TraverseFroward() {
	if l.Head == nil {
		fmt.Println(ListEmptyError)
		return
	}
	temp := l.Head
	for temp != nil {
		fmt.Println("data ->", temp.Data)
		temp = temp.Next
	}

}

// TraverseBackward will print the data inside
// list in the order of tail to head
func (l *List) TraverseBackward() {
	if l.Tail == nil {
		fmt.Println(ListEmptyError)
		return
	}
	temp := l.Tail
	for temp != nil {
		fmt.Println("data ->", temp.Data)
		temp = temp.Prev
	}

}

// Read from a array and insert data into list
func (l *List) ReadFromArray(arr []int) {
	for _, v := range arr {
		l.Insert(v)
	}
}

// Convert linked list into an array and
// return the data in list as an array
func (l *List) ConvertToArray() ([]int, error) {
	var result []int
	if l.Head == nil {
		return []int{}, errors.New(ListEmptyError)
	}
	temp := l.Head
	for temp != nil {
		result = append(result, temp.Data)
		temp = temp.Next
	}
	return result, nil
}

// Search performs a search in list and
// returns a boolen along with a error
func (l *List) Search(key int) (bool, error) {
	var result bool
	if l.Head == nil {
		return false, errors.New(ListEmptyError)
	}
	temp := l.Head
	for temp != nil {
		if temp.Data == key {
			return true, nil
		}
		temp = temp.Next
	}
	return result, nil
}

// establish a circluar connection between nodes in list
// use with caution because if you make the list into circular
// the provided methods will give un-expected results
func (l *List) MakeLinkedlistIntoCircular() {
	l.Head.Prev = l.Tail
	l.Tail.Next = l.Head
}

// sort will sort the list in ascending order
func (l *List) Sort() {
	temp := l.Head
	if l.Head == nil {
		fmt.Println(ListEmptyError)
		return
	}
	for temp != nil {
		comparable := temp.Next
		for comparable != nil {
			if temp.Data > comparable.Data {
				swap := temp.Data
				temp.Data = comparable.Data
				comparable.Data = swap
			}
			comparable = comparable.Next
		}
		temp = temp.Next
	}
}
