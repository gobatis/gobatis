package gobatis

import (
	"reflect"
	"strings"
)

func _len(v interface{}) int {
	return reflect.ValueOf(v).Len()
}

func _int(v interface{}) int {
	a := reflect.ValueOf(v).Int()
	r := int(a)
	if int64(r) != a {
		panic("int convert not equal")
	}
	return r
}

func _int32(v interface{}) int32 {
	a := reflect.ValueOf(v).Int()
	r := int32(a)
	if int64(r) != a {
		panic("int convert not equal")
	}
	return r
}

func _int64(v interface{}) int64 {
	return reflect.ValueOf(v).Int()
}

func _float32(v interface{}) float32 {
	a := reflect.ValueOf(v).Float()
	r := float32(a)
	if float64(r) != a {
		panic("float convert not equal")
	}
	return r
}

func _float64(v interface{}) float64 {
	return reflect.ValueOf(v).Float()
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
