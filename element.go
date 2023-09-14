package batis

import (
	"fmt"
	"strings"

	"github.com/gobatis/gobatis/dialector"
	"github.com/gobatis/gobatis/executor"
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
var _ Element = (*fetch)(nil)
var _ Element = (*build)(nil)
var _ Element = (*innerSQL)(nil)

type Namer func(name string) string

type Element interface {
	Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error)
}

func newInnerSQL(sql string, params ...executor.Param) innerSQL {
	return innerSQL{sql: sql, params: params}
}

type innerSQL struct {
	sql    string
	params []executor.Param
}

func (i innerSQL) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	raw = &executor.Raw{
		SQL:    i.sql,
		Params: i.params,
	}
	return
}

type onConflict struct {
	fields string
	sql    string
	params []executor.Param
}

func (o onConflict) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	raw = &executor.Raw{
		SQL:    fmt.Sprintf("on conflict(%s) %s", executor.TrimColumns(o.fields), o.sql),
		Params: o.params,
	}
	return
}

func OnConflict(fields string, sql string, params ...executor.Param) Element {
	return &onConflict{
		fields: fields,
		sql:    sql,
		params: params,
	}
}

type where struct {
	sql    string
	params []executor.Param
}

func (w where) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	raw = &executor.Raw{
		SQL:    fmt.Sprintf("where %s", strings.TrimSpace(w.sql)),
		Params: w.params,
	}
	return
}

func Where(sql string, params ...executor.Param) Element {
	return &where{
		sql:    sql,
		params: params,
	}
}

type update struct {
	table     string
	data      map[string]any
	elems     []Element
	where     *where
	returning *returning
}

func (u update) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	for _, v := range u.elems {
		switch vv := v.(type) {
		case *where:
			if u.where != nil {
				err = fmt.Errorf("batis.Where() should be invoked no more than once")
				return
			}
			u.where = vv
		case *returning:
			if u.returning != nil {
				err = fmt.Errorf("batis.Returning() should be invoked no more than once")
				return
			}
			u.returning = vv
		default:
			err = fmt.Errorf("method db.Update() accept elements use of batis.Where()")
			return
		}
	}

	raw = &executor.Raw{
		Query: u.returning != nil,
	}

	var sqls []string
	sqls = append(sqls, fmt.Sprintf("update %s set", namer.TableName(u.table)))

	var sets []string
	for k := range u.data {
		sets = append(sets, fmt.Sprintf("%s=#{%s}", namer.ColumnName(k), k))
	}
	sqls = append(sqls, strings.Join(sets, ","))
	for k, v := range u.data {
		raw.Params = append(raw.Params, executor.Param{
			Name:  k,
			Value: v,
		})
	}
	if u.where != nil {
		var r *executor.Raw
		r, err = u.where.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, fmt.Sprintf("%s", r.SQL))
		raw.Params = append(raw.Params, r.Params...)
	}
	raw.SQL = strings.Join(sqls, space)

	return
}

type insert struct {
	table      string
	data       any
	onConflict *onConflict
	returning  *returning
	elems      []Element
}

func (i insert) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("build insert sql error: %w", err)
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
	raw = &executor.Raw{
		Query: i.returning != nil,
	}
	var rows []executor.Row
	switch vv := i.data.(type) {
	case Rows:
		rows, err = vv.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = executor.ReflectRows(i.data, namer, tag)
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
		strings.Join(executor.RowColumns(rows[0], namer), ","),
		strings.Join(executor.RowVars(rows[0]), ","),
	))
	raw.Params = append(raw.Params, executor.RowParams(rows[0])...)

	if i.onConflict != nil {
		var r *executor.Raw
		r, err = i.onConflict.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, r.SQL)
		raw.Params = append(raw.Params, r.Params...)
	}

	if i.returning != nil {
		var r *executor.Raw
		r, err = i.returning.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, r.SQL)
		raw.Params = append(raw.Params, r.Params...)
	}

	raw.SQL = strings.Join(sqls, space)

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

func (i insertBatch) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("build insert batch sql error: %w", err)
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

	raw = &executor.Raw{
		Query: i.returning != nil,
	}

	var rows []executor.Row
	switch vv := i.data.(type) {
	case Rows:
		rows, err = vv.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = executor.ReflectRows(i.data, namer, tag)
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
		strings.Join(executor.RowColumns(rows[0], namer), ","),
		strings.Join(executor.RowsVars(rows), ","),
	))
	raw.Params = append(raw.Params, executor.RowsParams(rows)...)

	if i.onConflict != nil {
		var r *executor.Raw
		r, err = i.onConflict.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, r.SQL)
		r.Params = append(r.Params, r.Params...)
	}

	if i.returning != nil {
		var r *executor.Raw
		r, err = i.returning.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, r.SQL)
		r.Params = append(r.Params, r.Params...)
	}

	raw.SQL = strings.Join(sqls, space)

	return
}

type fetch struct {
}

func (f fetch) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	//TODO implement me
	panic("implement me")
}

type del struct {
	table     string
	elems     []Element
	where     *where
	returning *returning
}

func (d del) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
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

	raw = &executor.Raw{}

	var sqls []string
	sqls = append(sqls, fmt.Sprintf("delete from %s", namer.TableName(strings.TrimSpace(d.table))))

	if d.where != nil {
		var r *executor.Raw
		r, err = d.where.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, fmt.Sprintf("%s", r.SQL))
		r.Params = append(r.Params, r.Params...)
	}
	raw.SQL = strings.Join(sqls, space)

	return
}

type query struct {
	sql    string
	params []executor.Param
}

func (q query) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	raw = &executor.Raw{
		Query:  true,
		SQL:    q.sql,
		Params: q.params,
	}
	return
}

type exec struct {
	sql    string
	params []executor.Param
}

func (e exec) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	raw = &executor.Raw{
		SQL:    strings.TrimSpace(e.sql),
		Params: e.params,
	}
	return
}

type build struct {
}

func (b build) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	panic("todo")
}

type returning struct {
	sql string
}

func (r returning) Raw(namer dialector.Namer, tag string) (raw *executor.Raw, err error) {
	raw = &executor.Raw{
		SQL: fmt.Sprintf("returning %s", executor.TrimColumns(r.sql)),
	}
	return
}

func Returning(fields string) Element {
	return &returning{
		sql: fields,
	}
}
