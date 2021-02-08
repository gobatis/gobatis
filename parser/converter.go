package parser

import (
	"github.com/koyeo/gobatis/compiler"
	"github.com/koyeo/gobatis/dtd"
)

func ParseConfiguration(nodes []*compiler.XMLNode) (configuration *Node, err error) {

	configuration = NewNode(dtd.CONFIGURATION)
	for _, node := range nodes {
		if node.Name == dtd.CONFIGURATION {
			err = parseConfigurationNode(configuration, node, dtd.Configuration)
			if err != nil {
				return
			}
			break
		}
	}
	return
}

func parseXMLAttribute(node *Node, xmlNode *compiler.XMLNode, elem *dtd.Element) (err error) {
	for _, attr := range xmlNode.Attributes {
		// 属性不支持
		if !elem.HasAttribute(attr.Name) {
			return newAttributeNotSupportErr(attr)
		}
		// 属性重复
		if node.hasAttribute(attr.Name) {
			return newAttributeDuplicateErr(attr)
		}
		node.addAttribute(attr.Name, attr.Value)
	}
	return
}

func parseConfigurationNode(node *Node, xmlNode *compiler.XMLNode, elem *dtd.Element) (err error) {
	err = parseXMLAttribute(node, xmlNode, elem)
	if err != nil {
		return
	}
	for _, childXmlNode := range xmlNode.Body {
		// 子节点不支持
		if !elem.HasNode(childXmlNode.Name) {
			return newNodeNotSupportErr(childXmlNode)
		}
		// 子节点重复错误
		if elem.GetNodeCount(childXmlNode.Name) == dtd.AT_MOST_ONCE &&
			node.count(childXmlNode.Name) > 0 {
			return newNodeDuplicateErr(childXmlNode)
		}
		// 判断是否解析 SQL
		childNode := NewNode(childXmlNode.Name)
		node.addNode(childNode)

	}

	// 判断是否包换必填属性
	// 判断是否包含必填节点

	return nil
}
