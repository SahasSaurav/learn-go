package main

import (
	"fmt"
)

// DataStructure defines the common interface for stack and queue
type DataStructure[T any] interface {
	Push(element T)
	Pop() (T, error)
	IsEmpty() bool
	Print()
}

// Stack represents a stack data structure
type Stack[T any] struct {
	elements []T
}

// Queue represents a queue data structure
type Queue[T any] struct {
	elements []T
}

// CreateNewStack initializes a new stack
func CreateNewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the stack
func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

// Pop removes and returns the top element from the stack
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("stack is empty")
	}
	topElement := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return topElement, nil
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Print displays the elements in the stack
func (s *Stack[T]) Print() {
	fmt.Printf("Stack: %v\n", s.elements)
}

// CreateNewQueue initializes a new queue
func CreateNewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Push adds an element to the queue
func (q *Queue[T]) Push(element T) {
	q.elements = append([]T{element}, q.elements...)
}

// Pop removes and returns the front element from the queue
func (q *Queue[T]) Pop() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	frontElement := q.elements[len(q.elements)-1]
	q.elements = q.elements[:len(q.elements)-1]
	return frontElement, nil
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Print displays the elements in the queue
func (q *Queue[T]) Print() {
	fmt.Printf("Queue: %v\n", q.elements)
}

func main() {
	
	stack := CreateNewStack[int]()
	fmt.Println("Is stack empty?", stack.IsEmpty())
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Print()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Print()

	queue := CreateNewQueue[int]()
	fmt.Println("Is queue empty?", queue.IsEmpty())
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	queue.Push(6)
	queue.Print()
	queue.Pop()
	queue.Pop()
	queue.Print()
}
