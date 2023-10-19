package commons

import (
	"container/list"
)

type Stack[T any] struct {
	elems []T
	list  list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Peek() (v T) {
	v = s.elems[s.Len()-1]
	return
}

func (s *Stack[T]) Pop() (v T) {
	l := s.Len()
	v = s.elems[l-1]
	s.elems = s.elems[:l-1]
	return v
}

func (s *Stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}

func (s *Stack[T]) Len() int {
	return len(s.elems)
}

func (s *Stack[T]) Index(i int) T {
	return s.elems[i]
}
