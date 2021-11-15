package generator

type Type struct {
	Var      string
	Type     string
	Types    []string
	Array    bool
	Default  string
	Accept   []string
	Reject   []string
	Packages []string
}
