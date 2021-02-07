package compiler

import (
	"encoding/json"
)

const (
	TT_TEXT         = "text"
	TT_START_TAG    = "tag:start"
	TT_END_TAG      = "tag:end"
	TT_SELF_END_TAG = "tag:self-end"
	TT_ATTR_VALUE   = "attr:value"
	TT_ATTR_NAME    = "attr:name"
)

const (
	TT_ID           = "id"
	TT_COMMA        = ","
	TT_EQUAL        = "="
	TT_SQL_STRUCT   = "sql:struct"
	TT_SQL_DOT      = "sql:."
	TT_SQL_VAR      = "sql:var"
	TT_PROPERTY_VAR = "property:var"
	TT_PROPERTY_DS  = "property:ds"
	TT_PROPERTY_VAL = "property:val"
)

type Token struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Start *Point `json:"start"`
	End   *Point `json:"end"`
}

func (p *Token) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

type Point struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
