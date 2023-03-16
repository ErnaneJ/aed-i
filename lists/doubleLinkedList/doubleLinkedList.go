package list

import (
	"errors"
	"fmt"
)

type DoubleLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (doubleLinkedList *DoubleLinkedList) Init() {
	doubleLinkedList.head = nil
	doubleLinkedList.tail = nil
	doubleLinkedList.length = 0
}

func (doubleLinkedList *DoubleLinkedList) Add(value int) {
	if doubleLinkedList.head == nil {
		newNode := &Node{value: value, next: nil, prev: nil}
		doubleLinkedList.head = newNode
		doubleLinkedList.tail = newNode
	} else {
		aux := doubleLinkedList.tail

		newNode := &Node{value: value, next: nil, prev: aux}
		aux.next = newNode
		doubleLinkedList.tail = newNode
	}

	doubleLinkedList.length++
}

func (doubleLinkedList *DoubleLinkedList) insertElementAtStart(value int) {
	if doubleLinkedList.head == nil {
		newNode := Node{value: value, next: nil}
		doubleLinkedList.head = &newNode
		doubleLinkedList.tail = &newNode
	} else {
		newNode := &Node{value: value, next: doubleLinkedList.head, prev: nil}
		doubleLinkedList.head = newNode
	}
}

func (doubleLinkedList *DoubleLinkedList) AddOnIndex(index int, value int) error {
	if index < 0 || index > doubleLinkedList.length {
		return errors.New("the index is invalid")
	}

	if index == 0 {
		doubleLinkedList.insertElementAtStart(value)
	} else {
		element := doubleLinkedList.head
		for i := 0; i < index-1; i++ {
			element = element.next
		}

		newNode := &Node{value: value, next: element.next, prev: element}
		element.next = newNode
	}

	doubleLinkedList.length++

	return nil
}

func (doubleLinkedList *DoubleLinkedList) Remove() error {
	if doubleLinkedList.length <= 0 {
		return errors.New("there is no element to be removed")
	}

	if doubleLinkedList.length == 1 {
		doubleLinkedList.head = nil
		doubleLinkedList.tail = nil
		doubleLinkedList.length--
		return nil
	}

	doubleLinkedList.tail = doubleLinkedList.tail.prev
	doubleLinkedList.tail.next = nil
	doubleLinkedList.length--
	return nil
}

func (doubleLinkedList *DoubleLinkedList) RemoveOnIndex(index int) error {
	if doubleLinkedList.length <= 0 || index < 0 || index > doubleLinkedList.length {
		return errors.New("the index is invalid")
	}

	element := doubleLinkedList.head
	for i := 0; i < index-1; i++ {
		element = element.next
	}

	aux := element.next
	element.next = aux.next
	doubleLinkedList.length--

	return nil
}

func (doubleLinkedList *DoubleLinkedList) Get(index int) (int, error) {
	if index < 0 || index > doubleLinkedList.length {
		return -1, errors.New("the index is invalid")
	}

	element := doubleLinkedList.head
	for i := 0; i < index; i++ {
		element = element.next
	}

	return element.value, nil
}

func (doubleLinkedList *DoubleLinkedList) Set(index int, value int) error {
	if index < 0 || index > doubleLinkedList.length {
		return errors.New("the index is invalid")
	}

	element := doubleLinkedList.head
	for i := 0; i < index; i++ {
		element = element.next
	}

	element.value = value
	return nil
}

func (doubleLinkedList *DoubleLinkedList) Size() int {
	return doubleLinkedList.length
}

func (doubleLinkedList *DoubleLinkedList) ShowData() {
	element := doubleLinkedList.head
	fmt.Println(".:: Length:", doubleLinkedList.length)
	fmt.Print(".:: Values: [")
	for i := 0; i < doubleLinkedList.length; i++ {
		fmt.Print(element.value)
		if i < doubleLinkedList.length-1 {
			fmt.Print(", ")
		}
		element = element.next
	}
	fmt.Println("]")
}
