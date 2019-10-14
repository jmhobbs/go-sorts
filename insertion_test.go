package sorts

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	input := []int{5, 1, 4, 2, 8}
	expected := []int{1, 2, 4, 5, 8}

	actual := InsertionSort().Sort(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("bad sort.\nexpected: %v\n  actual: %v", expected, actual)
	}
}
