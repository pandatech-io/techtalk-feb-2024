package stack

import "errors"

var ErrEmptyStack error = errors.New("stack is empty")

type St struct {
	stack []int
}

func New() *St {
	return &St{
		stack: make([]int, 0),
	}
}

func (s *St) IsEmpty() bool {
	return s.Size() == 0
}

func (s *St) Push(num int) {
	s.stack = append(s.stack, num)
}

func (s *St) Size() int {
	return len(s.stack)
}

func (s *St) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}
	latestIndex := s.Size() - 1
	latestValue := s.stack[latestIndex]
	s.stack = s.stack[:latestIndex]
	return latestValue, nil
}
