package katatest

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Fatalf("%s is not equal to %s!", x, y)
	}
}
