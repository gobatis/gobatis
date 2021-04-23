package gobatis

import (
	"container/list"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/koyeo/gobatis/dtd"
	"github.com/koyeo/gobatis/parser/expr"
	"github.com/koyeo/gobatis/parser/xml"
	"github.com/spf13/cast"
	"reflect"
	"strings"
	"sync"
)

func parseConfig(engine *Engine, file, content string) (err error) {
	listener := &xmlParser{
		file:          file,
		stack:         newXMLStack(),
		rootElement:   dtd.Configuration,
		elementGetter: dtd.ConfigElement,
	}
	err = parseNode(listener, content)
	if err != nil {
		return
	}
	
	//d, _ := json.MarshalIndent(listener.rootNode, "", "\t")
	//fmt.Println(string(d))
	return
}

func parseMapper(engine *Engine, file, content string) (err error) {
	
	listener := &xmlParser{
		file:          file,
		stack:         newXMLStack(),
		rootElement:   dtd.Mapper,
		elementGetter: dtd.MapperElement,
	}
	err = parseNode(listener, content)
	if err != nil {
		return
	}
	
	if listener.rootNode == nil {
		engine.logger.Warnf("empty mapperCache file: %s", file)
		return
	}
	
	for _, v := range listener.rootNode.Nodes {
		if v.Name == dtd.SELECT ||
			v.Name == dtd.INSERT ||
			v.Name == dtd.DELETE ||
			v.Name == dtd.UPDATE ||
			v.Name == dtd.SQL {
			id := v.GetAttribute(dtd.ID)
			if id == "" {
				err = parseError(file, v.start, fmt.Sprintf("element: %s miss id", v.Name))
				return
			}
			err = engine.addStatement(file, v.start, id, v)
			if err != nil {
				return
			}
		}
	}
	
	return
}

func parseNode(listener *xmlParser, content string) (err error) {
	lexer := xml.NewXMLLexer(antlr.NewInputStream(strings.TrimSpace(content)))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := xml.NewXMLParser(stream)
	parser.BuildParseTrees = true
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	//parser.SetErrorHandler()
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Document())
	if listener.error != nil {
		err = listener.error
		return
	}
	return
}

func parseError(file string, token antlr.Token, msg string) error {
	return fmt.Errorf("%s line %d:%d, parse error: %s", file, token.GetLine(), token.GetColumn(), msg)
}

func newXMLNode(file, name string, token antlr.Token) *xmlNode {
	return &xmlNode{
		File:  file,
		Name:  name,
		start: token,
	}
}

type xmlNode struct {
	File       string `json:"-"`
	Name       string
	Text       string
	Attributes map[string]*xmlNodeAttribute
	Nodes      []*xmlNode
	nodesCount map[string]int
	start      antlr.Token
	textOnly   bool
}

type xmlNodeAttribute struct {
	File  string      `json:"-"`
	Start antlr.Token `json:"-"`
	Value string
}

func (p *xmlNode) GetAttribute(name string) string {
	if p.Attributes == nil {
		return ""
	}
	v, ok := p.Attributes[name]
	if ok {
		return v.Value
	}
	return ""
}

func (p *xmlNode) AddAttribute(name string, value *xmlNodeAttribute) {
	if p.Attributes == nil {
		p.Attributes = map[string]*xmlNodeAttribute{}
	}
	p.Attributes[name] = value
}

func (p *xmlNode) HasAttribute(name string) bool {
	if p.Attributes == nil {
		return false
	}
	_, ok := p.Attributes[name]
	return ok
}

func (p *xmlNode) AddNode(node *xmlNode) {
	if p.nodesCount == nil {
		p.nodesCount = map[string]int{}
	}
	p.nodesCount[node.Name]++
	p.Nodes = append(p.Nodes, node)
}

func (p *xmlNode) countNode(name string) int {
	if p.nodesCount == nil {
		return 0
	}
	return p.nodesCount[name]
}

type xmlNodeStack struct {
	list *list.List
	lock *sync.RWMutex
}

func newXMLStack() *xmlNodeStack {
	l := list.New()
	lock := &sync.RWMutex{}
	return &xmlNodeStack{l, lock}
}

func (p *xmlNodeStack) Push(value *xmlNode) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.list.PushBack(value)
}

func (p *xmlNodeStack) Pop() *xmlNode {
	p.lock.Lock()
	defer p.lock.Unlock()
	e := p.list.Back()
	if e != nil {
		p.list.Remove(e)
		return e.Value.(*xmlNode)
	}
	return nil
}

func (p *xmlNodeStack) Peak() *xmlNode {
	p.lock.RLock()
	defer p.lock.Unlock()
	e := p.list.Back()
	if e != nil {
		return e.Value.(*xmlNode)
	}
	return nil
}

func (p *xmlNodeStack) Len() int {
	p.lock.RLock()
	defer p.lock.Unlock()
	return p.list.Len()
}

func (p *xmlNodeStack) Empty() bool {
	p.lock.RLock()
	defer p.lock.Unlock()
	return p.list.Len() == 0
}

type xmlParser struct {
	*antlr.BaseParseTreeListener
	file          string
	content       []byte
	depth         int
	error         error
	stack         *xmlNodeStack
	rootNode      *xmlNode
	rootElement   *dtd.Element
	elementGetter func(name string) (elem *dtd.Element, err error)
}

func (p *xmlParser) trimAttributeValueQuote(val string) string {
	if strings.HasPrefix(val, "'") || strings.HasSuffix(val, "'") {
		val = strings.TrimPrefix(val, "'")
		return strings.TrimSuffix(val, "'")
	}
	if strings.HasPrefix(val, "\"") || strings.HasSuffix(val, "\"") {
		val = strings.TrimPrefix(val, "\"")
		return strings.TrimSuffix(val, "\"")
	}
	return val
}

func (p *xmlParser) validateNode(node *xmlNode, elem *dtd.Element) {
	
	// check required attributes
	if elem.Attributes != nil {
		for k, v := range elem.Attributes {
			if v == dtd.REQUIRED && !node.HasAttribute(k) {
				p.setError(fmt.Sprintf("element: %s miss required attribute: %s", node.Name, k), node.start)
				return
			}
		}
	}
	
	for _, childNode := range node.Nodes {
		
		if childNode.textOnly {
			continue
		}
		
		// check not supported node
		if !elem.HasNode(childNode.Name) {
			p.setError(
				fmt.Sprintf("element: %s not support child element: %s", node.Name, childNode.Name),
				childNode.start,
			)
			return
		}
		
		// check at most once node
		if elem.GetNodeCount(childNode.Name) == dtd.AT_MOST_ONCE && node.countNode(childNode.Name) > 1 {
			p.setError(
				fmt.Sprintf("element: %s not support duplicate element: %s", node.Name, childNode.Name),
				childNode.start,
			)
			return
		}
		
		childElem, err := p.elementGetter(childNode.Name)
		if err != nil {
			p.setError(err.Error(), childNode.start)
			return
		}
		
		p.validateNode(childNode, childElem)
		if p.error != nil {
			return
		}
	}
	
	// check at least once node
	if elem.Nodes != nil {
		for k, v := range elem.Nodes {
			if v == dtd.AT_LEAST_ONCE && node.countNode(k) == 0 {
				p.setError(fmt.Sprintf("element %s miss required element %s", node.Name, k), node.start)
				return
			}
		}
	}
}

func (p *xmlParser) setError(msg string, token antlr.Token) {
	p.error = parseError(p.file, token, msg)
}

func (p *xmlParser) enterElement(c *xml.ElementContext) {
	if p.error != nil {
		return
	}
	name := c.Name(0)
	if p.depth == 0 {
		if name.GetText() != p.rootElement.Name {
			p.setError(fmt.Sprintf("first level tag %s unsupported", name.GetText()), name.GetSymbol())
			return
		}
	}
	p.stack.Push(newXMLNode(p.file, name.GetText(), name.GetSymbol()))
	p.depth++
}

func (p *xmlParser) exitElement(_ *xml.ElementContext) {
	if p.error != nil || p.stack.Peak() == nil {
		return
	}
	p.depth--
	child := p.stack.Pop()
	
	if p.stack.Len() > 0 {
		p.stack.Peak().AddNode(child)
	} else {
		p.rootNode = child
		p.validateNode(child, p.rootElement)
	}
}

func (p *xmlParser) enterAttribute(c *xml.AttributeContext) {
	// <?xml version="1.0" encoding="UTF-8" ?>
	if p.error != nil || p.stack.Peak() == nil {
		return
	}
	name := strings.TrimSpace(c.Name().GetText())
	value := strings.TrimSpace(p.trimAttributeValueQuote(c.STRING().GetText()))
	p.stack.Peak().AddAttribute(name, &xmlNodeAttribute{
		File:  p.file,
		Start: c.STRING().GetSymbol(),
		Value: value,
	})
}

func (p *xmlParser) enterChardata(c *xml.ChardataContext) {
	if p.error != nil || p.stack.Peak() == nil {
		return
	}
	text := strings.TrimSpace(c.GetText())
	if text != "" {
		p.stack.Peak().AddNode(&xmlNode{File: p.file, Text: text, start: c.GetStart(), textOnly: true})
	}
}

func (p *xmlParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.GetRuleIndex() {
	case xml.XMLParserRULE_element:
		p.enterElement(ctx.(*xml.ElementContext))
	case xml.XMLParserRULE_attribute:
		p.enterAttribute(ctx.(*xml.AttributeContext))
	case xml.XMLParserRULE_chardata:
		p.enterChardata(ctx.(*xml.ChardataContext))
	}
	
}

func (p *xmlParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.GetRuleIndex() {
	case xml.XMLParserRULE_element:
		p.exitElement(nil)
	}
}

type exprValue struct {
	value      interface{}
	paramValue reflect.Value
	aliasKind  reflect.Kind
}

func (p *exprValue) isExpected() bool {
	return p.aliasKind == reflect.Interface || p.paramValue.Kind() == p.aliasKind
}

func (p *exprValue) convertible() bool {
	return false
}

type exprStack struct {
	list *list.List
	lock *sync.RWMutex
}

func newExprStack() *exprStack {
	l := list.New()
	lock := &sync.RWMutex{}
	return &exprStack{l, lock}
}

func (p *exprStack) Push(value *exprValue) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.list.PushBack(value)
}

func (p *exprStack) Pop() (val *exprValue, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	if p.list.Len() < 1 {
		err = fmt.Errorf("stack is empty unable to pop")
		return
	}
	e := p.list.Back()
	if e != nil {
		p.list.Remove(e)
		val = e.Value.(*exprValue)
		return
	}
	return
}

func (p *exprStack) Peak() *exprValue {
	p.lock.RLock()
	defer p.lock.Unlock()
	e := p.list.Back()
	if e != nil {
		return e.Value.(*exprValue)
	}
	return nil
}

func (p *exprStack) Len() int {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.list.Len()
}

func (p *exprStack) Empty() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.list.Len() == 0
}

func newExprParams(params ...interface{}) *exprParams {
	r := &exprParams{}
	for _, v := range params {
		rv := reflect.ValueOf(v)
		r.values = append(r.values, exprValue{
			value:      v,
			paramValue: rv,
			aliasKind:  rv.Kind(),
		})
	}
	return r
}

type exprParams struct {
	values  []exprValue
	aliases map[string]int
}

// return exprValue not *exprValue to protect params
func (p *exprParams) get(name string) (val exprValue, ok bool) {
	if p.aliases == nil {
		return
	}
	index, ok := p.aliases[name]
	if !ok {
		return
	}
	if index > len(p.values)-1 {
		return
	}
	ok = true
	val = p.values[index]
	return
}

func (p *exprParams) alias(name, expected string, index int) error {
	if p.aliases == nil {
		p.aliases = map[string]int{}
	} else {
		_, ok := p.aliases[name]
		if ok {
			return fmt.Errorf("duplicated alias: %s", name)
		}
	}
	vl := len(p.values) - 1
	if index < 0 || index > vl {
		return fmt.Errorf("alias: %s index: %d out of params length: %d", name, index, vl)
	}
	if expected != "" {
		var err error
		ev := p.values[index]
		ev.aliasKind, err = p.toReflectKind(expected)
		if err != nil {
			return fmt.Errorf("convert alias: %s type: %s to %s error: %s", name, ev.aliasKind, expected, err)
		}
		if !ev.isExpected() && !ev.convertible() {
			return fmt.Errorf("param type: %s is not alias expteced type: %s",
				ev.paramValue.Kind(), ev.aliasKind)
		}
	}
	p.aliases[name] = index
	return nil
}

func (p *exprParams) toReflectKind(t string) (kind reflect.Kind, err error) {
	
	switch t {
	case "bool":
		kind = reflect.Bool
	case "int":
		kind = reflect.Int
	case "int8":
		kind = reflect.Int8
	case "int16":
		kind = reflect.Int16
	case "int32":
		kind = reflect.Int32
	case "int64":
		kind = reflect.Int64
	case "uint":
		kind = reflect.Uint
	case "uint8":
		kind = reflect.Uint8
	case "uint16":
		kind = reflect.Uint16
	case "uint32":
		kind = reflect.Uint32
	case "uint64":
		kind = reflect.Uint64
	case "float32":
		kind = reflect.Float32
	case "float64":
		kind = reflect.Float64
	case "complex64":
		kind = reflect.Complex64
	case "complex128":
		kind = reflect.Complex128
	case "array":
		kind = reflect.Array
	case "auto":
		kind = reflect.Interface
	case "map":
		kind = reflect.Map
	case "string":
		kind = reflect.String
	case "struct":
		kind = reflect.Struct
	default:
		err = fmt.Errorf("unsupported type:%s", t)
	}
	return
}

func (p *exprParams) index(index int) (val exprValue, ok bool) {
	if p.aliases == nil {
		return
	}
	if index > len(p.values)-1 {
		return
	}
	ok = true
	val = p.values[index]
	return
}

type exprParser struct {
	*expr.BaseExprParserListener
	*exprStack
	params     *exprParams
	error      error
	file       string
	aliasIndex int
}

func newExprParser(params ...interface{}) *exprParser {
	r := new(exprParser)
	r.exprStack = newExprStack()
	r.params = newExprParams(params...)
	return r
}

func (p *exprParser) ExitExpression(ctx *expr.ExpressionContext) {
	if p.error != nil {
		return
	}
	
	if ctx.GetUnary_op() != nil {
		left, err := p.Pop()
		if err != nil {
			p.error = err
			return
		}
		err = p.unaryCalc(left, ctx.GetUnary_op())
		if err != nil {
			p.error = err
			return
		}
	} else if ctx.GetMul_op() != nil ||
		ctx.GetAdd_op() != nil ||
		ctx.GetRel_op() != nil ||
		ctx.LOGICAL_AND() != nil ||
		ctx.LOGICAL_OR() != nil {
		left, right, err := p.popBinaryOperands()
		if err != nil {
			p.error = err
			return
		}
		if ctx.GetAdd_op() != nil {
			err = p.numericStringCalc(left, right, ctx.GetAdd_op())
			if err != nil {
				p.error = err
				return
			}
		} else if ctx.GetMul_op() != nil {
			err = p.numericStringCalc(left, right, ctx.GetMul_op())
			if err != nil {
				p.error = err
				return
			}
		} else if ctx.GetRel_op() != nil {
			err = p.relationCalc(left, right, ctx.GetRel_op())
			if err != nil {
				p.error = err
				return
			}
		} else if ctx.LOGICAL_AND() != nil {
			err = p.logicCalc(left, right, ctx.LOGICAL_AND().GetSymbol())
			if err != nil {
				p.error = err
				return
			}
		} else if ctx.LOGICAL_OR() != nil {
			err = p.logicCalc(left, right, ctx.LOGICAL_OR().GetSymbol())
			if err != nil {
				p.error = err
				return
			}
		}
	}
}

func (p *exprParser) EnterParamDecl(ctx *expr.ParamDeclContext) {
	if p.error != nil {
		return
	}
	name := ctx.IDENTIFIER().GetText()
	var expected string
	if ctx.ParamType() != nil {
		expected = ctx.ParamType().GetText()
	}
	err := p.params.alias(name, expected, p.aliasIndex)
	if err != nil {
		p.error = parseError(p.file, ctx.GetStart(), err.Error())
		return
	}
	p.aliasIndex++
}

func (p *exprParser) ExitOperandName(ctx *expr.OperandNameContext) {
	if p.error != nil {
		return
	}
	alias := ctx.IDENTIFIER(0).GetText()
	val, ok := p.params.get(alias)
	if !ok {
		p.error = parseError(p.file, ctx.GetStart(), fmt.Sprintf("can't fetch alias: %s", alias))
		return
	}
	//ctx.AllIDENTIFIER()
	p.Push(&val)
}

func (p *exprParser) ExitInteger(ctx *expr.IntegerContext) {
	p.Push(&exprValue{
		value:     ctx.GetText(),
		aliasKind: reflect.Int,
	})
}

func (p *exprParser) ExitString_(ctx *expr.String_Context) {
	p.Push(&exprValue{
		value:     ctx.GetText(),
		aliasKind: reflect.String,
	})
}

func (p *exprParser) ExitFloat_(ctx *expr.Float_Context) {
	p.Push(&exprValue{
		value:     ctx.GetText(),
		aliasKind: reflect.Float64,
	})
}

func (p *exprParser) parseExpression(params, expresion string) (result interface{}, err error) {
	err = p.parseParam(params)
	if err != nil {
		return
	}
	parser, err := p.parser(expresion)
	if err != nil {
		return
	}
	
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Expressions())
	err = p.error
	if err != nil {
		return
	}
	if p.Len() != 1 {
		err = fmt.Errorf("unexpected stack length: %d", p.Len())
		return
	}
	
	v, _ := p.Pop()
	result = v.value
	
	return
}

func (p *exprParser) popBinaryOperands() (left, right *exprValue, err error) {
	right, err = p.Pop()
	if err != nil {
		return
	}
	left, err = p.Pop()
	if err != nil {
		return
	}
	return
}

func (p *exprParser) getReferKind(left, right *exprValue) (kind reflect.Kind) {
	if left.value == nil && right.value != nil {
		kind = right.aliasKind
	}
	kind = left.aliasKind
	return p.getMaxKind(kind)
}

func (p *exprParser) getMaxKind(kind reflect.Kind) reflect.Kind {
	if kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 {
		kind = reflect.Int64
	} else if kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 {
		kind = reflect.Uint64
	} else if kind == reflect.Float32 ||
		kind == reflect.Float64 {
		kind = reflect.Float64
	}
	return kind
}

func (p *exprParser) numericStringCalc(left, right *exprValue, op antlr.Token) error {
	kind := p.getReferKind(left, right)
	if !p.isNumeric(kind) && (op.GetTokenType() == expr.ExprParserPLUS && kind != reflect.String) {
		return parseError(p.file, op, fmt.Sprintf(
			"invalid operation: %s %s %s,unsupported type: %s", left.value, op.GetText(), right.value, kind,
		))
	}
	
	result := &exprValue{aliasKind: kind}
	switch kind {
	case reflect.Int64:
		a, err := cast.ToInt64E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		b, err := cast.ToInt64E(right.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserPLUS:
			result.value = a + b
		case expr.ExprParserMINUS:
			result.value = a - b
		case expr.ExprParserSTAR:
			result.value = a * b
		case expr.ExprParserDIV:
			if b == 0 {
				return p.divisionByZero(op)
			}
			result.value = a / b
		case expr.ExprParserCARET:
			result.value = a ^ b
		case expr.ExprParserOR:
			result.value = a | b
		case expr.ExprParserAMPERSAND:
			result.value = a & b
		case expr.ExprParserMOD:
			result.value = a % b
		case expr.ExprParserLSHIFT:
			result.value = a << b
		case expr.ExprParserRSHIFT:
			result.value = a >> b
		case expr.ExprParserBIT_CLEAR:
			result.value = a &^ b
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.Uint64:
		a, err := cast.ToInt8E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		b, err := cast.ToInt8E(right.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserPLUS:
			result.value = a + b
		case expr.ExprParserMINUS:
			result.value = a - b
		case expr.ExprParserSTAR:
			result.value = a * b
		case expr.ExprParserDIV:
			if b == 0 {
				return p.divisionByZero(op)
			}
			result.value = a / b
		case expr.ExprParserCARET:
			result.value = a ^ b
		case expr.ExprParserOR:
			result.value = a | b
		case expr.ExprParserAMPERSAND:
			result.value = a & b
		case expr.ExprParserMOD:
			result.value = a % b
		case expr.ExprParserLSHIFT:
			result.value = a << b
		case expr.ExprParserRSHIFT:
			result.value = a >> b
		case expr.ExprParserBIT_CLEAR:
			result.value = a &^ b
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.Float64:
		a, err := cast.ToIntE(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		b, err := cast.ToIntE(right.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserPLUS:
			result.value = a + b
		case expr.ExprParserMINUS:
			result.value = a - b
		case expr.ExprParserSTAR:
			result.value = a * b
		case expr.ExprParserDIV:
			if b == 0 {
				return p.divisionByZero(op)
			}
			result.value = a / b
		case expr.ExprParserCARET:
			result.value = a ^ b
		case expr.ExprParserOR:
			result.value = a | b
		case expr.ExprParserAMPERSAND:
			result.value = a & b
		case expr.ExprParserMOD:
			result.value = a % b
		case expr.ExprParserLSHIFT:
			result.value = a << b
		case expr.ExprParserRSHIFT:
			result.value = a >> b
		case expr.ExprParserBIT_CLEAR:
			result.value = a &^ b
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.String:
		a, err := cast.ToStringE(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		b, err := cast.ToStringE(right.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserPLUS:
			result.value = a + b
		default:
			return p.unsupportedOpError(op)
		}
	default:
		return parseError(p.file, op, fmt.Sprintf("unsupported numeric op type: %s", kind))
	}
	p.Push(result)
	
	return nil
}

func (p *exprParser) relationCalc(left, right *exprValue, op antlr.Token) error {
	kind := p.getReferKind(left, right)
	if !p.isNumeric(kind) && kind != reflect.String {
		return parseError(p.file, op, fmt.Sprintf(
			"invalid operation: %s %s %s, unsupported type: %s", left.value, op.GetText(), right.value, kind,
		))
	}
	result := &exprValue{aliasKind: reflect.Bool}
	switch kind {
	case reflect.Int64:
		a, err := cast.ToInt64E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		b, err := cast.ToInt64E(right.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserEQUALS:
			result.value = a == b
		case expr.ExprParserNOT_EQUALS:
			result.value = a != b
		case expr.ExprParserLESS:
			result.value = a < b
		case expr.ExprParserLESS_OR_EQUALS:
			result.value = a <= b
		case expr.ExprParserGREATER:
			result.value = a > b
		case expr.ExprParserGREATER_OR_EQUALS:
			result.value = a >= b
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.Uint64:
		a, err := cast.ToUint64E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		b, err := cast.ToUint64E(right.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserEQUALS:
			result.value = a == b
		case expr.ExprParserNOT_EQUALS:
			result.value = a != b
		case expr.ExprParserLESS:
			result.value = a < b
		case expr.ExprParserLESS_OR_EQUALS:
			result.value = a <= b
		case expr.ExprParserGREATER:
			result.value = a > b
		case expr.ExprParserGREATER_OR_EQUALS:
			result.value = a >= b
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.Float64:
		a, err := cast.ToFloat64E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		b, err := cast.ToFloat64E(right.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserEQUALS:
			result.value = a == b
		case expr.ExprParserNOT_EQUALS:
			result.value = a != b
		case expr.ExprParserLESS:
			result.value = a < b
		case expr.ExprParserLESS_OR_EQUALS:
			result.value = a <= b
		case expr.ExprParserGREATER:
			result.value = a > b
		case expr.ExprParserGREATER_OR_EQUALS:
			result.value = a >= b
		default:
			return p.unsupportedOpError(op)
		}
	default:
		return parseError(p.file, op, fmt.Sprintf("unsupported relation op type: %s", kind))
	}
	p.Push(result)
	
	return nil
}

func (p *exprParser) unaryCalc(left *exprValue, op antlr.Token) error {
	
	if !p.isNumeric(left.aliasKind) && left.aliasKind != reflect.Bool {
		return parseError(p.file, op, fmt.Sprintf(
			"invalid operation: %s%s , unsupported type: %s", op.GetText(), left.value, left.aliasKind,
		))
	}
	
	kind := p.getMaxKind(left.aliasKind)
	result := &exprValue{aliasKind: kind}
	switch kind {
	case reflect.Int64:
		a, err := cast.ToInt64E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserPLUS:
			result.value = +a
		case expr.ExprParserMINUS:
			result.value = -a
		case expr.ExprParserCARET:
			result.value = ^a
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.Uint64:
		a, err := cast.ToUint64E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserPLUS:
			result.value = +a
		case expr.ExprParserMINUS:
			result.value = -a
		case expr.ExprParserCARET:
			result.value = ^a
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.Float64:
		a, err := cast.ToFloat64E(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserPLUS:
			result.value = +a
		case expr.ExprParserMINUS:
			result.value = -a
		default:
			return p.unsupportedOpError(op)
		}
	case reflect.Bool:
		a, err := cast.ToBoolE(left.value)
		if err != nil {
			return p.castError(op, err)
		}
		switch op.GetTokenType() {
		case expr.ExprParserEXCLAMATION:
			result.value = !a
		default:
			return p.unsupportedOpError(op)
		}
	default:
		return parseError(p.file, op, fmt.Sprintf("unsupported unary op type: %s", left.aliasKind))
	}
	p.Push(result)
	
	return nil
}

func (p *exprParser) logicCalc(left, right *exprValue, op antlr.Token) error {
	a, err := cast.ToBoolE(left.value)
	if err != nil {
		return p.castError(op, err)
	}
	b, err := cast.ToBoolE(right.value)
	if err != nil {
		return p.castError(op, err)
	}
	result := &exprValue{aliasKind: reflect.Bool}
	switch op.GetTokenType() {
	case expr.ExprParserLOGICAL_AND:
		result.value = a && b
	case expr.ExprParserLOGICAL_OR:
		result.value = a || b
	}
	p.Push(result)
	return nil
}

func (p *exprParser) castError(op antlr.Token, err error) error {
	return parseError(p.file, op, fmt.Sprintf("convert error: %s", err))
}

func (p *exprParser) divisionByZero(op antlr.Token) error {
	return parseError(p.file, op, fmt.Sprintf("division by zero"))
}

func (p *exprParser) unsupportedOpError(op antlr.Token) error {
	return parseError(p.file, op, fmt.Sprintf("unsupported op: %s", op.GetText()))
}

func (p *exprParser) isNumeric(kind reflect.Kind) bool {
	if kind != reflect.Int &&
		kind != reflect.Int8 &&
		kind != reflect.Int16 &&
		kind != reflect.Int32 &&
		kind != reflect.Int64 &&
		kind != reflect.Uint &&
		kind != reflect.Uint8 &&
		kind != reflect.Uint16 &&
		kind != reflect.Uint32 &&
		kind != reflect.Uint64 &&
		kind != reflect.Float32 &&
		kind != reflect.Float64 {
		return true
	}
	return true
}

func (p *exprParser) parseParam(params string) (err error) {
	parser, err := p.parser(params)
	if err != nil {
		return
	}
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Parameters())
	return
}

func (p *exprParser) parser(data string) (parser *expr.ExprParser, err error) {
	lexer := expr.NewExprLexer(antlr.NewInputStream(data))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser = expr.NewExprParser(stream)
	parser.BuildParseTrees = true
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	return
}
