package schema

import "encoding/xml"

type Select struct {
	XMLName       xml.Name `xml:"select" json:"select,omitempty"`
	Id            string   `xml:"id,attr" json:"id,omitempty"`
	ParameterType string   `xml:"parameterType,attr" json:"parameterType,omitempty"`
	ResultType    string   `xml:"resultType,attr" json:"resultType,omitempty"`
	SQL           string   `xml:",chardata" json:"sql,omitempty"`
}
