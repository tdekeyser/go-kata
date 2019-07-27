package numbers

import (
	. "github.com/tdekeyser/go-kata/katatest"
	"testing"
)

func Test_ContainsPythagoreanTriplet(t *testing.T) {
	AssertEqual(t, containsPythagoreanTriplet([]int{3, 1, 4, 6, 5}), true)
	AssertEqual(t, containsPythagoreanTriplet([]int{10, 4, 6, 12, 5}), false)
}
