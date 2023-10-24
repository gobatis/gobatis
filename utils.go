package batis

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	//"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dialector"
	"github.com/stretchr/testify/require"
	"github.com/ttacon/chalk"
)

var errorType reflect.Type
var scannerType reflect.Type
var valuerType reflect.Type

type Valuer interface {
	Value() (driver.Value, error)
}

func init() {
	errorType = reflect.TypeOf((*error)(nil)).Elem()
	scannerType = reflect.TypeOf((*sql.Scanner)(nil)).Elem()
	valuerType = reflect.TypeOf((*Valuer)(nil)).Elem()
}

func isContext(v reflect.Type) bool {
	if v.Name() == "Context" && v.PkgPath() == "context" {
		return true
	}
	return false
}

func isTx(v reflect.Type) bool {
	if v.Kind() == reflect.Ptr && v.Elem().Name() == "Tx" &&
		(v.Elem().PkgPath() == "database/sql" ||
			v.Elem().PkgPath() == "github.com/gobatis/gobatis") {
		return true
	}
	return false
}

func isDB(v reflect.Type) bool {
	if v.Kind() == reflect.Ptr && v.Elem().Name() == "DB" && v.Elem().PkgPath() == "github.com/gobatis/gobatis" {
		return true
	}
	return false
}

func isError(t reflect.Type) bool {
	return t.Implements(reflect.TypeOf((*error)(nil)).Elem())
}

func isStructSlice(r reflect.Type) bool {
	if r.Kind() != reflect.Slice {
		return false
	}
	elem := r.Elem()
	for {
		if elem.Kind() != reflect.Ptr {
			break
		}
		elem = elem.Elem()
	}
	return elem.Kind() == reflect.Struct
}

func SplitStructSlice(data any, limit int) ([]any, error) {
	val := reflect.ValueOf(data)

	if !isStructSlice(val.Type()) {
		return nil, fmt.Errorf("expected a struct slice, got %s", val.Type())
	}

	sliceLen := val.Len()
	if sliceLen == 0 {
		return nil, fmt.Errorf("got empty slice")
	}

	var result []any
	for i := 0; i < sliceLen; i += limit {
		end := i + limit
		if end > sliceLen {
			end = sliceLen
		}

		subSlice := val.Slice(i, end)
		result = append(result, subSlice.Interface())
	}

	return result, nil
}

func reflectValueElem(vt reflect.Value) reflect.Value {
	for {
		if vt.Kind() != reflect.Ptr {
			break
		}
		vt = vt.Elem()
	}
	return vt
}

func reflectTypeElem(vt reflect.Type) reflect.Type {
	for {
		if vt.Kind() != reflect.Ptr {
			break
		}
		vt = vt.Elem()
	}
	return vt
}

func printVars(vars []interface{}) string {
	if len(vars) == 0 {
		return ""
	}
	r := "\n"
	for i, v := range vars {
		_t := ""
		if v != nil {
			_t = reflect.TypeOf(v).String()
		}
		r += fmt.Sprintf("   $%d %s (%s) %+v\n",
			i+1, chalk.Green.Color("=>"), chalk.Yellow.Color(_t), v)
	}
	return r
}

//func toSnakeCase(s string) string {
//	var re = regexp.MustCompile(`([^A-Z_])([A-Z])`)
//	snakeStr := re.ReplaceAllString(s, "${1}_${2}")
//	return strings.ToLower(snakeStr)
//}

//func Extract[T any, V any](items []T, fn func(item T) V) []V {
//	var r []V
//	for _, v := range items {
//		r = append(r, fn(v))
//	}
//	return r
//}

func reflectRow(columns []string, row []interface{}, pv reflect.Value, first bool) (bool, error) {

	switch pv.Kind() {
	case reflect.Slice, reflect.Array:
		return false, setArray(pv, newRowMap(columns, row))
	case reflect.Struct:
		return true, setStruct(pv, newRowMap(columns, row))
	}
	return true, setValue(pv, row[0])
}

func prepareFieldName(f reflect.StructField) string {
	field := f.Tag.Get("db")
	if field == "" {
		field = toSnakeCase(f.Name)
	}
	return trimComma(field)
}

func trimComma(field string) string {
	if strings.Contains(field, ",") {
		return strings.TrimSpace(strings.Split(field, ",")[0])
	}
	return field
}

func toSnakeCase(s string) string {
	var re = regexp.MustCompile(`([^A-Z_])([A-Z])`)
	snakeStr := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snakeStr)
}

func newRowMap(columns []string, values []interface{}) rowMap {
	m := rowMap{}
	for i, v := range columns {
		m[v] = values[i]
	}
	return m
}

type rowMap map[string]interface{}

//func reflectStructs(r rowMap, ptr reflect.Value) error {
//	var _type reflect.Type
//	if ptr.Type().Elem().Kind() != reflect.Ptr {
//		// var test []Test => Test
//		_type = ptr.Type().Elem().Elem()
//	} else {
//		// var test []*Test => Test
//		_type = ptr.Type().Elem().Elem()
//	}
//	elem := reflect.New(_type)
//	for i := 0; i < _type.NumField(); i++ {
//		field := _type.Field(i).Tag.Get("db")
//		field = trimComma(field)
//		v, ok := r[field]
//		if ok && v != nil {
//			if elem.Elem().Field(i).Kind() == reflect.Ptr {
//				elem.Elem().Field(i).Set(reflect.New(elem.Elem().Field(i).Type().Elem()))
//			}
//			err := SetValue(elem.Elem().Field(i), v)
//			if err != nil {
//				return err
//			}
//		}
//	}
//	if ptr.Type().Elem().Elem().Kind() != reflect.Ptr {
//		ptr.Elem().Set(reflect.Append(ptr.Elem(), elem.Elem()))
//	} else {
//		ptr.Elem().Set(reflect.Append(ptr.Elem(), elem))
//	}
//	return nil
//}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "gobatis: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "gobatis: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "gobatis: Unmarshal(nil " + e.Type.String() + ")"
}

func setValue(pv reflect.Value, v any) error {

	if t := reflect.New(pv.Type()); t.Type().Implements(scannerType) {
		s := t.Interface().(sql.Scanner)
		err := s.Scan(v)
		if err != nil {
			return err
		}
		pv.Set(t.Elem())
		return nil
	}

	vv := reflect.ValueOf(v)

	switch vv.Kind() {
	case reflect.Bool:
		return setBool(pv, vv)
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64:
		return setNumber(pv, vv)
	case reflect.String:
		return setString(pv, vv)
	}

	switch pv.Interface().(type) {
	case time.Time:
		switch v.(type) {
		case time.Time:
			pv.Set(vv)
		default:
			panic("todo")
		}
	}

	return nil
}

func setArray(pv reflect.Value, r rowMap) (err error) {

	t := pv.Type().Elem()
	ptr := false
	if t.Kind() == reflect.Ptr {
		ptr = true
		for {
			if t.Kind() != reflect.Ptr {
				break
			}
			t = t.Elem()
		}
	}
	switch t.Kind() {
	case reflect.Struct:
		if ptr {
			pv.Set(reflect.Append(pv, reflect.New(pv.Type().Elem().Elem())))
		} else {
			pv.Set(reflect.Append(pv, reflect.New(pv.Type().Elem()).Elem()))
		}
		err = setStruct(indirect(pv.Index(pv.Len()-1), false), r)
		if err != nil {
			return
		}
	default:
		// TODO 反射数组单值
		if ptr {
			pv.Set(reflect.Append(pv, reflect.New(pv.Type().Elem().Elem())))
		} else {
			pv.Set(reflect.Append(pv, reflect.New(pv.Type().Elem()).Elem()))
		}
		var a any
		for _, vv := range r {
			a = vv
		}
		err = setValue(indirect(pv.Index(pv.Len()-1), false), a)
		if err != nil {
			return
		}
		//err = fmt.Errorf("expect struct, got: %s", t.String())
		return
	}

	return
}

func setStruct(pv reflect.Value, r rowMap) (err error) {

	//var tags map[string]struct{}
	//if first {
	//	tags = map[string]struct{}{}
	//}

	t := pv.Type()
	for i := 0; i < t.NumField(); i++ {
		n := prepareFieldName(t.Field(i))
		//if first {
		//	if _, ok := tags[n]; ok {
		//		return fmt.Errorf("field tag: '%s' is duplicated in struct: '%s'", n, _type)
		//	}
		//	tags[n] = struct{}{}
		//}
		v, ok := r[n]
		if !ok {
			// TODO
			//if !false {
			//	return fmt.Errorf("no data for struct: '%s' field: '%s'", _type, _type.Field(i).Name)
			//}
		} else if v != nil {
			err = setValue(indirect(pv.Field(i), false), v)
			if err != nil {
				err = fmt.Errorf("parse column: %s value error: %w", n, err)
				return
			}
		}
	}

	return
}

func setNumber(pv reflect.Value, vv reflect.Value) (err error) {
	pv = indirect(pv, false)
	switch pv.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		var v int64
		if vv.CanInt() {
			v = vv.Int()
		} else if vv.CanUint() {
			v = int64(vv.Uint())
		} else if vv.CanFloat() {
			v = int64(vv.Float())
		}
		pv.SetInt(v)
		err = literalEqual(pv, vv)
		if err != nil {
			return
		}
		return
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		var v uint64
		if vv.CanUint() {
			v = vv.Uint()
		} else if vv.CanInt() {
			v = uint64(vv.Int())
		} else if vv.CanFloat() {
			v = uint64(vv.Float())
		}
		pv.SetUint(v)
		err = literalEqual(pv, vv)
		if err != nil {
			return
		}
		return
	case reflect.Float32,
		reflect.Float64:
		var v float64
		if vv.CanFloat() {
			v = vv.Float()
		} else if vv.CanInt() {
			v = float64(vv.Int())
		} else if vv.CanUint() {
			v = float64(vv.Uint())
		}
		pv.SetFloat(v)
		//err = literalEqual(pv, vv)
		//if err != nil {
		//	return
		//}
		return
	default:
		return fmt.Errorf("unsupport reflect type %s to %s", vv.Type(), pv.Type())
	}
}

func literalEqual(pv, vv reflect.Value) error {
	if fmt.Sprintf("%v", pv.Interface()) != fmt.Sprintf("%v", vv.Interface()) {
		return fmt.Errorf("convert type %s:%v to %s:%v miss data", vv.Type(), vv.Interface(), pv.Type(), pv.Interface())
	}
	return nil
}

func setString(pv reflect.Value, vv reflect.Value) error {
	pv = indirect(pv, false)
	if pv.Kind() == reflect.String {
		pv.Set(vv)
		return nil
	}
	return fmt.Errorf("unsupport reflect type %s to %s", vv.Type(), pv.Type())
}

func setBool(pv reflect.Value, vv reflect.Value) (err error) {
	pv = indirect(pv, false)
	if pv.Kind() == reflect.Bool {
		pv.Set(vv)
		return nil
	}
	return fmt.Errorf("unsupport reflect type %s to %s", vv.Type(), pv.Type())
}

// indirect walks down v allocating pointers as needed,
// until it gets to a non-pointer.
// If decodingNull is true, indirect stops at the first settable pointer so it
// can be set to nil.
func indirect(v reflect.Value, decodingNull bool) reflect.Value {
	// Issue #24153 indicates that it is generally not a guaranteed property
	// that you may round-trip a reflect.Value by calling Value.Addr().Elem()
	// and expect the value to still be settable for values derived from
	// unexported embedded struct fields.
	//
	// The logic below effectively does this when it first addresses the value
	// (to satisfy possible pointer methods) and continues to dereference
	// subsequent pointers as necessary.
	//
	// After the first round-trip, we set v back to the original value to
	// preserve the original RW flags contained in reflect.Value.
	v0 := v
	haveAddr := false

	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Pointer && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Pointer && !e.IsNil() && (!decodingNull || e.Elem().Kind() == reflect.Pointer) {
				haveAddr = false
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Pointer {
			break
		}

		if decodingNull && v.CanSet() {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return v
}

type Row []*Column

type Column struct {
	column string
	value  any
}

func ReflectRows(v any, namer dialector.Namer, tag string) (rows []Row, err error) {

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	rt := rv.Type()
	multiple := false

	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		multiple = true
		rt = rv.Type().Elem()
	}

	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if rt.Kind() != reflect.Struct {
		err = fmt.Errorf("only accept struct, got: %s", rt.Kind())
		return
	}

	if multiple {
		for i := 0; i < rv.Len(); i++ {
			rows = append(rows, reflectStruct(rt, rv.Index(i), namer, tag))
		}
	} else {
		rows = append(rows, reflectStruct(rt, rv, namer, tag))
	}

	return
}

func TrimColumns(s string) string {
	return strings.TrimSuffix(s, ",")
}

func ExtractTag(s string) string {
	idx := strings.Index(s, ";")
	if idx == -1 {
		return s
	}
	return s[:idx]
}

func reflectStruct(rt reflect.Type, rv reflect.Value, namer dialector.Namer, tag string) Row {

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	var r Row
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		var n string
		if t := f.Tag.Get(tag); t != "" {
			n = ExtractTag(t)
		} else {
			n = namer.ColumnName(f.Name)
		}
		v := rv.Field(i)
		if !isNil(v) {
			r = append(r, &Column{
				column: n,
				value:  v.Interface(),
			})
		}
	}

	return r
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice, reflect.UnsafePointer:
		if v.IsNil() {
			return true
		}
	}
	return false
}

func RowColumns(row Row, namer dialector.Namer) (columns []string) {
	for _, v := range row {
		columns = append(columns, namer.ReservedName(v.column))
	}
	return
}

func RowVars(row Row) (vars []string) {
	for _, v := range row {
		vars = append(vars, fmt.Sprintf("#{%s}", v.column))
	}
	return
}

func RowParams(row Row) (params []NameValue) {
	for _, v := range row {
		params = append(params, NameValue{
			Name:  v.column,
			Value: v.value,
		})
	}
	return
}

func RowsVars(rows []Row) (vars []string) {
	for i, v := range rows {
		var s []string
		for _, vv := range v {
			s = append(s, fmt.Sprintf("#{%s%d}", vv.column, i))
		}
		vars = append(vars, fmt.Sprintf("(%s)", strings.Join(s, ",")))
	}
	return
}

func RowsParams(rows []Row) (params []NameValue) {
	for i, v := range rows {
		for _, vv := range v {
			params = append(params, NameValue{
				Name:  fmt.Sprintf("%s%d", vv.column, i),
				Value: vv.value,
			})
		}

	}
	return
}

func TestIndirect(t *testing.T) {

	var a *int
	rv := reflect.ValueOf(&a)

	t.Log(rv.Kind())

	rv = indirect(rv, false)

	t.Log(rv.Kind())

	rv.Set(reflect.ValueOf(3))

	t.Log(*a)
}

func TestSetValueBasic(t *testing.T) {

	var a *int
	rv := reflect.ValueOf(&a)
	err := setValue(rv, 1)
	require.NoError(t, err)
	t.Log(*a)

	var b int8
	rv = reflect.ValueOf(&b)
	err = setValue(rv, 9)
	require.NoError(t, err)
	t.Log(b)

	var c float32
	rv = reflect.ValueOf(&c)
	err = setValue(rv, 3)
	require.NoError(t, err)
	t.Log(c)

	var d string
	rv = reflect.ValueOf(&d)
	err = setValue(rv, "hello world")
	require.NoError(t, err)
	t.Log(d)
}
