package sorts

type insertionsort struct{}

func InsertionSort() insertionsort {
	return insertionsort{}
}

func (s insertionsort) Sort(in []int) []int {
	for i := 1; i < len(in); i++ {
		if in[i-1] <= in[i] {
			continue
		}
		for j := 0; j < i; j++ {
			if in[j] > in[i] {
				out := []int{}
				out = append(out, in[0:j]...)
				out = append(out, in[i])
				out = append(out, in[j:i]...)
				out = append(out, in[i+1:]...)
				in = out
				break
			}
		}
	}
	return in
}
