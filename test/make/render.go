package make

type Render interface {
	MakeFieldEntity(field Field) []string
	MakeFieldParamFuncXML(field Field) []string
	MakeFieldEntityFuncXML(field Field) []string
	MakeEntityFuncXML(fields []Field) []string
	MakeParamsFuncXML(fields []Field) []string
	MakeFieldMapperFunc(field Field) []string
	MakeEntityMapperFunc(fields []Field) []string
	MakeTestCases(fields []Field) []string
}

type PGRender struct {
}

func (P PGRender) MakeFieldEntity(field Field) []string {
	panic("implement me")
}

func (P PGRender) MakeFieldParamFuncXML(field Field) []string {
	panic("implement me")
}

func (P PGRender) MakeFieldEntityFuncXML(field Field) []string {
	panic("implement me")
}

func (P PGRender) MakeEntityFuncXML(field Field) []string {
	panic("implement me")
}

func (P PGRender) MakeParamsFuncXML(field Field) []string {
	panic("implement me")
}

func (P PGRender) MakeFieldMapperFunc(field Field) []string {
	var fs []string
	// InsertParameterBigintInt64
	// InsertParameterBigintInt64OriginalPointer
	// InsertParameterBigintInt64PointerOriginal
	// InsertParameterBigintInt64PointerPointer
	
	// InsertEntityBigintInt64
	// InsertEntityBigintInt64OriginalPointer
	// InsertEntityBigintInt64PointerOriginal
	// InsertEntityBigintInt64PointerPointer
	
	// InsertParameterBigintInt64OriginalOriginalTx
	// InsertParameterBigintInt64OriginalOriginalRows
	// MustInsertParameterBigintInt64OriginalOriginal
	
	// EmbedInsertParameterBigintInt64
	// InsertParameterBigintInt64OriginalOriginalContext
	// InsertParameterBigintInt64OriginalOriginalStmt
	
	// InsertParameterBigintInt8
	// InsertParameterBigintInt16
	// InsertParameterBigintInt32
	return fs
}

func (P PGRender) MakeEntityMapperFunc(field Field) []string {
	var fs []string
	
	return fs
}
