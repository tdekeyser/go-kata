package sorts

func BubbleSort(l []int) []int {
	var swapped bool

	for i := 0; i < len(l)-1; i++ {
		if big, small := l[i], l[i+1]; big > small {
			l[i] = small
			l[i+1] = big
			swapped = true
		}
	}

	if swapped {
		return BubbleSort(l)
	}
	return l
}
