package list

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	result := doubleLinkedList.Size()
	expected := 0

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAdd(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	doubleLinkedList.Add(1)

	result := doubleLinkedList.Size()
	expected := 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAddOnIndexWithValidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	doubleLinkedList.AddOnIndex(0, 1)
	doubleLinkedList.AddOnIndex(1, 2)
	doubleLinkedList.AddOnIndex(2, 3)

	result, _ := doubleLinkedList.Get(1)
	expected := 2

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAddOnIndexWithInvalidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	result, expected := doubleLinkedList.AddOnIndex(-584, -999), "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveWithEmptyList(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	result, expected := doubleLinkedList.Remove(), "there is no element to be removed"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveWithPopulatedList(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()
	doubleLinkedList.Add(1)

	result, expected := doubleLinkedList.Remove(), error(nil)

	if result != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveOnIndexWithValidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	doubleLinkedList.Add(10)
	doubleLinkedList.Add(50)

	doubleLinkedList.RemoveOnIndex(1)
	result, expected := doubleLinkedList.Size(), 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestRemoveOnIndexWithInvalidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	result, expected := doubleLinkedList.RemoveOnIndex(-1), "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestGetWithValidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()
	doubleLinkedList.Add(1)

	result, _ := doubleLinkedList.Get(0)
	expected := 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestGetWithInvalidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()
	doubleLinkedList.Add(1)

	_, result := doubleLinkedList.Get(-1)
	expected := "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestSetWithValidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()
	doubleLinkedList.Add(1)

	doubleLinkedList.Set(0, 10)

	result, _ := doubleLinkedList.Get(0)
	expected := 10

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestSetWithInvalidIndex(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()
	doubleLinkedList.Add(1)

	result := doubleLinkedList.Set(-1, 10)
	expected := "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestSizeWithEmptyList(t *testing.T) {
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.Init()

	result := doubleLinkedList.Size()
	expected := 0

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}
