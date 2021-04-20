package gobatis

func defaultOption() *Option {
	return &Option{
		ConvertNullToZeroValue: true,
	}
}

type Option struct {
	ConvertNullToZeroValue bool
}
