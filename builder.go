package gobatis

import "fmt"

type Builder struct {
	elems []element
}

func (b *Builder) addElement(e element) *Builder {
	b.elems = append(b.elems, e)
	return b
}

func (b *Builder) From(sql string, params ...NameValue) *Builder {
	b.addElement(element{
		name:   fromTag,
		sql:    sql,
		params: params,
	})
	return b
}

func (b *Builder) Where(sql string, params ...NameValue) *Builder {
	b.addElement(
		element{
			name:   whereTag,
			sql:    sql,
			params: params,
		})
	return b
}

func (b *Builder) Count(sql string) *Builder {
	b.addElement(element{
		name: countTag,
		sql:  sql,
	})
	return b
}

func (b *Builder) Page(page, limit int64) *Builder {
	b.addElement(element{
		name: pagingTag,
		sql:  fmt.Sprintf("limit %d offset %d", limit, page),
	})
	return b
}

func (b *Builder) Scroll(limit int64, and Element) *Builder {
	b.addElement(element{
		name: scrollTag,
		//sql:  fmt.Sprintf("%s limit %d", sql, limit),
	})
	return b
}
