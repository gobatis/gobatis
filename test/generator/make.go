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
	var entities []*Entity
	entities = append(entities, &Entity{Name: "TypeOriginal"})
	entities = append(entities, &Entity{Name: "TypePointer"})
	entities = append(entities, &Entity{Name: "ArrayTypeOriginal"})
	entities = append(entities, &Entity{Name: "ArrayTypePointer"})
	for _, v := range PostgresqlTypes {
		param := &Param{Name: fmt.Sprintf("T%s", strcase.ToCamel(v.Type)), Type: v.Default, Tag: fmt.Sprintf("t_%s", v.Type)}
		entities[0].Params = append(entities[0].Params, param)
		entities[1].Params = append(entities[1].Params, param.forkType("*"+param.Type))
		entities = append(entities, &Entity{
			Name:   fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Type)),
			Params: []*Param{param},
		})
		param.Type = "*" + param.Type
		entities = append(entities, &Entity{
			Name:   fmt.Sprintf("%sPointer", strcase.ToCamel(v.Type)),
			Params: []*Param{param},
		})
		
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
			Id:            iName.ParameterOriginal(false),
			ShowParameter: true,
			Params:        []Param{{Name: fmt.Sprintf("var_%s", v.Type)}},
			Sql:           fmt.Sprintf("insert into types(%s) values(#{%s});", fmt.Sprintf("t_%s", v.Type), fmt.Sprintf("var_%s", v.Type)),
		}
		sName := SName{Action: "Select", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		_select := &Statement{
			Tag:           "select",
			Id:            sName.ParameterOriginal(false),
			ShowParameter: true,
			ShowResult:    true,
			Params:        []Param{{Name: "id"}},
			Result:        []Param{{Name: fmt.Sprintf("t_%s", v.Type)}},
			Sql:           fmt.Sprintf("select t_%s from types where id = #{id};", v.Type),
		}
		uName := SName{Action: "Update", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		update := &Statement{
			Tag:           "update",
			Id:            uName.ParameterOriginal(false),
			ShowParameter: true,
			Params:        []Param{{Name: "id"}, {Name: "val"}},
			Result:        []Param{{Name: fmt.Sprintf("t_%s", v.Type)}},
			Sql:           fmt.Sprintf("update types set t_%s = #{ val } from types where id = #{id};", v.Type),
		}
		dName := SName{Action: "Delete", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		_delete := &Statement{
			Tag:           "delete",
			Id:            dName.ParameterOriginal(false),
			ShowParameter: true,
			Params:        []Param{{Name: "id"}},
			Sql:           fmt.Sprintf("delete from types where id = #{id};"),
		}
		insertStatements = append(insertStatements,
			insert,
			insert.ForkId(iName.ParameterPointerOriginal(false)),
			insert.ForkId(iName.EntityOriginal(false)),
			insert.ForkId(iName.EntityPointerOriginal(false)),
			insert.ForkId(iName.ParameterEmbed(false)).ForkSql(
				fmt.Sprintf("insert into types(%s) values(${%s});", fmt.Sprintf("t_%s", v.Type), fmt.Sprintf("var_%s", v.Type)),
			),
		)
		selectStatements = append(selectStatements,
			_select,
			_select.ForkId(sName.ParameterPointerOriginal(false)),
			_select.ForkId(sName.ParameterOriginalPointer(false)),
			_select.ForkId(sName.ParameterPointerPointer(false)),
			_select.ForkId(sName.EntityOriginal(false)),
			_select.ForkId(sName.EntityPointerOriginal(false)),
		)
		updateStatements = append(updateStatements,
			update,
			update.ForkId(uName.ParameterPointerOriginal(false)),
			update.ForkId(uName.EntityOriginal(false)),
			update.ForkId(uName.EntityPointerOriginal(false)),
		)
		deleteStatements = append(deleteStatements,
			_delete,
			_delete.ForkId(dName.ParameterPointerOriginal(false)),
			_delete.ForkId(dName.EntityOriginal(false)),
			_delete.ForkId(dName.EntityPointerOriginal(false)),
		)
		if !v.Array {
			continue
		}
		insert = insert.ForkSql(
			fmt.Sprintf("insert into array_types(%s) values(#{%s});", fmt.Sprintf("t_%s", v.Type), fmt.Sprintf("var_%s", v.Type)),
		)
		_select = _select.ForkSql(fmt.Sprintf("select t_%s from array_types where id = #{id};", v.Type))
		insertStatements = append(insertStatements,
			insert.ForkId(iName.ParameterOriginal(true)),
			insert.ForkId(iName.ParameterPointerOriginal(true)),
			insert.ForkId(iName.EntityOriginal(true)),
			insert.ForkId(iName.EntityPointerOriginal(true)),
			insert.ForkId(iName.ParameterEmbed(true)),
		)
		selectStatements = append(selectStatements,
			_select.ForkId(sName.ParameterOriginal(true)),
			_select.ForkId(sName.ParameterPointerOriginal(true)),
			_select.ForkId(sName.ParameterOriginalPointer(true)),
			_select.ForkId(sName.ParameterPointerPointer(true)),
			_select.ForkId(sName.EntityOriginal(true)),
			_select.ForkId(sName.EntityPointerOriginal(true)),
			_select.ForkId(sName.EntityOriginalPointer(true)),
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
			"database/sql",
			"github.com/gobatis/gobatis",
			"context",
		},
	}
	var (
		insertMethods []*Method
		selectMethods []*Method
	)
	for _, v := range PostgresqlTypes {
		iName := SName{Action: "Insert", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		sName := SName{Action: "Select", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		insertMethods = append(insertMethods,
			&Method{Name: iName.ParameterOriginal(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: iName.ParameterPointerOriginal(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: iName.EntityOriginal(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: iName.EntityPointerOriginal(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: iName.ParameterTx(false), In: []Param{{Name: "tx", Type: "sql.Tx"}, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: iName.ParameterRows(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "int"}, {Type: "error"}}},
			//&Method{Name: iName.ParameterContext(false), In: []Param{{Name: "ctx", Type: "context.Context"}, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: iName.ParameterStmt(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Name: "stmt", Type: "*gobatis.Stmt"}, {Name: "err", Type: "error"}}},
			//&Method{Name: iName.ParameterMust(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: iName.ParameterEmbed(false), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
		)
		selectMethods = append(selectMethods,
			&Method{Name: sName.ParameterOriginal(false), In: []Param{{Name: "id", Type: "int"}}, Out: []Param{{Type: v.Default}, {Type: "error"}}},
			&Method{Name: sName.ParameterOriginalPointer(false), In: []Param{{Name: "id", Type: "int"}}, Out: []Param{{Type: "*" + v.Default}, {Type: "error"}}},
		)
		if !v.Array {
			continue
		}
		insertMethods = append(insertMethods,
			&Method{Name: iName.ParameterOriginal(true), In: []Param{{Name: "items", Type: arrayType(v.Default)}}, Out: []Param{{Type: "error"}}},
			&Method{Name: iName.ParameterPointerOriginal(true), In: []Param{{Name: "items", Type: arrayType("*" + v.Default)}}, Out: []Param{{Type: "error"}}},
		)
		selectMethods = append(selectMethods,
			&Method{Name: sName.ParameterOriginal(true), In: []Param{{Name: "id", Type: "int"}}, Out: []Param{{Type: arrayType(v.Default)}, {Type: "error"}}},
			&Method{Name: sName.ParameterOriginalPointer(true), In: []Param{{Name: "id", Type: "int"}}, Out: []Param{{Type: arrayType("*" + v.Default)}, {Type: "error"}}},
		)
	}
	
	var methods []*Method
	methods = append(methods, insertMethods...)
	methods = append(methods, selectMethods...)
	RenderMapper("./test/postgresql/make_mapper.go", header, methods)
}

type SelectCaseData struct {
	MockFunc       string
	InsertOriginal string
	SelectOriginal string
}

const selectCaseTpl = `

id1 := dm.NextId()
v1 := _mock.{{ MockFunc }}()
err := mapper.{{ InsertOriginal }}(v1)
require.NoError(t, err, val)

r1,err := mapper.{{ SelectOriginal }}(id1)
require.NoError(t, err, id)
require.Equal(t, v1, r1)

id2 := dm.NextId()
v2 := _mock.{{ MockFunc }}()
err := mapper.{{ InsertPointer }}(&v2)
require.NoError(t, err, val1)

r2,err := mapper.{{ SelectOriginalPointer }}(id2)
require.NoError(t, err, id2)
require.Equal(t, val, r2)


`

func makePostgresqlCases() {
	header := GoHeader{
		Package: "postgresql",
		Imports: []string{
			"testing",
			"github.com/stretchr/testify/require",
			"github.com/gozelle/_mock",
			"github.com/gobatis/gobatis",
			"github.com/gobatis/gobatis/driver/postgresql",
			"github.com/AlekSi/pointer",
			"github.com/gobatis/gobatis/test/generator",
		},
	}
	var testCacses []*TestCase
	for _, v := range PostgresqlTypes {
		iName := SName{Action: "Insert", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		sName := SName{Action: "Select", Name: strcase.ToCamel(v.Type), Type: strcase.ToCamel(v.Default)}
		testCacses = append(testCacses,
			&TestCase{
				Code: RenderTpl(selectCaseTpl, SelectCaseData{
					MockFunc:       strings.Title(v.Default),
					InsertOriginal: iName.ParameterOriginal(false),
					SelectOriginal: sName.ParameterOriginal(false),
				}),
			},
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
