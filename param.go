package gobatis

type NameValue struct {
	name  string
	value any
}

func Param(name string, value any) NameValue {
	return NameValue{name: name, value: value}
}
