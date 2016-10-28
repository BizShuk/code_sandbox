package stack

type stack struct {
	stack []interface{}
}

func (s *stack) Push(e interface{}) error {
	s = append(s.stack, e)
	return nil
}

func (s *stack) Pop() error {
	if s.Size() <= 0 {
		return error.New("no size...")
	}
	len := len(s.stack) - 1
	popd := s.stack[len]
	s = s.stack[0:len]
	return popd
}

func (s *stack) Size() int {
	return len(s.stack)
}

func (s *stack) Peek() interface{} {
	var last interface{}
	last = s.stack[len(s.stack)-1]

	return last
}

func (s *stack) Reverse() {
	last := s.Pop()
	s.Reverse()
	s.Push(last)
}
