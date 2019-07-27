package numbers

import (
	. "github.com/tdekeyser/go-kata/katatest"
	"testing"
)

func Test_ContainsPythagoreanTriplet(t *testing.T) {
	AssertEqual(t, ContainsPythagoreanTriplet([]int{3, 1, 4, 6, 5}), true)
	AssertEqual(t, ContainsPythagoreanTriplet([]int{10, 4, 6, 12, 5}), false)
}
