package reflects

import "reflect"

func ValueElem(rv reflect.Value) reflect.Value {
	for {
		if rv.Kind() != reflect.Pointer {
			return rv
		}
		rv = rv.Elem()
	}
}

func TypeElem(rt reflect.Type) reflect.Type {
	for {
		if rt.Kind() != reflect.Pointer {
			return rt
		}
		rt = rt.Elem()
	}
}
