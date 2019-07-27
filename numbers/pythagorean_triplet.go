package numbers

import (
	"sort"
)

// Runtime
// O(n + n + (n*(n+1)/2)) = O(n^2), with n = len(arr)
func ContainsPythagoreanTriplet(arr []int) bool {
	square(arr)
	sort.Ints(arr)

	for i := len(arr) - 1; i > 1; i-- {
		if l, r := 0, i-1; foundSumInTheMiddle(l, r, i, arr) {
			return true
		}
	}
	return false
}

func square(arr []int) {
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
}

// Runtime
// O(right - left)
func foundSumInTheMiddle(left, right, target int, arr []int) bool {
	for left != right {

		sum := arr[left] + arr[right]

		if sum == arr[target] {
			return true
		}

		if sum < arr[target] {
			left++
		} else {
			right--
		}
	}
	return false
}
