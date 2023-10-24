package batis

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gobatis/gobatis/reflects"
)

type Params map[string]any

// Param This function takes in a name and a value,
// and returns a struct containing the name-value pair.
// The struct has two fields: "Name" and "Value".
// The "Name" field is set to the input name string,
// and the "Value" field is set to the input value of any type.
// This function can be useful for generating parameters to be passed into other functions or APIs.
func Param(name string, value any) NameValue {
	return NameValue{Name: name, Value: value}
}

func LooseDest(dest any, fields ...string) Dest {
	return Dest{loose: true, dest: dest, fields: fields}
}

type Dest struct {
	loose  bool
	dest   any
	fields []string
}

var mappingPathReg = regexp.MustCompile(`^\$(\.[a-zA-Z]\w*)+$`)

func Extract(v any, path string) (r []any, err error) {
	if !mappingPathReg.MatchString(path) {
		err = fmt.Errorf("invlaid extract path format: %s", path)
		return
	}
	paths := strings.Split(path, ".")[1:]
	if len(paths) > 0 {
		err = extract(&r, reflect.ValueOf(v), paths)
	}
	return
}

func extract(r *[]any, rv reflect.Value, paths []string) (err error) {
	rv = reflects.ValueElem(rv)
	if (rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array) && reflects.TypeElem(rv.Type().Elem()).Kind() != reflect.Struct &&
		rv.Kind() != reflect.Struct {
		err = fmt.Errorf("method Extract() accepts only the struct type and its slice or array form")
		return
	}
	name := paths[0]
	paths = paths[1:]
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		for i := 0; i < rv.Len(); i++ {
			v := rv.Index(i).FieldByName(name)
			if !v.IsValid() {
				err = fmt.Errorf("invalid field: %s in type: %s", name, rv.Type().Elem())
				return
			}
			if len(paths) > 0 {
				err = extract(r, v, paths)
				if err != nil {
					return
				}
			} else {
				*r = append(*r, v.Interface())
			}
		}
	} else {
		v := rv.FieldByName(name)
		if !v.IsValid() {
			err = fmt.Errorf("invalid field: %s in type: %s", name, rv.Type())
			return
		}
		if len(paths) > 0 {
			err = extract(r, v, paths)
			if err != nil {
				return
			}
		} else {
			*r = append(*r, v.Interface())
		}
	}
	return
}
