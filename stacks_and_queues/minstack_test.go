package stacks_and_queues

import (
	. "github.com/tdekeyser/go-kata/katatest"
	"testing"
)

type stackInt int

func (a stackInt) compare(b interface{}) int {
	i, ok := b.(stackInt)
	if !ok {
		return 1
	}

	if a < i {
		return -1
	}
	if a > i {
		return 1
	}
	return 0
}

func Test_len(t *testing.T) {
	s := new(stack)

	s.push(stackInt(10))
	s.push(stackInt(9))

	AssertEqual(t, s.len(), 2)
}

func Test_push_peek(t *testing.T) {
	s := new(stack)

	s.push(stackInt(10))
	s.push(stackInt(5))

	AssertEqual(t, s.peek(), stackInt(5))

	s.push(stackInt(24))

	AssertEqual(t, s.len(), 3)
	AssertEqual(t, s.peek(), stackInt(24))
}

func Test_push_pop(t *testing.T) {
	s := new(stack)

	s.push(stackInt(3))
	s.push(stackInt(1))
	s.push(stackInt(24))

	AssertEqual(t, s.len(), 3)
	AssertEqual(t, s.pop(), stackInt(24))
	AssertEqual(t, s.len(), 2)
	AssertEqual(t, s.pop(), stackInt(1))
}

func Test_push_getMin(t *testing.T) {
	s := new(stack)

	s.push(stackInt(3))
	s.push(stackInt(24))

	AssertEqual(t, s.getMin(), stackInt(3))

	s.push(stackInt(1))

	AssertEqual(t, s.len(), 3)
	AssertEqual(t, s.getMin(), stackInt(1))
}
