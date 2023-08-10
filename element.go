package batis

import (
	"fmt"
	"github.com/gobatis/gobatis/executor"
	"strings"
)

var _ Element = (*element)(nil)
var _ Element = (*onConflict)(nil)
var _ Element = (*where)(nil)
var _ Element = (*and)(nil)
var _ Element = (*update)(nil)
var _ Element = (*insert)(nil)
var _ Element = (*insertBatch)(nil)
var _ Element = (*fetch)(nil)
var _ Element = (*_delete)(nil)
var _ Element = (*query)(nil)
var _ Element = (*_executor)(nil)
var _ Element = (*build)(nil)

type Namer func(name string) string

type Element interface {
	SQL(n Namer) (string, error)
	Params() []executor.NameValue
}

type element struct {
	name   int
	fields []string
	sql    string
	params []executor.NameValue
}

func (e element) SQL(n Namer) (string, error) {
	return e.sql, nil
}

func (e element) Params() []executor.NameValue {
	return e.params
}

type onConflict struct {
	fields []string
	sql    string
	params []executor.NameValue
}

func (o onConflict) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (o onConflict) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

func OnConflict(fields []string, sql string, params ...executor.NameValue) Element {
	return &onConflict{
		fields: fields,
		sql:    sql,
		params: params,
	}
}

type where struct {
	sql    string
	params []executor.NameValue
}

func (w where) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (w where) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

func Where(sql string, params ...executor.NameValue) Element {
	return &where{
		sql:    sql,
		params: params,
	}
}

type and struct {
	sql    string
	params []executor.NameValue
}

func (a and) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a and) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

func And(sql string, params ...executor.NameValue) Element {
	return &and{
		sql:    sql,
		params: params,
	}
}

type update struct {
	table string
	data  map[string]executor.Executor
	where Element
}

func (u update) SQL(n Namer) (string, error) {
	var updateStmt strings.Builder
	updateStmt.WriteString(fmt.Sprintf("update %s set ", u.table))
	for key := range u.data {
		updateStmt.WriteString(fmt.Sprintf("%s=#{%s}, ", key, key))
	}
	return updateStmt.String(), nil
}

func (u update) Params() []executor.NameValue {
	var params []executor.NameValue
	for k, v := range u.data {
		params = append(params, executor.NameValue{
			Name:  k,
			Value: v,
		})
	}
	return params
}

type insert struct {
	table string
	data  any
	elems []Element
}

func (i insert) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (i insert) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

type insertBatch struct {
	table string
	data  any
	batch int
	elems []Element
}

func (i insertBatch) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (i insertBatch) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

type fetch struct {
}

func (f fetch) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (f fetch) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

type _delete struct {
	table string
	where Element
}

func (del _delete) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (del _delete) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

type query struct {
}

func (q query) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (q query) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

type _executor struct {
}

func (e _executor) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (e _executor) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}

type build struct {
}

func (b build) SQL(n Namer) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b build) Params() []executor.NameValue {
	//TODO implement me
	panic("implement me")
}
