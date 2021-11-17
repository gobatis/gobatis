package generator

import (
	"fmt"
	"strings"
)

type SName struct {
	Action string
	Name   string
	Type   string
}

func (p SName) arrayType(a bool) string {
	if a {
		return "Array"
	}
	return ""
}

func (p SName) ParameterOriginal(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%s", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterOriginalPointer(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sOriginalPointer", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterPointerOriginal(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sPointerOriginal", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterPointerPointer(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sPointerPointer", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) EntityOriginal(array bool) string {
	return fmt.Sprintf("%s%sEntity%s%s", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) EntityOriginalPointer(array bool) string {
	return fmt.Sprintf("%s%sEntity%s%sOriginalPointer", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) EntityPointerOriginal(array bool) string {
	return fmt.Sprintf("%s%sEntity%s%sPointerOriginal", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) EntityPointerPointer(array bool) string {
	return fmt.Sprintf("%s%sEntity%s%sPointerPointer", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterTx(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sTx", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterRows(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sRows", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterMust(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sMust", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterEmbed(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sEmbed", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterContext(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sContext", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterStmt(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sStmt", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterID(array bool) string {
	return fmt.Sprintf("%s%sParameter%s%sID", p.Action, p.arrayType(array), p.Name, p.Type)
}

func (p SName) ParameterOtherType(array bool, t string) string {
	return fmt.Sprintf("%s%sParameter%s%sOther", p.Action, p.arrayType(array), p.Name, strings.Title(t))
}
