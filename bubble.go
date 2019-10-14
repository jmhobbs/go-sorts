package sorts

type bubblesort struct{}

func BubbleSort() Sort {
	return bubblesort{}
}

func (bs bubblesort) Sort(in []int) []int {
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(in)-1; i++ {
			if in[i] > in[i+1] {
				h := in[i]
				in[i] = in[i+1]
				in[i+1] = h
				changed = true
			}
		}
	}

	return in
}
