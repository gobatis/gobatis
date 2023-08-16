package paging

import (
	"fmt"
	"strings"

	"github.com/gobatis/gobatis/executor"
)

const (
	selectTag = iota + 101
	selectExceptTag
	whereTag
	countTag
	fromTag
	pageTag
	scrollTag
	onConflictTag
	raw
	tableTag
)

func NewPaging() *Paging {
	return &Paging{}
}

type Paging struct {
	elems map[int][]executor.Element
}

func joinElements(items []executor.Element) (string, []executor.NameValue) {

	var sqls []string
	var params []executor.NameValue
	for _, v := range items {
		sqls = append(sqls, v.SQL)
		params = append(params, v.Params...)
	}

	return strings.Join(sqls, " "), params
}

func (b *Paging) Build() (executors []executor.Executor, err error) {

	if b.elems == nil {
		return
	}

	w := executor.Executor{}
	c := executor.Executor{}
	s := executor.Executor{}

	if v, ok := b.elems[selectTag]; ok {
		sql, params := joinElements(v)
		s.SQL += sql + " "
		s.Params = append(s.Params, params...)
	}

	if v, ok := b.elems[countTag]; ok {
		sql, params := joinElements(v)
		c.SQL += sql + " "
		c.Params = append(c.Params, params...)
	}

	if v, ok := b.elems[fromTag]; ok {
		sql, params := joinElements(v)
		w.SQL += sql + " "
		w.Params = append(w.Params, params...)
	}

	if v, ok := b.elems[whereTag]; ok {
		sql, params := joinElements(v)
		w.SQL += sql + " "
		w.Params = append(w.Params, params...)
	}

	lo := executor.Executor{}
	if v, ok := b.elems[pageTag]; ok {
		sql, params := joinElements(v)
		lo.SQL += sql + " "
		lo.Params = append(lo.Params, params...)
	}

	if s.SQL != "" {
		s.Merge(w)
		s.Merge(lo)
		s.Type = executor.Query
		executors = append(executors, s)
	}
	if c.SQL != "" {
		c.Merge(w)
		c.Type = executor.Query
		executors = append(executors, c)
	}

	return
}

func (b *Paging) addElement(e executor.Element) *Paging {
	if b.elems == nil {
		b.elems = map[int][]executor.Element{}
	}

	if _, ok := b.elems[e.Name]; !ok {
		b.elems[e.Name] = make([]executor.Element, 0)
	}
	b.elems[e.Name] = append(b.elems[e.Name], e)
	return b
}

func Select(sql string, params ...executor.NameValue) *Paging {
	b := &Paging{}
	b.addElement(executor.Element{
		Name: selectTag,
		SQL:  fmt.Sprintf("select %s", sql),
	})
	return b
}

func SelectExcept(fields ...string) *Paging {
	b := &Paging{}
	b.addElement(executor.Element{
		Name: selectExceptTag,
		SQL:  strings.Join(fields, ","),
	})
	return b
}

func Raw(sql string, params ...executor.NameValue) *Paging {
	b := &Paging{}
	b.addElement(executor.Element{
		Name: selectExceptTag,
		SQL:  sql,
	})
	return b
}

func (b *Paging) From(sql string, params ...executor.NameValue) *Paging {
	b.addElement(executor.Element{
		Name:   fromTag,
		SQL:    fmt.Sprintf("from %s", sql),
		Params: params,
	})
	return b
}

func (b *Paging) Where(sql string, params ...executor.NameValue) *Paging {
	b.addElement(
		executor.Element{
			Name:   whereTag,
			SQL:    fmt.Sprintf("where %s", sql),
			Params: params,
		})
	return b
}

func (b *Paging) Count(sql string) *Paging {
	b.addElement(executor.Element{
		Name: countTag,
		SQL:  fmt.Sprintf("select count(%s)", sql),
	})
	return b
}

func (b *Paging) Page(page, limit int64) *Paging {
	b.addElement(executor.Element{
		Name: pageTag,
		SQL:  fmt.Sprintf("limit %d offset %d", limit, limit*page),
	})
	return b
}

func (b *Paging) Scroll(limit int64, and executor.Element) *Paging {
	b.addElement(executor.Element{
		Name: scrollTag,
		//sql:  fmt.Sprintf("%s limit %d", sql, limit),
	})
	return b
}
