package list

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	values []int
	length int
}

func (arraylist *ArrayList) Init() {
	arraylist.values = make([]int, 10)
}

func (arraylist *ArrayList) double() {
	doubledValues := make([]int, len(arraylist.values)*2)
	for i := 0; i < len(arraylist.values); i++ {
		doubledValues[i] = arraylist.values[i]
	}
	arraylist.values = doubledValues
}

func (arraylist *ArrayList) Add(value int) {
	if arraylist.length == len(arraylist.values) {
		arraylist.double()
	}
	arraylist.values[arraylist.length] = value
	arraylist.length++
}

func (arraylist *ArrayList) AddOnIndex(index int, value int) error {
	if index < 0 || index > 2*len(arraylist.values) {
		return errors.New("the index is invalid")
	}

	if arraylist.length == len(arraylist.values) {
		arraylist.double()
	}

	for i := arraylist.length; i > index; i-- {
		arraylist.values[i] = arraylist.values[i-1]
	}

	arraylist.values[index] = value
	arraylist.length++

	return nil
}

func (arraylist *ArrayList) Remove() error {
	if arraylist.length <= 0 {
		return errors.New("there is no element to be removed")
	}

	arraylist.values[arraylist.Size()-1] = 0
	arraylist.length--
	return nil
}

func (arraylist *ArrayList) RemoveOnIndex(index int) error {
	if index < 0 || index >= arraylist.length || arraylist.length <= 0 {
		return errors.New("the index is invalid")
	}

	for i := index; i < arraylist.length; i++ {
		arraylist.values[i] = arraylist.values[i+1]
	}
	arraylist.length--

	return nil
}

func (arraylist *ArrayList) Get(index int) (int, error) {
	if index < 0 || index > arraylist.length {
		return -1, errors.New("the index is invalid")
	}

	return arraylist.values[index], nil
}

func (arraylist *ArrayList) Set(index int, value int) error {
	if index < 0 || index > arraylist.length {
		return errors.New("the index is invalid")
	}

	arraylist.values[index] = value

	return nil
}

func (arraylist *ArrayList) Size() int {
	return arraylist.length
}

func (arrayList *ArrayList) ShowData() {
	fmt.Println(".:: Length:", arrayList.length)
	fmt.Print(".:: Values: [")
	for i := 0; i < arrayList.length; i++ {
		fmt.Print(arrayList.values[i])
		if i < arrayList.length-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}
