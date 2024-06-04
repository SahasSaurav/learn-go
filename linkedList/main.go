package main

import "fmt"

type node[T any] struct {
	data T
	next *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	size int
}

func (l *LinkedList[T]) Append(element T) {
	newNode := &node[T]{data: element, next: nil}
	if l.IsEmpty() {
		l.head = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	l.size++
}

func (l *LinkedList[T]) Prepend(element T) {
	newNode := &node[T]{data: element, next: nil}
	if l.IsEmpty() {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
	l.size++
}

func (l *LinkedList[T]) RemoveFromLast() error {
	if l.IsEmpty() {
		return fmt.Errorf("linked list is empty")
	}
	if l.head.next == nil {
		l.head = nil
	} else {
		currentNode := l.head
		for currentNode.next.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = nil
	}
	l.size--
	return nil
}

func (l *LinkedList[T]) RemoveFromFront() error {
	if l.IsEmpty() {
		return fmt.Errorf("linked list is empty")
	}
	if l.head.next == nil {
		l.head = nil
	} else {
		l.head = l.head.next
	}
	l.size--
	return nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l *LinkedList[T]) Print() {
	data := make([]T, 0)
	currentNode := l.head
	for currentNode != nil {
		data = append(data, currentNode.data)
		currentNode = currentNode.next
	}

	fmt.Printf("Linked list %v\n", data)
}

func main() {
	list := LinkedList[int]{}
	list.Prepend(10)
	list.Print()
	list.RemoveFromLast()
	list.Print()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Prepend(0)
	list.Print()
	list.RemoveFromLast()
	list.Print()
	list.RemoveFromFront()
	list.Print()
}
