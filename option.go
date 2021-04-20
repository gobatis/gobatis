package gobatis

func defaultOption() *Option {
	return &Option{
		ConvertNullToZeroValue: true,
	}
}

type Option struct {
	ConvertNullToZeroValue bool
	Host                   string
	Port                   uint16
	User                   string
	Password               string
	Database               string
}
