package compiler

import (
	"fmt"
	"github.com/koyeo/gobatis/dtd"
)

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

type Node struct {
	Name       string
	attributes map[string]string
	nodes      []*Node
	countMap   map[string]int
}

func (p *Node) addAttribute(name, value string) {
	if p.attributes == nil {
		p.attributes = map[string]string{}
	}
	p.attributes[name] = value
}

func (p *Node) hasAttribute(name string) bool {
	if p.attributes == nil {
		return false
	}
	_, ok := p.attributes[name]
	return ok
}

func (p *Node) addNode(node *Node) {
	if p.countMap == nil {
		p.countMap = map[string]int{}
	}
	p.countMap[node.Name]++
	p.nodes = append(p.nodes, node)
}

func (p *Node) count(name string) int {
	if p.countMap == nil {
		return 0
	}
	return p.countMap[name]
}

type NodeParser struct {
	file string
}

func (p *NodeParser) ParseConfigurationFile(nodes []*XMLNode) (configuration *Node, err error) {

	configuration = NewNode(dtd.CONFIGURATION)
	for _, node := range nodes {
		if node.Name == dtd.CONFIGURATION {
			err = p.parseConfigurationNode(configuration, node, dtd.Configuration)
			if err != nil {
				return
			}
			break
		}
	}
	return
}

func (p *NodeParser) ParseMapperFile() (err error) {
	return
}

func (p *NodeParser) parseXMLAttribute(node *Node, xmlNode *XMLNode, elem *dtd.Element) (err error) {
	for _, attr := range xmlNode.Attributes {
		// 属性不支持
		if !elem.HasAttribute(attr.Name) {
			return p.newAttributeNotSupportErr(attr)
		}
		// 属性重复
		if node.hasAttribute(attr.Name) {
			return p.newAttributeDuplicateErr(attr)
		}
		node.addAttribute(attr.Name, attr.Value)
	}
	return
}

func (p *NodeParser) parseConfigurationNode(node *Node, xmlNode *XMLNode, elem *dtd.Element) (err error) {
	err = p.parseXMLAttribute(node, xmlNode, elem)
	if err != nil {
		return
	}
	for _, childXmlNode := range xmlNode.Body {
		// 子节点不支持
		if !elem.HasNode(childXmlNode.Name) {
			return p.newNodeNotSupportErr(childXmlNode)
		}
		// 子节点重复错误
		if elem.GetNodeCount(childXmlNode.Name) == dtd.AT_MOST_ONCE &&
			node.count(childXmlNode.Name) > 0 {
			return p.newNodeDuplicateErr(childXmlNode)
		}
		// 判断是否解析 SQL
		childNode := NewNode(childXmlNode.Name)
		node.addNode(childNode)

	}

	// 判断是否包换必填属性
	// 判断是否包含必填节点

	return nil
}

func (p *NodeParser) newAttributeDuplicateErr(attr *XMLAttribute) error {
	return fmt.Errorf(
		"duplicate attribute: %s at line:%d column: %d",
		attr.Name,
		attr.Start.Line,
		attr.Start.Column,
	)
}

func (p *NodeParser) newAttributeNotSupportErr(attr *XMLAttribute) error {
	return fmt.Errorf(
		" attribute: %s not support at line:%d column: %d",
		attr.Name,
		attr.Start.Line,
		attr.Start.Column,
	)
}

func (p *NodeParser) newNodeNotSupportErr(xmlNode *XMLNode) error {
	return fmt.Errorf(
		"node: %s not suport at line:%d column: %d",
		xmlNode.Name,
		xmlNode.Start.Line,
		xmlNode.Start.Column,
	)
}

func (p *NodeParser) newNodeDuplicateErr(xmlNode *XMLNode) error {
	return fmt.Errorf(
		"duplicate node: %s not suport at line:%d column: %d",
		xmlNode.Name,
		xmlNode.Start.Line,
		xmlNode.Start.Column,
	)
}
