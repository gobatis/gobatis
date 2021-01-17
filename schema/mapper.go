package schema

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func UnmarshalMapper(data []byte) (mapper *Mapper, err error) {
	mapper = new(Mapper)
	err = xml.Unmarshal(data, mapper)
	if err != nil {
		err = fmt.Errorf("parse sql xml error: %s", err)
		return
	}
	return
}

type Mapper struct {
	XMLName   xml.Name  `xml:"mapper" json:"mapper,omitempty"`
	Resource  string    `xml:"resource,attr" json:"resource,omitempty"`
	Namespace string    `xml:"namespace,attr" json:"namespace,omitempty"`
	Selects   []*Select `xml:"select" json:"selects,omitempty"`
}

func (p Mapper) String() string {
	d, _ := json.MarshalIndent(p, "", "\t")
	return string(d)
}
