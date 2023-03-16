package list

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	result := linkedList.Size()
	expected := 0

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAdd(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	linkedList.Add(1)

	result := linkedList.Size()
	expected := 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAddOnIndexWithValidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	linkedList.AddOnIndex(0, 1)
	linkedList.AddOnIndex(1, 2)
	linkedList.AddOnIndex(2, 3)

	result, _ := linkedList.Get(1)
	expected := 2

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAddOnIndexWithInvalidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	result, expected := linkedList.AddOnIndex(-584, -999), "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveWithEmptyList(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	result, expected := linkedList.Remove(), "there is no element to be removed"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveWithPopulatedList(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()
	linkedList.Add(1)

	result, expected := linkedList.Remove(), error(nil)

	if result != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveOnIndexWithValidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	linkedList.Add(10)
	linkedList.Add(50)

	linkedList.RemoveOnIndex(1)
	result, expected := linkedList.Size(), 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestRemoveOnIndexWithInvalidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	result, expected := linkedList.RemoveOnIndex(-1), "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestGetWithValidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()
	linkedList.Add(1)

	result, _ := linkedList.Get(0)
	expected := 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestGetWithInvalidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()
	linkedList.Add(1)

	_, result := linkedList.Get(-1)
	expected := "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestSetWithValidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()
	linkedList.Add(1)

	linkedList.Set(0, 10)

	result, _ := linkedList.Get(0)
	expected := 10

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestSetWithInvalidIndex(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()
	linkedList.Add(1)

	result := linkedList.Set(-1, 10)
	expected := "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestSizeWithEmptyList(t *testing.T) {
	linkedList := LinkedList{}
	linkedList.Init()

	result := linkedList.Size()
	expected := 0

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}
