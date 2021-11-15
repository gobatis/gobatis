package generator

import (
	"github.com/flosch/pongo2/v4"
	"github.com/jinzhu/copier"
)

type GoHeader struct {
	Package string
	Imports []string
}

type Param struct {
	Name string
	Type string
	Tag  string
}

func (p Param) forkType(t string) *Param {
	n := new(Param)
	err := copier.Copy(n, p)
	if err != nil {
		panic(err)
	}
	n.Type = t
	return n
}

type Statement struct {
	Tag           string
	Id            string
	ShowParameter bool
	Params        []Param
	Sql           string
}

func (p Statement) ForkId(id string) *Statement {
	n := new(Statement)
	err := copier.Copy(n, p)
	if err != nil {
		panic(err)
	}
	n.Id = id
	return n
}

func (p Statement) ForkSql(sql string) *Statement {
	n := new(Statement)
	err := copier.Copy(n, p)
	if err != nil {
		panic(err)
	}
	n.Sql = sql
	return n
}

func (p Statement) ForkParams(params []Param) *Statement {
	n := new(Statement)
	err := copier.Copy(n, p)
	if err != nil {
		panic(err)
	}
	n.Params = params
	return n
}

func RenderStatements(dist string, data []*Statement) {
	tpl, err := pongo2.FromFile(resolve("./test/generator/tpl/statement.tpl"))
	if err != nil {
		panic(err)
	}
	res, err := tpl.Execute(map[string]interface{}{
		"Statements": data,
	})
	if err != nil {
		panic(err)
	}
	err = writeFile(dist, res)
	if err != nil {
		panic(err)
	}
}

type Method struct {
	Name string
	In   []Param
	Out  []Param
}

type Mapper struct {
}

type TestCase struct {
	Name string
	Code string
}

type Entity struct {
	Name   string
	Params []*Param
}

func RenderEntity(dist string, header GoHeader, entities []*Entity) {
	tpl, err := pongo2.FromFile(resolve("./test/generator/tpl/entity.tpl"))
	if err != nil {
		panic(err)
	}
	res, err := tpl.Execute(map[string]interface{}{
		"Header":   header,
		"Entities": entities,
	})
	if err != nil {
		panic(err)
	}
	err = writeFile(dist, gofmt(res))
	if err != nil {
		panic(err)
	}
}

func RenderMapper(dist string, header GoHeader, methods []*Method) {
	tpl, err := pongo2.FromFile(resolve("./test/generator/tpl/mapper.tpl"))
	if err != nil {
		panic(err)
	}
	res, err := tpl.Execute(map[string]interface{}{
		"Header":  header,
		"Methods": methods,
	})
	if err != nil {
		panic(err)
	}
	err = writeFile(dist, gofmt(res))
	if err != nil {
		panic(err)
	}
}

func RenderTestcases(dist string, header GoHeader, testCases []*TestCase) {
	tpl, err := pongo2.FromFile(resolve("./test/generator/tpl/test_case.tpl"))
	if err != nil {
		panic(err)
	}
	res, err := tpl.Execute(map[string]interface{}{
		"Header": header,
		"Cases":  testCases,
	})
	if err != nil {
		panic(err)
	}
	err = writeFile(dist, gofmt(res))
	if err != nil {
		panic(err)
	}
}
