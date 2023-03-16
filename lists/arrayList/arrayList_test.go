package list

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	result := arrayList.Size()
	expected := 0

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAdd(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	arrayList.Add(1)

	result := arrayList.Size()
	expected := 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAddWithFullList(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	for i := 0; i < 10; i++ {
		arrayList.Add(1)
	}

	arrayList.Add(1)

	result := arrayList.Size()
	expected := 11

	if result != expected && len(arrayList.values) == 20 {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestAddOnIndexWithValidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	arrayList.AddOnIndex(1, -999)

	resultCase1, expectedCase1 := arrayList.Size(), 1

	resultCase2, _ := arrayList.Get(1)
	expectedCase2 := -999

	if resultCase1 != expectedCase1 || resultCase2 != expectedCase2 {
		t.Errorf("The result was [%d, %d] but the expected was [%d, %d].", resultCase1, resultCase2, expectedCase1, expectedCase2)
	}
}

func TestAddOnIndexWithInvalidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	result, expected := arrayList.AddOnIndex(-584, -999), "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveOnIndexWithValidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	arrayList.Add(10)
	arrayList.Add(50)

	arrayList.RemoveOnIndex(1)
	result, expected := arrayList.Size(), 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestRemoveOnIndexWithInvalidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	result, expected := arrayList.RemoveOnIndex(-1), "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveWithEmptyList(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	result, expected := arrayList.Remove(), "there is no element to be removed"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestRemoveWithPopulatedList(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	arrayList.Add(1)

	result, expected := arrayList.Remove(), error(nil)

	if result != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestGetWithValidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	arrayList.Add(1)

	result, _ := arrayList.Get(0)
	expected := 1

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestGetWithInvalidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	arrayList.Add(1)

	_, result := arrayList.Get(-1)
	expected := "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestSetWithValidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	arrayList.Add(1)

	arrayList.Set(0, 10)

	result, _ := arrayList.Get(0)
	expected := 10

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}

func TestSetWithInvalidIndex(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()
	arrayList.Add(1)

	result := arrayList.Set(-1, 10)
	expected := "the index is invalid"

	if fmt.Sprint(result) != expected {
		t.Errorf("The result was %s but the expected was %s.", result, expected)
	}
}

func TestSizeWithEmptyList(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Init()

	result := arrayList.Size()
	expected := 0

	if result != expected {
		t.Errorf("The result was %d but the expected was %d.", result, expected)
	}
}
