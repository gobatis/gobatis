package make

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"io/ioutil"
)

type Database struct {
	Name   string
	Dist   string
	Fields []Field
}

type Field struct {
	Var      string
	Tag      string
	Types    []string
	Array    bool
	Default  string
	Accept   []string
	Reject   []string
	Packages []string
}

func NewMaker(database Database) *Maker {
	return &Maker{database: database}
}

type Maker struct {
	database Database
}

func Make() (err error) {
	//err = MakeEntity()
	//if err != nil {
	//	return
	//}
	return
}

func (p Maker) MakeEntity() (err error) {
	c := "type GoType struct {\n"
	for _, v := range p.database.Fields {
		c += fmt.Sprintf("%4s%-30s%-30s`sql:\"%s\"`\n", "", strcase.ToCamel(v.Tag), v.Tag, v.Tag)
	}
	c += "}"
	fmt.Println(c)
	return
}

func makeMapper() {

}

func makeXML() {

}

func makeTestCase() {

}

func writeFile(path, content string) (err error) {
	err = ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return
	}
	return
}
