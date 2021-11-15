package generator

import (
	"fmt"
	"github.com/iancoleman/strcase"
)

var PostgresqlTypes = []Field{
	{Tag: "bigint", Array: true, Default: "int64", Accept: []string{"int*", "uint*"}, Reject: []string{"float*", "string"}},
	{Tag: "int8", Array: true, Default: "int8", Accept: []string{}, Reject: []string{}},
	{Tag: "bigserial", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "serial8", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "bit", Array: true, Default: "byte", Accept: []string{}, Reject: []string{}},
	{Tag: "bit_varying", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "boolean", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "bool", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "box", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "bytea", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "character", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "char", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "character_varying", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "varchar", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "cidr", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "circle", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "date", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "double_precision", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "float8", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "inet", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "integer", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "int", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "int4", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "interval", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "json", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "jsonb", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "line", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "lseg", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "macaddr", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "macaddr8", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "money", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "numeric", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "decimal", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "path", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "pg_lsn", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "pg_snapshot", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "point", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "polygon", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "real", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "float4", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "smallint", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "int2", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "smallserial", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "serial2", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "serial", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "serial4", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "text", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "time", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "time_with_timezone", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "timez", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "timestamp", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "timestamp_with_timezone", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "timestampz", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "tsquery", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "tsvector", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "txid_snapshot", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "uuid", Array: true, Default: "int64", Accept: []string{}, Reject: []string{}},
	{Tag: "xml", Array: true, Default: "string", Accept: []string{}, Reject: []string{}},
}

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
	entities = append(entities, &Entity{Name: "AllType"})
	for _, v := range PostgresqlTypes {
		param := Param{
			Name: fmt.Sprintf("T%s", strcase.ToCamel(v.Tag)),
			Type: v.Default,
			Tag:  fmt.Sprintf("t_%s", v.Tag),
		}
		entities[0].Params = append(entities[0].Params, param)
		entities = append(entities, &Entity{
			Name:   fmt.Sprintf("%sOriginal", strcase.ToCamel(v.Tag)),
			Params: []Param{param},
		})
		param.Type = "*" + param.Type
		entities = append(entities, &Entity{
			Name:   fmt.Sprintf("%sPointer", strcase.ToCamel(v.Tag)),
			Params: []Param{param},
		})
		
	}
	RenderEntity("./test/postgresql/entity.make.go", header, entities)
}

func makePostgresqlXML() {
	var statements []*Statement
	for _, v := range PostgresqlTypes {
		sname := SName{
			Action: "Insert",
			Name:   strcase.ToCamel(v.Tag),
			Type:   strcase.ToCamel(v.Default),
		}
		base := &Statement{
			Tag:           "insert",
			Id:            sname.ParameterOriginal(),
			ShowParameter: true,
			Params: []Param{
				{
					Name: fmt.Sprintf("var_%s", v.Tag),
				},
			},
			Sql: fmt.Sprintf("insert into types(%s) values(#{%s});", fmt.Sprintf("t_%s", v.Tag), fmt.Sprintf("var_%s", v.Tag)),
		}
		statements = append(statements,
			base,
			base.ForkId(sname.ParameterPointerOriginal()),
			base.ForkId(sname.EntityOriginal()),
			base.ForkId(sname.EntityPointerOriginal()),
			base.ForkId(sname.ParameterEmbed()).ForkSql(fmt.Sprintf("insert into types(%s) values(${%s});", fmt.Sprintf("t_%s", v.Tag), fmt.Sprintf("var_%s", v.Tag))),
		)
	}
	RenderStatements("./test/postgresql/sql/insert.make.xml", statements)
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
			Name:   strcase.ToCamel(v.Tag),
			Type:   strcase.ToCamel(v.Default),
		}
		methods = append(methods,
			&Method{Name: sname.ParameterOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.ParameterPointerOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.EntityOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.EntityPointerOriginal(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: "*" + v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.ParameterTx(), In: []Param{{Name: "tx", Type: "sql.Tx"}, {Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.ParameterRows(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "int"}, {Type: "error"}}},
			&Method{Name: sname.ParameterContext(), In: []Param{{Name: "ctx", Type: "context.Context"}, {Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.ParameterStmt(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Name: "stmt", Type: "*gobatis.Stmt"}, {Name: "err", Type: "error"}}},
			&Method{Name: sname.ParameterMust(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "error"}}},
			&Method{Name: sname.ParameterEmbed(), In: []Param{{Name: fmt.Sprintf("var_%s", v.Tag), Type: v.Default}}, Out: []Param{{Type: "error"}}},
		)
	}
	RenderMapper("./test/postgresql/mapper.make.go", header, methods)
}

func makePostgresqlCases() {
	header := GoHeader{
		Package: "postgresql",
		Imports: []string{
			"testing",
			"github.com/stretchr/testify/require",
			"github.com/gozelle/_mock",
		},
	}
	var testCacses []*TestCase
	for _, v := range PostgresqlTypes {
		sname := SName{
			Action: "Insert",
			Name:   strcase.ToCamel(v.Tag),
			Type:   strcase.ToCamel(v.Default),
		}
		testCacses = append(testCacses,
			&TestCase{
				Code: fmt.Sprintf("err = mapper.%s(_mock.%s())\n%4srequire.NoError(t, err)",
					sname.ParameterOriginal(), strcase.ToCamel(v.Default), " "),
			},
		)
	}
	RenderTestcases("./test/postgresql/mapper.make_test.go", header, testCacses)
}
