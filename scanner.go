package gobatis

type Scanner struct {
	err error
	ctx *Context
}

func (s Scanner) Scan(ptr ...any) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s Scanner) Error() error {
	return s.err
}

func (s Scanner) AffectRows() (int, error) {
	return 0, nil
}
