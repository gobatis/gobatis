package schema

import "encoding/xml"

type Property struct {
	XMLName xml.Name `xml:"property" json:"property,omitempty"`
	Name    string   `xml:"name,attr" json:"name,omitempty"`
	Value   string   `xml:"value,attr" json:"value,omitempty"`
}


