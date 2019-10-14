package sorts

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testSamples("Insertion Sort", InsertionSort(), t)
}
