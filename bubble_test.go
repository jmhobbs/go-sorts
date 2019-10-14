package sorts

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	testSamples("Bubble Sort", BubbleSort(), t)
}
