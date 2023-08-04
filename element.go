package gobatis

import (
	"strings"
)

const (
	selectTag = iota
	selectExceptTag
	whereTag
	countTag
	fromTag
	pagingTag
	scrollTag
	onConflictTag
	raw
	tableTag
)

type Element interface {
	SQL() string
	Params() []NameValue
}

var _ Element = (*element)(nil)

type element struct {
	name   int
	sql    string
	params []NameValue
}

func (e element) SQL() string {
	return e.sql
}

func (e element) Params() []NameValue {
	return e.params
}

func OnConflict(fields []string, sql string, params ...NameValue) Element {
	return &element{
		name:   onConflictTag,
		sql:    sql,
		params: params,
	}
}

func Where(sql string, params ...NameValue) Element {
	return &element{
		name:   whereTag,
		sql:    sql,
		params: params,
	}
}

func And(sql string, params ...NameValue) Element {
	return &element{
		name:   whereTag,
		sql:    sql,
		params: params,
	}
}

// Select TODO wrap field
func Select(fields ...string) *Builder {
	b := &Builder{}
	b.addElement(element{
		name: selectTag,
		sql:  strings.Join(fields, ","),
	})
	return b
}

func SelectAllExcept(fields ...string) *Builder {
	b := &Builder{}
	b.addElement(element{
		name: selectExceptTag,
		sql:  strings.Join(fields, ","),
	})
	return b
}

func Raw(sql string, params ...NameValue) *Builder {
	b := &Builder{}
	b.addElement(element{
		name: selectExceptTag,
		sql:  sql,
	})
	return b
}
