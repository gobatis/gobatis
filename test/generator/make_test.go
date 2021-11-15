package generator

import "testing"

func TestNewMaker(t *testing.T) {
	maker := NewMaker(Postgresql)
	maker.MakeEntity()
}
