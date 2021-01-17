package engine

import "encoding/json"

type Method struct {
	Name          string
	ParameterType string
	ResultType    string
	SQL           string
}

func (p *Method) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}
