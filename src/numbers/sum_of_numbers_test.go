package numbers

import (
	"reflect"
	"testing"
)

func TestSumOfNumbers_zero(t *testing.T) {
	assertEqual(t, SumOfNumbers(0), 0)
}

func TestSumOfNumbers_singleDigit(t *testing.T) {
	assertEqual(t, SumOfNumbers(9), 9)
}

func TestSumOfNumbers_multipleDigits(t *testing.T) {
	assertEqual(t, SumOfNumbers(345), 12)
}

func TestSumOfNumbers_negative(t *testing.T) {
	assertEqual(t, SumOfNumbers(-345), 12)
}

func assertEqual(t *testing.T, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Fatalf("%s is not equal to %s!", x, y)
	}
}
