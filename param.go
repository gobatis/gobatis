package batis

import "github.com/gobatis/gobatis/executor"

// Param This function takes in a name and a value, 
// and returns a struct containing the name-value pair. 
// The struct has two fields: "Name" and "Value". 
// The "Name" field is set to the input name string, 
// and the "Value" field is set to the input value of any type. 
// This function can be useful for generating parameters to be passed into other functions or APIs. 
func Param(name string, value any) executor.NameValue {
	return executor.NameValue{Name: name, Value: value}
}
