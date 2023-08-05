package builder

import "github.com/gobatis/gobatis/executor"

type Builder interface {
	Build() ([]executor.Executor, error)
}
