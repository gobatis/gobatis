package batis

import (
	"fmt"
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

func OnConflict(sql string, params ...NameValue) Element {
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

// Select TODO wrap field
func Select(fields ...string) Element {
	return &element{
		name: selectTag,
		sql:  strings.Join(fields, ","),
	}
}

func SelectExcept(fields ...string) Element {
	return &element{
		name: selectExceptTag,
		sql:  strings.Join(fields, ","),
	}
}

func Count(sql string) Element {
	return &element{
		name: countTag,
		sql:  sql,
	}
}

func From(sql string, params ...NameValue) Element {
	return &element{
		name:   fromTag,
		sql:    sql,
		params: params,
	}
}

func Paging(page, limit int64) Element {
	return &element{
		name: pagingTag,
		sql:  fmt.Sprintf("limit %d offset %d", limit, page),
	}
}

func Scroll(limit int64, sql string) Element {
	return &element{
		name: scrollTag,
		sql:  fmt.Sprintf("%s limit %d", sql, limit),
	}
}
