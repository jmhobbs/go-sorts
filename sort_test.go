package sorts

import (
	"fmt"
	"reflect"
	"testing"
)

func testSamples(name string, sort Sort, t *testing.T) {
	samples := []struct {
		Input  []int
		Sorted []int
	}{
		{
			[]int{5, 1, 5, 4, 2, 8},
			[]int{1, 2, 4, 5, 5, 8},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1, 1, 1, 1, 1},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
	}

	for i, sample := range samples {
		t.Run(fmt.Sprintf("%s:%d", name, i), func(t *testing.T) {
			actual := sort.Sort(sample.Input)

			if !reflect.DeepEqual(sample.Sorted, actual) {
				t.Errorf("bad sort.\nexpected: %v\n  actual: %v", sample.Sorted, actual)
			}
		})
	}
}
