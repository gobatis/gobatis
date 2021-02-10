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
	Tokens     []*Token
	attributes map[string]string
	nodes      []*Node
	count      map[string]int
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
	if p.count == nil {
		p.count = map[string]int{}
	}
	p.count[node.Name]++
	p.nodes = append(p.nodes, node)
}

func (p *Node) countNode(name string) int {
	if p.count == nil {
		return 0
	}
	return p.count[name]
}

func NewNodeParser(file string) *NodeParser {
	return &NodeParser{file: file}
}

type NodeParser struct {
	file string
}

func (p *NodeParser) ParseConfiguration(content []byte) (configuration *Node, err error) {

	nodes, err := NewXMLParser().Parse(content)
	if err != nil {
		return
	}
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

func (p *NodeParser) ParseMapper(content []byte) (err error) {
	//nodes := p.parseXmlNodes(content)
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

	// 判断是否包换必填属性
	if elem.Attributes != nil {
		for k, v := range elem.Attributes {
			if v == dtd.REQUIRED && node.hasAttribute(k) {
				return p.newNodeMissRequiredAttributeErr(xmlNode, k)
			}
		}
	}

	// 判断是否解析 SQL Tokens
	if elem.HasNode(dtd.PCDATA) {
		node.Tokens = xmlNode.Tokens
	}

	for _, childXmlNode := range xmlNode.Body {
		// 子节点不支持
		if !elem.HasNode(childXmlNode.Name) {
			return p.newNodeNotSupportErr(childXmlNode)
		}
		// 子节点重复错误
		if elem.GetNodeCount(childXmlNode.Name) == dtd.AT_MOST_ONCE &&
			node.countNode(childXmlNode.Name) > 0 {
			return p.newNodeDuplicateErr(childXmlNode)
		}
		childNode := NewNode(childXmlNode.Name)
		node.addNode(childNode)
	}

	// 判断是否包含必填节点
	if elem.Nodes != nil {
		for k, v := range elem.Nodes {
			if v == dtd.AT_LEAST_ONCE && node.countNode(k) == 0 {
				return p.newNodeMissRequiredChildNodeErr(xmlNode, k)
			}
		}
	}

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

func (p *NodeParser) newNodeMissRequiredChildNodeErr(xmlNode *XMLNode, attrName string) error {
	return fmt.Errorf(
		"node: %s miss required child node %s :%d column: %d",
		xmlNode.Name,
		attrName,
		xmlNode.Start.Line,
		xmlNode.Start.Column,
	)
}

func (p *NodeParser) newNodeMissRequiredAttributeErr(xmlNode *XMLNode, attrName string) error {
	return fmt.Errorf(
		"node: %s miss required attribute %s :%d column: %d",
		xmlNode.Name,
		attrName,
		xmlNode.Start.Line,
		xmlNode.Start.Column,
	)
}
