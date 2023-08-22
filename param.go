package batis

type KeyValue struct {
	Name  string
	Value any
}

// Param This function takes in a name and a value, 
// and returns a struct containing the name-value pair. 
// The struct has two fields: "Name" and "Value". 
// The "Name" field is set to the input name string, 
// and the "Value" field is set to the input value of any type. 
// This function can be useful for generating parameters to be passed into other functions or APIs. 
func Param(name string, value any) KeyValue {
	return KeyValue{Name: name, Value: value}
}

func Select(data any, columns string) Rows {
	return &selectColumns{
		columns: columns,
		data:    data,
	}
}

func Except(data any, columns string) Rows {
	return &exceptColumns{
		data:    data,
		columns: columns,
	}
}
