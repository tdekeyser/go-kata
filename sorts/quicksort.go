package sorts

// Divide and conquer algorithm : O(n log n)
func QuickSort(l []int) []int {
	if len(l) <= 1 {
		return l
	}

	pivot := l[0]
	left, right := partition(l[1:], pivot)

	leftResult, rightResult := QuickSort(left), QuickSort(right)

	return append(append(leftResult, pivot), rightResult...)
}

func partition(l []int, pivot int) (left, right []int) {
	for _, e := range l {
		if e < pivot {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}
	return left, right
}
