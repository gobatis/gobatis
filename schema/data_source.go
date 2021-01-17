package schema

import "encoding/xml"

type DataSource struct {
	XMLName     xml.Name    `xml:"dataSource" json:"xmlName,omitempty"`
	Type        string      `xml:"type,attr" json:"type,omitempty"`
	Properties  []*Property `xml:"property" json:"properties,omitempty"`
	propertyMap *PropertyMap
}

func (p *DataSource) PropertyMap() *PropertyMap {
	if p.propertyMap == nil {
		p.propertyMap = NewPropertyMap()
		p.propertyMap.Add(p.Properties...)
	}
	return p.propertyMap
}

func (p *DataSource) Driver() string {
	v := p.PropertyMap().Get(DRIVER)
	if v != nil {
		return v.Value
	}
	return ""
}

func (p *DataSource) URL() string {
	v := p.PropertyMap().Get(URL)
	if v != nil {
		return v.Value
	}
	return ""
}

func (p *DataSource) Username() string {
	v := p.PropertyMap().Get(USERNAME)
	if v != nil {
		return v.Value
	}
	return ""
}

func (p *DataSource) Password() string {
	v := p.PropertyMap().Get(PASSWORD)
	if v != nil {
		return v.Value
	}
	return ""
}

func (p *DataSource) Database() string {
	v := p.PropertyMap().Get(DATABASE)
	if v != nil {
		return v.Value
	}
	return ""
}
