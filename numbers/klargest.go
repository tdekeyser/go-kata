package numbers

import "sort"

// Runtime
// O(n*log(n)), for n = len(arr)
func kLargest(k int, arr []int) []int {
	sort.Ints(arr)

	r := make([]int, 3)

	for i := 0; i < k; i++ {
		r[i] = arr[len(arr)-i-1]
	}

	return r
}
