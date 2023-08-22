package batis

type Builder interface {
	Build() ([]executor, error)
}
