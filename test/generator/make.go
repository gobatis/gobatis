package generator

import (
	"fmt"
	"github.com/iancoleman/strcase"
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
	for _, v := range PostgresqlTypes {
		param := &Param{
			Name: fmt.Sprintf("T%s", strcase.ToCamel(v.Type)),
			Type: v.Default,
			Tag:  fmt.Sprintf("t_%s", v.Type),
		}
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
	)
	
	for _, v := range PostgresqlTypes {
		iname := SName{
			Action: "Insert",
			Name:   strcase.ToCamel(v.Type),
			Type:   strcase.ToCamel(v.Default),
		}
		insert := &Statement{
			Tag:           "insert",
			Id:            iname.ParameterOriginal(),
			ShowParameter: true,
			Params: []Param{
				{
					Name: fmt.Sprintf("var_%s", v.Type),
				},
			},
			Sql: fmt.Sprintf("insert into types(%s) values(#{%s});", fmt.Sprintf("t_%s", v.Type), fmt.Sprintf("var_%s", v.Type)),
		}
		sname := SName{
			Action: "Select",
			Name:   strcase.ToCamel(v.Type),
			Type:   strcase.ToCamel(v.Default),
		}
		_select := &Statement{
			Tag:           "select",
			Id:            sname.ParameterOriginal(),
			ShowParameter: true,
			Params: []Param{
				{
					Name: "id",
				},
			},
			Sql: fmt.Sprintf("select t_%s from types where id = #{id};", v.Type),
		}
		insertStatements = append(insertStatements,
			insert,
			insert.ForkId(iname.ParameterPointerOriginal()),
			insert.ForkId(iname.EntityOriginal()),
			insert.ForkId(iname.EntityPointerOriginal()),
			insert.ForkId(iname.ParameterEmbed()).ForkSql(fmt.Sprintf("insert into types(%s) values(${%s});", fmt.Sprintf("t_%s", v.Type), fmt.Sprintf("var_%s", v.Type))),
		)
		selectStatements = append(selectStatements,
			_select,
			_select.ForkId(sname.ParameterPointerOriginal()),
			_select.ForkId(sname.EntityOriginal()),
			_select.ForkId(sname.EntityPointerOriginal()),
		)
	}
	RenderStatements("./test/postgresql/sql/make_insert.xml", insertStatements)
	RenderStatements("./test/postgresql/sql/make_select.xml", selectStatements)
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
	var methods []*Method
	for _, v := range PostgresqlTypes {
		sname := SName{
			Action: "Insert",
			Name:   strcase.ToCamel(v.Type),
			Type:   strcase.ToCamel(v.Default),
		}
		methods = append(methods,
			&Method{Name: sname.ParameterOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.ParameterPointerOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: sname.EntityOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: sname.EntityPointerOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: "*" + v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: sname.ParameterTx(), In: []Param{{Name: "tx", Type: "sql.Tx"}, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: sname.ParameterRows(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "int"}, {Type: "error"}}},
			//&Method{Name: sname.ParameterContext(), In: []Param{{Name: "ctx", Type: "context.Context"}, {Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: sname.ParameterStmt(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Name: "stmt", Type: "*gobatis.Stmt"}, {Name: "err", Type: "error"}}},
			//&Method{Name: sname.ParameterMust(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			//&Method{Name: sname.ParameterEmbed(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Type), Type: v.Default}}, Out: []Param{{Type: "error"}}},
		)
	}
	RenderMapper("./test/postgresql/make_mapper.go", header, methods)
}

func makePostgresqlCases() {
	header := GoHeader{
		Package: "postgresql",
		Imports: []string{
			"testing",
			"github.com/stretchr/testify/require",
			"github.com/gozelle/_mock",
			"github.com/gobatis/gobatis",
			"github.com/AlekSi/pointer",
		},
	}
	var testCacses []*TestCase
	for _, v := range PostgresqlTypes {
		sname := SName{
			Action: "Insert",
			Name:   strcase.ToCamel(v.Type),
			Type:   strcase.ToCamel(v.Default),
		}
		testCacses = append(testCacses,
			&TestCase{
				Code: fmt.Sprintf("err = mapper.%s(_mock.%s())\n%4srequire.NoError(t, err)",
					sname.ParameterOriginal(), strcase.ToCamel(v.Default), " "),
			},
			&TestCase{
				Code: fmt.Sprintf("err = mapper.%s(pointer.To%s(_mock.%s()))\n%4srequire.NoError(t, err)",
					sname.ParameterPointerOriginal(), strcase.ToCamel(v.Default), strcase.ToCamel(v.Default), " "),
			},
		)
	}
	RenderTestcases("./test/postgresql/make_mapper_test.go", header, testCacses)
}
