package batis

import "github.com/gobatis/gobatis/executor"

type Params map[string]any

// Param This function takes in a name and a value,
// and returns a struct containing the name-value pair.
// The struct has two fields: "Name" and "Value".
// The "Name" field is set to the input name string,
// and the "Value" field is set to the input value of any type.
// This function can be useful for generating parameters to be passed into other functions or APIs.
func Param(name string, value any) executor.Param {
	return executor.Param{Name: name, Value: value}
}

func LooseDest(dest any, fields ...string) Dest {
	return Dest{loose: true, dest: dest, fields: fields}
}

type Dest struct {
	loose  bool
	dest   any
	fields []string
}

func Extract(v any, path string) (r any, err error) {

	return
}
