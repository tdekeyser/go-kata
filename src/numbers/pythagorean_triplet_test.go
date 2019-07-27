package numbers

import "testing"

func Test_ContainsPythagoreanTriplet(t *testing.T) {
	assertEqual(t, ContainsPythagoreanTriplet([]int{3, 1, 4, 6, 5}), true)
	assertEqual(t, ContainsPythagoreanTriplet([]int{10, 4, 6, 12, 5}), false)
}
