package postgresql

type Mapper struct {
	ResetTable func() error
	*MakeMapper
}
