package parser

import (
	"fmt"
	"github.com/koyeo/gobatis/compiler"
)

func newAttributeDuplicateErr(attr *compiler.XMLAttribute) error {
	return fmt.Errorf(
		"duplicate attribute: %s at line:%d column: %d",
		attr.Name,
		attr.Start.Line,
		attr.Start.Column,
	)
}

func newAttributeNotSupportErr(attr *compiler.XMLAttribute) error {
	return fmt.Errorf(
		" attribute: %s not support at line:%d column: %d",
		attr.Name,
		attr.Start.Line,
		attr.Start.Column,
	)
}

func newNodeNotSupportErr(xmlNode *compiler.XMLNode) error {
	return fmt.Errorf(
		"node: %s not suport at line:%d column: %d",
		xmlNode.Name,
		xmlNode.Start.Line,
		xmlNode.Start.Column,
	)
}

func newNodeDuplicateErr(xmlNode *compiler.XMLNode) error {
	return fmt.Errorf(
		"duplicate node: %s not suport at line:%d column: %d",
		xmlNode.Name,
		xmlNode.Start.Line,
		xmlNode.Start.Column,
	)
}
