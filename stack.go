package tdd_go

import "errors"

type stack struct {
	size     int
	elements []int
}

func (s *stack) Pop() (int, error) {
	if s.size == 0 {
		return -1, errors.New("stack underflow")
	}
	s.size -= 1
	popped := s.elements[s.size]
	s.elements = s.elements[:len(s.elements)-1]
	return popped, nil
}

func (s stack) IsEmpty() bool {
	return s.size == 0
}

func (s *stack) Push(element int) {
	s.elements = append(s.elements, element)
	s.size += 1
}

func (s *stack) Size() int {
	return s.size
}

func NewStack() *stack {
	return &stack{
		size: 0,
	}
}
