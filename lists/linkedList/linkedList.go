package list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head   *Node
	length int
}

type Node struct {
	value int
	next  *Node
}

func (linkedList *LinkedList) Init() {
	linkedList.head = nil
	linkedList.length = 0
}

func (linkedList *LinkedList) Add(value int) {
	if linkedList.head == nil {
		linkedList.head = &Node{value: value, next: nil}
	} else {
		aux := linkedList.head
		for i := 0; aux.next != nil; i++ {
			aux = aux.next
		}

		aux.next = &Node{value: value, next: nil}
	}

	linkedList.length++
}

func (linkedList *LinkedList) insertElementAtStart(value int) {
	if linkedList.head == nil {
		linkedList.head = &Node{value: value, next: nil}
	} else {
		newNode := &Node{value: value, next: linkedList.head}
		linkedList.head = newNode
	}
}

func (linkedList *LinkedList) AddOnIndex(index int, value int) error {
	if index < 0 || index > linkedList.length {
		return errors.New("the index is invalid")
	}

	if index == 0 {
		linkedList.insertElementAtStart(value)
	} else {
		element := linkedList.head
		for i := 0; i < index-1; i++ {
			element = element.next
		}

		newNode := &Node{value: value, next: element.next}
		element.next = newNode
	}

	linkedList.length++

	return nil
}

func (linkedList *LinkedList) Remove() error {
	if linkedList.length <= 0 {
		return errors.New("there is no element to be removed")
	}

	linkedList.length--
	return nil
}

func (linkedList *LinkedList) RemoveOnIndex(index int) error {
	if linkedList.length <= 0 || index < 0 || index > linkedList.length {
		return errors.New("the index is invalid")
	}

	element := linkedList.head
	for i := 0; i < index-1; i++ {
		element = element.next
	}

	aux := element.next
	element.next = aux.next
	linkedList.length--

	return nil
}

func (linkedList *LinkedList) Get(index int) (int, error) {
	if index < 0 || index > linkedList.length {
		return -1, errors.New("the index is invalid")
	}

	element := linkedList.head
	for i := 0; i < index; i++ {
		element = element.next
	}

	return element.value, nil
}

func (linkedList *LinkedList) Set(index int, value int) error {
	if index < 0 || index > linkedList.length {
		return errors.New("the index is invalid")
	}

	element := linkedList.head
	for i := 0; i < index; i++ {
		element = element.next
	}

	element.value = value
	return nil
}

func (linkedList *LinkedList) Size() int {
	return linkedList.length
}

func (linkedList *LinkedList) ShowData() {
	element := linkedList.head
	fmt.Println(".:: Length:", linkedList.length)
	fmt.Print(".:: Values: [")
	for i := 0; i < linkedList.length; i++ {
		fmt.Print(element.value)
		if i < linkedList.length-1 {
			fmt.Print(", ")
		}
		element = element.next
	}
	fmt.Println("]")
}
