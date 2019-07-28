package stacks_and_queues

type comparable interface {
	compare(c interface{}) int
}

type minstack interface {
	len() int
	push(data comparable)
	peek() comparable
	pop() comparable
	getMin() comparable
}

type stack struct {
	data comparable
	size int
	min  comparable
	next *stack
}

// Runtime
// O(1)
func (s *stack) len() int {
	return s.size
}

// Runtime
// O(1)
func (s *stack) getMin() comparable {
	return s.min
}

func (s *stack) peek() comparable {
	return s.data
}

func (s *stack) pop() comparable {
	d := s.data
	*s = *s.next
	return d
}

func (s *stack) push(data comparable) {
	tmp := *s
	*s = stack{
		data: data,
		size: tmp.size + 1,
		min:  minimum(s.min, data),
		next: &tmp,
	}
}

func minimum(a, b comparable) comparable {
	if a == nil {
		return b
	}
	if a.compare(b) > 0 {
		return b
	}
	return a
}
