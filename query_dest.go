package batis

type queryDest struct {
	sql    string
	params []NameValue
	dest   any
}
