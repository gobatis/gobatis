package batis

import (
	"fmt"
	"strings"
	
	"github.com/gobatis/gobatis/dialector"
	"github.com/pkg/errors"
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
	SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error)
}

type onConflict struct {
	fields string
	sql    string
	params []KeyValue
}

func (o onConflict) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	sql = fmt.Sprintf("on confilct(%s) %s", TrimColumns(o.fields), o.sql)
	params = o.params
	return
}

func OnConflict(fields string, sql string, params ...KeyValue) Element {
	return &onConflict{
		fields: fields,
		sql:    sql,
		params: params,
	}
}

type where struct {
	sql    string
	params []KeyValue
}

func (w where) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	sql = fmt.Sprintf("where %s", strings.TrimSpace(w.sql))
	params = w.params
	return
}

func Where(sql string, params ...KeyValue) Element {
	return &where{
		sql:    sql,
		params: params,
	}
}

type and struct {
	sql    string
	params []KeyValue
}

func (a and) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	//TODO implement me
	panic("implement me")
}

func And(sql string, params ...KeyValue) Element {
	return &and{
		sql:    sql,
		params: params,
	}
}

type update struct {
	table string
	data  map[string]any
	elems []Element
	where *where
}

func (u update) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	for _, v := range u.elems {
		switch vv := v.(type) {
		case *where:
			if u.where != nil {
				err = fmt.Errorf("batis.Where() should be invoked no more than once")
				return
			}
			u.where = vv
		default:
			err = fmt.Errorf("method db.Update() accept elements use of batis.Where()")
			return
		}
	}
	
	var sqls []string
	sqls = append(sqls, fmt.Sprintf("update %s set", namer.TableName(u.table)))
	
	var sets []string
	for k := range u.data {
		sets = append(sets, fmt.Sprintf("%s=#{%s}", namer.ColumnName(k), k))
	}
	sqls = append(sqls, strings.Join(sets, ","))
	for k, v := range u.data {
		params = append(params, KeyValue{
			Name:  k,
			Value: v,
		})
	}
	if u.where != nil {
		var _s string
		var _p []KeyValue
		_s, _p, err = u.where.SQL(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, fmt.Sprintf("%s", _s))
		params = append(params, _p...)
	}
	sql = strings.Join(sqls, space)
	
	return
}

type insert struct {
	table      string
	data       any
	onConflict *onConflict
	returning  *returning
	elems      []Element
}

func (i insert) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "build insert sql error")
		}
	}()
	
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
	
	var rows []Row
	switch vv := i.data.(type) {
	case Rows:
		rows, err = vv.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = reflectRows(i.data, namer, tag)
		if err != nil {
			return
		}
	}
	
	if l := len(rows); l != 1 {
		err = fmt.Errorf("expect 1 row, got: %d", l)
		return
	}
	
	var sqls []string
	
	sqls = append(sqls, fmt.Sprintf("insert into %s(%s) values(%s)",
		namer.TableName(i.table),
		strings.Join(rowColumns(rows[0], namer), ","),
		strings.Join(rowVars(rows[0]), ","),
	))
	params = append(params, rowParams(rows[0])...)
	
	if i.onConflict != nil {
		var _s string
		var _p []KeyValue
		_s, _p, err = i.onConflict.SQL(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, _s)
		params = append(params, _p...)
	}
	
	if i.returning != nil {
		var _s string
		var _p []KeyValue
		_s, _p, err = i.returning.SQL(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, _s)
		params = append(params, _p...)
	}
	
	sql = strings.Join(sqls, space)
	
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

func (i insertBatch) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	
	defer func() {
		if err != nil {
			err = errors.Wrap(err, "build insert batch sql error")
		}
	}()
	
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
	
	var rows []Row
	switch vv := i.data.(type) {
	case Rows:
		rows, err = vv.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = reflectRows(i.data, namer, tag)
		if err != nil {
			return
		}
	}
	
	if l := len(rows); l == 0 {
		err = fmt.Errorf("expect rows legnth > 0, got: %d", l)
		return
	}
	
	var sqls []string
	
	sqls = append(sqls, fmt.Sprintf("insert into %s(%s) values%s",
		namer.TableName(i.table),
		strings.Join(rowColumns(rows[0], namer), ","),
		strings.Join(rowsVars(rows), ","),
	))
	params = append(params, rowsParams(rows)...)
	
	if i.onConflict != nil {
		var _s string
		var _p []KeyValue
		_s, _p, err = i.onConflict.SQL(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, _s)
		params = append(params, _p...)
	}
	
	if i.returning != nil {
		var _s string
		var _p []KeyValue
		_s, _p, err = i.returning.SQL(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, _s)
		params = append(params, _p...)
	}
	
	sql = strings.Join(sqls, space)
	
	return
}

type fetch struct {
}

func (f fetch) SQL(namer dialector.Namer, tag string) (s string, params []KeyValue, err error) {
	//TODO implement me
	panic("implement me")
}

type del struct {
	table string
	elems []Element
	where *where
}

func (d del) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	for _, v := range d.elems {
		switch vv := v.(type) {
		case *where:
			if d.where != nil {
				err = fmt.Errorf("batis.Where() should be invoked no more than once")
				return
			}
			d.where = vv
		default:
			err = fmt.Errorf("method db.Delete() accept elements use of batis.Where()")
			return
		}
	}
	
	var sqls []string
	sqls = append(sqls, fmt.Sprintf("delete from %s", namer.TableName(strings.TrimSpace(d.table))))
	
	if d.where != nil {
		var _s string
		var _p []KeyValue
		_s, _p, err = d.where.SQL(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, fmt.Sprintf("%s", _s))
		params = append(params, _p...)
	}
	sql = strings.Join(sqls, space)
	
	return
}

type query struct {
	sql    string
	params []KeyValue
}

func (q query) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	sql = strings.TrimSpace(sql)
	params = q.params
	return
}

type exec struct {
	sql    string
	params []KeyValue
}

func (e exec) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	sql = strings.TrimSpace(sql)
	params = e.params
	return
}

type build struct {
}

func (b build) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	panic("todo")
}

type returning struct {
	sql string
}

func (r returning) SQL(namer dialector.Namer, tag string) (sql string, params []KeyValue, err error) {
	sql = fmt.Sprintf("returning %s", TrimColumns(r.sql))
	return
}

func Returning(fields string) Element {
	return &returning{
		sql: fields,
	}
}
