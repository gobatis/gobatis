package cast

import "reflect"

type Value struct {
	reflect.Value
}

func (p Value) Add(a Value) (err error) {
	
	switch p.Kind() {
	case reflect.Int8:
	
	}
	
	return
}
