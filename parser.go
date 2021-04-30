package gobatis

import (
	"container/list"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gobatis/gobatis/cast"
	"github.com/gobatis/gobatis/dtd"
	"github.com/gobatis/gobatis/parser/expr"
	"github.com/gobatis/gobatis/parser/xml"
	"reflect"
	"strings"
	"sync"
)

func parseConfig(engine *Engine, file, content string) (err error) {
	l := &xmlParser{
		file:          file,
		stack:         newXMLStack(),
		coverage:      newCoverage(),
		rootElement:   dtd.Configuration,
		elementGetter: dtd.ConfigElement,
	}
	err = parseNode(l, content)
	if err != nil {
		return
	}
	
	if !l.coverage.covered() {
		err = fmt.Errorf("parse config token not coverd: %d/%d", l.coverage.len(), l.coverage.total)
		return
	}
	
	return
}

func parseMapper(engine *Engine, file, content string) (err error) {
	
	l := &xmlParser{
		file:          file,
		stack:         newXMLStack(),
		coverage:      newCoverage(),
		rootElement:   dtd.Mapper,
		elementGetter: dtd.MapperElement,
	}
	err = parseNode(l, content)
	if err != nil {
		return
	}
	
	if !l.coverage.covered() {
		err = fmt.Errorf("parse mapper token not coverd: %d/%d", l.coverage.len(), l.coverage.total)
		return
	}
	
	if l.rootNode == nil {
		engine.logger.Warnf("empty mapperCache file: %s", file)
		return
	}
	
	for _, v := range l.rootNode.Nodes {
		if v.Name == dtd.SELECT ||
			v.Name == dtd.INSERT ||
			v.Name == dtd.DELETE ||
			v.Name == dtd.UPDATE ||
			v.Name == dtd.SQL {
			id := v.GetAttribute(dtd.ID)
			if id == "" {
				err = parseError(file, v.ctx, fmt.Sprintf("element: %s miss id", v.Name))
				return
			}
			err = engine.addStatement(file, v.ctx, v.start, id, v)
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

func parseError(file string, ctx antlr.ParserRuleContext, msg string) error {
	return fmt.Errorf("%s line %d:%d: %s\nparse error: %s", file, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), ctx.GetText(), msg)
}

func newCoverage() *coverage {
	return &coverage{
		scan: map[int]bool{},
	}
}

type coverage struct {
	total int
	scan  map[int]bool
}

func (p *coverage) setTotal(total int) {
	p.total = total
}

func (p *coverage) add(ctx antlr.ParserRuleContext) {
	for i := ctx.GetStart().GetTokenIndex(); i <= ctx.GetStop().GetTokenIndex(); i++ {
		p.scan[i] = true
	}
}

func (p *coverage) len() int {
	return len(p.scan)
}

func (p *coverage) covered() bool {
	return p.total == p.len()
}

func (p *coverage) notCovered() (indexes []int) {
	for i := 0; i <= p.total; i++ {
		if _, ok := p.scan[i]; !ok {
			indexes = append(indexes, i)
		}
	}
	return
}

func newXMLNode(file, name string, ctx antlr.ParserRuleContext, token antlr.Token) *xmlNode {
	return &xmlNode{
		File:  file,
		Name:  name,
		ctx:   ctx,
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
	ctx        antlr.ParserRuleContext
	textOnly   bool
}

type xmlNodeAttribute struct {
	File  string      `json:"-"`
	Start antlr.Token `json:"-"`
	Value string
	ctx   antlr.ParserRuleContext
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
	defer p.lock.RUnlock()
	e := p.list.Back()
	if e != nil {
		return e.Value.(*xmlNode)
	}
	return nil
}

func (p *xmlNodeStack) Len() int {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.list.Len()
}

func (p *xmlNodeStack) Empty() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.list.Len() == 0
}

type xmlParser struct {
	*antlr.BaseParseTreeListener
	coverage      *coverage
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
				p.setError(fmt.Sprintf("element: %s miss required attribute: %s", node.Name, k), node.ctx)
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
				childNode.ctx,
			)
			return
		}
		
		// check at most once node
		if elem.GetNodeCount(childNode.Name) == dtd.AT_MOST_ONCE && node.countNode(childNode.Name) > 1 {
			p.setError(
				fmt.Sprintf("element: %s not support duplicate element: %s", node.Name, childNode.Name),
				childNode.ctx,
			)
			return
		}
		
		childElem, err := p.elementGetter(childNode.Name)
		if err != nil {
			p.setError(err.Error(), childNode.ctx)
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
				p.setError(fmt.Sprintf("element %s miss required element %s", node.Name, k), node.ctx)
				return
			}
		}
	}
}

func (p *xmlParser) setError(msg string, ctx antlr.ParserRuleContext) {
	p.error = parseError(p.file, ctx, msg)
}

func (p *xmlParser) enterElement(c *xml.ElementContext) {
	if p.error != nil {
		return
	}
	name := c.Name(0)
	if p.depth == 0 {
		if name.GetText() != p.rootElement.Name {
			p.setError(fmt.Sprintf("first level tag %s unsupported", name.GetText()), c)
			return
		}
	}
	p.stack.Push(newXMLNode(p.file, name.GetText(), c, name.GetSymbol()))
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
		Value: value,
		Start: c.STRING().GetSymbol(),
		ctx:   c,
	})
}

func (p *xmlParser) enterChardata(c *xml.ChardataContext) {
	if p.error != nil || p.stack.Peak() == nil {
		return
	}
	text := strings.TrimSpace(c.GetText())
	if text != "" {
		p.stack.Peak().AddNode(&xmlNode{File: p.file, Text: text, ctx: c, start: c.GetStart(), textOnly: true})
	}
}

func (p *xmlParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.GetRuleIndex() {
	case xml.XMLParserRULE_document:
		p.coverage.setTotal(ctx.GetStop().GetTokenIndex() + 1)
	case xml.XMLParserRULE_element:
		p.enterElement(ctx.(*xml.ElementContext))
		p.coverage.add(ctx)
	case xml.XMLParserRULE_attribute:
		p.enterAttribute(ctx.(*xml.AttributeContext))
		p.coverage.add(ctx)
	case xml.XMLParserRULE_chardata:
		p.enterChardata(ctx.(*xml.ChardataContext))
		p.coverage.add(ctx)
	case xml.XMLParserRULE_prolog, xml.XMLParserRULE_misc:
		p.coverage.add(ctx)
	}
	
}

func (p *xmlParser) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.GetRuleIndex() {
	case xml.XMLParserRULE_element:
		p.exitElement(nil)
	}
}

type exprValue struct {
	value   interface{} // required
	source  string      // required at number type conversion and built-in function
	alias   string      // optional, display var path
	builtIn bool        // mark if built in function or package
}

func (p *exprValue) accept(received, expected reflect.Kind) bool {
	return expected == reflect.Interface || received == expected
}

func (p *exprValue) convertible(received, expected reflect.Kind) bool {
	if expected == reflect.Array && received == reflect.Slice {
		return true
	}
	return false
}

func (p *exprValue) int() (v int, err error) {
	return cast.ToIntE(p.value)
}

func (p *exprValue) reflectElem() (elem reflect.Value) {
	v := reflect.ValueOf(p.value)
	if v.Kind() == reflect.Ptr {
		elem = v.Elem()
	} else {
		elem = v
	}
	return elem
}

func (p *exprValue) visitMember(name string) (r *exprValue, err error) {
	
	elem := p.reflectElem()
	if elem.Kind() != reflect.Struct {
		err = fmt.Errorf("visit '%s.%s' is not struct", p.alias, name)
		return
	}
	
	mv := elem.FieldByName(name)
	if mv.Kind() == reflect.Invalid {
		mv = elem.MethodByName(name)
		if mv.Kind() == reflect.Invalid {
			err = fmt.Errorf("visit member '%s.%s' is invalid", p.alias, name)
			return
		}
	}
	
	r = &exprValue{
		value: mv.Interface(),
	}
	
	return
}

func (p *exprValue) visitArrayIndex(index int) (r *exprValue, err error) {
	
	elem := p.reflectElem()
	if elem.Kind() != reflect.Array && elem.Kind() != reflect.Slice {
		err = fmt.Errorf("visit var is not array or slice")
		return
	}
	
	mv := elem.Index(index)
	if mv.Kind() == reflect.Invalid {
		err = fmt.Errorf("visit array index '%d' is invalid", index)
		return
	}
	r = &exprValue{
		value: mv.Interface(),
	}
	return
}

func (p *exprValue) visitMapIndex(index reflect.Value) (r *exprValue, err error) {
	
	elem := p.reflectElem()
	if elem.Kind() != reflect.Map {
		err = fmt.Errorf("visit '%s' is not map", elem.Kind())
		return
	}
	
	mv := elem.MapIndex(index)
	if mv.Kind() == reflect.Invalid {
		err = fmt.Errorf("visit map index '%s' is invalid", index)
		return
	}
	r = &exprValue{
		value: mv.Interface(),
	}
	return
}

func (p *exprValue) call(ellipsis bool, params []reflect.Value) (r *exprValue, err error) {
	
	elem := p.reflectElem()
	if elem.Kind() != reflect.Func {
		err = fmt.Errorf("visit '%s' is not func", elem.Kind())
		return
	}
	var out []reflect.Value
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("function call error: %v", e)
			return
		}
	}()
	for i, v := range params {
		if i < elem.Type().NumIn() {
			va := elem.Type().In(i)
			if va.Kind() != reflect.Interface && va.String() != v.Type().String() {
				var vt interface{}
				vt, err = cast.ToReflectTypeE(va, v.Interface())
				if err == nil {
					params[i] = reflect.ValueOf(vt)
				}
			}
		}
	}
	if ellipsis {
		out = elem.CallSlice(params)
	} else {
		out = elem.Call(params)
	}
	
	if len(out) == 0 {
		err = fmt.Errorf("function must return one reuslt")
		return
	} else if len(out) > 1 {
		err = fmt.Errorf("function only support one result return")
		return
	}
	
	r = &exprValue{
		value: out[0].Interface(),
	}
	return
}

func (p *exprValue) visitSlice(format string, indexes ...int) (r *exprValue, err error) {
	var v reflect.Value
	switch len(indexes) {
	case 2:
		v = p.reflectElem().Slice(indexes[0], indexes[1])
	case 3:
		v = p.reflectElem().Slice3(indexes[0], indexes[1], indexes[2])
	default:
		err = fmt.Errorf("unsuppoted slice range index '%s'", format)
		return
	}
	r = &exprValue{value: v.Interface()}
	return
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
	r.set(params...)
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

func (p *exprParams) set(params ...interface{}) {
	for _, v := range params {
		p.values = append(p.values, exprValue{
			value: v,
		})
	}
}

func (p *exprParams) alias(name, _type string, index int) error {
	if p.aliases == nil {
		p.aliases = map[string]int{}
	} else {
		_, ok := p.aliases[name]
		if ok {
			return fmt.Errorf("duplicated alias '%s'", name)
		}
	}
	vl := len(p.values) - 1
	if index < 0 || index > vl {
		return fmt.Errorf("alias '%s' index '%d' out of params length '%d'", name, index, vl)
	}
	if _type != "" {
		ev := p.values[index]
		kind := ev.reflectElem().Kind()
		var err error
		var expected reflect.Kind
		expected, err = p.toReflectKind(_type)
		if err != nil {
			return fmt.Errorf("convert alias '%s' type error: %s", name, err)
		}
		if !ev.accept(kind, expected) && !ev.convertible(kind, expected) {
			return fmt.Errorf("param type '%s' is not alias '%s' expteced type '%s'", kind, name, expected)
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
		err = fmt.Errorf("unsupported alias type '%s'", t)
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

func newExprParser(params ...interface{}) *exprParser {
	r := new(exprParser)
	r.baseParams = newExprParams(params...)
	r.coverage = newCoverage()
	r.initBuiltIn()
	return r
}

type exprParser struct {
	*expr.BaseExprParserListener
	stack         *exprStack
	coverage      *coverage
	baseParams    *exprParams
	foreachParams *exprParams
	error         error
	file          string
	paramIndex    int
	builtIn       map[string]interface{}
}

func (p *exprParser) initBuiltIn() {
	p.builtIn = map[string]interface{}{}
	p.builtIn["len"] = _len
	p.builtIn["int"] = _int
	p.builtIn["int8"] = _int8
	p.builtIn["int16"] = _int16
	p.builtIn["int32"] = _int32
	p.builtIn["int64"] = _int64
	p.builtIn["uint"] = _uint
	p.builtIn["uint8"] = _uint8
	p.builtIn["uint16"] = _uint16
	p.builtIn["uint32"] = _uint32
	p.builtIn["uint64"] = _uint64
	p.builtIn["decimal"] = _decimal
	p.builtIn["bool"] = _bool
	p.builtIn["strings"] = _strings{}
}

func (p *exprParser) isBuiltIn(name string) bool {
	_, ok := p.builtIn[name]
	return ok
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
	
	if p.isBuiltIn(name) {
		p.error = parseError(p.file, ctx,
			fmt.Sprintf("alias '%s' conflict with built-in functions or objects", name))
		return
	}
	var err error
	if p.foreachParams != nil {
		err = p.foreachParams.alias(name, expected, p.paramIndex)
	} else {
		err = p.baseParams.alias(name, expected, p.paramIndex)
	}
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	p.paramIndex++
}

func (p *exprParser) EnterExpressions(ctx *expr.ExpressionsContext) {
	p.coverage.setTotal(ctx.GetStop().GetTokenIndex() + 1)
}

func (p *exprParser) ExitMisc(ctx *expr.MiscContext) {
	p.coverage.add(ctx)
}

func (p *exprParser) ExitExpression(ctx *expr.ExpressionContext) {
	if p.error != nil {
		return
	}
	
	if ctx.GetUnary_op() != nil {
		left, err := p.stack.Pop()
		if err != nil {
			p.error = parseError(p.file, ctx, err.Error())
			return
		}
		err = p.unaryCalc(left, ctx, ctx.GetUnary_op())
		if err != nil {
			p.error = err
			return
		}
		p.coverage.add(ctx)
	} else if ctx.GetMul_op() != nil ||
		ctx.GetAdd_op() != nil ||
		ctx.GetRel_op() != nil ||
		ctx.LOGICAL_AND() != nil ||
		ctx.LOGICAL_OR() != nil {
		left, right, err := p.popBinaryOperands()
		if err != nil {
			p.error = parseError(p.file, ctx, err.Error())
			return
		}
		if ctx.GetAdd_op() != nil {
			err = p.numericStringCalc(left, right, ctx, ctx.GetAdd_op())
			if err != nil {
				p.error = err
				return
			}
			p.coverage.add(ctx)
		} else if ctx.GetMul_op() != nil {
			err = p.numericStringCalc(left, right, ctx, ctx.GetMul_op())
			if err != nil {
				p.error = err
				return
			}
			p.coverage.add(ctx)
		} else if ctx.GetRel_op() != nil {
			err = p.relationCalc(left, right, ctx, ctx.GetRel_op())
			if err != nil {
				p.error = err
				return
			}
			p.coverage.add(ctx)
		} else if ctx.LOGICAL_AND() != nil {
			err = p.logicCalc(left, right, ctx, ctx.LOGICAL_AND().GetSymbol())
			if err != nil {
				p.error = err
				return
			}
			p.coverage.add(ctx)
		} else if ctx.LOGICAL_OR() != nil {
			err = p.logicCalc(left, right, ctx, ctx.LOGICAL_OR().GetSymbol())
			if err != nil {
				p.error = err
				return
			}
			p.coverage.add(ctx)
		}
	}
}

func (p *exprParser) ExitVar_(ctx *expr.Var_Context) {
	if p.error != nil {
		return
	}
	alias := ctx.IDENTIFIER().GetText()
	var val exprValue
	if p.isBuiltIn(alias) {
		val = exprValue{value: p.builtIn[alias], alias: alias}
	} else {
		var ok bool
		if p.foreachParams != nil {
			val, ok = p.foreachParams.get(alias)
			if !ok {
				val, ok = p.baseParams.get(alias)
			}
		} else {
			val, ok = p.baseParams.get(alias)
		}
		if !ok {
			p.error = parseError(p.file, ctx, fmt.Sprintf("can't find var '%s'", alias))
			return
		}
	}
	p.stack.Push(&val)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitMember(ctx *expr.MemberContext) {
	if p.error != nil {
		return
	}
	name := ctx.IDENTIFIER().GetText()
	obj, err := p.stack.Pop()
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	var mev *exprValue
	mev, err = obj.visitMember(name)
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	
	p.stack.Push(mev)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitIndex(ctx *expr.IndexContext) {
	
	if p.error != nil {
		return
	}
	
	index, err := p.stack.Pop()
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	
	object, err := p.stack.Pop()
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	objectReflectElem := object.reflectElem()
	var ev *exprValue
	if objectReflectElem.Kind() == reflect.Map {
		ev, err = object.visitMapIndex(index.reflectElem())
		if err != nil {
			p.error = parseError(p.file, ctx, err.Error())
			return
		}
	} else if objectReflectElem.Kind() == reflect.Slice || objectReflectElem.Kind() == reflect.Array {
		var i int
		i, err = index.int()
		if err != nil {
			p.error = parseError(p.file, ctx, fmt.Sprintf("array index error: %s", err))
			return
		}
		
		ev, err = object.visitArrayIndex(i)
		if err != nil {
			p.error = parseError(p.file, ctx, fmt.Sprintf("array index error: %s", err))
			return
		}
		
	} else {
		p.error = parseError(p.file, ctx, "unsupported index object")
		return
	}
	p.stack.Push(ev)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitSlice_(ctx *expr.Slice_Context) {
	if p.error != nil {
		return
	}
	l := len(ctx.AllExpression())
	args := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		arg, err := p.stack.Pop()
		if err != nil {
			p.error = parseError(p.file, ctx, err.Error())
			return
		}
		args[i], err = arg.int()
		if err != nil {
			p.error = parseError(p.file, ctx, err.Error())
			return
		}
	}
	target, err := p.stack.Pop()
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	r, err := target.visitSlice(ctx.GetText(), args...)
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	p.stack.Push(r)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitCall(ctx *expr.CallContext) {
	if p.error != nil {
		return
	}
	l := len(ctx.ExpressionList().(*expr.ExpressionListContext).AllExpression())
	args := make([]reflect.Value, l)
	for i := l - 1; i >= 0; i-- {
		arg, err := p.stack.Pop()
		if err != nil {
			p.error = parseError(p.file, ctx, err.Error())
			return
		}
		args[i] = arg.reflectElem()
	}
	f, err := p.stack.Pop()
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	r, err := f.call(ctx.ELLIPSIS() != nil, args)
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	p.stack.Push(r)
	p.coverage.add(ctx)
	return
}

func (p *exprParser) ExitInteger(ctx *expr.IntegerContext) {
	if p.error != nil {
		return
	}
	v, err := cast.ToIntE(ctx.GetText())
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	p.stack.Push(&exprValue{
		value:  v,
		source: ctx.GetText(),
	})
	p.coverage.add(ctx)
}

func (p *exprParser) ExitString_(ctx *expr.String_Context) {
	if p.error != nil {
		return
	}
	v := strings.TrimSuffix(strings.TrimPrefix(ctx.GetText(), "\""), "\"")
	p.stack.Push(&exprValue{
		value: v,
	})
	p.coverage.add(ctx)
}

func (p *exprParser) ExitFloat_(ctx *expr.Float_Context) {
	if p.error != nil {
		return
	}
	v, err := cast.ToDecimalE(ctx.GetText())
	if err != nil {
		p.error = parseError(p.file, ctx, err.Error())
		return
	}
	p.stack.Push(&exprValue{
		value:  v,
		source: ctx.GetText(),
	})
	p.coverage.add(ctx)
}

func (p *exprParser) parseParameter(params string) (err error) {
	parser, err := p.parser(params)
	if err != nil {
		return
	}
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Parameters())
	return
}

func (p *exprParser) parseExpression(expresion string) (result interface{}, err error) {
	
	p.stack = newExprStack()
	p.coverage = newCoverage()
	
	parser, err := p.parser(expresion)
	if err != nil {
		return
	}
	
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Expressions())
	
	err = p.error
	if err != nil {
		return
	}
	if p.stack.Len() != 1 {
		err = fmt.Errorf("unexpected reslut stack length: %d", p.stack.Len())
		return
	}
	if !p.coverage.covered() {
		err = fmt.Errorf("parse expression token not coverd: %d/%d", p.coverage.len(), p.coverage.total)
		return
	}
	
	v, _ := p.stack.Pop()
	result = v.value
	
	return
}

func (p *exprParser) popBinaryOperands() (left, right *exprValue, err error) {
	right, err = p.stack.Pop()
	if err != nil {
		return
	}
	left, err = p.stack.Pop()
	if err != nil {
		return
	}
	return
}

func (p *exprParser) numericStringCalc(left, right *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) error {
	
	var err error
	var result interface{}
	switch op.GetTokenType() {
	case expr.ExprParserPLUS:
		result, err = cast.AddAnyE(left.value, right.value)
	case expr.ExprParserMINUS:
		result, err = cast.SubAnyE(left.value, right.value)
	case expr.ExprParserSTAR:
		result, err = cast.MulAnyE(left.value, right.value)
	case expr.ExprParserDIV:
		result, err = cast.DivAnyE(left.value, right.value)
	case expr.ExprParserCARET:
		result, err = cast.CaretAnyE(left.value, right.value)
	case expr.ExprParserOR:
		result, err = cast.OrAnyE(left.value, right.value)
	case expr.ExprParserMOD:
		result, err = cast.ModAnyE(left.value, right.value)
	case expr.ExprParserLSHIFT:
		result, err = cast.LeftShiftAnyE(left.value, right.value)
	case expr.ExprParserRSHIFT:
		result, err = cast.RightShiftAnyE(left.value, right.value)
	case expr.ExprParserBIT_CLEAR:
		result, err = cast.BitClearAnyE(left.value, right.value)
	default:
		return p.unsupportedOpError(ctx)
	}
	if err != nil {
		return parseError(p.file, ctx, err.Error())
	}
	p.stack.Push(&exprValue{value: result})
	return nil
}

func (p *exprParser) relationCalc(left, right *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) error {
	
	var err error
	var result bool
	switch op.GetTokenType() {
	case expr.ExprParserEQUALS:
		result, err = cast.EqualAnyE(left.value, right.value)
	case expr.ExprParserNOT_EQUALS:
		result, err = cast.NotEqualAnyE(left.value, right.value)
	case expr.ExprParserLESS:
		result, err = cast.LessAnyE(left.value, right.value)
	case expr.ExprParserLESS_OR_EQUALS:
		result, err = cast.LessOrEqualAnyE(left.value, right.value)
	case expr.ExprParserGREATER:
		result, err = cast.GreaterAnyE(left.value, right.value)
	case expr.ExprParserGREATER_OR_EQUALS:
		result, err = cast.GreaterOrEqualAnyE(left.value, right.value)
	default:
		return p.unsupportedOpError(ctx)
	}
	if err != nil {
		return parseError(p.file, ctx, err.Error())
	}
	p.stack.Push(&exprValue{value: result})
	
	return nil
}

func (p *exprParser) unaryCalc(left *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) error {
	var err error
	var result interface{}
	switch op.GetTokenType() {
	case expr.ExprParserPLUS:
		result, err = cast.UnaryPlusAnyE(left.value)
	case expr.ExprParserMINUS:
		result, err = cast.UnaryMinusAnyE(left.value)
	case expr.ExprParserCARET:
		result, err = cast.UnaryCaretAnyE(left.value)
	case expr.ExprParserEXCLAMATION:
		result, err = cast.UnaryNotAnyE(left.value)
	default:
		return p.unsupportedOpError(ctx)
	}
	if err != nil {
		return parseError(p.file, ctx, err.Error())
	}
	p.stack.Push(&exprValue{value: result})
	return nil
}

func (p *exprParser) logicCalc(left, right *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) error {
	var err error
	var result interface{}
	switch op.GetTokenType() {
	case expr.ExprParserLOGICAL_AND:
		result, err = cast.LogicAndAnyE(left.value, right.value)
	case expr.ExprParserLOGICAL_OR:
		result, err = cast.LogicOrAnyE(left.value, right.value)
	}
	if err != nil {
		return parseError(p.file, ctx, err.Error())
	}
	p.stack.Push(&exprValue{value: result})
	
	return nil
}

func (p *exprParser) unsupportedOpError(ctx antlr.ParserRuleContext) error {
	return parseError(p.file, ctx, fmt.Sprintf("unsupported operation"))
}

func (p *exprParser) parser(data string) (parser *expr.ExprParser, err error) {
	lexer := expr.NewExprLexer(antlr.NewInputStream(data))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser = expr.NewExprParser(stream)
	parser.BuildParseTrees = true
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	return
}
