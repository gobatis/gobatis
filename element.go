package batis

import (
	"fmt"
	"strings"

	"github.com/gobatis/gobatis/dialector"
)

var _ Elem = (*query)(nil)
var _ Elem = (*exec)(nil)
var _ Elem = (*insert)(nil)
var _ Elem = (*insertBatch)(nil)
var _ Elem = (*returning)(nil)
var _ Elem = (*update)(nil)
var _ Elem = (*del)(nil)
var _ Elem = (*onConflict)(nil)
var _ Elem = (*where)(nil)
var _ Elem = (*fetch)(nil)
var _ Elem = (*build)(nil)
var _ Elem = (*innerSQL)(nil)

type Namer func(name string) string

type Elem interface {
	Raw(namer dialector.Namer, tag string) (raw *raw, err error)
}

func newInnerSQL(sql string, params ...NameValue) innerSQL {
	return innerSQL{sql: sql, params: params}
}

type innerSQL struct {
	sql    string
	params []NameValue
}

func (i innerSQL) Raw(namer dialector.Namer, tag string) (*raw, error) {
	return newRaw(false, i.sql, i.params), nil
}

type onConflict struct {
	fields string
	sql    string
	params []NameValue
}

func (o onConflict) Raw(namer dialector.Namer, tag string) (*raw, error) {
	return newRaw(false, fmt.Sprintf("on conflict(%s) %s", TrimColumns(o.fields), o.sql), o.params), nil
}

func OnConflict(fields string, sql string, params ...NameValue) Elem {
	return &onConflict{
		fields: fields,
		sql:    sql,
		params: params,
	}
}

type where struct {
	sql    string
	params []NameValue
}

func (w where) Raw(namer dialector.Namer, tag string) (raw *raw, err error) {
	return newRaw(false, fmt.Sprintf("where %s", strings.TrimSpace(w.sql)), w.params), nil
}

func Where(sql string, params ...NameValue) Elem {
	return &where{
		sql:    sql,
		params: params,
	}
}

type update struct {
	table     string
	data      map[string]any
	elems     []Elem
	where     *where
	returning *returning
}

func (u update) Raw(namer dialector.Namer, tag string) (r *raw, err error) {
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

	r = &raw{
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
		r.setVar(k, v)
	}
	r.Vars = u.data
	if u.where != nil {
		var rr *raw
		rr, err = u.where.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, fmt.Sprintf("%s", rr.SQL))
		for kk, vv := range rr.Vars {
			r.setVar(kk, vv)
		}
	}
	r.SQL = strings.Join(sqls, space)

	return
}

type insert struct {
	table      string
	data       any
	onConflict *onConflict
	returning  *returning
	elems      []Elem
}

func (i insert) Raw(namer dialector.Namer, tag string) (r *raw, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("build insert sql error: %w", err)
		}
	}()

	if i.data == nil {
		err = fmt.Errorf("insert data is nil")
		return
	}

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
	r = &raw{
		Query: i.returning != nil,
	}
	var rows []Row
	switch vv := i.data.(type) {
	case Rows:
		rows, err = vv.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = ReflectRows(i.data, namer, tag)
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
		strings.Join(RowColumns(rows[0], namer), ","),
		strings.Join(RowVars(rows[0]), ","),
	))
	r.setParams(RowParams(rows[0])...)

	if i.onConflict != nil {
		var rr *raw
		rr, err = i.onConflict.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, rr.SQL)
		r.mergeVars(rr.Vars)
	}

	if i.returning != nil {
		var rr *raw
		rr, err = i.returning.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, rr.SQL)
		r.mergeVars(rr.Vars)
	}

	r.SQL = strings.Join(sqls, space)

	return
}

type insertBatch struct {
	table      string
	data       any
	batch      int
	elems      []Elem
	onConflict *onConflict
	returning  *returning
}

func (i insertBatch) Raw(namer dialector.Namer, tag string) (r *raw, err error) {

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

	r = &raw{
		Query: i.returning != nil,
	}

	var rows []Row
	switch vv := i.data.(type) {
	case Rows:
		rows, err = vv.Reflect(namer, tag)
		if err != nil {
			return
		}
	default:
		rows, err = ReflectRows(i.data, namer, tag)
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
		strings.Join(RowColumns(rows[0], namer), ","),
		strings.Join(RowsVars(rows), ","),
	))
	r.setParams(RowsParams(rows)...)

	if i.onConflict != nil {
		var rr *raw
		rr, err = i.onConflict.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, rr.SQL)
		r.mergeVars(rr.Vars)
	}

	if i.returning != nil {
		var rr *raw
		rr, err = i.returning.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, rr.SQL)
		r.mergeVars(rr.Vars)
	}

	r.SQL = strings.Join(sqls, space)

	return
}

type fetch struct {
}

func (f fetch) Raw(namer dialector.Namer, tag string) (raw *raw, err error) {
	//TODO implement me
	panic("implement me")
}

type del struct {
	table     string
	elems     []Elem
	where     *where
	returning *returning
}

func (d del) Raw(namer dialector.Namer, tag string) (r *raw, err error) {
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

	r = &raw{}

	var sqls []string
	sqls = append(sqls, fmt.Sprintf("delete from %s", namer.TableName(strings.TrimSpace(d.table))))

	if d.where != nil {
		var rr *raw
		rr, err = d.where.Raw(namer, tag)
		if err != nil {
			return
		}
		sqls = append(sqls, fmt.Sprintf("%s", rr.SQL))
		r.mergeVars(rr.Vars)
	}
	r.SQL = strings.Join(sqls, space)

	return
}

type query struct {
	sql    string
	params []NameValue
}

func (q query) Raw(namer dialector.Namer, tag string) (*raw, error) {
	return newRaw(true, q.sql, q.params), nil
}

type exec struct {
	sql    string
	params []NameValue
}

func (e exec) Raw(namer dialector.Namer, tag string) (*raw, error) {
	return newRaw(false, strings.TrimSpace(e.sql), e.params), nil
}

type build struct {
}

func (b build) Raw(namer dialector.Namer, tag string) (raw *raw, err error) {
	panic("todo")
}

type returning struct {
	sql string
}

func (r returning) Raw(namer dialector.Namer, tag string) (raw *raw, err error) {
	return newRaw(false, fmt.Sprintf("returning %s", TrimColumns(r.sql)), nil), nil
}

func Returning(fields string) Elem {
	return &returning{
		sql: fields,
	}
}
