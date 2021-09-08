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

const (
	config_xml = "gobatis.xml"
)

func parseConfig(engine *Engine, file, content string) (err error) {
	defer func() {
		e := recover()
		err = castRecoverError(file, e)
	}()
	l := &xmlParser{
		file:          file,
		stack:         newXMLStack(),
		coverage:      newCoverage(),
		rootElement:   dtd.Configuration,
		elementGetter: dtd.ConfigElement,
	}
	walkXMLNodes(l, content)
	if !l.coverage.covered() {
		throw(file, nil, parseCoveredErr).
			format("parse config token not coverd: %d/%d", l.coverage.len(), l.coverage.total)
	}
	return
}

func parseMapper(engine *Engine, file, content string) (err error) {
	
	defer func() {
		e := recover()
		err = castRecoverError(file, e)
	}()
	
	l := &xmlParser{
		file:          file,
		stack:         newXMLStack(),
		coverage:      newCoverage(),
		rootElement:   dtd.Mapper,
		elementGetter: dtd.MapperElement,
	}
	walkXMLNodes(l, content)
	if !l.coverage.covered() {
		throw(file, nil, parseCoveredErr).
			format("parse mapper token not covered: %d/%d", l.coverage.len(), l.coverage.total)
	}
	
	if l.rootNode == nil {
		engine.logger.Warnf("empty mapperCache file: %s", file)
		return
	}
	for _, v := range l.rootNode.Nodes {
		switch v.Name {
		case dtd.SELECT, dtd.INSERT, dtd.DELETE, dtd.UPDATE, dtd.SQL:
			id := v.GetAttribute(dtd.ID)
			if id == "" {
				throw(file, v.ctx, parseMapperErr).format("fragment: %s miss id", v.Name)
			}
			engine.addFragment(file, v.ctx, id, v)
		}
	}
	return
}

func walkXMLNodes(listener antlr.ParseTreeListener, tokens string) {
	lexer := xml.NewXMLLexer(antlr.NewInputStream(strings.TrimSpace(tokens)))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(newErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := xml.NewXMLParser(stream)
	parser.BuildParseTrees = true
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	parser.SetErrorHandler(newParserErrorStrategy())
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Document())
}

func initExprParser(tokens string) (parser *expr.ExprParser) {
	lexer := expr.NewExprLexer(antlr.NewInputStream(tokens))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(newErrorListener())
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser = expr.NewExprParser(stream)
	parser.BuildParseTrees = true
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(false))
	parser.SetErrorHandler(newParserErrorStrategy())
	return
}

func parseFragment(db *DB, file, id string, node *xmlNode) (frag *fragment, err error) {
	
	defer func() {
		e := recover()
		err = castRecoverError(file, e)
	}()
	
	frag = &fragment{
		db:   db,
		id:   id,
		node: node,
	}
	
	frag.setResultAttribute()
	
	if node.HasAttribute(dtd.PARAMETER) {
		frag.in = frag.parseParams(node.GetAttribute(dtd.PARAMETER))
	}
	if node.HasAttribute(dtd.RESULT) {
		frag.out = frag.parseParams(node.GetAttribute(dtd.RESULT))
	}
	return
}

func toReflectValueElem(source interface{}) reflect.Value {
	v := reflect.ValueOf(source)
	for {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		} else {
			return v
		}
	}
}

func trimValueQuote(s string) string {
	if len(s) > 1 {
		switch s[0] {
		case 34, 39:
			return s[1 : len(s)-1]
		}
	}
	return s
}

func varToReflectKind(t string) (kind reflect.Kind, err error) {
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
	case "slice":
		kind = reflect.Slice
	case "interface":
		kind = reflect.Interface
	case "map":
		kind = reflect.Map
	case "string":
		kind = reflect.String
	case "struct":
		kind = reflect.Struct
	default:
		err = fmt.Errorf("unsupported var type '%s'", t)
	}
	return
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
	stack         *xmlNodeStack
	rootNode      *xmlNode
	rootElement   *dtd.Element
	elementGetter func(name string) (elem *dtd.Element, err error)
}

func (p *xmlParser) validateNode(node *xmlNode, elem *dtd.Element) {
	
	// check required attributes
	if elem.Attributes != nil {
		for k, v := range elem.Attributes {
			if v == dtd.REQUIRED && !node.HasAttribute(k) {
				throw(p.file, node.ctx, validateXMLNodeErr).
					format("element: %s miss required attribute: %s", node.Name, k)
			}
		}
	}
	
	switch node.Name {
	case dtd.SELECT, dtd.INSERT, dtd.UPDATE, dtd.DELETE:
		p.checkResultConflict(node)
	}
	
	for _, childNode := range node.Nodes {
		
		if childNode.textOnly {
			continue
		}
		
		// check not supported node
		if !elem.HasNode(childNode.Name) {
			throw(p.file, childNode.ctx, validateXMLNodeErr).
				format("element: %s not support child element: %s", node.Name, childNode.Name)
		}
		
		// check at most once node
		if elem.GetNodeCount(childNode.Name) == dtd.AT_MOST_ONCE && node.countNode(childNode.Name) > 1 {
			throw(p.file, childNode.ctx, validateXMLNodeErr).
				format("element: %s not support duplicate element: %s", node.Name, childNode.Name)
		}
		
		childElem, err := p.elementGetter(childNode.Name)
		if err != nil {
			throw(p.file, childNode.ctx, validateXMLNodeErr).with(err)
		}
		
		p.validateNode(childNode, childElem)
	}
	
	// check at least once node
	if elem.Nodes != nil {
		for k, v := range elem.Nodes {
			if v == dtd.AT_LEAST_ONCE && node.countNode(k) == 0 {
				throw(p.file, node.ctx, validateXMLNodeErr).
					format("element %s miss required element %s", node.Name, k)
			}
		}
	}
}

func (p *xmlParser) checkResultConflict(node *xmlNode) {
	l := map[string]bool{}
	if node.HasAttribute(dtd.RESULT) {
		l[dtd.RESULT] = true
	}
	if node.HasAttribute(dtd.RESULT_MAP) {
		l[dtd.RESULT_MAP] = true
	}
	if len(l) > 1 {
		attrs := ""
		for k := range l {
			attrs += k + ", "
		}
		attrs = strings.TrimSuffix(attrs, ", ")
		throw(p.file, node.ctx, resultAttributeConflictErr).format("%s attribute conflict", attrs)
	}
}

func (p *xmlParser) enterElement(c *xml.ElementContext) {
	name := c.Name(0)
	if p.depth == 0 {
		if name.GetText() != p.rootElement.Name {
			throw(p.file, c, parseMapperErr).format("top level element %s unsupported", name.GetText())
		}
	}
	p.stack.Push(newXMLNode(p.file, name.GetText(), c, name.GetSymbol()))
	p.depth++
}

func (p *xmlParser) enterReference(c *xml.ReferenceContext) {
	if c.EntityRef() != nil {
		v := ""
		switch c.EntityRef().GetText() {
		case "&lt;":
			v = "<"
		case "&gt":
			v = ">"
		case "&amp;":
			v = "&"
		case "&apos;":
			v = "'"
		case "&quot;":
			v = "\""
		}
		if v != "" {
			p.stack.Peak().AddNode(&xmlNode{File: p.file, Text: v, ctx: c, start: c.GetStart(), textOnly: true})
			p.coverage.add(c)
		}
	}
}

func (p *xmlParser) exitElement(_ *xml.ElementContext) {
	if p.stack.Peak() == nil {
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
	if p.stack.Peak() == nil {
		return
	}
	name := strings.TrimSpace(c.Name().GetText())
	value := strings.TrimSpace(trimValueQuote(c.STRING().GetText()))
	p.stack.Peak().AddAttribute(name, &xmlNodeAttribute{
		File:  p.file,
		Value: value,
		Start: c.STRING().GetSymbol(),
		ctx:   c,
	})
}

func (p *xmlParser) enterChardata(c *xml.ChardataContext) {
	if p.stack.Peak() == nil {
		return
	}
	p.stack.Peak().AddNode(&xmlNode{File: p.file, Text: c.GetText(), ctx: c, start: c.GetStart(), textOnly: true})
}

func (p *xmlParser) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.GetRuleIndex() {
	case xml.XMLParserRULE_document:
		p.coverage.setTotal(ctx.GetStop().GetTokenIndex() + 1)
	case xml.XMLParserRULE_reference:
		p.enterReference(ctx.(*xml.ReferenceContext))
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

type param struct {
	name  string
	_type string
	slice bool
}

func (p param) Type() string {
	if p.slice {
		return fmt.Sprintf("[]%s", p._type)
	}
	return p._type
}

func (p param) expected(vt reflect.Type) bool {
	
	vt = reflectTypeElem(vt)
	if p.slice {
		if vt.Kind() != reflect.Slice && vt.Kind() != reflect.Array {
			return false
		}
		vt = vt.Elem()
	}
	vt = reflectTypeElem(vt)
	
	switch p._type {
	case reflect.Interface.String():
		return true
	case reflect.Struct.String():
		if vt.Kind() == reflect.Struct {
			return true
		} else {
			return false
		}
	case reflect.Map.String():
		if vt.Kind() == reflect.Map {
			return true
		} else {
			return false
		}
	case reflect.Slice.String(), reflect.Array.String():
		if vt.Kind() == reflect.Slice || vt.Kind() == reflect.Array {
			return true
		} else {
			return false
		}
	default:
		if vt.String() == p._type {
			return true
		}
		return false
	}
}

func handleSlice(_type string) (string, bool) {
	slice := strings.HasPrefix(_type, "[]")
	if slice {
		return strings.TrimSpace(strings.TrimPrefix(_type, "[]")), true
	}
	return _type, false
}

func newParamParser(file string) *paramParser {
	r := new(paramParser)
	r.coverage = newCoverage()
	r.file = file
	return r
}

type paramParser struct {
	*expr.BaseExprParserListener
	coverage *coverage
	file     string
	index    int
	params   []*param
	check    map[string]bool
}

func (p *paramParser) EnterParamDecl(ctx *expr.ParamDeclContext) {
	var name string
	if ctx.IDENTIFIER() != nil {
		name = ctx.IDENTIFIER().GetText()
		if _builtin.is(name) {
			throw(p.file, ctx, parameterConflictWithBuiltInErr).format("'%s' conflict with builtin", name)
		}
	}
	var _type string
	if ctx.ParamType() != nil {
		_type = ctx.ParamType().GetText()
	}
	slice := false
	if _type == "" {
		_type = reflect.Interface.String()
	} else {
		_type, slice = handleSlice(_type)
	}
	p.addParam(ctx, name, _type, slice)
	p.index++
}

func (p *paramParser) walkMethods(parser *expr.ExprParser) {
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Parameters())
	if !p.coverage.covered() {
		throw(p.file, nil, parseCoveredErr).
			format("parse mapper method not covered: %d/%d", p.coverage.len(), p.coverage.total)
	}
}

func (p *paramParser) addParam(ctx antlr.ParserRuleContext, name, _type string, slice bool) {
	if p.check == nil {
		p.check = map[string]bool{}
	}
	if _, ok := p.check[name]; ok {
		throw(p.file, ctx, checkParameterErr).format("duplicated parameter '%s'", name)
	}
	p.check[name] = true
	p.params = append(p.params, &param{name: name, _type: _type, slice: slice})
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

func (p *exprValue) visitMember(name string) (r *exprValue, err error) {
	
	elem := toReflectValueElem(p.value)
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

func (p *exprValue) visitArray(index int) (r *exprValue, err error) {
	
	elem := toReflectValueElem(p.value)
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

func (p *exprValue) visitMap(index reflect.Value) (r *exprValue, err error) {
	
	elem := toReflectValueElem(p.value)
	if elem.Kind() != reflect.Map {
		err = fmt.Errorf("visit '%s' is not map", elem.Kind())
		return
	}
	
	mv := elem.MapIndex(index)
	r = &exprValue{}
	if mv.Kind() == reflect.Invalid {
		r.value = nil
	} else {
		r.value = mv.Interface()
	}
	return
}

func (p *exprValue) call(ellipsis bool, params []reflect.Value) (r *exprValue, err error) {
	
	elem := toReflectValueElem(p.value)
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
		v = toReflectValueElem(p.value).Slice(indexes[0], indexes[1])
	case 3:
		v = toReflectValueElem(p.value).Slice3(indexes[0], indexes[1], indexes[2])
	default:
		err = fmt.Errorf("unsuppoted slice range index '%s'", format)
		return
	}
	r = &exprValue{value: v.Interface()}
	return
}

type valueStack struct {
	list *list.List
}

func newValueStack() *valueStack {
	l := list.New()
	return &valueStack{l}
}

func (p *valueStack) push(value *exprValue) {
	p.list.PushBack(value)
}

func (p *valueStack) pop() (val *exprValue, err error) {
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

func (p *valueStack) peak() *exprValue {
	e := p.list.Back()
	if e != nil {
		return e.Value.(*exprValue)
	}
	return nil
}

func (p *valueStack) len() int {
	return p.list.Len()
}

func (p *valueStack) empty() bool {
	return p.list.Len() == 0
}

func newExprParams(params ...reflect.Value) *exprParams {
	r := &exprParams{}
	r.set(params...)
	return r
}

type exprParams struct {
	values []exprValue
	check  map[string]int
}

func (p *exprParams) get(name string) (val exprValue, ok bool) {
	if p.check == nil {
		return
	}
	index, ok := p.check[name]
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

func (p *exprParams) set(params ...reflect.Value) {
	p.values = make([]exprValue, 0)
	for _, v := range params {
		p.values = append(p.values, exprValue{
			value: v.Interface(),
		})
	}
}

func (p *exprParams) bind(expected *param, index int) error {
	
	if expected.name == "" {
		return fmt.Errorf("parameter name is empty")
	}
	
	if p.check == nil {
		p.check = map[string]int{}
	} else {
		_, ok := p.check[expected.name]
		if ok {
			return fmt.Errorf("duplicated parameter '%s'", expected.name)
		}
	}
	
	vl := len(p.values) - 1
	if index < 0 || index > vl {
		return fmt.Errorf("parameter '%s' index %d out of parameters length %d", expected.name, index, vl)
	}
	
	ev := p.values[index]
	p.check[expected.name] = index
	
	if cast.IsNil(ev.value) {
		return nil
	}
	
	elem := toReflectValueElem(ev.value)
	if !expected.expected(elem.Type()) {
		return fmt.Errorf("parameter '%s' expected '%s', got '%s'", expected.name, expected.Type(), elem.Type())
	}
	
	return nil
}

type paramsStack struct {
	list *list.List
}

func newParamsStack() *paramsStack {
	l := list.New()
	return &paramsStack{l}
}

func (p *paramsStack) push(value *exprParams) {
	p.list.PushBack(value)
}

func (p *paramsStack) pop() (val *exprParams, err error) {
	if p.list.Len() < 1 {
		err = fmt.Errorf("stack is empty unable to pop")
		return
	}
	e := p.list.Back()
	if e != nil {
		p.list.Remove(e)
		val = e.Value.(*exprParams)
		return
	}
	return
}

func (p *paramsStack) peak() *exprParams {
	e := p.list.Back()
	if e != nil {
		return e.Value.(*exprParams)
	}
	return nil
}

func (p *paramsStack) len() int {
	return p.list.Len()
}

func (p *paramsStack) empty() bool {
	return p.list.Len() == 0
}

func (p *paramsStack) getVar(name string) (exprValue, bool) {
	for i := p.list.Back(); i != nil; i = i.Prev() {
		v := i.Value.(*exprParams)
		if r, ok := v.get(name); ok {
			return r, true
		}
	}
	return exprValue{}, false
}

func newExprParser(params ...reflect.Value) *exprParser {
	r := new(exprParser)
	r.paramsStack = newParamsStack()
	r.paramsStack.push(r.builtParams())
	r.paramsStack.push(newExprParams(params...))
	return r
}

type exprParser struct {
	*expr.BaseExprParserListener
	nodeCtx     antlr.ParserRuleContext
	valueStack  *valueStack
	paramsStack *paramsStack
	coverage    *coverage
	file        string
	vars        []interface{}
	exprs       []string
	varIndex    int
	static      bool
}

func (p *exprParser) realVars() ([]interface{}, error) {
	for i, v := range p.vars {
		n, ok := v.(Valuer)
		if ok {
			vv, err := n.Value()
			if err != nil {
				return nil, err
			}
			p.vars[i] = vv
		}
	}
	return p.vars, nil
}

func (p *exprParser) builtParams() *exprParams {
	return newExprParams()
}

func (p *exprParser) EnterExpressions(ctx *expr.ExpressionsContext) {
	p.coverage.setTotal(ctx.GetStop().GetTokenIndex() + 1)
}

func (p *exprParser) ExitMisc(ctx *expr.MiscContext) {
	p.coverage.add(ctx)
}

func (p *exprParser) EnterOperand(ctx *expr.OperandContext) {
	p.coverage.add(ctx)
}

func (p *exprParser) ExitExpression(ctx *expr.ExpressionContext) {
	if ctx.GetUnary_op() != nil {
		left, err := p.valueStack.pop()
		if err != nil {
			p.throw(ctx, popValueStackErr).with(err)
		}
		p.unaryCalc(left, ctx, ctx.GetUnary_op())
		p.coverage.add(ctx)
	} else if ctx.GetMul_op() != nil ||
		ctx.GetAdd_op() != nil ||
		ctx.GetRel_op() != nil ||
		ctx.Logical() != nil {
		left, right, err := p.popBinaryOperands()
		if err != nil {
			p.throw(ctx, popBinaryOperandsErr).with(err)
		}
		if ctx.GetAdd_op() != nil {
			p.numericStringCalc(left, right, ctx, ctx.GetAdd_op())
			p.coverage.add(ctx)
		} else if ctx.GetMul_op() != nil {
			p.numericStringCalc(left, right, ctx, ctx.GetMul_op())
			p.coverage.add(ctx)
		} else if ctx.GetRel_op() != nil {
			p.relationCalc(left, right, ctx, ctx.GetRel_op())
			p.coverage.add(ctx)
		} else if ctx.Logical() != nil {
			p.static = false
			p.logicCalc(left, right, ctx, ctx.Logical().GetStart())
			p.coverage.add(ctx)
		}
	} else if ctx.GetTertiary() != nil {
		p.static = false
		condition, left, right, err := p.popTertiaryOperands()
		if err != nil {
			p.throw(ctx, popTertiaryOperandsErr).with(err)
		}
		p.coverage.add(ctx)
		p.tertiaryCalc(condition, left, right, ctx)
	}
}

func (p *exprParser) ExitVar_(ctx *expr.Var_Context) {
	p.static = false
	alias := ctx.IDENTIFIER().GetText()
	var val exprValue
	if _builtin.is(alias) {
		val = exprValue{value: _builtin.get(alias), alias: alias}
	} else {
		var ok bool
		val, ok = p.paramsStack.getVar(alias)
		if !ok {
			p.throw(ctx, parameterNotFoundErr).format("var '%s' not found", alias)
			return
		}
	}
	p.valueStack.push(&val)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitMember(ctx *expr.MemberContext) {
	name := ctx.IDENTIFIER().GetText()
	obj, err := p.valueStack.pop()
	if err != nil {
		p.throw(ctx, popValueStackErr).with(err)
	}
	var mev *exprValue
	mev, err = obj.visitMember(name)
	if err != nil {
		p.throw(ctx, visitMemberErr).with(err)
	}
	
	p.valueStack.push(mev)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitIndex(ctx *expr.IndexContext) {
	
	index, err := p.valueStack.pop()
	if err != nil {
		p.throw(ctx, popValueStackErr).with(err)
	}
	
	object, err := p.valueStack.pop()
	if err != nil {
		p.throw(ctx, popValueStackErr).with(err)
	}
	objectReflectElem := toReflectValueElem(object.value)
	var ev *exprValue
	if objectReflectElem.Kind() == reflect.Map {
		ev, err = object.visitMap(toReflectValueElem(index.value))
		if err != nil {
			p.throw(ctx, visitMapErr).with(err)
		}
	} else if objectReflectElem.Kind() == reflect.Slice || objectReflectElem.Kind() == reflect.Array {
		var i int
		i, err = index.int()
		if err != nil {
			p.throw(ctx, visitArrayErr).format("parse array index error: %s", err)
		}
		
		ev, err = object.visitArray(i)
		if err != nil {
			p.throw(ctx, visitArrayErr).format("visit array error: %s", err)
		}
		
	} else {
		p.throw(ctx, indexErr).format("parameter unsupported index")
	}
	p.valueStack.push(ev)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitSlice_(ctx *expr.Slice_Context) {
	l := len(ctx.AllExpression())
	args := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		arg, err := p.valueStack.pop()
		if err != nil {
			p.throw(ctx, popValueStackErr).with(err)
		}
		args[i], err = arg.int()
		if err != nil {
			p.throw(ctx, visitArrayErr).with(err)
		}
	}
	target, err := p.valueStack.pop()
	if err != nil {
		p.throw(ctx, popValueStackErr).with(err)
	}
	r, err := target.visitSlice(ctx.GetText(), args...)
	if err != nil {
		p.throw(ctx, visitArrayErr).with(err)
	}
	p.valueStack.push(r)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitCall(ctx *expr.CallContext) {
	var l = 0
	if ctx.ExpressionList() != nil {
		l = len(ctx.ExpressionList().(*expr.ExpressionListContext).AllExpression())
	}
	args := make([]reflect.Value, l)
	for i := l - 1; i >= 0; i-- {
		arg, err := p.valueStack.pop()
		if err != nil {
			p.throw(ctx, popValueStackErr).with(err)
		}
		args[i] = toReflectValueElem(arg.value)
	}
	f, err := p.valueStack.pop()
	if err != nil {
		p.throw(ctx, popValueStackErr).with(err)
	}
	r, err := f.call(ctx.ELLIPSIS() != nil, args)
	if err != nil {
		p.throw(ctx, callErr).with(err)
	}
	p.valueStack.push(r)
	p.coverage.add(ctx)
}

func (p *exprParser) ExitInteger(ctx *expr.IntegerContext) {
	v, err := cast.ToIntE(ctx.GetText())
	if err != nil {
		p.throw(ctx, parseIntegerErr).with(err)
	}
	p.valueStack.push(&exprValue{
		value:  v,
		source: ctx.GetText(),
	})
	p.coverage.add(ctx)
}

func (p *exprParser) ExitString_(ctx *expr.String_Context) {
	p.valueStack.push(&exprValue{
		value: trimValueQuote(ctx.GetText()),
	})
	p.coverage.add(ctx)
}

func (p *exprParser) ExitFloat_(ctx *expr.Float_Context) {
	v, err := cast.ToDecimalE(ctx.GetText())
	if err != nil {
		p.throw(ctx, parseDecimalErr).with(err)
	}
	p.valueStack.push(&exprValue{
		value:  v,
		source: ctx.GetText(),
	})
	p.coverage.add(ctx)
}

func (p *exprParser) ExitNil_(ctx *expr.Nil_Context) {
	p.valueStack.push(&exprValue{
		value: nil,
	})
	p.coverage.add(ctx)
}

func (p *exprParser) parseExpression(nodeCtx antlr.ParserRuleContext,
	expression string) (result interface{}, static bool, err error) {
	
	defer func() {
		e := recover()
		err = castRecoverError(p.file, e)
	}()
	
	p.nodeCtx = nodeCtx
	p.static = false
	p.valueStack = newValueStack()
	p.coverage = newCoverage()
	
	parser := initExprParser(expression)
	antlr.ParseTreeWalkerDefault.Walk(p, parser.Expressions())
	
	if !p.coverage.covered() {
		p.throw(nil, parseCoveredErr).format(
			"parse expression token not covered: %d/%d",
			p.coverage.len(), p.coverage.total,
		)
	}
	
	if p.valueStack.len() != 1 {
		p.throw(nil, popResultErr).format("expect result stack length: 1, got %d", p.valueStack.len())
	}
	v, err := p.valueStack.pop()
	if err != nil {
		p.throw(nil, popResultErr).with(err)
	}
	result = v.value
	static = p.static
	
	return
}

func (p *exprParser) addVar(_expr string, _var interface{}) {
	_var = cast.Indirect(_var)
	p.vars = append(p.vars, _var)
	p.exprs = append(p.exprs, _expr)
}

func (p *exprParser) throw(ctx antlr.ParserRuleContext, code int) *_error {
	return throw(p.file, ctx, code).setParent(p.nodeCtx)
}

func (p *exprParser) popBinaryOperands() (left, right *exprValue, err error) {
	right, err = p.valueStack.pop()
	if err != nil {
		return
	}
	left, err = p.valueStack.pop()
	if err != nil {
		return
	}
	return
}

func (p *exprParser) popTertiaryOperands() (condition, left, right *exprValue, err error) {
	right, err = p.valueStack.pop()
	if err != nil {
		return
	}
	left, err = p.valueStack.pop()
	if err != nil {
		return
	}
	condition, err = p.valueStack.pop()
	if err != nil {
		return
	}
	return
}

func (p *exprParser) numericStringCalc(left, right *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) {
	
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
		p.throw(ctx, unsupportedNumericCalc).format("unsupported numeric calc")
	}
	if err != nil {
		p.throw(ctx, numericCalcErr).with(err)
	}
	
	p.valueStack.push(&exprValue{value: result})
}

func (p *exprParser) relationCalc(left, right *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) {
	
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
		p.throw(ctx, unsupportedRelationCalcErr).format("unsupported relation calc")
	}
	if err != nil {
		p.throw(ctx, relationCalcError).with(err)
		
	}
	p.valueStack.push(&exprValue{value: result})
}

func (p *exprParser) unaryCalc(left *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) {
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
		p.throw(ctx, unsupportedUnaryCalc).format("unsupported %s%v", ctx.GetText(), left.value)
	}
	if err != nil {
		p.throw(ctx, unaryCalcError).with(err)
	}
	p.valueStack.push(&exprValue{value: result})
}

func (p *exprParser) logicCalc(left, right *exprValue, ctx antlr.ParserRuleContext, op antlr.Token) {
	
	var err error
	var result interface{}
	switch op.GetTokenType() {
	case expr.ExprParserLOGICAL_AND, expr.ExprParserLOGICAL_AND_LOWER, expr.ExprParserLOGICAL_AND_UPPER:
		result, err = cast.LogicAndAnyE(left.value, right.value)
	case expr.ExprParserLOGICAL_OR, expr.ExprParserLOGICAL_OR_LOWER, expr.ExprParserLOGICAL_OR_UPPER:
		result, err = cast.LogicOrAnyE(left.value, right.value)
	}
	if err != nil {
		p.throw(ctx, logicCalcErr).with(err)
	}
	p.valueStack.push(&exprValue{value: result})
}

func (p *exprParser) tertiaryCalc(condition, left, right *exprValue, ctx antlr.ParserRuleContext) {
	
	ok, err := cast.ToBoolE(condition.value)
	if err != nil {
		p.throw(ctx, castBoolErr).with(err)
	}
	if ok {
		p.valueStack.push(left)
	} else {
		p.valueStack.push(right)
	}
}
