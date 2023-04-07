package list

import (
	"errors"
	"fmt"
)

type ArrayList struct {
	values []int
	length int
}

func (arraylist *ArrayList) Init(size int) error {
	if arraylist.length > 0 {
		return errors.New("arrayList already initialized")
	}

	if size > 0 {
		arraylist.values = make([]int, size)
		arraylist.length = 0
		return nil
	} else {
		return errors.New("can't initialize ArrayList with size <= 0")
	}
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

func (arraylist *ArrayList) AddOnIndex(value int, index int) error {
	if index < 0 {
		return errors.New("index cannot be less than zero")
	} else if index > arraylist.length {
		return errors.New("index cannot be greater than the length of the list")
	} // TODO: Would it be good to run recursively until the index is valid? index > 2*len(arraylist.values)

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
	if index < 0 {
		return errors.New("index cannot be less than zero")
	} else if index >= arraylist.length {
		return errors.New("index cannot be greater than or equal to the length of the list")
	} else if arraylist.length <= 0 {
		return errors.New("there is no element to be removed")
	}

	for i := index; i < arraylist.length-1; i++ {
		arraylist.values[i] = arraylist.values[i+1]
	}
	arraylist.length--

	return nil
}

func (arraylist *ArrayList) Get(index int) (int, error) {
	if index < 0 {
		return index, errors.New("index cannot be less than zero")
	} else if index > arraylist.length {
		return index, errors.New("index cannot be greater than to the length of the list")
	} else if arraylist.length <= 0 {
		return index, errors.New("there is no element to get")
	}

	return arraylist.values[index], nil
}

func (arraylist *ArrayList) Set(value int, index int) error {
	if index < 0 {
		return errors.New("index cannot be less than zero")
	} else if index > arraylist.length {
		return errors.New("index cannot be greater than to the length of the list")
	} else if arraylist.length <= 0 {
		return errors.New("there is no element to set")
	}

	arraylist.values[index] = value

	return nil
}

func (arraylist *ArrayList) Size() int {
	return arraylist.length
}

func (arrayList *ArrayList) ShowList() {
	fmt.Printf(".:: Element: %T\n", arrayList)
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
