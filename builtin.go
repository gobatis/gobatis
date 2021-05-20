package gobatis

import (
	"github.com/gobatis/gobatis/cast"
	"github.com/shopspring/decimal"
	"reflect"
	"strings"
	"sync"
)

var _builtin *builtin

func init() {
	_builtin = &builtin{
		mu: sync.RWMutex{},
		_map: map[string]interface{}{
			"len":     _len,
			"int":     _int,
			"int8":    _int8,
			"int16":   _int16,
			"int32":   _int32,
			"int64":   _int64,
			"uint":    _uint,
			"uint8":   _uint8,
			"uint16":  _uint16,
			"uint32":  _uint32,
			"uint64":  _uint64,
			"decimal": _decimal,
			"bool":    _bool,
			"string":  _string,
			"strings": _strings{},
		},
	}
}

type builtin struct {
	mu   sync.RWMutex
	_map map[string]interface{}
}

func (p *builtin) get(name string) interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()
	v, _ := p._map[name]
	return v
}

func (p *builtin) is(name string) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	_, ok := p._map[name]
	return ok
}

func _len(v interface{}) int {
	return reflect.ValueOf(v).Len()
}

func _int(v interface{}) int {
	r, err := cast.ToIntE(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _int8(v interface{}) int8 {
	r, err := cast.ToInt8E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _int16(v interface{}) int16 {
	r, err := cast.ToInt16E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _int32(v interface{}) int32 {
	r, err := cast.ToInt32E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _int64(v interface{}) int64 {
	r, err := cast.ToInt64E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _uint(v interface{}) uint {
	r, err := cast.ToUintE(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _uint8(v interface{}) uint8 {
	r, err := cast.ToUint8E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _uint16(v interface{}) uint16 {
	r, err := cast.ToUint16E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _uint32(v interface{}) uint32 {
	r, err := cast.ToUint32E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _uint64(v interface{}) uint64 {
	r, err := cast.ToUint64E(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _decimal(v interface{}) decimal.Decimal {
	r, err := cast.ToDecimalE(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _bool(v interface{}) bool {
	r, err := cast.ToBoolE(v)
	if err != nil {
		panic(err)
	}
	return r
}

func _string(v interface{}) string {
	r, err := cast.ToStringE(v)
	if err != nil {
		panic(err)
	}
	return r
}

type _strings struct{}

func (p _strings) Count(s, substr string) int                { return strings.Count(s, substr) }
func (p _strings) Contains(s, substr string) bool            { return strings.Contains(s, substr) }
func (p _strings) ContainsAny(s, chars string) bool          { return strings.ContainsAny(s, chars) }
func (p _strings) ContainsRune(s string, r rune) bool        { return strings.ContainsRune(s, r) }
func (p _strings) LastIndex(s, substr string) int            { return strings.LastIndex(s, substr) }
func (p _strings) IndexAny(s, chars string) int              { return strings.IndexAny(s, chars) }
func (p _strings) LastIndexAny(s, chars string) int          { return strings.LastIndexAny(s, chars) }
func (p _strings) SplitN(s, sep string, n int) []string      { return strings.SplitN(s, sep, n) }
func (p _strings) SplitAfterN(s, sep string, n int) []string { return strings.SplitAfterN(s, sep, n) }
func (p _strings) Split(s, sep string) []string              { return strings.Split(s, sep) }
func (p _strings) SplitAfter(s, sep string) []string         { return strings.SplitAfter(s, sep) }
func (p _strings) Fields(s string) []string                  { return strings.Fields(s) }
func (p _strings) Join(elems []string, sep string) string    { return strings.Join(elems, sep) }
func (p _strings) HasPrefix(s, prefix string) bool           { return strings.HasPrefix(s, prefix) }
func (p _strings) HasSuffix(s, suffix string) bool           { return strings.HasSuffix(s, suffix) }
func (p _strings) Repeat(s string, count int) string         { return strings.Repeat(s, count) }
func (p _strings) ToUpper(s string) string                   { return strings.ToUpper(s) }
func (p _strings) ToLower(s string) string                   { return strings.ToLower(s) }
func (p _strings) ToTitle(s string) string                   { return strings.ToTitle(s) }
func (p _strings) Title(s string) string                     { return strings.Title(s) }
func (p _strings) Trim(s, cutset string) string              { return strings.Trim(s, cutset) }
func (p _strings) TrimLeft(s, cutset string) string          { return strings.TrimLeft(s, cutset) }
func (p _strings) TrimRight(s, cutset string) string         { return strings.TrimRight(s, cutset) }
func (p _strings) TrimSpace(s string) string                 { return strings.TrimSpace(s) }
func (p _strings) TrimPrefix(s, prefix string) string        { return strings.TrimPrefix(s, prefix) }
func (p _strings) TrimSuffix(s, suffix string) string        { return strings.TrimSuffix(s, suffix) }
func (p _strings) Replace(s, old, new string, n int) string  { return strings.Replace(s, old, new, n) }
func (p _strings) ReplaceAll(s, old, new string) string      { return strings.ReplaceAll(s, old, new) }
func (p _strings) EqualFold(s, t string) bool                { return strings.EqualFold(s, t) }
func (p _strings) Index(s, substr string) int                { return strings.Index(s, substr) }
