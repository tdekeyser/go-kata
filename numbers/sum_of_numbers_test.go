package numbers

import (
	. "github.com/tdekeyser/go-kata/katatest"
	"testing"
)

func TestSumOfNumbers_zero(t *testing.T) {
	AssertEqual(t, SumOfNumbers(0), 0)
}

func TestSumOfNumbers_singleDigit(t *testing.T) {
	AssertEqual(t, SumOfNumbers(9), 9)
}

func TestSumOfNumbers_multipleDigits(t *testing.T) {
	AssertEqual(t, SumOfNumbers(345), 12)
}

func TestSumOfNumbers_negative(t *testing.T) {
	AssertEqual(t, SumOfNumbers(-345), 12)
}
