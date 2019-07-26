package numbers

import "testing"

func Test_k_largest(t *testing.T) {
	assertEqual(t,
		KLargest(3, []int{1, 23, 12, 9, 30, 2, 50}),
		[]int{50, 30, 23})
}
