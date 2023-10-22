package xsql

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack[int]()
	t.Log(s.Peek())
	s.Push(1)
	t.Log(s.Peek())
	s.Push(2)
	t.Log(s.Peek())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Peek())
}
