package main

// Stack a simple byte stack
type Stack struct {
	// the stack
	stack []byte
	// the stack pointer
	sp int
}

// NewStack create a new stack
func NewStack() *Stack {
	return &Stack{
		stack: make([]byte, 0),
		sp:    0,
	}
}

// Push a byte onto the stack
func (s *Stack) Push(b byte) {
	s.stack = append(s.stack, b)
	s.sp++
}

// Pop a byte off the stack
func (s *Stack) Pop() byte {
	s.sp--

	poppedValue := s.stack[s.sp]

	// remove the last element
	s.stack = s.stack[:s.sp]

	return poppedValue
}
