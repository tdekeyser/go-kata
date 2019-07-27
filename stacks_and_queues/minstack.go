package stacks_and_queues

type MinStack interface {
	len() int
	push(data interface{})
	peek() interface{}
	pop() interface{}
	getMin() interface{}
}

type Stack struct {
	data interface{}
	next *Stack
}

func New() *Stack {
	return &Stack{nil, nil}
}

func (s *Stack) len() int {
	if s.next == nil {
		return 1
	}
	return 1 + s.next.len()
}
