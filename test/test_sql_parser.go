package main

import (
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func main() {
	stmt, err := sqlparser.Parse("SELECT e.employee_id, e.employee_name, d.dept_name FROM EmployeeTB AS e, DeptTB AS d WHERE e.dept_id=d.dept_id;")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("stmt = %+v\n", stmt)
	//d, _ := json.MarshalIndent(stmt, "", "\t")
	switch s := stmt.(type) {
	case *sqlparser.Select:
		fmt.Println(s.Where)
	default:
		fmt.Println("unknown type")
	}
	//fmt.Println(string(d))
}
