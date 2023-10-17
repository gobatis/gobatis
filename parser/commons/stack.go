package commons

import "container/list"

type Stack[T any] struct {
	elems []T
	list  list.List
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Peek() (v T, ok bool) {
	if l := s.Len(); l > 0 {
		v = s.elems[l-1]
		ok = true
		return
	}
	return
}

func (s *Stack[T]) Pop() (v T, ok bool) {
	if l := s.Len(); l > 0 {
		v = s.elems[l-1]
		s.elems = s.elems[:l-1]
		ok = true
		return
	}
	return
}

func (s *Stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}

func (s *Stack[T]) Len() int {
	return len(s.elems)
}
