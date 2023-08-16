package batis

import (
	"fmt"
	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/executor"
	"github.com/gobatis/gobatis/reflects"
	"strings"
)

var _ Element = (*query)(nil)
var _ Element = (*exec)(nil)
var _ Element = (*insert)(nil)
var _ Element = (*insertBatch)(nil)
var _ Element = (*returning)(nil)
var _ Element = (*update)(nil)
var _ Element = (*del)(nil)
var _ Element = (*onConflict)(nil)
var _ Element = (*where)(nil)
var _ Element = (*and)(nil)
var _ Element = (*fetch)(nil)
var _ Element = (*build)(nil)

type Namer func(name string) string

type Element interface {
	SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error)
}

type onConflict struct {
	fields string
	sql    string
	params []executor.NameValue
}

func (o onConflict) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	
	return
}

func OnConflict(fields string, sql string, params ...executor.NameValue) Element {
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

func (w where) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
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

func (a and) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
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
	data  map[string]any
	where Element
}

func (u update) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	var updateStmt strings.Builder
	updateStmt.WriteString(fmt.Sprintf("update %s set ", u.table))
	for key := range u.data {
		updateStmt.WriteString(fmt.Sprintf("%s=#{%s}, ", key, key))
	}
	return updateStmt.String(), nil, nil
}

//func (u update) Params() []executor.NameValue {
//	var params []executor.NameValue
//	for k, v := range u.data {
//		params = append(params, executor.NameValue{
//			Name:  k,
//			Value: v,
//		})
//	}
//	return params
//}

type insert struct {
	table      string
	data       any
	onConflict *onConflict
	returning  *returning
	elems      []Element
}

func (i insert) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	for _, v := range i.elems {
		switch vv := v.(type) {
		case *onConflict:
			if i.onConflict != nil {
				err = fmt.Errorf("batis.OnConflict() should be invoked no more than once")
				return
			}
			i.onConflict = vv
		case *returning:
			if i.returning != nil {
				err = fmt.Errorf("batis.Returning() should be invoked no more than once")
				return
			}
			i.returning = vv
		
		default:
			err = fmt.Errorf("method db.Insert() accept elements use of batis.OnConflict() or batis.Returning()")
			return
		}
	}
	
	var rows []reflects.Row
	switch vv := i.data.(type) {
	case reflects.Rows:
		rows, err = vv.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = reflects.ReflectRows(i.data, namer, tag)
		if err != nil {
			return
		}
	}
	
	var sqls []string
	sqls = append(sqls, fmt.Sprintf("insert into %s(%s) values(%s)",
		namer.TableName(i.table),
		strings.Join(reflects.RowsColumns(rows, namer), ","),
		strings.Join(reflects.RowsVars(rows), ","),
	))
	
	return
}

type insertBatch struct {
	table      string
	data       any
	batch      int
	elems      []Element
	onConflict *onConflict
	returning  *returning
}

func (i insertBatch) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	
	var rows []reflects.Row
	switch v := i.data.(type) {
	case reflects.Rows:
		rows, err = v.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = reflects.ReflectRows(i.data, namer, tag)
		if err != nil {
			return
		}
	}
	
	_ = rows
	
	return
}

type fetch struct {
}

func (f fetch) SQL(namer dialector.Namer, tag string) (s string, params []executor.NameValue, err error) {
	//TODO implement me
	panic("implement me")
}

type del struct {
	table string
	where Element
}

func (d del) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	_, ok := d.where.(*where)
	if !ok {
		return "", nil, fmt.Errorf("db.Delete excpet where element use batis.Where()")
	}
	return "", nil, nil
}

type query struct {
	sql    string
	params []executor.NameValue
}

func (q query) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	//TODO implement me
	panic("implement me")
}

type exec struct {
}

func (e exec) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	//TODO implement me
	panic("implement me")
}

type build struct {
}

func (b build) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	//TODO implement me
	panic("implement me")
}

type returning struct {
	sql string
}

func (r returning) SQL(namer dialector.Namer, tag string) (sql string, params []executor.NameValue, err error) {
	return fmt.Sprintf("returning %s", r), nil, nil
}

func Returning(fields string) Element {
	return &returning{
		sql: fields,
	}
}

func buildExecutor(namer dialector.Namer, tag string, et int, elems ...Element) (e *executor.Executor, err error) {
	
	var sqls []string
	params := map[string]executor.NameValue{}
	
	for _, v := range elems {
		var s string
		s, _, err = v.SQL(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, s)
		//for _, vv := range v.Params() {
		//	params[vv.Name] = vv
		//}
	}
	
	e = &executor.Executor{
		Type:   et,
		SQL:    strings.Join(sqls, space),
		Params: nil,
		Err:    nil,
		Conn:   nil,
	}
	
	for _, v := range params {
		e.Params = append(e.Params, v)
	}
	
	return
}
