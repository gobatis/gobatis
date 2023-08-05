package batis

import "github.com/gobatis/gobatis/executor"

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
	Params() []executor.NameValue
}

var _ Element = (*element)(nil)

type element struct {
	name   int
	sql    string
	params []executor.NameValue
}

func (e element) SQL() string {
	return e.sql
}

func (e element) Params() []executor.NameValue {
	return e.params
}

func OnConflict(sql string, params ...executor.NameValue) Element {
	return &element{
		name:   onConflictTag,
		sql:    sql,
		params: params,
	}
}

func Where(sql string, params ...executor.NameValue) Element {
	return &element{
		name:   whereTag,
		sql:    sql,
		params: params,
	}
}

func And(sql string, params ...executor.NameValue) Element {
	return &element{
		name:   whereTag,
		sql:    sql,
		params: params,
	}
}
