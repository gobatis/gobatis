package gobatis

type Params = map[string]interface{}

//func newMapperCache() *mapperCache {
//	return &mapperCache{}
//}
//
//type mapperCache struct {
//	sync.Mutex
//	Cache      *xmlNode
//	Statements map[string]*xmlNode
//}
//
//func (p *mapperCache) addStatement(file string, token antlr.Token, id string, node *xmlNode) (err error) {
//	p.Lock()
//	defer p.Unlock()
//	if p.Statements == nil {
//		p.Statements = map[string]*xmlNode{}
//	}
//	_, ok := p.Statements[id]
//	if ok {
//		err = parseError(file, token, fmt.Sprintf("duplicate statement: %s", id))
//		return
//	}
//	p.Statements[id] = node
//	return
//}

type Mapper struct {
	Engine *Engine
}
