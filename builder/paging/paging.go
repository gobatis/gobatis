package paging

import (
	"fmt"
	batis "github.com/gobatis/gobatis"
	"github.com/gobatis/gobatis/builder"
	"github.com/gobatis/gobatis/executor"
	"strings"
)

const (
	selectTag       = "select"
	selectExceptTag = "select_except"
	whereTag
	countTag
	fromTag
	pagingTag
	scrollTag
	onConflictTag
	raw
	tableTag
)

func NewPaging() *Paging {
	return &Paging{}
}

var _ builder.Builder = (*Paging)(nil)

type Paging struct {
	elems map[string][]executor.Element
}

func joinElements(items []executor.Element) (string, []batis.NameValue) {
	
	var sqls []string
	var params []batis.NameValue
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
	
	if v, ok := b.elems[countTag]; ok {
		sql, params := joinElements(v)
		c.SQL += sql
		c.Params = append(s.Params, params...)
	}
	
	if v, ok := b.elems[selectTag]; ok {
		sql, params := joinElements(v)
		s.SQL += sql
		s.Params = append(s.Params, params...)
	}
	
	if v, ok := b.elems[fromTag]; ok {
		sql, params := joinElements(v)
		w.SQL += sql
		w.Params = append(s.Params, params...)
	}
	
	if v, ok := b.elems[whereTag]; ok {
		sql, params := joinElements(v)
		w.SQL += sql
		w.Params = append(s.Params, params...)
	}
	
	if s.SQL != "" {
		s.Merge(w)
		executors = append(executors, s)
	}
	if c.SQL != "" {
		c.Merge(w)
		executors = append(executors, c)
	}
	
	return
}

func (b *Paging) addElement(e executor.Element) *Paging {
	if b.elems == nil {
		b.elems = map[string][]executor.Element{}
	}
	
	if _, ok := b.elems[e.Name]; !ok {
		b.elems[e.Name] = make([]executor.Element, 0)
	}
	b.elems[e.Name] = append(b.elems[e.Name], e)
	return b
}

func Select(sql string, params ...batis.NameValue) *Paging {
	b := &Paging{}
	b.addElement(executor.Element{
		Name: selectTag,
		SQL:  sql,
	})
	return b
}

func SelectAllExcept(fields ...string) *Paging {
	b := &Paging{}
	b.addElement(executor.Element{
		Name: selectExceptTag,
		SQL:  strings.Join(fields, ","),
	})
	return b
}

func Raw(sql string, params ...batis.NameValue) *Paging {
	b := &Paging{}
	b.addElement(executor.Element{
		Name: selectExceptTag,
		SQL:  sql,
	})
	return b
}

func (b *Paging) From(sql string, params ...batis.NameValue) *Paging {
	b.addElement(executor.Element{
		Name:   fromTag,
		SQL:    sql,
		Params: params,
	})
	return b
}

func (b *Paging) Where(sql string, params ...batis.NameValue) *Paging {
	b.addElement(
		executor.Element{
			Name:   whereTag,
			SQL:    sql,
			Params: params,
		})
	return b
}

func (b *Paging) Count(sql string) *Paging {
	b.addElement(executor.Element{
		Name: countTag,
		SQL:  sql,
	})
	return b
}

func (b *Paging) Page(page, limit int64) *Paging {
	b.addElement(executor.Element{
		Name: pagingTag,
		SQL:  fmt.Sprintf("limit %d offset %d", limit, page),
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
