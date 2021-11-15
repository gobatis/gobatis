package generator

type Render interface {
	MakeFieldEntity(field Type) []string
	MakeFieldParamFuncXML(field Type) []string
	MakeFieldEntityFuncXML(field Type) []string
	MakeEntityFuncXML(fields []Type) []string
	MakeParamsFuncXML(fields []Type) []string
	MakeFieldMapperFunc(field Type) []string
	MakeEntityMapperFunc(fields []Type) []string
	MakeTestCases(fields []Type) []string
}

type PGRender struct {
}

func (P PGRender) MakeFieldEntity(field Type) []string {
	panic("implement me")
}

func (P PGRender) MakeFieldParamFuncXML(field Type) []string {
	panic("implement me")
}

func (P PGRender) MakeFieldEntityFuncXML(field Type) []string {
	panic("implement me")
}

func (P PGRender) MakeEntityFuncXML(field Type) []string {
	panic("implement me")
}

func (P PGRender) MakeParamsFuncXML(field Type) []string {
	panic("implement me")
}

func (P PGRender) MakeFieldMapperFunc(field Type) []string {
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
	// InsertParameterBigintInt64OriginalOriginalContext
	// InsertParameterBigintInt64OriginalOriginalStmt
	
	// MustInsertParameterBigintInt64OriginalOriginal
	
	// EmbedInsertParameterBigintInt64
	
	// InsertParameterBigintInt8
	// InsertParameterBigintInt16
	// InsertParameterBigintInt32
	return fs
}

func (P PGRender) MakeEntityMapperFunc(field Type) []string {
	var fs []string
	
	return fs
}
