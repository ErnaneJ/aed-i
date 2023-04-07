package list

import (
	"errors"
	"fmt"
)

func (doublyLinkedList *DoublyLinkedList) Init() {
	doublyLinkedList.head = nil
	doublyLinkedList.tail = nil
	doublyLinkedList.length = 0
}

func (doublyLinkedList *DoublyLinkedList) Add(value int) {
	if doublyLinkedList.head == nil {
		newNode := &NodeDoublyLinkedList{value: value, next: nil, prev: nil}
		doublyLinkedList.head = newNode
		doublyLinkedList.tail = newNode
	} else {
		aux := doublyLinkedList.tail

		newNode := &NodeDoublyLinkedList{value: value, next: nil, prev: aux}
		aux.next = newNode
		doublyLinkedList.tail = newNode
	}

	doublyLinkedList.length++
}

func (doublyLinkedList *DoublyLinkedList) insertElementAtStart(value int) {
	if doublyLinkedList.head == nil {
		newNode := NodeDoublyLinkedList{value: value, next: nil}
		doublyLinkedList.head = &newNode
		doublyLinkedList.tail = &newNode
	} else {
		newNode := &NodeDoublyLinkedList{value: value, next: doublyLinkedList.head, prev: nil}
		doublyLinkedList.head = newNode
	}
}

func (doublyLinkedList *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index < 0 || index > doublyLinkedList.length {
		return errors.New("the index is invalid")
	}

	if index == 0 {
		doublyLinkedList.insertElementAtStart(value)
	} else {
		element := doublyLinkedList.head
		for i := 0; i < index-1; i++ {
			element = element.next
		}

		newNode := &NodeDoublyLinkedList{value: value, next: element.next, prev: element}
		element.next = newNode
	}

	doublyLinkedList.length++

	return nil
}

func (doublyLinkedList *DoublyLinkedList) Remove() error {
	if doublyLinkedList.length <= 0 {
		return errors.New("there is no element to be removed")
	}

	if doublyLinkedList.length == 1 {
		doublyLinkedList.head = nil
		doublyLinkedList.tail = nil
		doublyLinkedList.length--
		return nil
	}

	doublyLinkedList.tail = doublyLinkedList.tail.prev
	doublyLinkedList.tail.next = nil
	doublyLinkedList.length--
	return nil
}

func (doublyLinkedList *DoublyLinkedList) RemoveOnIndex(index int) error {
	if doublyLinkedList.length <= 0 || index < 0 || index > doublyLinkedList.length {
		return errors.New("the index is invalid")
	}

	if index == 0 {
		doublyLinkedList.head = doublyLinkedList.head.next
	} else {
		itemToDelete := doublyLinkedList.head
		previousItem := doublyLinkedList.head
		for i := 0; i < index; i++ {
			previousItem = itemToDelete
			itemToDelete = itemToDelete.next
		}
		previousItem.next = itemToDelete.next
		if itemToDelete.next != nil {
			itemToDelete.next.prev = previousItem
		}
	}

	doublyLinkedList.length--
	return nil
}

func (doublyLinkedList *DoublyLinkedList) Get(index int) (int, error) {
	if index < 0 || index > doublyLinkedList.length {
		return -1, errors.New("the index is invalid")
	}

	element := doublyLinkedList.head
	for i := 0; i < index; i++ {
		element = element.next
	}

	return element.value, nil
}

func (doublyLinkedList *DoublyLinkedList) Set(value int, index int) error {
	if index < 0 || index > doublyLinkedList.length {
		return errors.New("the index is invalid")
	}

	element := doublyLinkedList.head
	for i := 0; i < index; i++ {
		element = element.next
	}

	element.value = value
	return nil
}

func (doublyLinkedList *DoublyLinkedList) Size() int {
	return doublyLinkedList.length
}

func (doublyLinkedList *DoublyLinkedList) ShowList() {
	element := doublyLinkedList.head
	fmt.Printf(".:: Element: %T\n", doublyLinkedList)
	fmt.Println(".:: Length:", doublyLinkedList.length)
	fmt.Print(".:: Values: [")
	for i := 0; i < doublyLinkedList.length; i++ {
		fmt.Print(element.value)
		if i < doublyLinkedList.length-1 {
			fmt.Print(", ")
		}
		element = element.next
	}
	fmt.Println("]")
}
