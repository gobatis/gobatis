package generator

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"strings"
)

func MakePostgresqlTestCases() {
	makePostgresqlEntity()
	makePostgresqlXML()
	makePostgresqlMapper()
	makePostgresqlCases()
}

func makePostgresqlEntity() {
	header := GoHeader{
		Package: "postgresql",
		Imports: []string{},
	}
	sid := &Param{Name: "Sid", Type: "string", Tag: "sid"}
	source := &Param{Name: "Source", Type: "string", Tag: "source"}
	var entities []*Entity
	entities = append(entities, &Entity{Name: "TypeOriginal", Params: []*Param{sid, source}})
	entities = append(entities, &Entity{Name: "TypePointer", Params: []*Param{sid.pointer(), source.pointer()}})
	entities = append(entities, &Entity{Name: "ArrayTypeOriginal", Params: []*Param{sid, source}})
	entities = append(entities, &Entity{Name: "ArrayTypePointer", Params: []*Param{sid.pointer(), source.pointer()}})
	
	for _, v := range PostgresqlTypes {
		param := &Param{Name: fmt.Sprintf("T%s", strcase.ToCamel(v.Type)), Type: v.Default, Tag: fmt.Sprintf("t_%s", v.Type)}
		entities[0].Params = append(entities[0].Params, param)
		entities[1].Params = append(entities[1].Params, param.forkType("*"+param.Type))
		entities[2].Params = append(entities[2].Params, param.forkType("[]"+param.Type))
		entities[3].Params = append(entities[3].Params, param.forkType("[]*"+param.Type))
		entities = append(entities, &Entity{Name: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type)), Params: []*Param{sid, source, param}})
		entities = append(entities, &Entity{Name: fmt.Sprintf("%sArrayOriginal", strcase.ToCamel(v.Type)), Params: []*Param{sid, source, param.forkType("[]" + param.Type)}})
		entities = append(entities, &Entity{Name: fmt.Sprintf("%sPointer", strcase.ToCamel(v.Type)), Params: []*Param{sid.forkType("*" + sid.Type), source.forkType("*" + source.Type), param.forkType("*" + param.Type)}})
		entities = append(entities, &Entity{Name: fmt.Sprintf("%sArrayPointer", strcase.ToCamel(v.Type)), Params: []*Param{sid.forkType("*" + sid.Type), source.forkType("*" + source.Type), param.forkType("[]*" + param.Type)}})
	}
	RenderEntity("./test/postgresql/make_entity.go", header, entities)
}

func makePostgresqlXML() {
	var (
		insertStatements []*Statement
		selectStatements []*Statement
		updateStatements []*Statement
		deleteStatements []*Statement
	)
	for _, v := range PostgresqlTypes {
		iName := SName{Action: "Insert", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		insert := &Statement{
			Tag:           "insert",
			ShowParameter: true,
			Params:        []*Param{{Name: "sid"}, {Name: "source"}, {Name: fmt.Sprintf("var_%s", v.Type)}},
			Sql:           fmt.Sprintf("insert into types(sid,source,%s) values(#{sid},#{source},#{%s});", fmt.Sprintf("t_%s", v.Type), fmt.Sprintf("var_%s", v.Type)),
		}
		sName := SName{Action: "Select", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		_select := &Statement{
			Tag:           "select",
			ShowParameter: true,
			ShowResult:    true,
			Params:        []*Param{{Name: "sid"}},
			Result:        []*Param{{Name: fmt.Sprintf("t_%s", v.Type)}},
			Sql:           fmt.Sprintf("select t_%s from types where sid = #{sid};", v.Type),
		}
		uName := SName{Action: "Update", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		update := &Statement{
			Tag:           "update",
			ShowParameter: true,
			Params:        []*Param{{Name: "sid"}, {Name: "source"}, {Name: "val"}},
			Result:        []*Param{{Name: fmt.Sprintf("t_%s", v.Type)}},
			Sql:           fmt.Sprintf("update types set t_%s = #{ val } where sid = #{sid};", v.Type),
		}
		dName := SName{Action: "Delete", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		_delete := &Statement{
			Tag:           "delete",
			ShowParameter: true,
			Params:        []*Param{{Name: "sid"}},
			Sql:           fmt.Sprintf("delete from types where sid = #{sid};"),
		}
		item := &Param{Name: "item:struct"}
		insertStatements = append(insertStatements,
			insert.ForkId(iName.ParameterOriginal(false)),
			insert.ForkId(iName.ParameterOriginalPointer(false)),
			insert.ForkId(iName.ParameterPointerOriginal(false)),
			insert.ForkId(iName.ParameterPointerPointer(false)),
			insert.ForkId(iName.EntityOriginal(false)).ForkParams([]*Param{item}),
			insert.ForkId(iName.EntityOriginalPointer(false)).ForkParams([]*Param{item}),
			insert.ForkId(iName.EntityPointerOriginal(false)).ForkParams([]*Param{item}),
			insert.ForkId(iName.EntityPointerPointer(false)).ForkParams([]*Param{item}),
		)
		selectStatements = append(selectStatements,
			_select.ForkId(sName.ParameterOriginal(false)),
			_select.ForkId(sName.ParameterPointerOriginal(false)),
			_select.ForkId(sName.ParameterOriginalPointer(false)),
			_select.ForkId(sName.ParameterPointerPointer(false)),
			_select.ForkId(sName.EntityOriginal(false)).ForkParams([]*Param{item}),
			_select.ForkId(sName.EntityOriginalPointer(false)).ForkParams([]*Param{item}),
			_select.ForkId(sName.EntityPointerOriginal(false)).ForkParams([]*Param{item}),
			_select.ForkId(sName.EntityPointerPointer(false)).ForkParams([]*Param{item}),
		)
		updateStatements = append(updateStatements,
			update.ForkId(uName.ParameterOriginal(false)),
			update.ForkId(uName.ParameterOriginalPointer(false)),
			update.ForkId(uName.ParameterPointerOriginal(false)),
			update.ForkId(uName.ParameterPointerPointer(false)),
			update.ForkId(uName.EntityOriginal(false)).ForkParams([]*Param{item}),
			update.ForkId(uName.EntityOriginalPointer(false)).ForkParams([]*Param{item}),
			update.ForkId(uName.EntityPointerOriginal(false)).ForkParams([]*Param{item}),
			update.ForkId(uName.EntityPointerPointer(false)).ForkParams([]*Param{item}),
		)
		deleteStatements = append(deleteStatements,
			_delete.ForkId(dName.ParameterOriginal(false)),
			_delete.ForkId(dName.ParameterOriginalPointer(false)),
			_delete.ForkId(dName.ParameterPointerOriginal(false)),
			_delete.ForkId(dName.ParameterPointerPointer(false)),
			_delete.ForkId(dName.EntityOriginal(false)).ForkParams([]*Param{item}),
			_delete.ForkId(dName.EntityOriginalPointer(false)).ForkParams([]*Param{item}),
			_delete.ForkId(dName.EntityPointerOriginal(false)).ForkParams([]*Param{item}),
			_delete.ForkId(dName.EntityPointerPointer(false)).ForkParams([]*Param{item}),
		)
		if !v.Array {
			continue
		}
		insert = insert.ForkSql(fmt.Sprintf("insert into array_types(sid,source,%s) values(#{sid},#{source},#{%s});", fmt.Sprintf("t_%s", v.Type), fmt.Sprintf("var_%s", v.Type)))
		_select = _select.ForkSql(fmt.Sprintf("select t_%s from array_types where sid = #{sid};", v.Type))
		update = update.ForkSql(fmt.Sprintf("update array_types set source =#{source}, t_%s = #{ val } where sid = #{sid};", v.Type))
		_delete = _delete.ForkSql(fmt.Sprintf("delete from array_types where sid = #{sid};"))
		insertStatements = append(insertStatements,
			insert.ForkId(iName.ParameterOriginal(true)),
			insert.ForkId(iName.ParameterOriginalPointer(true)),
			insert.ForkId(iName.ParameterPointerOriginal(true)),
			insert.ForkId(iName.ParameterPointerPointer(true)),
			insert.ForkId(iName.EntityOriginal(true)),
			insert.ForkId(iName.EntityOriginalPointer(true)),
			insert.ForkId(iName.EntityPointerOriginal(true)),
			insert.ForkId(iName.EntityPointerPointer(true)),
		)
		selectStatements = append(selectStatements,
			_select.ForkId(sName.ParameterOriginal(true)),
			_select.ForkId(sName.ParameterOriginalPointer(true)),
			_select.ForkId(sName.ParameterPointerOriginal(true)),
			_select.ForkId(sName.ParameterPointerPointer(true)),
			_select.ForkId(sName.EntityOriginal(true)),
			_select.ForkId(sName.EntityOriginalPointer(true)),
			_select.ForkId(sName.EntityPointerOriginal(true)),
			_select.ForkId(sName.EntityPointerPointer(true)),
		)
		updateStatements = append(updateStatements,
			update.ForkId(uName.ParameterOriginal(true)),
			update.ForkId(uName.ParameterOriginalPointer(true)),
			update.ForkId(uName.ParameterPointerOriginal(true)),
			update.ForkId(uName.ParameterPointerPointer(true)),
			update.ForkId(uName.EntityOriginal(true)).ForkParams([]*Param{item}),
			update.ForkId(uName.EntityOriginalPointer(true)).ForkParams([]*Param{item}),
			update.ForkId(uName.EntityPointerOriginal(true)).ForkParams([]*Param{item}),
			update.ForkId(uName.EntityPointerPointer(true)).ForkParams([]*Param{item}),
		)
		deleteStatements = append(deleteStatements,
			_delete.ForkId(dName.ParameterOriginal(true)),
			_delete.ForkId(dName.ParameterOriginalPointer(true)),
			_delete.ForkId(dName.ParameterPointerOriginal(true)),
			_delete.ForkId(dName.ParameterPointerPointer(true)),
			_delete.ForkId(dName.EntityOriginal(true)).ForkParams([]*Param{item}),
			_delete.ForkId(dName.EntityOriginalPointer(true)).ForkParams([]*Param{item}),
			_delete.ForkId(dName.EntityPointerOriginal(true)).ForkParams([]*Param{item}),
			_delete.ForkId(dName.EntityPointerPointer(true)).ForkParams([]*Param{item}),
		)
	}
	RenderStatements("./test/postgresql/sql/make/make_insert.xml", insertStatements)
	RenderStatements("./test/postgresql/sql/make/make_select.xml", selectStatements)
	RenderStatements("./test/postgresql/sql/make/make_update.xml", updateStatements)
	RenderStatements("./test/postgresql/sql/make/make_delete.xml", deleteStatements)
}

func makePostgresqlMapper() {
	header := GoHeader{
		Package: "postgresql",
		Imports: []string{
			//"database/sql",
			//"github.com/gobatis/gobatis",
			//"context",
		},
	}
	var (
		insertMethods []*Method
		selectMethods []*Method
		updateMethods []*Method
		deleteMethods []*Method
	)
	for _, v := range PostgresqlTypes {
		iName := SName{Action: "Insert", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		sName := SName{Action: "Select", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		uName := SName{Action: "Update", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		dName := SName{Action: "Delete", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		sid := &Param{Name: "sid", Type: "string"}
		source := &Param{Name: "source", Type: "string"}
		_int := &Param{Type: "int"}
		_err := &Param{Type: "error"}
		insertMethods = append(insertMethods,
			&Method{Name: iName.ParameterOriginal(false), In: []*Param{sid, source, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []*Param{_int, _err}},
			&Method{Name: iName.ParameterOriginalPointer(false), In: []*Param{sid, source, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []*Param{_int.pointer(), _err}},
			&Method{Name: iName.ParameterPointerOriginal(false), In: []*Param{sid.pointer(), source.pointer(), {Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []*Param{_int, _err}},
			&Method{Name: iName.ParameterPointerPointer(false), In: []*Param{sid, source, {Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []*Param{_int.pointer(), _err}},
			&Method{Name: iName.EntityOriginal(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: iName.EntityOriginalPointer(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: iName.EntityPointerOriginal(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("*%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: iName.EntityPointerPointer(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("*%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
		)
		selectMethods = append(selectMethods,
			&Method{Name: sName.ParameterOriginal(false), In: []*Param{sid}, Out: []*Param{{Type: v.Default}, {Type: "error"}}},
			&Method{Name: sName.ParameterOriginalPointer(false), In: []*Param{sid}, Out: []*Param{{Type: "*" + v.Default}, {Type: "error"}}},
			&Method{Name: sName.ParameterPointerOriginal(false), In: []*Param{sid}, Out: []*Param{{Type: v.Default}, {Type: "error"}}},
			&Method{Name: sName.ParameterPointerPointer(false), In: []*Param{sid}, Out: []*Param{{Type: "*" + v.Default}, {Type: "error"}}},
			&Method{Name: sName.EntityOriginal(false), In: []*Param{sid}, Out: []*Param{{Type: "*" + v.Default}, {Type: "error"}}},
			&Method{Name: sName.EntityOriginalPointer(false), In: []*Param{sid}, Out: []*Param{{Type: "*" + v.Default}, {Type: "error"}}},
			&Method{Name: sName.EntityPointerOriginal(false), In: []*Param{sid}, Out: []*Param{{Type: "*" + v.Default}, {Type: "error"}}},
			&Method{Name: sName.EntityPointerPointer(false), In: []*Param{sid}, Out: []*Param{{Type: "*" + v.Default}, {Type: "error"}}},
		)
		updateMethods = append(updateMethods,
			&Method{Name: uName.ParameterOriginal(false), In: []*Param{sid, source, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []*Param{_int, _err}},
			&Method{Name: uName.ParameterOriginalPointer(false), In: []*Param{sid, source, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []*Param{_int.pointer(), _err}},
			&Method{Name: uName.ParameterPointerOriginal(false), In: []*Param{sid.pointer(), source.pointer(), {Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []*Param{_int, _err}},
			&Method{Name: uName.ParameterPointerPointer(false), In: []*Param{sid, source, {Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []*Param{_int.pointer(), _err}},
			&Method{Name: uName.EntityOriginal(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: uName.EntityOriginalPointer(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: uName.EntityPointerOriginal(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("*%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: uName.EntityPointerPointer(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("*%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
		)
		deleteMethods = append(deleteMethods,
			&Method{Name: dName.ParameterOriginal(false), In: []*Param{sid}, Out: []*Param{_int, _err}},
			&Method{Name: dName.ParameterOriginalPointer(false), In: []*Param{sid}, Out: []*Param{_int.pointer(), _err}},
			&Method{Name: dName.ParameterPointerOriginal(false), In: []*Param{sid.pointer()}, Out: []*Param{_int, _err}},
			&Method{Name: dName.ParameterPointerPointer(false), In: []*Param{sid}, Out: []*Param{_int.pointer(), _err}},
			&Method{Name: dName.EntityOriginal(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: dName.EntityOriginalPointer(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: dName.EntityPointerOriginal(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("*%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
			&Method{Name: dName.EntityPointerPointer(false), In: []*Param{{Name: "item", Type: fmt.Sprintf("*%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{_err}},
		)
		if !v.Array {
			continue
		}
		insertMethods = append(insertMethods,
			&Method{Name: iName.ParameterOriginal(true), In: []*Param{sid, source, {Name: "items", Type: arrayType(v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: iName.ParameterOriginalPointer(true), In: []*Param{sid, source, {Name: "items", Type: arrayType(v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: iName.ParameterPointerOriginal(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: iName.ParameterPointerPointer(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: iName.EntityOriginal(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: iName.EntityOriginalPointer(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: iName.EntityPointerOriginal(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: iName.EntityPointerPointer(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
		)
		selectMethods = append(selectMethods,
			&Method{Name: sName.ParameterOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType(v.Default)}, {Type: "error"}}},
			&Method{Name: sName.ParameterOriginalPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType(v.Default)}, {Type: "error"}}},
			&Method{Name: sName.ParameterPointerOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType("*" + v.Default)}, {Type: "error"}}},
			&Method{Name: sName.ParameterPointerPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType("*" + v.Default)}, {Type: "error"}}},
			&Method{Name: sName.EntityOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType("*" + v.Default)}, {Type: "error"}}},
			&Method{Name: sName.EntityOriginalPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType("*" + v.Default)}, {Type: "error"}}},
			&Method{Name: sName.EntityPointerOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType("*" + v.Default)}, {Type: "error"}}},
			&Method{Name: sName.EntityPointerPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: arrayType("*" + v.Default)}, {Type: "error"}}},
		)
		updateMethods = append(updateMethods,
			&Method{Name: uName.ParameterOriginal(true), In: []*Param{sid, source, {Name: "items", Type: arrayType(v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: uName.ParameterOriginalPointer(true), In: []*Param{sid, source, {Name: "items", Type: arrayType(v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: uName.ParameterPointerOriginal(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: uName.ParameterPointerPointer(true), In: []*Param{sid, source, {Name: "items", Type: arrayType("*" + v.Default)}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: uName.EntityOriginal(true), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: uName.EntityOriginalPointer(true), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: uName.EntityPointerOriginal(true), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type))}}, Out: []*Param{{Type: "error"}}},
			&Method{Name: uName.EntityPointerPointer(true), In: []*Param{{Name: "item", Type: fmt.Sprintf("%sPointer", strcase.ToCamel(v.Type))}}, Out: []*Param{{Type: "error"}}},
		)
		deleteMethods = append(deleteMethods,
			&Method{Name: dName.ParameterOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: dName.ParameterOriginalPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: dName.ParameterPointerOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: dName.ParameterPointerPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: dName.EntityOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: dName.EntityOriginalPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: dName.EntityPointerOriginal(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: dName.EntityPointerPointer(true), In: []*Param{{Name: "id", Type: "int"}}, Out: []*Param{{Type: "int"}, {Type: "error"}}},
		)
	}
	
	var methods []*Method
	methods = append(methods, insertMethods...)
	methods = append(methods, selectMethods...)
	methods = append(methods, updateMethods...)
	methods = append(methods, deleteMethods...)
	RenderMapper("./test/postgresql/make_mapper.go", header, methods)
}

type SelectCaseData struct {
	MockFunc   string
	OutPointer bool
	InPointer  bool
	IsEntity   bool
	InsertFunc string
	UpdateFunc string
	SelectFunc string
	DeleteFunc string
}

const selectCaseTpl = `
{% autoescape off %}
sid := manager.NextId()
v := _mock.{{ MockFunc }}()
rows, err := mapper.{{ InsertFunc }}(sid, "{{ InsertFunc }}",v)
require.NoError(t, err, v)
require.Equal(t, 1, {% if OutPointer %}*{%endif%}rows)

r, err := mapper.{{ SelectFunc }}(sid)
require.NoError(t, err, sid)
require.Equal(t, v, {% if OutPointer %}*{%endif%}r)

v = _mock.{{ MockFunc }}()
rows, err =  mapper.{{ UpdateFunc }}(sid, "{{ UpdateFunc }}", v)
require.NoError(t, err, sid, v)
require.Equal(t, 1, {% if OutPointer %}*{%endif%}rows)

r, err = mapper.{{ SelectFunc }}(sid)
require.NoError(t, err, sid)
require.Equal(t, v, {% if OutPointer %}*{%endif%}r)

rows, err = mapper.{{ DeleteFunc }}(sid)
require.NoError(t, err, sid)
require.Equal(t, 1, {% if OutPointer %}*{%endif%}rows)
{% endautoescape %}
`

func makePostgresqlCases() {
	header := GoHeader{
		Package: "postgresql",
		Imports: []string{
			"testing",
			"github.com/stretchr/testify/require",
			"github.com/gozelle/_mock",
			//"github.com/gobatis/gobatis",
			//"github.com/gobatis/gobatis/driver/postgresql",
			//"github.com/AlekSi/pointer",
			"github.com/gobatis/gobatis/test/generator",
		},
	}
	var testCacses []*TestCase
	for _, v := range PostgresqlTypes {
		iName := SName{Action: "Insert", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		sName := SName{Action: "Select", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		uName := SName{Action: "Update", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		dName := SName{Action: "Delete", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		testCacses = append(testCacses,
			&TestCase{Code: RenderTpl(selectCaseTpl, SelectCaseData{
				MockFunc:   strings.Title(v.Default),
				InsertFunc: iName.ParameterOriginal(false),
				SelectFunc: sName.ParameterOriginal(false),
				UpdateFunc: uName.ParameterOriginal(false),
				DeleteFunc: dName.ParameterOriginal(false),
			})},
			&TestCase{Code: RenderTpl(selectCaseTpl, SelectCaseData{
				MockFunc:   strings.Title(v.Default),
				OutPointer: true,
				InsertFunc: iName.ParameterOriginalPointer(false),
				SelectFunc: sName.ParameterOriginalPointer(false),
				UpdateFunc: uName.ParameterOriginalPointer(false),
				DeleteFunc: dName.ParameterOriginalPointer(false),
			})},
		)
		//if v.Array {
		//	testCacses = append(testCacses,
		//		&TestCase{
		//			Code: fmt.Sprintf("err = mapper.%s([]%s{_mock.%s(),_mock.%s(),_mock.%s()})\n%4srequire.NoError(t, err)",
		//				iName.ParameterOriginal(true),
		//				v.Default,
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				" ",
		//			),
		//		},
		//		&TestCase{
		//			Code: fmt.Sprintf("err = mapper.%s([]%s{_mock.%s(),_mock.%s(),_mock.%s()})\n%4srequire.NoError(t, err)",
		//				iName.ParameterOriginal(true),
		//				v.Default,
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				" ",
		//			),
		//		},
		//		&TestCase{
		//			Code: fmt.Sprintf("err = mapper.%s([]*%s{pointer.To%s(_mock.%s()),pointer.To%s(_mock.%s()),pointer.To%s(_mock.%s()),})\n%4srequire.NoError(t, err)",
		//				iName.ParameterPointerOriginal(true),
		//				v.Default,
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				strcase.ToCamel(v.Default),
		//				" ",
		//			),
		//		},
		//	)
		//}
	}
	RenderTestcases("./test/postgresql/make_mapper_test.go", header, testCacses)
}
