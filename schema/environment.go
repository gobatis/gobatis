package schema

import "encoding/xml"

type Environment struct {
	XMLName            xml.Name            `xml:"environment" json:"environment,omitempty"`
	Id                 string              `xml:"id,attr" json:"id,omitempty"`
	TransactionManager *TransactionManager `xml:"transactionManager" json:"transactionManager,omitempty"`
	DataSource         *DataSource         `xml:"dataSource" json:"dataSource,omitempty"`
}
