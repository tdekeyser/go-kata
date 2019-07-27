package sorts

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	some_list := []int{5, 6, 3, 4, 1, 0, 9, 8, 7, 2}
	target_list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	assertArraysEqual(t, QuickSort([]int{}), []int{})
	assertArraysEqual(t, QuickSort([]int{2, 1, 0}), []int{0, 1, 2})
	assertArraysEqual(t, QuickSort(some_list), target_list)
}

func TestBubblesort(t *testing.T) {
	some_list := []int{5, 6, 3, 4, 1, 0, 9, 8, 7, 2}
	target_list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	BubbleSort(some_list) // in-place algorithm

	assertArraysEqual(t, some_list, target_list)
}

func assertArraysEqual(t *testing.T, x, y []int) {
	if !reflect.DeepEqual(x, y) {
		t.Errorf("Not equal: %d %d", x, y)
	}
}
