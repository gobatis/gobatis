module github.com/gobatis/gobatis

go 1.13

require (
	github.com/antlr/antlr4 v0.0.0-20210311221813-5e5b6d35b418
	github.com/fatih/structs v1.1.0
	github.com/flosch/pongo2/v4 v4.0.2
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/gozelle/_mock v0.0.0-20211114021104-dc75ce553720
	github.com/gozelle/decimal v1.3.2-0.20211117124208-0a2df7a603e8
	github.com/iancoleman/strcase v0.2.0
	github.com/jackc/pgconn v1.10.0
	github.com/jackc/pgtype v1.8.1
	github.com/jackc/pgx/v4 v4.13.0
	github.com/jinzhu/copier v0.3.2
	github.com/koyeo/_log v0.0.0-20211006130730-4b7537dd3986
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749 // indirect
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/stretchr/testify v1.7.0
	github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/tools v0.0.0-20201125231158-b5590deeca9b // indirect
)

replace github.com/gozelle/_mock v0.0.0-20211114021104-dc75ce553720 => ../_mock
