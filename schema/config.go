package schema

import (
	"encoding/json"
	"encoding/xml"
)

func UnmarshalConfiguration(data []byte) (configuration *Configuration, err error) {
	configuration = new(Configuration)
	err = xml.Unmarshal(data, configuration)
	if err != nil {
		return
	}
	return
}

type Configuration struct {
	XMLName      xml.Name      `xml:"configuration" json:"xmlName,omitempty"`
	Properties   *Properties   `xml:"properties" json:"properties,omitempty"`
	Environments *Environments `xml:"environments" json:"environments,omitempty"`
	TypeAliases  *TypeAliases  `xml:"typeAliases" json:"typeAliases,omitempty"`
	Mappers      *Mappers      `xml:"mappers" json:"mappers,omitempty"`
}

func (p Configuration) String() string {
	d, _ := json.MarshalIndent(p, "", "\t")
	return string(d)
}

func (p *Configuration) Module() string {
	if p.Properties != nil && p.Properties.PropertyMap().Has(MODULE) {
		return p.Properties.PropertyMap().Get(MODULE).Value
	}
	return ""
}

type Environments struct {
	XMLName  xml.Name       `xml:"environments" json:"xmlName,omitempty"`
	Default  string         `xml:"default,attr" json:"default,omitempty"`
	Children []*Environment `xml:"environment" json:"children,omitempty"`
}

type TransactionManager struct {
	XMLName xml.Name `xml:"transactionManager" json:"transactionManager,omitempty"`
	Type    string   `xml:"type,attr" json:"type"`
}

type Properties struct {
	XMLName     xml.Name    `xml:"properties" json:"properties,omitempty"`
	Children    []*Property `xml:"property" json:"children,omitempty"`
	propertyMap *PropertyMap
}

func (p *Properties) PropertyMap() *PropertyMap {
	if p.propertyMap == nil {
		p.propertyMap = NewPropertyMap()
		p.propertyMap.Add(p.Children...)
	}
	return p.propertyMap
}

type Mappers struct {
	XMLName  xml.Name `xml:"mappers" json:"mappers,omitempty"`
	Children []Mapper `xml:"mapper" json:"children,omitempty"`
}

type TypeAliases struct {
	XMLName  xml.Name     `xml:"typeAliases" json:"typeAliases,omitempty"`
	Children []*TypeAlias `xml:"typeAlias" json:"children,omitempty"`
}

type TypeAlias struct {
	XMLName xml.Name `xml:"typeAlias" json:"typeAlias,omitempty"`
	Alias   string   `xml:"alias,attr" json:"alias,omitempty"`
	Type    string   `xml:"type,attr" json:"type,omitempty"`
}
