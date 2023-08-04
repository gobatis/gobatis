package gobatis

import (
	"fmt"
	"strings"
)

type Builder interface {
	Build() []session
}

func NewPaging() *Paging {
	return &Paging{}
}

type Paging struct {
	elems []element
}

func (b *Paging) Build() []session {
	return nil
}

func (b *Paging) addElement(e element) *Paging {
	b.elems = append(b.elems, e)
	return b
}

// Select TODO wrap field
func (b *Paging) Select(fields ...string) *Paging {
	b.addElement(element{
		name: selectTag,
		sql:  strings.Join(fields, ","),
	})
	return b
}

func (b *Paging) SelectAllExcept(fields ...string) *Paging {
	b.addElement(element{
		name: selectExceptTag,
		sql:  strings.Join(fields, ","),
	})
	return b
}

func (b *Paging) Raw(sql string, params ...NameValue) *Paging {
	b.addElement(element{
		name: selectExceptTag,
		sql:  sql,
	})
	return b
}

func (b *Paging) From(sql string, params ...NameValue) *Paging {
	b.addElement(element{
		name:   fromTag,
		sql:    sql,
		params: params,
	})
	return b
}

func (b *Paging) Where(sql string, params ...NameValue) *Paging {
	b.addElement(
		element{
			name:   whereTag,
			sql:    sql,
			params: params,
		})
	return b
}

func (b *Paging) Count(sql string) *Paging {
	b.addElement(element{
		name: countTag,
		sql:  sql,
	})
	return b
}

func (b *Paging) Page(page, limit int64) *Paging {
	b.addElement(element{
		name: pagingTag,
		sql:  fmt.Sprintf("limit %d offset %d", limit, page),
	})
	return b
}

func (b *Paging) Scroll(limit int64, and Element) *Paging {
	b.addElement(element{
		name: scrollTag,
		//sql:  fmt.Sprintf("%s limit %d", sql, limit),
	})
	return b
}
