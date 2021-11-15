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

func (p SName) ParameterOriginal() string {
	return fmt.Sprintf("%sParameter%s%s", p.Action, p.Name, p.Type)
}

func (p SName) ParameterOriginalPointer() string {
	return fmt.Sprintf("%sParameter%s%sOriginalPointer", p.Action, p.Name, p.Type)
}

func (p SName) ParameterPointerOriginal() string {
	return fmt.Sprintf("%sParameter%s%sPointerOriginal", p.Action, p.Name, p.Type)
}

func (p SName) ParameterPointerPointer() string {
	return fmt.Sprintf("%sParameter%s%sPointerPointer", p.Action, p.Name, p.Type)
}

func (p SName) EntityOriginal() string {
	return fmt.Sprintf("%sEntity%s%s", p.Action, p.Name, p.Type)
}

func (p SName) EntityOriginalPointer() string {
	return fmt.Sprintf("%sEntity%s%sOriginalPointer", p.Action, p.Name, p.Type)
}

func (p SName) EntityPointerOriginal() string {
	return fmt.Sprintf("%sEntity%s%sPointerOriginal", p.Action, p.Name, p.Type)
}

func (p SName) EntityPointerPointer() string {
	return fmt.Sprintf("%sEntity%s%sPointerPointer", p.Action, p.Name, p.Type)
}

func (p SName) ParameterTx() string {
	return fmt.Sprintf("%sParameter%s%sTx", p.Action, p.Name, p.Type)
}

func (p SName) ParameterRows() string {
	return fmt.Sprintf("%sParameter%s%sRows", p.Action, p.Name, p.Type)
}

func (p SName) ParameterMust() string {
	return fmt.Sprintf("%sParameter%s%sMust", p.Action, p.Name, p.Type)
}

func (p SName) ParameterEmbed() string {
	return fmt.Sprintf("%sParameter%s%sEmbed", p.Action, p.Name, p.Type)
}

func (p SName) ParameterContext() string {
	return fmt.Sprintf("%sParameter%s%sContext", p.Action, p.Name, p.Type)
}

func (p SName) ParameterStmt() string {
	return fmt.Sprintf("%sParameter%s%sStmt", p.Action, p.Name, p.Type)
}

func (p SName) ParameterOtherType(t string) string {
	return fmt.Sprintf("%sParameter%s%sOther", p.Action, p.Name, strings.Title(t))
}
