package batis

type Scanner struct {
	ctx *Context
}

func (s Scanner) Scan(ptr ...any) error {
	return nil
}

func (s Scanner) Error() error {
	return nil
}

func (s Scanner) AffectRows() int {
	return 0
}
