package batis

type NameValue struct {
	Name  string
	Value any
}

type raw struct {
	Query bool
	SQL   string
	Vars  map[string]any
}

func (r *raw) setVar(name string, value any) {
	if r.Vars == nil {
		r.Vars = map[string]any{}
	}
	r.Vars[name] = value
}

func (r *raw) setParams(params ...NameValue) {
	for _, v := range params {
		r.setVar(v.Name, v.Value)
	}
}

func (r *raw) mergeVars(vars map[string]any) {
	for k, v := range vars {
		r.setVar(k, v)
	}
}

func newRaw(query bool, sql string, params []NameValue) *raw {
	r := &raw{
		Query: query,
		SQL:   sql,
	}
	r.setParams(params...)
	return r
}
