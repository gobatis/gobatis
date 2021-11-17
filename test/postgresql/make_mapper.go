package postgresql

import ()

type MakeMapper struct {
	InsertParameterBigintInt64                                   func(sid string, source string, var_bigint int64) (int, error)
	InsertParameterBigintInt64OriginalPointer                    func(sid string, source string, var_bigint int64) (*int, error)
	InsertParameterBigintInt64PointerOriginal                    func(sid *string, source *string, var_bigint *int64) (int, error)
	InsertParameterBigintInt64PointerPointer                     func(sid string, source string, var_bigint *int64) (*int, error)
	InsertEntityBigintInt64                                      func(item BigintOriginal) error
	InsertEntityBigintInt64OriginalPointer                       func(item BigintOriginal) error
	InsertEntityBigintInt64PointerOriginal                       func(item *BigintPointer) error
	InsertEntityBigintInt64PointerPointer                        func(item *BigintPointer) error
	InsertArrayParameterBigintInt64                              func(sid string, source string, items []int64) error
	InsertArrayParameterBigintInt64OriginalPointer               func(sid string, source string, items []int64) error
	InsertArrayParameterBigintInt64PointerOriginal               func(sid string, source string, items []*int64) error
	InsertArrayParameterBigintInt64PointerPointer                func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64                                 func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64OriginalPointer                  func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64PointerOriginal                  func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64PointerPointer                   func(sid string, source string, items []*int64) error
	InsertParameterInt8Int8                                      func(sid string, source string, var_int8 int8) (int, error)
	InsertParameterInt8Int8OriginalPointer                       func(sid string, source string, var_int8 int8) (*int, error)
	InsertParameterInt8Int8PointerOriginal                       func(sid *string, source *string, var_int8 *int8) (int, error)
	InsertParameterInt8Int8PointerPointer                        func(sid string, source string, var_int8 *int8) (*int, error)
	InsertEntityInt8Int8                                         func(item Int8Original) error
	InsertEntityInt8Int8OriginalPointer                          func(item Int8Original) error
	InsertEntityInt8Int8PointerOriginal                          func(item *Int8Pointer) error
	InsertEntityInt8Int8PointerPointer                           func(item *Int8Pointer) error
	InsertArrayParameterInt8Int8                                 func(sid string, source string, items []int8) error
	InsertArrayParameterInt8Int8OriginalPointer                  func(sid string, source string, items []int8) error
	InsertArrayParameterInt8Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	InsertArrayParameterInt8Int8PointerPointer                   func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8                                    func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8OriginalPointer                     func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8PointerOriginal                     func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8PointerPointer                      func(sid string, source string, items []*int8) error
	InsertParameterBooleanBool                                   func(sid string, source string, var_boolean bool) (int, error)
	InsertParameterBooleanBoolOriginalPointer                    func(sid string, source string, var_boolean bool) (*int, error)
	InsertParameterBooleanBoolPointerOriginal                    func(sid *string, source *string, var_boolean *bool) (int, error)
	InsertParameterBooleanBoolPointerPointer                     func(sid string, source string, var_boolean *bool) (*int, error)
	InsertEntityBooleanBool                                      func(item BooleanOriginal) error
	InsertEntityBooleanBoolOriginalPointer                       func(item BooleanOriginal) error
	InsertEntityBooleanBoolPointerOriginal                       func(item *BooleanPointer) error
	InsertEntityBooleanBoolPointerPointer                        func(item *BooleanPointer) error
	InsertArrayParameterBooleanBool                              func(sid string, source string, items []bool) error
	InsertArrayParameterBooleanBoolOriginalPointer               func(sid string, source string, items []bool) error
	InsertArrayParameterBooleanBoolPointerOriginal               func(sid string, source string, items []*bool) error
	InsertArrayParameterBooleanBoolPointerPointer                func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBool                                 func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBoolOriginalPointer                  func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBoolPointerOriginal                  func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBoolPointerPointer                   func(sid string, source string, items []*bool) error
	InsertParameterBoolBool                                      func(sid string, source string, var_bool bool) (int, error)
	InsertParameterBoolBoolOriginalPointer                       func(sid string, source string, var_bool bool) (*int, error)
	InsertParameterBoolBoolPointerOriginal                       func(sid *string, source *string, var_bool *bool) (int, error)
	InsertParameterBoolBoolPointerPointer                        func(sid string, source string, var_bool *bool) (*int, error)
	InsertEntityBoolBool                                         func(item BoolOriginal) error
	InsertEntityBoolBoolOriginalPointer                          func(item BoolOriginal) error
	InsertEntityBoolBoolPointerOriginal                          func(item *BoolPointer) error
	InsertEntityBoolBoolPointerPointer                           func(item *BoolPointer) error
	InsertArrayParameterBoolBool                                 func(sid string, source string, items []bool) error
	InsertArrayParameterBoolBoolOriginalPointer                  func(sid string, source string, items []bool) error
	InsertArrayParameterBoolBoolPointerOriginal                  func(sid string, source string, items []*bool) error
	InsertArrayParameterBoolBoolPointerPointer                   func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBool                                    func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBoolOriginalPointer                     func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBoolPointerOriginal                     func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBoolPointerPointer                      func(sid string, source string, items []*bool) error
	InsertParameterCharacterString                               func(sid string, source string, var_character string) (int, error)
	InsertParameterCharacterStringOriginalPointer                func(sid string, source string, var_character string) (*int, error)
	InsertParameterCharacterStringPointerOriginal                func(sid *string, source *string, var_character *string) (int, error)
	InsertParameterCharacterStringPointerPointer                 func(sid string, source string, var_character *string) (*int, error)
	InsertEntityCharacterString                                  func(item CharacterOriginal) error
	InsertEntityCharacterStringOriginalPointer                   func(item CharacterOriginal) error
	InsertEntityCharacterStringPointerOriginal                   func(item *CharacterPointer) error
	InsertEntityCharacterStringPointerPointer                    func(item *CharacterPointer) error
	InsertArrayParameterCharacterString                          func(sid string, source string, items []string) error
	InsertArrayParameterCharacterStringOriginalPointer           func(sid string, source string, items []string) error
	InsertArrayParameterCharacterStringPointerOriginal           func(sid string, source string, items []*string) error
	InsertArrayParameterCharacterStringPointerPointer            func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterString                             func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterStringOriginalPointer              func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterStringPointerOriginal              func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterStringPointerPointer               func(sid string, source string, items []*string) error
	InsertParameterCharString                                    func(sid string, source string, var_char string) (int, error)
	InsertParameterCharStringOriginalPointer                     func(sid string, source string, var_char string) (*int, error)
	InsertParameterCharStringPointerOriginal                     func(sid *string, source *string, var_char *string) (int, error)
	InsertParameterCharStringPointerPointer                      func(sid string, source string, var_char *string) (*int, error)
	InsertEntityCharString                                       func(item CharOriginal) error
	InsertEntityCharStringOriginalPointer                        func(item CharOriginal) error
	InsertEntityCharStringPointerOriginal                        func(item *CharPointer) error
	InsertEntityCharStringPointerPointer                         func(item *CharPointer) error
	InsertArrayParameterCharString                               func(sid string, source string, items []string) error
	InsertArrayParameterCharStringOriginalPointer                func(sid string, source string, items []string) error
	InsertArrayParameterCharStringPointerOriginal                func(sid string, source string, items []*string) error
	InsertArrayParameterCharStringPointerPointer                 func(sid string, source string, items []*string) error
	InsertArrayEntityCharString                                  func(sid string, source string, items []*string) error
	InsertArrayEntityCharStringOriginalPointer                   func(sid string, source string, items []*string) error
	InsertArrayEntityCharStringPointerOriginal                   func(sid string, source string, items []*string) error
	InsertArrayEntityCharStringPointerPointer                    func(sid string, source string, items []*string) error
	InsertParameterCharacterVaryingString                        func(sid string, source string, var_character_varying string) (int, error)
	InsertParameterCharacterVaryingStringOriginalPointer         func(sid string, source string, var_character_varying string) (*int, error)
	InsertParameterCharacterVaryingStringPointerOriginal         func(sid *string, source *string, var_character_varying *string) (int, error)
	InsertParameterCharacterVaryingStringPointerPointer          func(sid string, source string, var_character_varying *string) (*int, error)
	InsertEntityCharacterVaryingString                           func(item CharacterVaryingOriginal) error
	InsertEntityCharacterVaryingStringOriginalPointer            func(item CharacterVaryingOriginal) error
	InsertEntityCharacterVaryingStringPointerOriginal            func(item *CharacterVaryingPointer) error
	InsertEntityCharacterVaryingStringPointerPointer             func(item *CharacterVaryingPointer) error
	InsertArrayParameterCharacterVaryingString                   func(sid string, source string, items []string) error
	InsertArrayParameterCharacterVaryingStringOriginalPointer    func(sid string, source string, items []string) error
	InsertArrayParameterCharacterVaryingStringPointerOriginal    func(sid string, source string, items []*string) error
	InsertArrayParameterCharacterVaryingStringPointerPointer     func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingString                      func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingStringOriginalPointer       func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingStringPointerOriginal       func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingStringPointerPointer        func(sid string, source string, items []*string) error
	InsertParameterVarcharString                                 func(sid string, source string, var_varchar string) (int, error)
	InsertParameterVarcharStringOriginalPointer                  func(sid string, source string, var_varchar string) (*int, error)
	InsertParameterVarcharStringPointerOriginal                  func(sid *string, source *string, var_varchar *string) (int, error)
	InsertParameterVarcharStringPointerPointer                   func(sid string, source string, var_varchar *string) (*int, error)
	InsertEntityVarcharString                                    func(item VarcharOriginal) error
	InsertEntityVarcharStringOriginalPointer                     func(item VarcharOriginal) error
	InsertEntityVarcharStringPointerOriginal                     func(item *VarcharPointer) error
	InsertEntityVarcharStringPointerPointer                      func(item *VarcharPointer) error
	InsertArrayParameterVarcharString                            func(sid string, source string, items []string) error
	InsertArrayParameterVarcharStringOriginalPointer             func(sid string, source string, items []string) error
	InsertArrayParameterVarcharStringPointerOriginal             func(sid string, source string, items []*string) error
	InsertArrayParameterVarcharStringPointerPointer              func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharString                               func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharStringOriginalPointer                func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharStringPointerOriginal                func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharStringPointerPointer                 func(sid string, source string, items []*string) error
	InsertParameterFloat8Float32                                 func(sid string, source string, var_float8 float32) (int, error)
	InsertParameterFloat8Float32OriginalPointer                  func(sid string, source string, var_float8 float32) (*int, error)
	InsertParameterFloat8Float32PointerOriginal                  func(sid *string, source *string, var_float8 *float32) (int, error)
	InsertParameterFloat8Float32PointerPointer                   func(sid string, source string, var_float8 *float32) (*int, error)
	InsertEntityFloat8Float32                                    func(item Float8Original) error
	InsertEntityFloat8Float32OriginalPointer                     func(item Float8Original) error
	InsertEntityFloat8Float32PointerOriginal                     func(item *Float8Pointer) error
	InsertEntityFloat8Float32PointerPointer                      func(item *Float8Pointer) error
	InsertArrayParameterFloat8Float32                            func(sid string, source string, items []float32) error
	InsertArrayParameterFloat8Float32OriginalPointer             func(sid string, source string, items []float32) error
	InsertArrayParameterFloat8Float32PointerOriginal             func(sid string, source string, items []*float32) error
	InsertArrayParameterFloat8Float32PointerPointer              func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat8Float32                               func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat8Float32OriginalPointer                func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat8Float32PointerOriginal                func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat8Float32PointerPointer                 func(sid string, source string, items []*float32) error
	InsertParameterIntegerInt8                                   func(sid string, source string, var_integer int8) (int, error)
	InsertParameterIntegerInt8OriginalPointer                    func(sid string, source string, var_integer int8) (*int, error)
	InsertParameterIntegerInt8PointerOriginal                    func(sid *string, source *string, var_integer *int8) (int, error)
	InsertParameterIntegerInt8PointerPointer                     func(sid string, source string, var_integer *int8) (*int, error)
	InsertEntityIntegerInt8                                      func(item IntegerOriginal) error
	InsertEntityIntegerInt8OriginalPointer                       func(item IntegerOriginal) error
	InsertEntityIntegerInt8PointerOriginal                       func(item *IntegerPointer) error
	InsertEntityIntegerInt8PointerPointer                        func(item *IntegerPointer) error
	InsertArrayParameterIntegerInt8                              func(sid string, source string, items []int8) error
	InsertArrayParameterIntegerInt8OriginalPointer               func(sid string, source string, items []int8) error
	InsertArrayParameterIntegerInt8PointerOriginal               func(sid string, source string, items []*int8) error
	InsertArrayParameterIntegerInt8PointerPointer                func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8                                 func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8OriginalPointer                  func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8PointerOriginal                  func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8PointerPointer                   func(sid string, source string, items []*int8) error
	InsertParameterIntInt8                                       func(sid string, source string, var_int int8) (int, error)
	InsertParameterIntInt8OriginalPointer                        func(sid string, source string, var_int int8) (*int, error)
	InsertParameterIntInt8PointerOriginal                        func(sid *string, source *string, var_int *int8) (int, error)
	InsertParameterIntInt8PointerPointer                         func(sid string, source string, var_int *int8) (*int, error)
	InsertEntityIntInt8                                          func(item IntOriginal) error
	InsertEntityIntInt8OriginalPointer                           func(item IntOriginal) error
	InsertEntityIntInt8PointerOriginal                           func(item *IntPointer) error
	InsertEntityIntInt8PointerPointer                            func(item *IntPointer) error
	InsertArrayParameterIntInt8                                  func(sid string, source string, items []int8) error
	InsertArrayParameterIntInt8OriginalPointer                   func(sid string, source string, items []int8) error
	InsertArrayParameterIntInt8PointerOriginal                   func(sid string, source string, items []*int8) error
	InsertArrayParameterIntInt8PointerPointer                    func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8                                     func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8OriginalPointer                      func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8PointerOriginal                      func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8PointerPointer                       func(sid string, source string, items []*int8) error
	InsertParameterInt4Int8                                      func(sid string, source string, var_int4 int8) (int, error)
	InsertParameterInt4Int8OriginalPointer                       func(sid string, source string, var_int4 int8) (*int, error)
	InsertParameterInt4Int8PointerOriginal                       func(sid *string, source *string, var_int4 *int8) (int, error)
	InsertParameterInt4Int8PointerPointer                        func(sid string, source string, var_int4 *int8) (*int, error)
	InsertEntityInt4Int8                                         func(item Int4Original) error
	InsertEntityInt4Int8OriginalPointer                          func(item Int4Original) error
	InsertEntityInt4Int8PointerOriginal                          func(item *Int4Pointer) error
	InsertEntityInt4Int8PointerPointer                           func(item *Int4Pointer) error
	InsertArrayParameterInt4Int8                                 func(sid string, source string, items []int8) error
	InsertArrayParameterInt4Int8OriginalPointer                  func(sid string, source string, items []int8) error
	InsertArrayParameterInt4Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	InsertArrayParameterInt4Int8PointerPointer                   func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8                                    func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8OriginalPointer                     func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8PointerOriginal                     func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8PointerPointer                      func(sid string, source string, items []*int8) error
	InsertParameterNumericDecimal                                func(sid string, source string, var_numeric decimal) (int, error)
	InsertParameterNumericDecimalOriginalPointer                 func(sid string, source string, var_numeric decimal) (*int, error)
	InsertParameterNumericDecimalPointerOriginal                 func(sid *string, source *string, var_numeric *decimal) (int, error)
	InsertParameterNumericDecimalPointerPointer                  func(sid string, source string, var_numeric *decimal) (*int, error)
	InsertEntityNumericDecimal                                   func(item NumericOriginal) error
	InsertEntityNumericDecimalOriginalPointer                    func(item NumericOriginal) error
	InsertEntityNumericDecimalPointerOriginal                    func(item *NumericPointer) error
	InsertEntityNumericDecimalPointerPointer                     func(item *NumericPointer) error
	InsertArrayParameterNumericDecimal                           func(sid string, source string, items []decimal) error
	InsertArrayParameterNumericDecimalOriginalPointer            func(sid string, source string, items []decimal) error
	InsertArrayParameterNumericDecimalPointerOriginal            func(sid string, source string, items []*decimal) error
	InsertArrayParameterNumericDecimalPointerPointer             func(sid string, source string, items []*decimal) error
	InsertArrayEntityNumericDecimal                              func(sid string, source string, items []*decimal) error
	InsertArrayEntityNumericDecimalOriginalPointer               func(sid string, source string, items []*decimal) error
	InsertArrayEntityNumericDecimalPointerOriginal               func(sid string, source string, items []*decimal) error
	InsertArrayEntityNumericDecimalPointerPointer                func(sid string, source string, items []*decimal) error
	InsertParameterDecimalDecimal                                func(sid string, source string, var_decimal decimal) (int, error)
	InsertParameterDecimalDecimalOriginalPointer                 func(sid string, source string, var_decimal decimal) (*int, error)
	InsertParameterDecimalDecimalPointerOriginal                 func(sid *string, source *string, var_decimal *decimal) (int, error)
	InsertParameterDecimalDecimalPointerPointer                  func(sid string, source string, var_decimal *decimal) (*int, error)
	InsertEntityDecimalDecimal                                   func(item DecimalOriginal) error
	InsertEntityDecimalDecimalOriginalPointer                    func(item DecimalOriginal) error
	InsertEntityDecimalDecimalPointerOriginal                    func(item *DecimalPointer) error
	InsertEntityDecimalDecimalPointerPointer                     func(item *DecimalPointer) error
	InsertArrayParameterDecimalDecimal                           func(sid string, source string, items []decimal) error
	InsertArrayParameterDecimalDecimalOriginalPointer            func(sid string, source string, items []decimal) error
	InsertArrayParameterDecimalDecimalPointerOriginal            func(sid string, source string, items []*decimal) error
	InsertArrayParameterDecimalDecimalPointerPointer             func(sid string, source string, items []*decimal) error
	InsertArrayEntityDecimalDecimal                              func(sid string, source string, items []*decimal) error
	InsertArrayEntityDecimalDecimalOriginalPointer               func(sid string, source string, items []*decimal) error
	InsertArrayEntityDecimalDecimalPointerOriginal               func(sid string, source string, items []*decimal) error
	InsertArrayEntityDecimalDecimalPointerPointer                func(sid string, source string, items []*decimal) error
	InsertParameterFloat4Float32                                 func(sid string, source string, var_float4 float32) (int, error)
	InsertParameterFloat4Float32OriginalPointer                  func(sid string, source string, var_float4 float32) (*int, error)
	InsertParameterFloat4Float32PointerOriginal                  func(sid *string, source *string, var_float4 *float32) (int, error)
	InsertParameterFloat4Float32PointerPointer                   func(sid string, source string, var_float4 *float32) (*int, error)
	InsertEntityFloat4Float32                                    func(item Float4Original) error
	InsertEntityFloat4Float32OriginalPointer                     func(item Float4Original) error
	InsertEntityFloat4Float32PointerOriginal                     func(item *Float4Pointer) error
	InsertEntityFloat4Float32PointerPointer                      func(item *Float4Pointer) error
	InsertArrayParameterFloat4Float32                            func(sid string, source string, items []float32) error
	InsertArrayParameterFloat4Float32OriginalPointer             func(sid string, source string, items []float32) error
	InsertArrayParameterFloat4Float32PointerOriginal             func(sid string, source string, items []*float32) error
	InsertArrayParameterFloat4Float32PointerPointer              func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat4Float32                               func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat4Float32OriginalPointer                func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat4Float32PointerOriginal                func(sid string, source string, items []*float32) error
	InsertArrayEntityFloat4Float32PointerPointer                 func(sid string, source string, items []*float32) error
	InsertParameterSmallintInt8                                  func(sid string, source string, var_smallint int8) (int, error)
	InsertParameterSmallintInt8OriginalPointer                   func(sid string, source string, var_smallint int8) (*int, error)
	InsertParameterSmallintInt8PointerOriginal                   func(sid *string, source *string, var_smallint *int8) (int, error)
	InsertParameterSmallintInt8PointerPointer                    func(sid string, source string, var_smallint *int8) (*int, error)
	InsertEntitySmallintInt8                                     func(item SmallintOriginal) error
	InsertEntitySmallintInt8OriginalPointer                      func(item SmallintOriginal) error
	InsertEntitySmallintInt8PointerOriginal                      func(item *SmallintPointer) error
	InsertEntitySmallintInt8PointerPointer                       func(item *SmallintPointer) error
	InsertArrayParameterSmallintInt8                             func(sid string, source string, items []int8) error
	InsertArrayParameterSmallintInt8OriginalPointer              func(sid string, source string, items []int8) error
	InsertArrayParameterSmallintInt8PointerOriginal              func(sid string, source string, items []*int8) error
	InsertArrayParameterSmallintInt8PointerPointer               func(sid string, source string, items []*int8) error
	InsertArrayEntitySmallintInt8                                func(sid string, source string, items []*int8) error
	InsertArrayEntitySmallintInt8OriginalPointer                 func(sid string, source string, items []*int8) error
	InsertArrayEntitySmallintInt8PointerOriginal                 func(sid string, source string, items []*int8) error
	InsertArrayEntitySmallintInt8PointerPointer                  func(sid string, source string, items []*int8) error
	InsertParameterInt2Int8                                      func(sid string, source string, var_int2 int8) (int, error)
	InsertParameterInt2Int8OriginalPointer                       func(sid string, source string, var_int2 int8) (*int, error)
	InsertParameterInt2Int8PointerOriginal                       func(sid *string, source *string, var_int2 *int8) (int, error)
	InsertParameterInt2Int8PointerPointer                        func(sid string, source string, var_int2 *int8) (*int, error)
	InsertEntityInt2Int8                                         func(item Int2Original) error
	InsertEntityInt2Int8OriginalPointer                          func(item Int2Original) error
	InsertEntityInt2Int8PointerOriginal                          func(item *Int2Pointer) error
	InsertEntityInt2Int8PointerPointer                           func(item *Int2Pointer) error
	InsertArrayParameterInt2Int8                                 func(sid string, source string, items []int8) error
	InsertArrayParameterInt2Int8OriginalPointer                  func(sid string, source string, items []int8) error
	InsertArrayParameterInt2Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	InsertArrayParameterInt2Int8PointerPointer                   func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8                                    func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8OriginalPointer                     func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8PointerOriginal                     func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8PointerPointer                      func(sid string, source string, items []*int8) error
	InsertParameterTextString                                    func(sid string, source string, var_text string) (int, error)
	InsertParameterTextStringOriginalPointer                     func(sid string, source string, var_text string) (*int, error)
	InsertParameterTextStringPointerOriginal                     func(sid *string, source *string, var_text *string) (int, error)
	InsertParameterTextStringPointerPointer                      func(sid string, source string, var_text *string) (*int, error)
	InsertEntityTextString                                       func(item TextOriginal) error
	InsertEntityTextStringOriginalPointer                        func(item TextOriginal) error
	InsertEntityTextStringPointerOriginal                        func(item *TextPointer) error
	InsertEntityTextStringPointerPointer                         func(item *TextPointer) error
	InsertArrayParameterTextString                               func(sid string, source string, items []string) error
	InsertArrayParameterTextStringOriginalPointer                func(sid string, source string, items []string) error
	InsertArrayParameterTextStringPointerOriginal                func(sid string, source string, items []*string) error
	InsertArrayParameterTextStringPointerPointer                 func(sid string, source string, items []*string) error
	InsertArrayEntityTextString                                  func(sid string, source string, items []*string) error
	InsertArrayEntityTextStringOriginalPointer                   func(sid string, source string, items []*string) error
	InsertArrayEntityTextStringPointerOriginal                   func(sid string, source string, items []*string) error
	InsertArrayEntityTextStringPointerPointer                    func(sid string, source string, items []*string) error
	InsertParameterTimeTime                                      func(sid string, source string, var_time time) (int, error)
	InsertParameterTimeTimeOriginalPointer                       func(sid string, source string, var_time time) (*int, error)
	InsertParameterTimeTimePointerOriginal                       func(sid *string, source *string, var_time *time) (int, error)
	InsertParameterTimeTimePointerPointer                        func(sid string, source string, var_time *time) (*int, error)
	InsertEntityTimeTime                                         func(item TimeOriginal) error
	InsertEntityTimeTimeOriginalPointer                          func(item TimeOriginal) error
	InsertEntityTimeTimePointerOriginal                          func(item *TimePointer) error
	InsertEntityTimeTimePointerPointer                           func(item *TimePointer) error
	InsertArrayParameterTimeTime                                 func(sid string, source string, items []time) error
	InsertArrayParameterTimeTimeOriginalPointer                  func(sid string, source string, items []time) error
	InsertArrayParameterTimeTimePointerOriginal                  func(sid string, source string, items []*time) error
	InsertArrayParameterTimeTimePointerPointer                   func(sid string, source string, items []*time) error
	InsertArrayEntityTimeTime                                    func(sid string, source string, items []*time) error
	InsertArrayEntityTimeTimeOriginalPointer                     func(sid string, source string, items []*time) error
	InsertArrayEntityTimeTimePointerOriginal                     func(sid string, source string, items []*time) error
	InsertArrayEntityTimeTimePointerPointer                      func(sid string, source string, items []*time) error
	InsertParameterTimeWithTimezoneTime                          func(sid string, source string, var_time_with_timezone time) (int, error)
	InsertParameterTimeWithTimezoneTimeOriginalPointer           func(sid string, source string, var_time_with_timezone time) (*int, error)
	InsertParameterTimeWithTimezoneTimePointerOriginal           func(sid *string, source *string, var_time_with_timezone *time) (int, error)
	InsertParameterTimeWithTimezoneTimePointerPointer            func(sid string, source string, var_time_with_timezone *time) (*int, error)
	InsertEntityTimeWithTimezoneTime                             func(item TimeWithTimezoneOriginal) error
	InsertEntityTimeWithTimezoneTimeOriginalPointer              func(item TimeWithTimezoneOriginal) error
	InsertEntityTimeWithTimezoneTimePointerOriginal              func(item *TimeWithTimezonePointer) error
	InsertEntityTimeWithTimezoneTimePointerPointer               func(item *TimeWithTimezonePointer) error
	InsertArrayParameterTimeWithTimezoneTime                     func(sid string, source string, items []time) error
	InsertArrayParameterTimeWithTimezoneTimeOriginalPointer      func(sid string, source string, items []time) error
	InsertArrayParameterTimeWithTimezoneTimePointerOriginal      func(sid string, source string, items []*time) error
	InsertArrayParameterTimeWithTimezoneTimePointerPointer       func(sid string, source string, items []*time) error
	InsertArrayEntityTimeWithTimezoneTime                        func(sid string, source string, items []*time) error
	InsertArrayEntityTimeWithTimezoneTimeOriginalPointer         func(sid string, source string, items []*time) error
	InsertArrayEntityTimeWithTimezoneTimePointerOriginal         func(sid string, source string, items []*time) error
	InsertArrayEntityTimeWithTimezoneTimePointerPointer          func(sid string, source string, items []*time) error
	InsertParameterTimetzTime                                    func(sid string, source string, var_timetz time) (int, error)
	InsertParameterTimetzTimeOriginalPointer                     func(sid string, source string, var_timetz time) (*int, error)
	InsertParameterTimetzTimePointerOriginal                     func(sid *string, source *string, var_timetz *time) (int, error)
	InsertParameterTimetzTimePointerPointer                      func(sid string, source string, var_timetz *time) (*int, error)
	InsertEntityTimetzTime                                       func(item TimetzOriginal) error
	InsertEntityTimetzTimeOriginalPointer                        func(item TimetzOriginal) error
	InsertEntityTimetzTimePointerOriginal                        func(item *TimetzPointer) error
	InsertEntityTimetzTimePointerPointer                         func(item *TimetzPointer) error
	InsertArrayParameterTimetzTime                               func(sid string, source string, items []time) error
	InsertArrayParameterTimetzTimeOriginalPointer                func(sid string, source string, items []time) error
	InsertArrayParameterTimetzTimePointerOriginal                func(sid string, source string, items []*time) error
	InsertArrayParameterTimetzTimePointerPointer                 func(sid string, source string, items []*time) error
	InsertArrayEntityTimetzTime                                  func(sid string, source string, items []*time) error
	InsertArrayEntityTimetzTimeOriginalPointer                   func(sid string, source string, items []*time) error
	InsertArrayEntityTimetzTimePointerOriginal                   func(sid string, source string, items []*time) error
	InsertArrayEntityTimetzTimePointerPointer                    func(sid string, source string, items []*time) error
	InsertParameterTimestampTime                                 func(sid string, source string, var_timestamp time) (int, error)
	InsertParameterTimestampTimeOriginalPointer                  func(sid string, source string, var_timestamp time) (*int, error)
	InsertParameterTimestampTimePointerOriginal                  func(sid *string, source *string, var_timestamp *time) (int, error)
	InsertParameterTimestampTimePointerPointer                   func(sid string, source string, var_timestamp *time) (*int, error)
	InsertEntityTimestampTime                                    func(item TimestampOriginal) error
	InsertEntityTimestampTimeOriginalPointer                     func(item TimestampOriginal) error
	InsertEntityTimestampTimePointerOriginal                     func(item *TimestampPointer) error
	InsertEntityTimestampTimePointerPointer                      func(item *TimestampPointer) error
	InsertArrayParameterTimestampTime                            func(sid string, source string, items []time) error
	InsertArrayParameterTimestampTimeOriginalPointer             func(sid string, source string, items []time) error
	InsertArrayParameterTimestampTimePointerOriginal             func(sid string, source string, items []*time) error
	InsertArrayParameterTimestampTimePointerPointer              func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampTime                               func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampTimeOriginalPointer                func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampTimePointerOriginal                func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampTimePointerPointer                 func(sid string, source string, items []*time) error
	InsertParameterTimestampWithTimezoneTime                     func(sid string, source string, var_timestamp_with_timezone time) (int, error)
	InsertParameterTimestampWithTimezoneTimeOriginalPointer      func(sid string, source string, var_timestamp_with_timezone time) (*int, error)
	InsertParameterTimestampWithTimezoneTimePointerOriginal      func(sid *string, source *string, var_timestamp_with_timezone *time) (int, error)
	InsertParameterTimestampWithTimezoneTimePointerPointer       func(sid string, source string, var_timestamp_with_timezone *time) (*int, error)
	InsertEntityTimestampWithTimezoneTime                        func(item TimestampWithTimezoneOriginal) error
	InsertEntityTimestampWithTimezoneTimeOriginalPointer         func(item TimestampWithTimezoneOriginal) error
	InsertEntityTimestampWithTimezoneTimePointerOriginal         func(item *TimestampWithTimezonePointer) error
	InsertEntityTimestampWithTimezoneTimePointerPointer          func(item *TimestampWithTimezonePointer) error
	InsertArrayParameterTimestampWithTimezoneTime                func(sid string, source string, items []time) error
	InsertArrayParameterTimestampWithTimezoneTimeOriginalPointer func(sid string, source string, items []time) error
	InsertArrayParameterTimestampWithTimezoneTimePointerOriginal func(sid string, source string, items []*time) error
	InsertArrayParameterTimestampWithTimezoneTimePointerPointer  func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampWithTimezoneTime                   func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampWithTimezoneTimeOriginalPointer    func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampWithTimezoneTimePointerOriginal    func(sid string, source string, items []*time) error
	InsertArrayEntityTimestampWithTimezoneTimePointerPointer     func(sid string, source string, items []*time) error
	InsertParameterTimestamptzTime                               func(sid string, source string, var_timestamptz time) (int, error)
	InsertParameterTimestamptzTimeOriginalPointer                func(sid string, source string, var_timestamptz time) (*int, error)
	InsertParameterTimestamptzTimePointerOriginal                func(sid *string, source *string, var_timestamptz *time) (int, error)
	InsertParameterTimestamptzTimePointerPointer                 func(sid string, source string, var_timestamptz *time) (*int, error)
	InsertEntityTimestamptzTime                                  func(item TimestamptzOriginal) error
	InsertEntityTimestamptzTimeOriginalPointer                   func(item TimestamptzOriginal) error
	InsertEntityTimestamptzTimePointerOriginal                   func(item *TimestamptzPointer) error
	InsertEntityTimestamptzTimePointerPointer                    func(item *TimestamptzPointer) error
	InsertArrayParameterTimestamptzTime                          func(sid string, source string, items []time) error
	InsertArrayParameterTimestamptzTimeOriginalPointer           func(sid string, source string, items []time) error
	InsertArrayParameterTimestamptzTimePointerOriginal           func(sid string, source string, items []*time) error
	InsertArrayParameterTimestamptzTimePointerPointer            func(sid string, source string, items []*time) error
	InsertArrayEntityTimestamptzTime                             func(sid string, source string, items []*time) error
	InsertArrayEntityTimestamptzTimeOriginalPointer              func(sid string, source string, items []*time) error
	InsertArrayEntityTimestamptzTimePointerOriginal              func(sid string, source string, items []*time) error
	InsertArrayEntityTimestamptzTimePointerPointer               func(sid string, source string, items []*time) error
	SelectParameterBigintInt64                                   func(sid string) (int64, error)
	SelectParameterBigintInt64OriginalPointer                    func(sid string) (*int64, error)
	SelectParameterBigintInt64PointerOriginal                    func(sid string) (int64, error)
	SelectParameterBigintInt64PointerPointer                     func(sid string) (*int64, error)
	SelectEntityBigintInt64                                      func(sid string) (*int64, error)
	SelectEntityBigintInt64OriginalPointer                       func(sid string) (*int64, error)
	SelectEntityBigintInt64PointerOriginal                       func(sid string) (*int64, error)
	SelectEntityBigintInt64PointerPointer                        func(sid string) (*int64, error)
	SelectArrayParameterBigintInt64                              func(id int) ([]int64, error)
	SelectArrayParameterBigintInt64OriginalPointer               func(id int) ([]int64, error)
	SelectArrayParameterBigintInt64PointerOriginal               func(id int) ([]*int64, error)
	SelectArrayParameterBigintInt64PointerPointer                func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64                                 func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64OriginalPointer                  func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64PointerOriginal                  func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64PointerPointer                   func(id int) ([]*int64, error)
	SelectParameterInt8Int8                                      func(sid string) (int8, error)
	SelectParameterInt8Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectParameterInt8Int8PointerOriginal                       func(sid string) (int8, error)
	SelectParameterInt8Int8PointerPointer                        func(sid string) (*int8, error)
	SelectEntityInt8Int8                                         func(sid string) (*int8, error)
	SelectEntityInt8Int8OriginalPointer                          func(sid string) (*int8, error)
	SelectEntityInt8Int8PointerOriginal                          func(sid string) (*int8, error)
	SelectEntityInt8Int8PointerPointer                           func(sid string) (*int8, error)
	SelectArrayParameterInt8Int8                                 func(id int) ([]int8, error)
	SelectArrayParameterInt8Int8OriginalPointer                  func(id int) ([]int8, error)
	SelectArrayParameterInt8Int8PointerOriginal                  func(id int) ([]*int8, error)
	SelectArrayParameterInt8Int8PointerPointer                   func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8                                    func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8OriginalPointer                     func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8PointerOriginal                     func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8PointerPointer                      func(id int) ([]*int8, error)
	SelectParameterBooleanBool                                   func(sid string) (bool, error)
	SelectParameterBooleanBoolOriginalPointer                    func(sid string) (*bool, error)
	SelectParameterBooleanBoolPointerOriginal                    func(sid string) (bool, error)
	SelectParameterBooleanBoolPointerPointer                     func(sid string) (*bool, error)
	SelectEntityBooleanBool                                      func(sid string) (*bool, error)
	SelectEntityBooleanBoolOriginalPointer                       func(sid string) (*bool, error)
	SelectEntityBooleanBoolPointerOriginal                       func(sid string) (*bool, error)
	SelectEntityBooleanBoolPointerPointer                        func(sid string) (*bool, error)
	SelectArrayParameterBooleanBool                              func(id int) ([]bool, error)
	SelectArrayParameterBooleanBoolOriginalPointer               func(id int) ([]bool, error)
	SelectArrayParameterBooleanBoolPointerOriginal               func(id int) ([]*bool, error)
	SelectArrayParameterBooleanBoolPointerPointer                func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBool                                 func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBoolOriginalPointer                  func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBoolPointerOriginal                  func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBoolPointerPointer                   func(id int) ([]*bool, error)
	SelectParameterBoolBool                                      func(sid string) (bool, error)
	SelectParameterBoolBoolOriginalPointer                       func(sid string) (*bool, error)
	SelectParameterBoolBoolPointerOriginal                       func(sid string) (bool, error)
	SelectParameterBoolBoolPointerPointer                        func(sid string) (*bool, error)
	SelectEntityBoolBool                                         func(sid string) (*bool, error)
	SelectEntityBoolBoolOriginalPointer                          func(sid string) (*bool, error)
	SelectEntityBoolBoolPointerOriginal                          func(sid string) (*bool, error)
	SelectEntityBoolBoolPointerPointer                           func(sid string) (*bool, error)
	SelectArrayParameterBoolBool                                 func(id int) ([]bool, error)
	SelectArrayParameterBoolBoolOriginalPointer                  func(id int) ([]bool, error)
	SelectArrayParameterBoolBoolPointerOriginal                  func(id int) ([]*bool, error)
	SelectArrayParameterBoolBoolPointerPointer                   func(id int) ([]*bool, error)
	SelectArrayEntityBoolBool                                    func(id int) ([]*bool, error)
	SelectArrayEntityBoolBoolOriginalPointer                     func(id int) ([]*bool, error)
	SelectArrayEntityBoolBoolPointerOriginal                     func(id int) ([]*bool, error)
	SelectArrayEntityBoolBoolPointerPointer                      func(id int) ([]*bool, error)
	SelectParameterCharacterString                               func(sid string) (string, error)
	SelectParameterCharacterStringOriginalPointer                func(sid string) (*string, error)
	SelectParameterCharacterStringPointerOriginal                func(sid string) (string, error)
	SelectParameterCharacterStringPointerPointer                 func(sid string) (*string, error)
	SelectEntityCharacterString                                  func(sid string) (*string, error)
	SelectEntityCharacterStringOriginalPointer                   func(sid string) (*string, error)
	SelectEntityCharacterStringPointerOriginal                   func(sid string) (*string, error)
	SelectEntityCharacterStringPointerPointer                    func(sid string) (*string, error)
	SelectArrayParameterCharacterString                          func(id int) ([]string, error)
	SelectArrayParameterCharacterStringOriginalPointer           func(id int) ([]string, error)
	SelectArrayParameterCharacterStringPointerOriginal           func(id int) ([]*string, error)
	SelectArrayParameterCharacterStringPointerPointer            func(id int) ([]*string, error)
	SelectArrayEntityCharacterString                             func(id int) ([]*string, error)
	SelectArrayEntityCharacterStringOriginalPointer              func(id int) ([]*string, error)
	SelectArrayEntityCharacterStringPointerOriginal              func(id int) ([]*string, error)
	SelectArrayEntityCharacterStringPointerPointer               func(id int) ([]*string, error)
	SelectParameterCharString                                    func(sid string) (string, error)
	SelectParameterCharStringOriginalPointer                     func(sid string) (*string, error)
	SelectParameterCharStringPointerOriginal                     func(sid string) (string, error)
	SelectParameterCharStringPointerPointer                      func(sid string) (*string, error)
	SelectEntityCharString                                       func(sid string) (*string, error)
	SelectEntityCharStringOriginalPointer                        func(sid string) (*string, error)
	SelectEntityCharStringPointerOriginal                        func(sid string) (*string, error)
	SelectEntityCharStringPointerPointer                         func(sid string) (*string, error)
	SelectArrayParameterCharString                               func(id int) ([]string, error)
	SelectArrayParameterCharStringOriginalPointer                func(id int) ([]string, error)
	SelectArrayParameterCharStringPointerOriginal                func(id int) ([]*string, error)
	SelectArrayParameterCharStringPointerPointer                 func(id int) ([]*string, error)
	SelectArrayEntityCharString                                  func(id int) ([]*string, error)
	SelectArrayEntityCharStringOriginalPointer                   func(id int) ([]*string, error)
	SelectArrayEntityCharStringPointerOriginal                   func(id int) ([]*string, error)
	SelectArrayEntityCharStringPointerPointer                    func(id int) ([]*string, error)
	SelectParameterCharacterVaryingString                        func(sid string) (string, error)
	SelectParameterCharacterVaryingStringOriginalPointer         func(sid string) (*string, error)
	SelectParameterCharacterVaryingStringPointerOriginal         func(sid string) (string, error)
	SelectParameterCharacterVaryingStringPointerPointer          func(sid string) (*string, error)
	SelectEntityCharacterVaryingString                           func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringOriginalPointer            func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringPointerOriginal            func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringPointerPointer             func(sid string) (*string, error)
	SelectArrayParameterCharacterVaryingString                   func(id int) ([]string, error)
	SelectArrayParameterCharacterVaryingStringOriginalPointer    func(id int) ([]string, error)
	SelectArrayParameterCharacterVaryingStringPointerOriginal    func(id int) ([]*string, error)
	SelectArrayParameterCharacterVaryingStringPointerPointer     func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingString                      func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringOriginalPointer       func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringPointerOriginal       func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringPointerPointer        func(id int) ([]*string, error)
	SelectParameterVarcharString                                 func(sid string) (string, error)
	SelectParameterVarcharStringOriginalPointer                  func(sid string) (*string, error)
	SelectParameterVarcharStringPointerOriginal                  func(sid string) (string, error)
	SelectParameterVarcharStringPointerPointer                   func(sid string) (*string, error)
	SelectEntityVarcharString                                    func(sid string) (*string, error)
	SelectEntityVarcharStringOriginalPointer                     func(sid string) (*string, error)
	SelectEntityVarcharStringPointerOriginal                     func(sid string) (*string, error)
	SelectEntityVarcharStringPointerPointer                      func(sid string) (*string, error)
	SelectArrayParameterVarcharString                            func(id int) ([]string, error)
	SelectArrayParameterVarcharStringOriginalPointer             func(id int) ([]string, error)
	SelectArrayParameterVarcharStringPointerOriginal             func(id int) ([]*string, error)
	SelectArrayParameterVarcharStringPointerPointer              func(id int) ([]*string, error)
	SelectArrayEntityVarcharString                               func(id int) ([]*string, error)
	SelectArrayEntityVarcharStringOriginalPointer                func(id int) ([]*string, error)
	SelectArrayEntityVarcharStringPointerOriginal                func(id int) ([]*string, error)
	SelectArrayEntityVarcharStringPointerPointer                 func(id int) ([]*string, error)
	SelectParameterFloat8Float32                                 func(sid string) (float32, error)
	SelectParameterFloat8Float32OriginalPointer                  func(sid string) (*float32, error)
	SelectParameterFloat8Float32PointerOriginal                  func(sid string) (float32, error)
	SelectParameterFloat8Float32PointerPointer                   func(sid string) (*float32, error)
	SelectEntityFloat8Float32                                    func(sid string) (*float32, error)
	SelectEntityFloat8Float32OriginalPointer                     func(sid string) (*float32, error)
	SelectEntityFloat8Float32PointerOriginal                     func(sid string) (*float32, error)
	SelectEntityFloat8Float32PointerPointer                      func(sid string) (*float32, error)
	SelectArrayParameterFloat8Float32                            func(id int) ([]float32, error)
	SelectArrayParameterFloat8Float32OriginalPointer             func(id int) ([]float32, error)
	SelectArrayParameterFloat8Float32PointerOriginal             func(id int) ([]*float32, error)
	SelectArrayParameterFloat8Float32PointerPointer              func(id int) ([]*float32, error)
	SelectArrayEntityFloat8Float32                               func(id int) ([]*float32, error)
	SelectArrayEntityFloat8Float32OriginalPointer                func(id int) ([]*float32, error)
	SelectArrayEntityFloat8Float32PointerOriginal                func(id int) ([]*float32, error)
	SelectArrayEntityFloat8Float32PointerPointer                 func(id int) ([]*float32, error)
	SelectParameterIntegerInt8                                   func(sid string) (int8, error)
	SelectParameterIntegerInt8OriginalPointer                    func(sid string) (*int8, error)
	SelectParameterIntegerInt8PointerOriginal                    func(sid string) (int8, error)
	SelectParameterIntegerInt8PointerPointer                     func(sid string) (*int8, error)
	SelectEntityIntegerInt8                                      func(sid string) (*int8, error)
	SelectEntityIntegerInt8OriginalPointer                       func(sid string) (*int8, error)
	SelectEntityIntegerInt8PointerOriginal                       func(sid string) (*int8, error)
	SelectEntityIntegerInt8PointerPointer                        func(sid string) (*int8, error)
	SelectArrayParameterIntegerInt8                              func(id int) ([]int8, error)
	SelectArrayParameterIntegerInt8OriginalPointer               func(id int) ([]int8, error)
	SelectArrayParameterIntegerInt8PointerOriginal               func(id int) ([]*int8, error)
	SelectArrayParameterIntegerInt8PointerPointer                func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8                                 func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8OriginalPointer                  func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8PointerOriginal                  func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8PointerPointer                   func(id int) ([]*int8, error)
	SelectParameterIntInt8                                       func(sid string) (int8, error)
	SelectParameterIntInt8OriginalPointer                        func(sid string) (*int8, error)
	SelectParameterIntInt8PointerOriginal                        func(sid string) (int8, error)
	SelectParameterIntInt8PointerPointer                         func(sid string) (*int8, error)
	SelectEntityIntInt8                                          func(sid string) (*int8, error)
	SelectEntityIntInt8OriginalPointer                           func(sid string) (*int8, error)
	SelectEntityIntInt8PointerOriginal                           func(sid string) (*int8, error)
	SelectEntityIntInt8PointerPointer                            func(sid string) (*int8, error)
	SelectArrayParameterIntInt8                                  func(id int) ([]int8, error)
	SelectArrayParameterIntInt8OriginalPointer                   func(id int) ([]int8, error)
	SelectArrayParameterIntInt8PointerOriginal                   func(id int) ([]*int8, error)
	SelectArrayParameterIntInt8PointerPointer                    func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8                                     func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8OriginalPointer                      func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8PointerOriginal                      func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8PointerPointer                       func(id int) ([]*int8, error)
	SelectParameterInt4Int8                                      func(sid string) (int8, error)
	SelectParameterInt4Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectParameterInt4Int8PointerOriginal                       func(sid string) (int8, error)
	SelectParameterInt4Int8PointerPointer                        func(sid string) (*int8, error)
	SelectEntityInt4Int8                                         func(sid string) (*int8, error)
	SelectEntityInt4Int8OriginalPointer                          func(sid string) (*int8, error)
	SelectEntityInt4Int8PointerOriginal                          func(sid string) (*int8, error)
	SelectEntityInt4Int8PointerPointer                           func(sid string) (*int8, error)
	SelectArrayParameterInt4Int8                                 func(id int) ([]int8, error)
	SelectArrayParameterInt4Int8OriginalPointer                  func(id int) ([]int8, error)
	SelectArrayParameterInt4Int8PointerOriginal                  func(id int) ([]*int8, error)
	SelectArrayParameterInt4Int8PointerPointer                   func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8                                    func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8OriginalPointer                     func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8PointerOriginal                     func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8PointerPointer                      func(id int) ([]*int8, error)
	SelectParameterNumericDecimal                                func(sid string) (decimal, error)
	SelectParameterNumericDecimalOriginalPointer                 func(sid string) (*decimal, error)
	SelectParameterNumericDecimalPointerOriginal                 func(sid string) (decimal, error)
	SelectParameterNumericDecimalPointerPointer                  func(sid string) (*decimal, error)
	SelectEntityNumericDecimal                                   func(sid string) (*decimal, error)
	SelectEntityNumericDecimalOriginalPointer                    func(sid string) (*decimal, error)
	SelectEntityNumericDecimalPointerOriginal                    func(sid string) (*decimal, error)
	SelectEntityNumericDecimalPointerPointer                     func(sid string) (*decimal, error)
	SelectArrayParameterNumericDecimal                           func(id int) ([]decimal, error)
	SelectArrayParameterNumericDecimalOriginalPointer            func(id int) ([]decimal, error)
	SelectArrayParameterNumericDecimalPointerOriginal            func(id int) ([]*decimal, error)
	SelectArrayParameterNumericDecimalPointerPointer             func(id int) ([]*decimal, error)
	SelectArrayEntityNumericDecimal                              func(id int) ([]*decimal, error)
	SelectArrayEntityNumericDecimalOriginalPointer               func(id int) ([]*decimal, error)
	SelectArrayEntityNumericDecimalPointerOriginal               func(id int) ([]*decimal, error)
	SelectArrayEntityNumericDecimalPointerPointer                func(id int) ([]*decimal, error)
	SelectParameterDecimalDecimal                                func(sid string) (decimal, error)
	SelectParameterDecimalDecimalOriginalPointer                 func(sid string) (*decimal, error)
	SelectParameterDecimalDecimalPointerOriginal                 func(sid string) (decimal, error)
	SelectParameterDecimalDecimalPointerPointer                  func(sid string) (*decimal, error)
	SelectEntityDecimalDecimal                                   func(sid string) (*decimal, error)
	SelectEntityDecimalDecimalOriginalPointer                    func(sid string) (*decimal, error)
	SelectEntityDecimalDecimalPointerOriginal                    func(sid string) (*decimal, error)
	SelectEntityDecimalDecimalPointerPointer                     func(sid string) (*decimal, error)
	SelectArrayParameterDecimalDecimal                           func(id int) ([]decimal, error)
	SelectArrayParameterDecimalDecimalOriginalPointer            func(id int) ([]decimal, error)
	SelectArrayParameterDecimalDecimalPointerOriginal            func(id int) ([]*decimal, error)
	SelectArrayParameterDecimalDecimalPointerPointer             func(id int) ([]*decimal, error)
	SelectArrayEntityDecimalDecimal                              func(id int) ([]*decimal, error)
	SelectArrayEntityDecimalDecimalOriginalPointer               func(id int) ([]*decimal, error)
	SelectArrayEntityDecimalDecimalPointerOriginal               func(id int) ([]*decimal, error)
	SelectArrayEntityDecimalDecimalPointerPointer                func(id int) ([]*decimal, error)
	SelectParameterFloat4Float32                                 func(sid string) (float32, error)
	SelectParameterFloat4Float32OriginalPointer                  func(sid string) (*float32, error)
	SelectParameterFloat4Float32PointerOriginal                  func(sid string) (float32, error)
	SelectParameterFloat4Float32PointerPointer                   func(sid string) (*float32, error)
	SelectEntityFloat4Float32                                    func(sid string) (*float32, error)
	SelectEntityFloat4Float32OriginalPointer                     func(sid string) (*float32, error)
	SelectEntityFloat4Float32PointerOriginal                     func(sid string) (*float32, error)
	SelectEntityFloat4Float32PointerPointer                      func(sid string) (*float32, error)
	SelectArrayParameterFloat4Float32                            func(id int) ([]float32, error)
	SelectArrayParameterFloat4Float32OriginalPointer             func(id int) ([]float32, error)
	SelectArrayParameterFloat4Float32PointerOriginal             func(id int) ([]*float32, error)
	SelectArrayParameterFloat4Float32PointerPointer              func(id int) ([]*float32, error)
	SelectArrayEntityFloat4Float32                               func(id int) ([]*float32, error)
	SelectArrayEntityFloat4Float32OriginalPointer                func(id int) ([]*float32, error)
	SelectArrayEntityFloat4Float32PointerOriginal                func(id int) ([]*float32, error)
	SelectArrayEntityFloat4Float32PointerPointer                 func(id int) ([]*float32, error)
	SelectParameterSmallintInt8                                  func(sid string) (int8, error)
	SelectParameterSmallintInt8OriginalPointer                   func(sid string) (*int8, error)
	SelectParameterSmallintInt8PointerOriginal                   func(sid string) (int8, error)
	SelectParameterSmallintInt8PointerPointer                    func(sid string) (*int8, error)
	SelectEntitySmallintInt8                                     func(sid string) (*int8, error)
	SelectEntitySmallintInt8OriginalPointer                      func(sid string) (*int8, error)
	SelectEntitySmallintInt8PointerOriginal                      func(sid string) (*int8, error)
	SelectEntitySmallintInt8PointerPointer                       func(sid string) (*int8, error)
	SelectArrayParameterSmallintInt8                             func(id int) ([]int8, error)
	SelectArrayParameterSmallintInt8OriginalPointer              func(id int) ([]int8, error)
	SelectArrayParameterSmallintInt8PointerOriginal              func(id int) ([]*int8, error)
	SelectArrayParameterSmallintInt8PointerPointer               func(id int) ([]*int8, error)
	SelectArrayEntitySmallintInt8                                func(id int) ([]*int8, error)
	SelectArrayEntitySmallintInt8OriginalPointer                 func(id int) ([]*int8, error)
	SelectArrayEntitySmallintInt8PointerOriginal                 func(id int) ([]*int8, error)
	SelectArrayEntitySmallintInt8PointerPointer                  func(id int) ([]*int8, error)
	SelectParameterInt2Int8                                      func(sid string) (int8, error)
	SelectParameterInt2Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectParameterInt2Int8PointerOriginal                       func(sid string) (int8, error)
	SelectParameterInt2Int8PointerPointer                        func(sid string) (*int8, error)
	SelectEntityInt2Int8                                         func(sid string) (*int8, error)
	SelectEntityInt2Int8OriginalPointer                          func(sid string) (*int8, error)
	SelectEntityInt2Int8PointerOriginal                          func(sid string) (*int8, error)
	SelectEntityInt2Int8PointerPointer                           func(sid string) (*int8, error)
	SelectArrayParameterInt2Int8                                 func(id int) ([]int8, error)
	SelectArrayParameterInt2Int8OriginalPointer                  func(id int) ([]int8, error)
	SelectArrayParameterInt2Int8PointerOriginal                  func(id int) ([]*int8, error)
	SelectArrayParameterInt2Int8PointerPointer                   func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8                                    func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8OriginalPointer                     func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8PointerOriginal                     func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8PointerPointer                      func(id int) ([]*int8, error)
	SelectParameterTextString                                    func(sid string) (string, error)
	SelectParameterTextStringOriginalPointer                     func(sid string) (*string, error)
	SelectParameterTextStringPointerOriginal                     func(sid string) (string, error)
	SelectParameterTextStringPointerPointer                      func(sid string) (*string, error)
	SelectEntityTextString                                       func(sid string) (*string, error)
	SelectEntityTextStringOriginalPointer                        func(sid string) (*string, error)
	SelectEntityTextStringPointerOriginal                        func(sid string) (*string, error)
	SelectEntityTextStringPointerPointer                         func(sid string) (*string, error)
	SelectArrayParameterTextString                               func(id int) ([]string, error)
	SelectArrayParameterTextStringOriginalPointer                func(id int) ([]string, error)
	SelectArrayParameterTextStringPointerOriginal                func(id int) ([]*string, error)
	SelectArrayParameterTextStringPointerPointer                 func(id int) ([]*string, error)
	SelectArrayEntityTextString                                  func(id int) ([]*string, error)
	SelectArrayEntityTextStringOriginalPointer                   func(id int) ([]*string, error)
	SelectArrayEntityTextStringPointerOriginal                   func(id int) ([]*string, error)
	SelectArrayEntityTextStringPointerPointer                    func(id int) ([]*string, error)
	SelectParameterTimeTime                                      func(sid string) (time, error)
	SelectParameterTimeTimeOriginalPointer                       func(sid string) (*time, error)
	SelectParameterTimeTimePointerOriginal                       func(sid string) (time, error)
	SelectParameterTimeTimePointerPointer                        func(sid string) (*time, error)
	SelectEntityTimeTime                                         func(sid string) (*time, error)
	SelectEntityTimeTimeOriginalPointer                          func(sid string) (*time, error)
	SelectEntityTimeTimePointerOriginal                          func(sid string) (*time, error)
	SelectEntityTimeTimePointerPointer                           func(sid string) (*time, error)
	SelectArrayParameterTimeTime                                 func(id int) ([]time, error)
	SelectArrayParameterTimeTimeOriginalPointer                  func(id int) ([]time, error)
	SelectArrayParameterTimeTimePointerOriginal                  func(id int) ([]*time, error)
	SelectArrayParameterTimeTimePointerPointer                   func(id int) ([]*time, error)
	SelectArrayEntityTimeTime                                    func(id int) ([]*time, error)
	SelectArrayEntityTimeTimeOriginalPointer                     func(id int) ([]*time, error)
	SelectArrayEntityTimeTimePointerOriginal                     func(id int) ([]*time, error)
	SelectArrayEntityTimeTimePointerPointer                      func(id int) ([]*time, error)
	SelectParameterTimeWithTimezoneTime                          func(sid string) (time, error)
	SelectParameterTimeWithTimezoneTimeOriginalPointer           func(sid string) (*time, error)
	SelectParameterTimeWithTimezoneTimePointerOriginal           func(sid string) (time, error)
	SelectParameterTimeWithTimezoneTimePointerPointer            func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTime                             func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTimeOriginalPointer              func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTimePointerOriginal              func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTimePointerPointer               func(sid string) (*time, error)
	SelectArrayParameterTimeWithTimezoneTime                     func(id int) ([]time, error)
	SelectArrayParameterTimeWithTimezoneTimeOriginalPointer      func(id int) ([]time, error)
	SelectArrayParameterTimeWithTimezoneTimePointerOriginal      func(id int) ([]*time, error)
	SelectArrayParameterTimeWithTimezoneTimePointerPointer       func(id int) ([]*time, error)
	SelectArrayEntityTimeWithTimezoneTime                        func(id int) ([]*time, error)
	SelectArrayEntityTimeWithTimezoneTimeOriginalPointer         func(id int) ([]*time, error)
	SelectArrayEntityTimeWithTimezoneTimePointerOriginal         func(id int) ([]*time, error)
	SelectArrayEntityTimeWithTimezoneTimePointerPointer          func(id int) ([]*time, error)
	SelectParameterTimetzTime                                    func(sid string) (time, error)
	SelectParameterTimetzTimeOriginalPointer                     func(sid string) (*time, error)
	SelectParameterTimetzTimePointerOriginal                     func(sid string) (time, error)
	SelectParameterTimetzTimePointerPointer                      func(sid string) (*time, error)
	SelectEntityTimetzTime                                       func(sid string) (*time, error)
	SelectEntityTimetzTimeOriginalPointer                        func(sid string) (*time, error)
	SelectEntityTimetzTimePointerOriginal                        func(sid string) (*time, error)
	SelectEntityTimetzTimePointerPointer                         func(sid string) (*time, error)
	SelectArrayParameterTimetzTime                               func(id int) ([]time, error)
	SelectArrayParameterTimetzTimeOriginalPointer                func(id int) ([]time, error)
	SelectArrayParameterTimetzTimePointerOriginal                func(id int) ([]*time, error)
	SelectArrayParameterTimetzTimePointerPointer                 func(id int) ([]*time, error)
	SelectArrayEntityTimetzTime                                  func(id int) ([]*time, error)
	SelectArrayEntityTimetzTimeOriginalPointer                   func(id int) ([]*time, error)
	SelectArrayEntityTimetzTimePointerOriginal                   func(id int) ([]*time, error)
	SelectArrayEntityTimetzTimePointerPointer                    func(id int) ([]*time, error)
	SelectParameterTimestampTime                                 func(sid string) (time, error)
	SelectParameterTimestampTimeOriginalPointer                  func(sid string) (*time, error)
	SelectParameterTimestampTimePointerOriginal                  func(sid string) (time, error)
	SelectParameterTimestampTimePointerPointer                   func(sid string) (*time, error)
	SelectEntityTimestampTime                                    func(sid string) (*time, error)
	SelectEntityTimestampTimeOriginalPointer                     func(sid string) (*time, error)
	SelectEntityTimestampTimePointerOriginal                     func(sid string) (*time, error)
	SelectEntityTimestampTimePointerPointer                      func(sid string) (*time, error)
	SelectArrayParameterTimestampTime                            func(id int) ([]time, error)
	SelectArrayParameterTimestampTimeOriginalPointer             func(id int) ([]time, error)
	SelectArrayParameterTimestampTimePointerOriginal             func(id int) ([]*time, error)
	SelectArrayParameterTimestampTimePointerPointer              func(id int) ([]*time, error)
	SelectArrayEntityTimestampTime                               func(id int) ([]*time, error)
	SelectArrayEntityTimestampTimeOriginalPointer                func(id int) ([]*time, error)
	SelectArrayEntityTimestampTimePointerOriginal                func(id int) ([]*time, error)
	SelectArrayEntityTimestampTimePointerPointer                 func(id int) ([]*time, error)
	SelectParameterTimestampWithTimezoneTime                     func(sid string) (time, error)
	SelectParameterTimestampWithTimezoneTimeOriginalPointer      func(sid string) (*time, error)
	SelectParameterTimestampWithTimezoneTimePointerOriginal      func(sid string) (time, error)
	SelectParameterTimestampWithTimezoneTimePointerPointer       func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTime                        func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTimeOriginalPointer         func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTimePointerOriginal         func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTimePointerPointer          func(sid string) (*time, error)
	SelectArrayParameterTimestampWithTimezoneTime                func(id int) ([]time, error)
	SelectArrayParameterTimestampWithTimezoneTimeOriginalPointer func(id int) ([]time, error)
	SelectArrayParameterTimestampWithTimezoneTimePointerOriginal func(id int) ([]*time, error)
	SelectArrayParameterTimestampWithTimezoneTimePointerPointer  func(id int) ([]*time, error)
	SelectArrayEntityTimestampWithTimezoneTime                   func(id int) ([]*time, error)
	SelectArrayEntityTimestampWithTimezoneTimeOriginalPointer    func(id int) ([]*time, error)
	SelectArrayEntityTimestampWithTimezoneTimePointerOriginal    func(id int) ([]*time, error)
	SelectArrayEntityTimestampWithTimezoneTimePointerPointer     func(id int) ([]*time, error)
	SelectParameterTimestamptzTime                               func(sid string) (time, error)
	SelectParameterTimestamptzTimeOriginalPointer                func(sid string) (*time, error)
	SelectParameterTimestamptzTimePointerOriginal                func(sid string) (time, error)
	SelectParameterTimestamptzTimePointerPointer                 func(sid string) (*time, error)
	SelectEntityTimestamptzTime                                  func(sid string) (*time, error)
	SelectEntityTimestamptzTimeOriginalPointer                   func(sid string) (*time, error)
	SelectEntityTimestamptzTimePointerOriginal                   func(sid string) (*time, error)
	SelectEntityTimestamptzTimePointerPointer                    func(sid string) (*time, error)
	SelectArrayParameterTimestamptzTime                          func(id int) ([]time, error)
	SelectArrayParameterTimestamptzTimeOriginalPointer           func(id int) ([]time, error)
	SelectArrayParameterTimestamptzTimePointerOriginal           func(id int) ([]*time, error)
	SelectArrayParameterTimestamptzTimePointerPointer            func(id int) ([]*time, error)
	SelectArrayEntityTimestamptzTime                             func(id int) ([]*time, error)
	SelectArrayEntityTimestamptzTimeOriginalPointer              func(id int) ([]*time, error)
	SelectArrayEntityTimestamptzTimePointerOriginal              func(id int) ([]*time, error)
	SelectArrayEntityTimestamptzTimePointerPointer               func(id int) ([]*time, error)
	UpdateParameterBigintInt64                                   func(sid string, source string, var_bigint int64) (int, error)
	UpdateParameterBigintInt64OriginalPointer                    func(sid string, source string, var_bigint int64) (*int, error)
	UpdateParameterBigintInt64PointerOriginal                    func(sid *string, source *string, var_bigint *int64) (int, error)
	UpdateParameterBigintInt64PointerPointer                     func(sid string, source string, var_bigint *int64) (*int, error)
	UpdateEntityBigintInt64                                      func(item BigintOriginal) error
	UpdateEntityBigintInt64OriginalPointer                       func(item BigintOriginal) error
	UpdateEntityBigintInt64PointerOriginal                       func(item *BigintPointer) error
	UpdateEntityBigintInt64PointerPointer                        func(item *BigintPointer) error
	UpdateArrayParameterBigintInt64                              func(sid string, source string, items []int64) error
	UpdateArrayParameterBigintInt64OriginalPointer               func(sid string, source string, items []int64) error
	UpdateArrayParameterBigintInt64PointerOriginal               func(sid string, source string, items []*int64) error
	UpdateArrayParameterBigintInt64PointerPointer                func(sid string, source string, items []*int64) error
	UpdateArrayEntityBigintInt64                                 func(item BigintOriginal) error
	UpdateArrayEntityBigintInt64OriginalPointer                  func(item BigintPointer) error
	UpdateArrayEntityBigintInt64PointerOriginal                  func(item BigintOriginal) error
	UpdateArrayEntityBigintInt64PointerPointer                   func(item BigintPointer) error
	UpdateParameterInt8Int8                                      func(sid string, source string, var_int8 int8) (int, error)
	UpdateParameterInt8Int8OriginalPointer                       func(sid string, source string, var_int8 int8) (*int, error)
	UpdateParameterInt8Int8PointerOriginal                       func(sid *string, source *string, var_int8 *int8) (int, error)
	UpdateParameterInt8Int8PointerPointer                        func(sid string, source string, var_int8 *int8) (*int, error)
	UpdateEntityInt8Int8                                         func(item Int8Original) error
	UpdateEntityInt8Int8OriginalPointer                          func(item Int8Original) error
	UpdateEntityInt8Int8PointerOriginal                          func(item *Int8Pointer) error
	UpdateEntityInt8Int8PointerPointer                           func(item *Int8Pointer) error
	UpdateArrayParameterInt8Int8                                 func(sid string, source string, items []int8) error
	UpdateArrayParameterInt8Int8OriginalPointer                  func(sid string, source string, items []int8) error
	UpdateArrayParameterInt8Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	UpdateArrayParameterInt8Int8PointerPointer                   func(sid string, source string, items []*int8) error
	UpdateArrayEntityInt8Int8                                    func(item Int8Original) error
	UpdateArrayEntityInt8Int8OriginalPointer                     func(item Int8Pointer) error
	UpdateArrayEntityInt8Int8PointerOriginal                     func(item Int8Original) error
	UpdateArrayEntityInt8Int8PointerPointer                      func(item Int8Pointer) error
	UpdateParameterBooleanBool                                   func(sid string, source string, var_boolean bool) (int, error)
	UpdateParameterBooleanBoolOriginalPointer                    func(sid string, source string, var_boolean bool) (*int, error)
	UpdateParameterBooleanBoolPointerOriginal                    func(sid *string, source *string, var_boolean *bool) (int, error)
	UpdateParameterBooleanBoolPointerPointer                     func(sid string, source string, var_boolean *bool) (*int, error)
	UpdateEntityBooleanBool                                      func(item BooleanOriginal) error
	UpdateEntityBooleanBoolOriginalPointer                       func(item BooleanOriginal) error
	UpdateEntityBooleanBoolPointerOriginal                       func(item *BooleanPointer) error
	UpdateEntityBooleanBoolPointerPointer                        func(item *BooleanPointer) error
	UpdateArrayParameterBooleanBool                              func(sid string, source string, items []bool) error
	UpdateArrayParameterBooleanBoolOriginalPointer               func(sid string, source string, items []bool) error
	UpdateArrayParameterBooleanBoolPointerOriginal               func(sid string, source string, items []*bool) error
	UpdateArrayParameterBooleanBoolPointerPointer                func(sid string, source string, items []*bool) error
	UpdateArrayEntityBooleanBool                                 func(item BooleanOriginal) error
	UpdateArrayEntityBooleanBoolOriginalPointer                  func(item BooleanPointer) error
	UpdateArrayEntityBooleanBoolPointerOriginal                  func(item BooleanOriginal) error
	UpdateArrayEntityBooleanBoolPointerPointer                   func(item BooleanPointer) error
	UpdateParameterBoolBool                                      func(sid string, source string, var_bool bool) (int, error)
	UpdateParameterBoolBoolOriginalPointer                       func(sid string, source string, var_bool bool) (*int, error)
	UpdateParameterBoolBoolPointerOriginal                       func(sid *string, source *string, var_bool *bool) (int, error)
	UpdateParameterBoolBoolPointerPointer                        func(sid string, source string, var_bool *bool) (*int, error)
	UpdateEntityBoolBool                                         func(item BoolOriginal) error
	UpdateEntityBoolBoolOriginalPointer                          func(item BoolOriginal) error
	UpdateEntityBoolBoolPointerOriginal                          func(item *BoolPointer) error
	UpdateEntityBoolBoolPointerPointer                           func(item *BoolPointer) error
	UpdateArrayParameterBoolBool                                 func(sid string, source string, items []bool) error
	UpdateArrayParameterBoolBoolOriginalPointer                  func(sid string, source string, items []bool) error
	UpdateArrayParameterBoolBoolPointerOriginal                  func(sid string, source string, items []*bool) error
	UpdateArrayParameterBoolBoolPointerPointer                   func(sid string, source string, items []*bool) error
	UpdateArrayEntityBoolBool                                    func(item BoolOriginal) error
	UpdateArrayEntityBoolBoolOriginalPointer                     func(item BoolPointer) error
	UpdateArrayEntityBoolBoolPointerOriginal                     func(item BoolOriginal) error
	UpdateArrayEntityBoolBoolPointerPointer                      func(item BoolPointer) error
	UpdateParameterCharacterString                               func(sid string, source string, var_character string) (int, error)
	UpdateParameterCharacterStringOriginalPointer                func(sid string, source string, var_character string) (*int, error)
	UpdateParameterCharacterStringPointerOriginal                func(sid *string, source *string, var_character *string) (int, error)
	UpdateParameterCharacterStringPointerPointer                 func(sid string, source string, var_character *string) (*int, error)
	UpdateEntityCharacterString                                  func(item CharacterOriginal) error
	UpdateEntityCharacterStringOriginalPointer                   func(item CharacterOriginal) error
	UpdateEntityCharacterStringPointerOriginal                   func(item *CharacterPointer) error
	UpdateEntityCharacterStringPointerPointer                    func(item *CharacterPointer) error
	UpdateArrayParameterCharacterString                          func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterStringOriginalPointer           func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterStringPointerOriginal           func(sid string, source string, items []*string) error
	UpdateArrayParameterCharacterStringPointerPointer            func(sid string, source string, items []*string) error
	UpdateArrayEntityCharacterString                             func(item CharacterOriginal) error
	UpdateArrayEntityCharacterStringOriginalPointer              func(item CharacterPointer) error
	UpdateArrayEntityCharacterStringPointerOriginal              func(item CharacterOriginal) error
	UpdateArrayEntityCharacterStringPointerPointer               func(item CharacterPointer) error
	UpdateParameterCharString                                    func(sid string, source string, var_char string) (int, error)
	UpdateParameterCharStringOriginalPointer                     func(sid string, source string, var_char string) (*int, error)
	UpdateParameterCharStringPointerOriginal                     func(sid *string, source *string, var_char *string) (int, error)
	UpdateParameterCharStringPointerPointer                      func(sid string, source string, var_char *string) (*int, error)
	UpdateEntityCharString                                       func(item CharOriginal) error
	UpdateEntityCharStringOriginalPointer                        func(item CharOriginal) error
	UpdateEntityCharStringPointerOriginal                        func(item *CharPointer) error
	UpdateEntityCharStringPointerPointer                         func(item *CharPointer) error
	UpdateArrayParameterCharString                               func(sid string, source string, items []string) error
	UpdateArrayParameterCharStringOriginalPointer                func(sid string, source string, items []string) error
	UpdateArrayParameterCharStringPointerOriginal                func(sid string, source string, items []*string) error
	UpdateArrayParameterCharStringPointerPointer                 func(sid string, source string, items []*string) error
	UpdateArrayEntityCharString                                  func(item CharOriginal) error
	UpdateArrayEntityCharStringOriginalPointer                   func(item CharPointer) error
	UpdateArrayEntityCharStringPointerOriginal                   func(item CharOriginal) error
	UpdateArrayEntityCharStringPointerPointer                    func(item CharPointer) error
	UpdateParameterCharacterVaryingString                        func(sid string, source string, var_character_varying string) (int, error)
	UpdateParameterCharacterVaryingStringOriginalPointer         func(sid string, source string, var_character_varying string) (*int, error)
	UpdateParameterCharacterVaryingStringPointerOriginal         func(sid *string, source *string, var_character_varying *string) (int, error)
	UpdateParameterCharacterVaryingStringPointerPointer          func(sid string, source string, var_character_varying *string) (*int, error)
	UpdateEntityCharacterVaryingString                           func(item CharacterVaryingOriginal) error
	UpdateEntityCharacterVaryingStringOriginalPointer            func(item CharacterVaryingOriginal) error
	UpdateEntityCharacterVaryingStringPointerOriginal            func(item *CharacterVaryingPointer) error
	UpdateEntityCharacterVaryingStringPointerPointer             func(item *CharacterVaryingPointer) error
	UpdateArrayParameterCharacterVaryingString                   func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterVaryingStringOriginalPointer    func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterVaryingStringPointerOriginal    func(sid string, source string, items []*string) error
	UpdateArrayParameterCharacterVaryingStringPointerPointer     func(sid string, source string, items []*string) error
	UpdateArrayEntityCharacterVaryingString                      func(item CharacterVaryingOriginal) error
	UpdateArrayEntityCharacterVaryingStringOriginalPointer       func(item CharacterVaryingPointer) error
	UpdateArrayEntityCharacterVaryingStringPointerOriginal       func(item CharacterVaryingOriginal) error
	UpdateArrayEntityCharacterVaryingStringPointerPointer        func(item CharacterVaryingPointer) error
	UpdateParameterVarcharString                                 func(sid string, source string, var_varchar string) (int, error)
	UpdateParameterVarcharStringOriginalPointer                  func(sid string, source string, var_varchar string) (*int, error)
	UpdateParameterVarcharStringPointerOriginal                  func(sid *string, source *string, var_varchar *string) (int, error)
	UpdateParameterVarcharStringPointerPointer                   func(sid string, source string, var_varchar *string) (*int, error)
	UpdateEntityVarcharString                                    func(item VarcharOriginal) error
	UpdateEntityVarcharStringOriginalPointer                     func(item VarcharOriginal) error
	UpdateEntityVarcharStringPointerOriginal                     func(item *VarcharPointer) error
	UpdateEntityVarcharStringPointerPointer                      func(item *VarcharPointer) error
	UpdateArrayParameterVarcharString                            func(sid string, source string, items []string) error
	UpdateArrayParameterVarcharStringOriginalPointer             func(sid string, source string, items []string) error
	UpdateArrayParameterVarcharStringPointerOriginal             func(sid string, source string, items []*string) error
	UpdateArrayParameterVarcharStringPointerPointer              func(sid string, source string, items []*string) error
	UpdateArrayEntityVarcharString                               func(item VarcharOriginal) error
	UpdateArrayEntityVarcharStringOriginalPointer                func(item VarcharPointer) error
	UpdateArrayEntityVarcharStringPointerOriginal                func(item VarcharOriginal) error
	UpdateArrayEntityVarcharStringPointerPointer                 func(item VarcharPointer) error
	UpdateParameterFloat8Float32                                 func(sid string, source string, var_float8 float32) (int, error)
	UpdateParameterFloat8Float32OriginalPointer                  func(sid string, source string, var_float8 float32) (*int, error)
	UpdateParameterFloat8Float32PointerOriginal                  func(sid *string, source *string, var_float8 *float32) (int, error)
	UpdateParameterFloat8Float32PointerPointer                   func(sid string, source string, var_float8 *float32) (*int, error)
	UpdateEntityFloat8Float32                                    func(item Float8Original) error
	UpdateEntityFloat8Float32OriginalPointer                     func(item Float8Original) error
	UpdateEntityFloat8Float32PointerOriginal                     func(item *Float8Pointer) error
	UpdateEntityFloat8Float32PointerPointer                      func(item *Float8Pointer) error
	UpdateArrayParameterFloat8Float32                            func(sid string, source string, items []float32) error
	UpdateArrayParameterFloat8Float32OriginalPointer             func(sid string, source string, items []float32) error
	UpdateArrayParameterFloat8Float32PointerOriginal             func(sid string, source string, items []*float32) error
	UpdateArrayParameterFloat8Float32PointerPointer              func(sid string, source string, items []*float32) error
	UpdateArrayEntityFloat8Float32                               func(item Float8Original) error
	UpdateArrayEntityFloat8Float32OriginalPointer                func(item Float8Pointer) error
	UpdateArrayEntityFloat8Float32PointerOriginal                func(item Float8Original) error
	UpdateArrayEntityFloat8Float32PointerPointer                 func(item Float8Pointer) error
	UpdateParameterIntegerInt8                                   func(sid string, source string, var_integer int8) (int, error)
	UpdateParameterIntegerInt8OriginalPointer                    func(sid string, source string, var_integer int8) (*int, error)
	UpdateParameterIntegerInt8PointerOriginal                    func(sid *string, source *string, var_integer *int8) (int, error)
	UpdateParameterIntegerInt8PointerPointer                     func(sid string, source string, var_integer *int8) (*int, error)
	UpdateEntityIntegerInt8                                      func(item IntegerOriginal) error
	UpdateEntityIntegerInt8OriginalPointer                       func(item IntegerOriginal) error
	UpdateEntityIntegerInt8PointerOriginal                       func(item *IntegerPointer) error
	UpdateEntityIntegerInt8PointerPointer                        func(item *IntegerPointer) error
	UpdateArrayParameterIntegerInt8                              func(sid string, source string, items []int8) error
	UpdateArrayParameterIntegerInt8OriginalPointer               func(sid string, source string, items []int8) error
	UpdateArrayParameterIntegerInt8PointerOriginal               func(sid string, source string, items []*int8) error
	UpdateArrayParameterIntegerInt8PointerPointer                func(sid string, source string, items []*int8) error
	UpdateArrayEntityIntegerInt8                                 func(item IntegerOriginal) error
	UpdateArrayEntityIntegerInt8OriginalPointer                  func(item IntegerPointer) error
	UpdateArrayEntityIntegerInt8PointerOriginal                  func(item IntegerOriginal) error
	UpdateArrayEntityIntegerInt8PointerPointer                   func(item IntegerPointer) error
	UpdateParameterIntInt8                                       func(sid string, source string, var_int int8) (int, error)
	UpdateParameterIntInt8OriginalPointer                        func(sid string, source string, var_int int8) (*int, error)
	UpdateParameterIntInt8PointerOriginal                        func(sid *string, source *string, var_int *int8) (int, error)
	UpdateParameterIntInt8PointerPointer                         func(sid string, source string, var_int *int8) (*int, error)
	UpdateEntityIntInt8                                          func(item IntOriginal) error
	UpdateEntityIntInt8OriginalPointer                           func(item IntOriginal) error
	UpdateEntityIntInt8PointerOriginal                           func(item *IntPointer) error
	UpdateEntityIntInt8PointerPointer                            func(item *IntPointer) error
	UpdateArrayParameterIntInt8                                  func(sid string, source string, items []int8) error
	UpdateArrayParameterIntInt8OriginalPointer                   func(sid string, source string, items []int8) error
	UpdateArrayParameterIntInt8PointerOriginal                   func(sid string, source string, items []*int8) error
	UpdateArrayParameterIntInt8PointerPointer                    func(sid string, source string, items []*int8) error
	UpdateArrayEntityIntInt8                                     func(item IntOriginal) error
	UpdateArrayEntityIntInt8OriginalPointer                      func(item IntPointer) error
	UpdateArrayEntityIntInt8PointerOriginal                      func(item IntOriginal) error
	UpdateArrayEntityIntInt8PointerPointer                       func(item IntPointer) error
	UpdateParameterInt4Int8                                      func(sid string, source string, var_int4 int8) (int, error)
	UpdateParameterInt4Int8OriginalPointer                       func(sid string, source string, var_int4 int8) (*int, error)
	UpdateParameterInt4Int8PointerOriginal                       func(sid *string, source *string, var_int4 *int8) (int, error)
	UpdateParameterInt4Int8PointerPointer                        func(sid string, source string, var_int4 *int8) (*int, error)
	UpdateEntityInt4Int8                                         func(item Int4Original) error
	UpdateEntityInt4Int8OriginalPointer                          func(item Int4Original) error
	UpdateEntityInt4Int8PointerOriginal                          func(item *Int4Pointer) error
	UpdateEntityInt4Int8PointerPointer                           func(item *Int4Pointer) error
	UpdateArrayParameterInt4Int8                                 func(sid string, source string, items []int8) error
	UpdateArrayParameterInt4Int8OriginalPointer                  func(sid string, source string, items []int8) error
	UpdateArrayParameterInt4Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	UpdateArrayParameterInt4Int8PointerPointer                   func(sid string, source string, items []*int8) error
	UpdateArrayEntityInt4Int8                                    func(item Int4Original) error
	UpdateArrayEntityInt4Int8OriginalPointer                     func(item Int4Pointer) error
	UpdateArrayEntityInt4Int8PointerOriginal                     func(item Int4Original) error
	UpdateArrayEntityInt4Int8PointerPointer                      func(item Int4Pointer) error
	UpdateParameterNumericDecimal                                func(sid string, source string, var_numeric decimal) (int, error)
	UpdateParameterNumericDecimalOriginalPointer                 func(sid string, source string, var_numeric decimal) (*int, error)
	UpdateParameterNumericDecimalPointerOriginal                 func(sid *string, source *string, var_numeric *decimal) (int, error)
	UpdateParameterNumericDecimalPointerPointer                  func(sid string, source string, var_numeric *decimal) (*int, error)
	UpdateEntityNumericDecimal                                   func(item NumericOriginal) error
	UpdateEntityNumericDecimalOriginalPointer                    func(item NumericOriginal) error
	UpdateEntityNumericDecimalPointerOriginal                    func(item *NumericPointer) error
	UpdateEntityNumericDecimalPointerPointer                     func(item *NumericPointer) error
	UpdateArrayParameterNumericDecimal                           func(sid string, source string, items []decimal) error
	UpdateArrayParameterNumericDecimalOriginalPointer            func(sid string, source string, items []decimal) error
	UpdateArrayParameterNumericDecimalPointerOriginal            func(sid string, source string, items []*decimal) error
	UpdateArrayParameterNumericDecimalPointerPointer             func(sid string, source string, items []*decimal) error
	UpdateArrayEntityNumericDecimal                              func(item NumericOriginal) error
	UpdateArrayEntityNumericDecimalOriginalPointer               func(item NumericPointer) error
	UpdateArrayEntityNumericDecimalPointerOriginal               func(item NumericOriginal) error
	UpdateArrayEntityNumericDecimalPointerPointer                func(item NumericPointer) error
	UpdateParameterDecimalDecimal                                func(sid string, source string, var_decimal decimal) (int, error)
	UpdateParameterDecimalDecimalOriginalPointer                 func(sid string, source string, var_decimal decimal) (*int, error)
	UpdateParameterDecimalDecimalPointerOriginal                 func(sid *string, source *string, var_decimal *decimal) (int, error)
	UpdateParameterDecimalDecimalPointerPointer                  func(sid string, source string, var_decimal *decimal) (*int, error)
	UpdateEntityDecimalDecimal                                   func(item DecimalOriginal) error
	UpdateEntityDecimalDecimalOriginalPointer                    func(item DecimalOriginal) error
	UpdateEntityDecimalDecimalPointerOriginal                    func(item *DecimalPointer) error
	UpdateEntityDecimalDecimalPointerPointer                     func(item *DecimalPointer) error
	UpdateArrayParameterDecimalDecimal                           func(sid string, source string, items []decimal) error
	UpdateArrayParameterDecimalDecimalOriginalPointer            func(sid string, source string, items []decimal) error
	UpdateArrayParameterDecimalDecimalPointerOriginal            func(sid string, source string, items []*decimal) error
	UpdateArrayParameterDecimalDecimalPointerPointer             func(sid string, source string, items []*decimal) error
	UpdateArrayEntityDecimalDecimal                              func(item DecimalOriginal) error
	UpdateArrayEntityDecimalDecimalOriginalPointer               func(item DecimalPointer) error
	UpdateArrayEntityDecimalDecimalPointerOriginal               func(item DecimalOriginal) error
	UpdateArrayEntityDecimalDecimalPointerPointer                func(item DecimalPointer) error
	UpdateParameterFloat4Float32                                 func(sid string, source string, var_float4 float32) (int, error)
	UpdateParameterFloat4Float32OriginalPointer                  func(sid string, source string, var_float4 float32) (*int, error)
	UpdateParameterFloat4Float32PointerOriginal                  func(sid *string, source *string, var_float4 *float32) (int, error)
	UpdateParameterFloat4Float32PointerPointer                   func(sid string, source string, var_float4 *float32) (*int, error)
	UpdateEntityFloat4Float32                                    func(item Float4Original) error
	UpdateEntityFloat4Float32OriginalPointer                     func(item Float4Original) error
	UpdateEntityFloat4Float32PointerOriginal                     func(item *Float4Pointer) error
	UpdateEntityFloat4Float32PointerPointer                      func(item *Float4Pointer) error
	UpdateArrayParameterFloat4Float32                            func(sid string, source string, items []float32) error
	UpdateArrayParameterFloat4Float32OriginalPointer             func(sid string, source string, items []float32) error
	UpdateArrayParameterFloat4Float32PointerOriginal             func(sid string, source string, items []*float32) error
	UpdateArrayParameterFloat4Float32PointerPointer              func(sid string, source string, items []*float32) error
	UpdateArrayEntityFloat4Float32                               func(item Float4Original) error
	UpdateArrayEntityFloat4Float32OriginalPointer                func(item Float4Pointer) error
	UpdateArrayEntityFloat4Float32PointerOriginal                func(item Float4Original) error
	UpdateArrayEntityFloat4Float32PointerPointer                 func(item Float4Pointer) error
	UpdateParameterSmallintInt8                                  func(sid string, source string, var_smallint int8) (int, error)
	UpdateParameterSmallintInt8OriginalPointer                   func(sid string, source string, var_smallint int8) (*int, error)
	UpdateParameterSmallintInt8PointerOriginal                   func(sid *string, source *string, var_smallint *int8) (int, error)
	UpdateParameterSmallintInt8PointerPointer                    func(sid string, source string, var_smallint *int8) (*int, error)
	UpdateEntitySmallintInt8                                     func(item SmallintOriginal) error
	UpdateEntitySmallintInt8OriginalPointer                      func(item SmallintOriginal) error
	UpdateEntitySmallintInt8PointerOriginal                      func(item *SmallintPointer) error
	UpdateEntitySmallintInt8PointerPointer                       func(item *SmallintPointer) error
	UpdateArrayParameterSmallintInt8                             func(sid string, source string, items []int8) error
	UpdateArrayParameterSmallintInt8OriginalPointer              func(sid string, source string, items []int8) error
	UpdateArrayParameterSmallintInt8PointerOriginal              func(sid string, source string, items []*int8) error
	UpdateArrayParameterSmallintInt8PointerPointer               func(sid string, source string, items []*int8) error
	UpdateArrayEntitySmallintInt8                                func(item SmallintOriginal) error
	UpdateArrayEntitySmallintInt8OriginalPointer                 func(item SmallintPointer) error
	UpdateArrayEntitySmallintInt8PointerOriginal                 func(item SmallintOriginal) error
	UpdateArrayEntitySmallintInt8PointerPointer                  func(item SmallintPointer) error
	UpdateParameterInt2Int8                                      func(sid string, source string, var_int2 int8) (int, error)
	UpdateParameterInt2Int8OriginalPointer                       func(sid string, source string, var_int2 int8) (*int, error)
	UpdateParameterInt2Int8PointerOriginal                       func(sid *string, source *string, var_int2 *int8) (int, error)
	UpdateParameterInt2Int8PointerPointer                        func(sid string, source string, var_int2 *int8) (*int, error)
	UpdateEntityInt2Int8                                         func(item Int2Original) error
	UpdateEntityInt2Int8OriginalPointer                          func(item Int2Original) error
	UpdateEntityInt2Int8PointerOriginal                          func(item *Int2Pointer) error
	UpdateEntityInt2Int8PointerPointer                           func(item *Int2Pointer) error
	UpdateArrayParameterInt2Int8                                 func(sid string, source string, items []int8) error
	UpdateArrayParameterInt2Int8OriginalPointer                  func(sid string, source string, items []int8) error
	UpdateArrayParameterInt2Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	UpdateArrayParameterInt2Int8PointerPointer                   func(sid string, source string, items []*int8) error
	UpdateArrayEntityInt2Int8                                    func(item Int2Original) error
	UpdateArrayEntityInt2Int8OriginalPointer                     func(item Int2Pointer) error
	UpdateArrayEntityInt2Int8PointerOriginal                     func(item Int2Original) error
	UpdateArrayEntityInt2Int8PointerPointer                      func(item Int2Pointer) error
	UpdateParameterTextString                                    func(sid string, source string, var_text string) (int, error)
	UpdateParameterTextStringOriginalPointer                     func(sid string, source string, var_text string) (*int, error)
	UpdateParameterTextStringPointerOriginal                     func(sid *string, source *string, var_text *string) (int, error)
	UpdateParameterTextStringPointerPointer                      func(sid string, source string, var_text *string) (*int, error)
	UpdateEntityTextString                                       func(item TextOriginal) error
	UpdateEntityTextStringOriginalPointer                        func(item TextOriginal) error
	UpdateEntityTextStringPointerOriginal                        func(item *TextPointer) error
	UpdateEntityTextStringPointerPointer                         func(item *TextPointer) error
	UpdateArrayParameterTextString                               func(sid string, source string, items []string) error
	UpdateArrayParameterTextStringOriginalPointer                func(sid string, source string, items []string) error
	UpdateArrayParameterTextStringPointerOriginal                func(sid string, source string, items []*string) error
	UpdateArrayParameterTextStringPointerPointer                 func(sid string, source string, items []*string) error
	UpdateArrayEntityTextString                                  func(item TextOriginal) error
	UpdateArrayEntityTextStringOriginalPointer                   func(item TextPointer) error
	UpdateArrayEntityTextStringPointerOriginal                   func(item TextOriginal) error
	UpdateArrayEntityTextStringPointerPointer                    func(item TextPointer) error
	UpdateParameterTimeTime                                      func(sid string, source string, var_time time) (int, error)
	UpdateParameterTimeTimeOriginalPointer                       func(sid string, source string, var_time time) (*int, error)
	UpdateParameterTimeTimePointerOriginal                       func(sid *string, source *string, var_time *time) (int, error)
	UpdateParameterTimeTimePointerPointer                        func(sid string, source string, var_time *time) (*int, error)
	UpdateEntityTimeTime                                         func(item TimeOriginal) error
	UpdateEntityTimeTimeOriginalPointer                          func(item TimeOriginal) error
	UpdateEntityTimeTimePointerOriginal                          func(item *TimePointer) error
	UpdateEntityTimeTimePointerPointer                           func(item *TimePointer) error
	UpdateArrayParameterTimeTime                                 func(sid string, source string, items []time) error
	UpdateArrayParameterTimeTimeOriginalPointer                  func(sid string, source string, items []time) error
	UpdateArrayParameterTimeTimePointerOriginal                  func(sid string, source string, items []*time) error
	UpdateArrayParameterTimeTimePointerPointer                   func(sid string, source string, items []*time) error
	UpdateArrayEntityTimeTime                                    func(item TimeOriginal) error
	UpdateArrayEntityTimeTimeOriginalPointer                     func(item TimePointer) error
	UpdateArrayEntityTimeTimePointerOriginal                     func(item TimeOriginal) error
	UpdateArrayEntityTimeTimePointerPointer                      func(item TimePointer) error
	UpdateParameterTimeWithTimezoneTime                          func(sid string, source string, var_time_with_timezone time) (int, error)
	UpdateParameterTimeWithTimezoneTimeOriginalPointer           func(sid string, source string, var_time_with_timezone time) (*int, error)
	UpdateParameterTimeWithTimezoneTimePointerOriginal           func(sid *string, source *string, var_time_with_timezone *time) (int, error)
	UpdateParameterTimeWithTimezoneTimePointerPointer            func(sid string, source string, var_time_with_timezone *time) (*int, error)
	UpdateEntityTimeWithTimezoneTime                             func(item TimeWithTimezoneOriginal) error
	UpdateEntityTimeWithTimezoneTimeOriginalPointer              func(item TimeWithTimezoneOriginal) error
	UpdateEntityTimeWithTimezoneTimePointerOriginal              func(item *TimeWithTimezonePointer) error
	UpdateEntityTimeWithTimezoneTimePointerPointer               func(item *TimeWithTimezonePointer) error
	UpdateArrayParameterTimeWithTimezoneTime                     func(sid string, source string, items []time) error
	UpdateArrayParameterTimeWithTimezoneTimeOriginalPointer      func(sid string, source string, items []time) error
	UpdateArrayParameterTimeWithTimezoneTimePointerOriginal      func(sid string, source string, items []*time) error
	UpdateArrayParameterTimeWithTimezoneTimePointerPointer       func(sid string, source string, items []*time) error
	UpdateArrayEntityTimeWithTimezoneTime                        func(item TimeWithTimezoneOriginal) error
	UpdateArrayEntityTimeWithTimezoneTimeOriginalPointer         func(item TimeWithTimezonePointer) error
	UpdateArrayEntityTimeWithTimezoneTimePointerOriginal         func(item TimeWithTimezoneOriginal) error
	UpdateArrayEntityTimeWithTimezoneTimePointerPointer          func(item TimeWithTimezonePointer) error
	UpdateParameterTimetzTime                                    func(sid string, source string, var_timetz time) (int, error)
	UpdateParameterTimetzTimeOriginalPointer                     func(sid string, source string, var_timetz time) (*int, error)
	UpdateParameterTimetzTimePointerOriginal                     func(sid *string, source *string, var_timetz *time) (int, error)
	UpdateParameterTimetzTimePointerPointer                      func(sid string, source string, var_timetz *time) (*int, error)
	UpdateEntityTimetzTime                                       func(item TimetzOriginal) error
	UpdateEntityTimetzTimeOriginalPointer                        func(item TimetzOriginal) error
	UpdateEntityTimetzTimePointerOriginal                        func(item *TimetzPointer) error
	UpdateEntityTimetzTimePointerPointer                         func(item *TimetzPointer) error
	UpdateArrayParameterTimetzTime                               func(sid string, source string, items []time) error
	UpdateArrayParameterTimetzTimeOriginalPointer                func(sid string, source string, items []time) error
	UpdateArrayParameterTimetzTimePointerOriginal                func(sid string, source string, items []*time) error
	UpdateArrayParameterTimetzTimePointerPointer                 func(sid string, source string, items []*time) error
	UpdateArrayEntityTimetzTime                                  func(item TimetzOriginal) error
	UpdateArrayEntityTimetzTimeOriginalPointer                   func(item TimetzPointer) error
	UpdateArrayEntityTimetzTimePointerOriginal                   func(item TimetzOriginal) error
	UpdateArrayEntityTimetzTimePointerPointer                    func(item TimetzPointer) error
	UpdateParameterTimestampTime                                 func(sid string, source string, var_timestamp time) (int, error)
	UpdateParameterTimestampTimeOriginalPointer                  func(sid string, source string, var_timestamp time) (*int, error)
	UpdateParameterTimestampTimePointerOriginal                  func(sid *string, source *string, var_timestamp *time) (int, error)
	UpdateParameterTimestampTimePointerPointer                   func(sid string, source string, var_timestamp *time) (*int, error)
	UpdateEntityTimestampTime                                    func(item TimestampOriginal) error
	UpdateEntityTimestampTimeOriginalPointer                     func(item TimestampOriginal) error
	UpdateEntityTimestampTimePointerOriginal                     func(item *TimestampPointer) error
	UpdateEntityTimestampTimePointerPointer                      func(item *TimestampPointer) error
	UpdateArrayParameterTimestampTime                            func(sid string, source string, items []time) error
	UpdateArrayParameterTimestampTimeOriginalPointer             func(sid string, source string, items []time) error
	UpdateArrayParameterTimestampTimePointerOriginal             func(sid string, source string, items []*time) error
	UpdateArrayParameterTimestampTimePointerPointer              func(sid string, source string, items []*time) error
	UpdateArrayEntityTimestampTime                               func(item TimestampOriginal) error
	UpdateArrayEntityTimestampTimeOriginalPointer                func(item TimestampPointer) error
	UpdateArrayEntityTimestampTimePointerOriginal                func(item TimestampOriginal) error
	UpdateArrayEntityTimestampTimePointerPointer                 func(item TimestampPointer) error
	UpdateParameterTimestampWithTimezoneTime                     func(sid string, source string, var_timestamp_with_timezone time) (int, error)
	UpdateParameterTimestampWithTimezoneTimeOriginalPointer      func(sid string, source string, var_timestamp_with_timezone time) (*int, error)
	UpdateParameterTimestampWithTimezoneTimePointerOriginal      func(sid *string, source *string, var_timestamp_with_timezone *time) (int, error)
	UpdateParameterTimestampWithTimezoneTimePointerPointer       func(sid string, source string, var_timestamp_with_timezone *time) (*int, error)
	UpdateEntityTimestampWithTimezoneTime                        func(item TimestampWithTimezoneOriginal) error
	UpdateEntityTimestampWithTimezoneTimeOriginalPointer         func(item TimestampWithTimezoneOriginal) error
	UpdateEntityTimestampWithTimezoneTimePointerOriginal         func(item *TimestampWithTimezonePointer) error
	UpdateEntityTimestampWithTimezoneTimePointerPointer          func(item *TimestampWithTimezonePointer) error
	UpdateArrayParameterTimestampWithTimezoneTime                func(sid string, source string, items []time) error
	UpdateArrayParameterTimestampWithTimezoneTimeOriginalPointer func(sid string, source string, items []time) error
	UpdateArrayParameterTimestampWithTimezoneTimePointerOriginal func(sid string, source string, items []*time) error
	UpdateArrayParameterTimestampWithTimezoneTimePointerPointer  func(sid string, source string, items []*time) error
	UpdateArrayEntityTimestampWithTimezoneTime                   func(item TimestampWithTimezoneOriginal) error
	UpdateArrayEntityTimestampWithTimezoneTimeOriginalPointer    func(item TimestampWithTimezonePointer) error
	UpdateArrayEntityTimestampWithTimezoneTimePointerOriginal    func(item TimestampWithTimezoneOriginal) error
	UpdateArrayEntityTimestampWithTimezoneTimePointerPointer     func(item TimestampWithTimezonePointer) error
	UpdateParameterTimestamptzTime                               func(sid string, source string, var_timestamptz time) (int, error)
	UpdateParameterTimestamptzTimeOriginalPointer                func(sid string, source string, var_timestamptz time) (*int, error)
	UpdateParameterTimestamptzTimePointerOriginal                func(sid *string, source *string, var_timestamptz *time) (int, error)
	UpdateParameterTimestamptzTimePointerPointer                 func(sid string, source string, var_timestamptz *time) (*int, error)
	UpdateEntityTimestamptzTime                                  func(item TimestamptzOriginal) error
	UpdateEntityTimestamptzTimeOriginalPointer                   func(item TimestamptzOriginal) error
	UpdateEntityTimestamptzTimePointerOriginal                   func(item *TimestamptzPointer) error
	UpdateEntityTimestamptzTimePointerPointer                    func(item *TimestamptzPointer) error
	UpdateArrayParameterTimestamptzTime                          func(sid string, source string, items []time) error
	UpdateArrayParameterTimestamptzTimeOriginalPointer           func(sid string, source string, items []time) error
	UpdateArrayParameterTimestamptzTimePointerOriginal           func(sid string, source string, items []*time) error
	UpdateArrayParameterTimestamptzTimePointerPointer            func(sid string, source string, items []*time) error
	UpdateArrayEntityTimestamptzTime                             func(item TimestamptzOriginal) error
	UpdateArrayEntityTimestamptzTimeOriginalPointer              func(item TimestamptzPointer) error
	UpdateArrayEntityTimestamptzTimePointerOriginal              func(item TimestamptzOriginal) error
	UpdateArrayEntityTimestamptzTimePointerPointer               func(item TimestamptzPointer) error
	DeleteParameterBigintInt64                                   func(sid string) (int, error)
	DeleteParameterBigintInt64OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterBigintInt64PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterBigintInt64PointerPointer                     func(sid string) (*int, error)
	DeleteEntityBigintInt64                                      func(item BigintOriginal) error
	DeleteEntityBigintInt64OriginalPointer                       func(item BigintOriginal) error
	DeleteEntityBigintInt64PointerOriginal                       func(item *BigintPointer) error
	DeleteEntityBigintInt64PointerPointer                        func(item *BigintPointer) error
	DeleteArrayParameterBigintInt64                              func(id int) (int, error)
	DeleteArrayParameterBigintInt64OriginalPointer               func(id int) (int, error)
	DeleteArrayParameterBigintInt64PointerOriginal               func(id int) (int, error)
	DeleteArrayParameterBigintInt64PointerPointer                func(id int) (int, error)
	DeleteArrayEntityBigintInt64                                 func(id int) (int, error)
	DeleteArrayEntityBigintInt64OriginalPointer                  func(id int) (int, error)
	DeleteArrayEntityBigintInt64PointerOriginal                  func(id int) (int, error)
	DeleteArrayEntityBigintInt64PointerPointer                   func(id int) (int, error)
	DeleteParameterInt8Int8                                      func(sid string) (int, error)
	DeleteParameterInt8Int8OriginalPointer                       func(sid string) (*int, error)
	DeleteParameterInt8Int8PointerOriginal                       func(sid *string) (int, error)
	DeleteParameterInt8Int8PointerPointer                        func(sid string) (*int, error)
	DeleteEntityInt8Int8                                         func(item Int8Original) error
	DeleteEntityInt8Int8OriginalPointer                          func(item Int8Original) error
	DeleteEntityInt8Int8PointerOriginal                          func(item *Int8Pointer) error
	DeleteEntityInt8Int8PointerPointer                           func(item *Int8Pointer) error
	DeleteArrayParameterInt8Int8                                 func(id int) (int, error)
	DeleteArrayParameterInt8Int8OriginalPointer                  func(id int) (int, error)
	DeleteArrayParameterInt8Int8PointerOriginal                  func(id int) (int, error)
	DeleteArrayParameterInt8Int8PointerPointer                   func(id int) (int, error)
	DeleteArrayEntityInt8Int8                                    func(id int) (int, error)
	DeleteArrayEntityInt8Int8OriginalPointer                     func(id int) (int, error)
	DeleteArrayEntityInt8Int8PointerOriginal                     func(id int) (int, error)
	DeleteArrayEntityInt8Int8PointerPointer                      func(id int) (int, error)
	DeleteParameterBooleanBool                                   func(sid string) (int, error)
	DeleteParameterBooleanBoolOriginalPointer                    func(sid string) (*int, error)
	DeleteParameterBooleanBoolPointerOriginal                    func(sid *string) (int, error)
	DeleteParameterBooleanBoolPointerPointer                     func(sid string) (*int, error)
	DeleteEntityBooleanBool                                      func(item BooleanOriginal) error
	DeleteEntityBooleanBoolOriginalPointer                       func(item BooleanOriginal) error
	DeleteEntityBooleanBoolPointerOriginal                       func(item *BooleanPointer) error
	DeleteEntityBooleanBoolPointerPointer                        func(item *BooleanPointer) error
	DeleteArrayParameterBooleanBool                              func(id int) (int, error)
	DeleteArrayParameterBooleanBoolOriginalPointer               func(id int) (int, error)
	DeleteArrayParameterBooleanBoolPointerOriginal               func(id int) (int, error)
	DeleteArrayParameterBooleanBoolPointerPointer                func(id int) (int, error)
	DeleteArrayEntityBooleanBool                                 func(id int) (int, error)
	DeleteArrayEntityBooleanBoolOriginalPointer                  func(id int) (int, error)
	DeleteArrayEntityBooleanBoolPointerOriginal                  func(id int) (int, error)
	DeleteArrayEntityBooleanBoolPointerPointer                   func(id int) (int, error)
	DeleteParameterBoolBool                                      func(sid string) (int, error)
	DeleteParameterBoolBoolOriginalPointer                       func(sid string) (*int, error)
	DeleteParameterBoolBoolPointerOriginal                       func(sid *string) (int, error)
	DeleteParameterBoolBoolPointerPointer                        func(sid string) (*int, error)
	DeleteEntityBoolBool                                         func(item BoolOriginal) error
	DeleteEntityBoolBoolOriginalPointer                          func(item BoolOriginal) error
	DeleteEntityBoolBoolPointerOriginal                          func(item *BoolPointer) error
	DeleteEntityBoolBoolPointerPointer                           func(item *BoolPointer) error
	DeleteArrayParameterBoolBool                                 func(id int) (int, error)
	DeleteArrayParameterBoolBoolOriginalPointer                  func(id int) (int, error)
	DeleteArrayParameterBoolBoolPointerOriginal                  func(id int) (int, error)
	DeleteArrayParameterBoolBoolPointerPointer                   func(id int) (int, error)
	DeleteArrayEntityBoolBool                                    func(id int) (int, error)
	DeleteArrayEntityBoolBoolOriginalPointer                     func(id int) (int, error)
	DeleteArrayEntityBoolBoolPointerOriginal                     func(id int) (int, error)
	DeleteArrayEntityBoolBoolPointerPointer                      func(id int) (int, error)
	DeleteParameterCharacterString                               func(sid string) (int, error)
	DeleteParameterCharacterStringOriginalPointer                func(sid string) (*int, error)
	DeleteParameterCharacterStringPointerOriginal                func(sid *string) (int, error)
	DeleteParameterCharacterStringPointerPointer                 func(sid string) (*int, error)
	DeleteEntityCharacterString                                  func(item CharacterOriginal) error
	DeleteEntityCharacterStringOriginalPointer                   func(item CharacterOriginal) error
	DeleteEntityCharacterStringPointerOriginal                   func(item *CharacterPointer) error
	DeleteEntityCharacterStringPointerPointer                    func(item *CharacterPointer) error
	DeleteArrayParameterCharacterString                          func(id int) (int, error)
	DeleteArrayParameterCharacterStringOriginalPointer           func(id int) (int, error)
	DeleteArrayParameterCharacterStringPointerOriginal           func(id int) (int, error)
	DeleteArrayParameterCharacterStringPointerPointer            func(id int) (int, error)
	DeleteArrayEntityCharacterString                             func(id int) (int, error)
	DeleteArrayEntityCharacterStringOriginalPointer              func(id int) (int, error)
	DeleteArrayEntityCharacterStringPointerOriginal              func(id int) (int, error)
	DeleteArrayEntityCharacterStringPointerPointer               func(id int) (int, error)
	DeleteParameterCharString                                    func(sid string) (int, error)
	DeleteParameterCharStringOriginalPointer                     func(sid string) (*int, error)
	DeleteParameterCharStringPointerOriginal                     func(sid *string) (int, error)
	DeleteParameterCharStringPointerPointer                      func(sid string) (*int, error)
	DeleteEntityCharString                                       func(item CharOriginal) error
	DeleteEntityCharStringOriginalPointer                        func(item CharOriginal) error
	DeleteEntityCharStringPointerOriginal                        func(item *CharPointer) error
	DeleteEntityCharStringPointerPointer                         func(item *CharPointer) error
	DeleteArrayParameterCharString                               func(id int) (int, error)
	DeleteArrayParameterCharStringOriginalPointer                func(id int) (int, error)
	DeleteArrayParameterCharStringPointerOriginal                func(id int) (int, error)
	DeleteArrayParameterCharStringPointerPointer                 func(id int) (int, error)
	DeleteArrayEntityCharString                                  func(id int) (int, error)
	DeleteArrayEntityCharStringOriginalPointer                   func(id int) (int, error)
	DeleteArrayEntityCharStringPointerOriginal                   func(id int) (int, error)
	DeleteArrayEntityCharStringPointerPointer                    func(id int) (int, error)
	DeleteParameterCharacterVaryingString                        func(sid string) (int, error)
	DeleteParameterCharacterVaryingStringOriginalPointer         func(sid string) (*int, error)
	DeleteParameterCharacterVaryingStringPointerOriginal         func(sid *string) (int, error)
	DeleteParameterCharacterVaryingStringPointerPointer          func(sid string) (*int, error)
	DeleteEntityCharacterVaryingString                           func(item CharacterVaryingOriginal) error
	DeleteEntityCharacterVaryingStringOriginalPointer            func(item CharacterVaryingOriginal) error
	DeleteEntityCharacterVaryingStringPointerOriginal            func(item *CharacterVaryingPointer) error
	DeleteEntityCharacterVaryingStringPointerPointer             func(item *CharacterVaryingPointer) error
	DeleteArrayParameterCharacterVaryingString                   func(id int) (int, error)
	DeleteArrayParameterCharacterVaryingStringOriginalPointer    func(id int) (int, error)
	DeleteArrayParameterCharacterVaryingStringPointerOriginal    func(id int) (int, error)
	DeleteArrayParameterCharacterVaryingStringPointerPointer     func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingString                      func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingStringOriginalPointer       func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingStringPointerOriginal       func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingStringPointerPointer        func(id int) (int, error)
	DeleteParameterVarcharString                                 func(sid string) (int, error)
	DeleteParameterVarcharStringOriginalPointer                  func(sid string) (*int, error)
	DeleteParameterVarcharStringPointerOriginal                  func(sid *string) (int, error)
	DeleteParameterVarcharStringPointerPointer                   func(sid string) (*int, error)
	DeleteEntityVarcharString                                    func(item VarcharOriginal) error
	DeleteEntityVarcharStringOriginalPointer                     func(item VarcharOriginal) error
	DeleteEntityVarcharStringPointerOriginal                     func(item *VarcharPointer) error
	DeleteEntityVarcharStringPointerPointer                      func(item *VarcharPointer) error
	DeleteArrayParameterVarcharString                            func(id int) (int, error)
	DeleteArrayParameterVarcharStringOriginalPointer             func(id int) (int, error)
	DeleteArrayParameterVarcharStringPointerOriginal             func(id int) (int, error)
	DeleteArrayParameterVarcharStringPointerPointer              func(id int) (int, error)
	DeleteArrayEntityVarcharString                               func(id int) (int, error)
	DeleteArrayEntityVarcharStringOriginalPointer                func(id int) (int, error)
	DeleteArrayEntityVarcharStringPointerOriginal                func(id int) (int, error)
	DeleteArrayEntityVarcharStringPointerPointer                 func(id int) (int, error)
	DeleteParameterFloat8Float32                                 func(sid string) (int, error)
	DeleteParameterFloat8Float32OriginalPointer                  func(sid string) (*int, error)
	DeleteParameterFloat8Float32PointerOriginal                  func(sid *string) (int, error)
	DeleteParameterFloat8Float32PointerPointer                   func(sid string) (*int, error)
	DeleteEntityFloat8Float32                                    func(item Float8Original) error
	DeleteEntityFloat8Float32OriginalPointer                     func(item Float8Original) error
	DeleteEntityFloat8Float32PointerOriginal                     func(item *Float8Pointer) error
	DeleteEntityFloat8Float32PointerPointer                      func(item *Float8Pointer) error
	DeleteArrayParameterFloat8Float32                            func(id int) (int, error)
	DeleteArrayParameterFloat8Float32OriginalPointer             func(id int) (int, error)
	DeleteArrayParameterFloat8Float32PointerOriginal             func(id int) (int, error)
	DeleteArrayParameterFloat8Float32PointerPointer              func(id int) (int, error)
	DeleteArrayEntityFloat8Float32                               func(id int) (int, error)
	DeleteArrayEntityFloat8Float32OriginalPointer                func(id int) (int, error)
	DeleteArrayEntityFloat8Float32PointerOriginal                func(id int) (int, error)
	DeleteArrayEntityFloat8Float32PointerPointer                 func(id int) (int, error)
	DeleteParameterIntegerInt8                                   func(sid string) (int, error)
	DeleteParameterIntegerInt8OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterIntegerInt8PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterIntegerInt8PointerPointer                     func(sid string) (*int, error)
	DeleteEntityIntegerInt8                                      func(item IntegerOriginal) error
	DeleteEntityIntegerInt8OriginalPointer                       func(item IntegerOriginal) error
	DeleteEntityIntegerInt8PointerOriginal                       func(item *IntegerPointer) error
	DeleteEntityIntegerInt8PointerPointer                        func(item *IntegerPointer) error
	DeleteArrayParameterIntegerInt8                              func(id int) (int, error)
	DeleteArrayParameterIntegerInt8OriginalPointer               func(id int) (int, error)
	DeleteArrayParameterIntegerInt8PointerOriginal               func(id int) (int, error)
	DeleteArrayParameterIntegerInt8PointerPointer                func(id int) (int, error)
	DeleteArrayEntityIntegerInt8                                 func(id int) (int, error)
	DeleteArrayEntityIntegerInt8OriginalPointer                  func(id int) (int, error)
	DeleteArrayEntityIntegerInt8PointerOriginal                  func(id int) (int, error)
	DeleteArrayEntityIntegerInt8PointerPointer                   func(id int) (int, error)
	DeleteParameterIntInt8                                       func(sid string) (int, error)
	DeleteParameterIntInt8OriginalPointer                        func(sid string) (*int, error)
	DeleteParameterIntInt8PointerOriginal                        func(sid *string) (int, error)
	DeleteParameterIntInt8PointerPointer                         func(sid string) (*int, error)
	DeleteEntityIntInt8                                          func(item IntOriginal) error
	DeleteEntityIntInt8OriginalPointer                           func(item IntOriginal) error
	DeleteEntityIntInt8PointerOriginal                           func(item *IntPointer) error
	DeleteEntityIntInt8PointerPointer                            func(item *IntPointer) error
	DeleteArrayParameterIntInt8                                  func(id int) (int, error)
	DeleteArrayParameterIntInt8OriginalPointer                   func(id int) (int, error)
	DeleteArrayParameterIntInt8PointerOriginal                   func(id int) (int, error)
	DeleteArrayParameterIntInt8PointerPointer                    func(id int) (int, error)
	DeleteArrayEntityIntInt8                                     func(id int) (int, error)
	DeleteArrayEntityIntInt8OriginalPointer                      func(id int) (int, error)
	DeleteArrayEntityIntInt8PointerOriginal                      func(id int) (int, error)
	DeleteArrayEntityIntInt8PointerPointer                       func(id int) (int, error)
	DeleteParameterInt4Int8                                      func(sid string) (int, error)
	DeleteParameterInt4Int8OriginalPointer                       func(sid string) (*int, error)
	DeleteParameterInt4Int8PointerOriginal                       func(sid *string) (int, error)
	DeleteParameterInt4Int8PointerPointer                        func(sid string) (*int, error)
	DeleteEntityInt4Int8                                         func(item Int4Original) error
	DeleteEntityInt4Int8OriginalPointer                          func(item Int4Original) error
	DeleteEntityInt4Int8PointerOriginal                          func(item *Int4Pointer) error
	DeleteEntityInt4Int8PointerPointer                           func(item *Int4Pointer) error
	DeleteArrayParameterInt4Int8                                 func(id int) (int, error)
	DeleteArrayParameterInt4Int8OriginalPointer                  func(id int) (int, error)
	DeleteArrayParameterInt4Int8PointerOriginal                  func(id int) (int, error)
	DeleteArrayParameterInt4Int8PointerPointer                   func(id int) (int, error)
	DeleteArrayEntityInt4Int8                                    func(id int) (int, error)
	DeleteArrayEntityInt4Int8OriginalPointer                     func(id int) (int, error)
	DeleteArrayEntityInt4Int8PointerOriginal                     func(id int) (int, error)
	DeleteArrayEntityInt4Int8PointerPointer                      func(id int) (int, error)
	DeleteParameterNumericDecimal                                func(sid string) (int, error)
	DeleteParameterNumericDecimalOriginalPointer                 func(sid string) (*int, error)
	DeleteParameterNumericDecimalPointerOriginal                 func(sid *string) (int, error)
	DeleteParameterNumericDecimalPointerPointer                  func(sid string) (*int, error)
	DeleteEntityNumericDecimal                                   func(item NumericOriginal) error
	DeleteEntityNumericDecimalOriginalPointer                    func(item NumericOriginal) error
	DeleteEntityNumericDecimalPointerOriginal                    func(item *NumericPointer) error
	DeleteEntityNumericDecimalPointerPointer                     func(item *NumericPointer) error
	DeleteArrayParameterNumericDecimal                           func(id int) (int, error)
	DeleteArrayParameterNumericDecimalOriginalPointer            func(id int) (int, error)
	DeleteArrayParameterNumericDecimalPointerOriginal            func(id int) (int, error)
	DeleteArrayParameterNumericDecimalPointerPointer             func(id int) (int, error)
	DeleteArrayEntityNumericDecimal                              func(id int) (int, error)
	DeleteArrayEntityNumericDecimalOriginalPointer               func(id int) (int, error)
	DeleteArrayEntityNumericDecimalPointerOriginal               func(id int) (int, error)
	DeleteArrayEntityNumericDecimalPointerPointer                func(id int) (int, error)
	DeleteParameterDecimalDecimal                                func(sid string) (int, error)
	DeleteParameterDecimalDecimalOriginalPointer                 func(sid string) (*int, error)
	DeleteParameterDecimalDecimalPointerOriginal                 func(sid *string) (int, error)
	DeleteParameterDecimalDecimalPointerPointer                  func(sid string) (*int, error)
	DeleteEntityDecimalDecimal                                   func(item DecimalOriginal) error
	DeleteEntityDecimalDecimalOriginalPointer                    func(item DecimalOriginal) error
	DeleteEntityDecimalDecimalPointerOriginal                    func(item *DecimalPointer) error
	DeleteEntityDecimalDecimalPointerPointer                     func(item *DecimalPointer) error
	DeleteArrayParameterDecimalDecimal                           func(id int) (int, error)
	DeleteArrayParameterDecimalDecimalOriginalPointer            func(id int) (int, error)
	DeleteArrayParameterDecimalDecimalPointerOriginal            func(id int) (int, error)
	DeleteArrayParameterDecimalDecimalPointerPointer             func(id int) (int, error)
	DeleteArrayEntityDecimalDecimal                              func(id int) (int, error)
	DeleteArrayEntityDecimalDecimalOriginalPointer               func(id int) (int, error)
	DeleteArrayEntityDecimalDecimalPointerOriginal               func(id int) (int, error)
	DeleteArrayEntityDecimalDecimalPointerPointer                func(id int) (int, error)
	DeleteParameterFloat4Float32                                 func(sid string) (int, error)
	DeleteParameterFloat4Float32OriginalPointer                  func(sid string) (*int, error)
	DeleteParameterFloat4Float32PointerOriginal                  func(sid *string) (int, error)
	DeleteParameterFloat4Float32PointerPointer                   func(sid string) (*int, error)
	DeleteEntityFloat4Float32                                    func(item Float4Original) error
	DeleteEntityFloat4Float32OriginalPointer                     func(item Float4Original) error
	DeleteEntityFloat4Float32PointerOriginal                     func(item *Float4Pointer) error
	DeleteEntityFloat4Float32PointerPointer                      func(item *Float4Pointer) error
	DeleteArrayParameterFloat4Float32                            func(id int) (int, error)
	DeleteArrayParameterFloat4Float32OriginalPointer             func(id int) (int, error)
	DeleteArrayParameterFloat4Float32PointerOriginal             func(id int) (int, error)
	DeleteArrayParameterFloat4Float32PointerPointer              func(id int) (int, error)
	DeleteArrayEntityFloat4Float32                               func(id int) (int, error)
	DeleteArrayEntityFloat4Float32OriginalPointer                func(id int) (int, error)
	DeleteArrayEntityFloat4Float32PointerOriginal                func(id int) (int, error)
	DeleteArrayEntityFloat4Float32PointerPointer                 func(id int) (int, error)
	DeleteParameterSmallintInt8                                  func(sid string) (int, error)
	DeleteParameterSmallintInt8OriginalPointer                   func(sid string) (*int, error)
	DeleteParameterSmallintInt8PointerOriginal                   func(sid *string) (int, error)
	DeleteParameterSmallintInt8PointerPointer                    func(sid string) (*int, error)
	DeleteEntitySmallintInt8                                     func(item SmallintOriginal) error
	DeleteEntitySmallintInt8OriginalPointer                      func(item SmallintOriginal) error
	DeleteEntitySmallintInt8PointerOriginal                      func(item *SmallintPointer) error
	DeleteEntitySmallintInt8PointerPointer                       func(item *SmallintPointer) error
	DeleteArrayParameterSmallintInt8                             func(id int) (int, error)
	DeleteArrayParameterSmallintInt8OriginalPointer              func(id int) (int, error)
	DeleteArrayParameterSmallintInt8PointerOriginal              func(id int) (int, error)
	DeleteArrayParameterSmallintInt8PointerPointer               func(id int) (int, error)
	DeleteArrayEntitySmallintInt8                                func(id int) (int, error)
	DeleteArrayEntitySmallintInt8OriginalPointer                 func(id int) (int, error)
	DeleteArrayEntitySmallintInt8PointerOriginal                 func(id int) (int, error)
	DeleteArrayEntitySmallintInt8PointerPointer                  func(id int) (int, error)
	DeleteParameterInt2Int8                                      func(sid string) (int, error)
	DeleteParameterInt2Int8OriginalPointer                       func(sid string) (*int, error)
	DeleteParameterInt2Int8PointerOriginal                       func(sid *string) (int, error)
	DeleteParameterInt2Int8PointerPointer                        func(sid string) (*int, error)
	DeleteEntityInt2Int8                                         func(item Int2Original) error
	DeleteEntityInt2Int8OriginalPointer                          func(item Int2Original) error
	DeleteEntityInt2Int8PointerOriginal                          func(item *Int2Pointer) error
	DeleteEntityInt2Int8PointerPointer                           func(item *Int2Pointer) error
	DeleteArrayParameterInt2Int8                                 func(id int) (int, error)
	DeleteArrayParameterInt2Int8OriginalPointer                  func(id int) (int, error)
	DeleteArrayParameterInt2Int8PointerOriginal                  func(id int) (int, error)
	DeleteArrayParameterInt2Int8PointerPointer                   func(id int) (int, error)
	DeleteArrayEntityInt2Int8                                    func(id int) (int, error)
	DeleteArrayEntityInt2Int8OriginalPointer                     func(id int) (int, error)
	DeleteArrayEntityInt2Int8PointerOriginal                     func(id int) (int, error)
	DeleteArrayEntityInt2Int8PointerPointer                      func(id int) (int, error)
	DeleteParameterTextString                                    func(sid string) (int, error)
	DeleteParameterTextStringOriginalPointer                     func(sid string) (*int, error)
	DeleteParameterTextStringPointerOriginal                     func(sid *string) (int, error)
	DeleteParameterTextStringPointerPointer                      func(sid string) (*int, error)
	DeleteEntityTextString                                       func(item TextOriginal) error
	DeleteEntityTextStringOriginalPointer                        func(item TextOriginal) error
	DeleteEntityTextStringPointerOriginal                        func(item *TextPointer) error
	DeleteEntityTextStringPointerPointer                         func(item *TextPointer) error
	DeleteArrayParameterTextString                               func(id int) (int, error)
	DeleteArrayParameterTextStringOriginalPointer                func(id int) (int, error)
	DeleteArrayParameterTextStringPointerOriginal                func(id int) (int, error)
	DeleteArrayParameterTextStringPointerPointer                 func(id int) (int, error)
	DeleteArrayEntityTextString                                  func(id int) (int, error)
	DeleteArrayEntityTextStringOriginalPointer                   func(id int) (int, error)
	DeleteArrayEntityTextStringPointerOriginal                   func(id int) (int, error)
	DeleteArrayEntityTextStringPointerPointer                    func(id int) (int, error)
	DeleteParameterTimeTime                                      func(sid string) (int, error)
	DeleteParameterTimeTimeOriginalPointer                       func(sid string) (*int, error)
	DeleteParameterTimeTimePointerOriginal                       func(sid *string) (int, error)
	DeleteParameterTimeTimePointerPointer                        func(sid string) (*int, error)
	DeleteEntityTimeTime                                         func(item TimeOriginal) error
	DeleteEntityTimeTimeOriginalPointer                          func(item TimeOriginal) error
	DeleteEntityTimeTimePointerOriginal                          func(item *TimePointer) error
	DeleteEntityTimeTimePointerPointer                           func(item *TimePointer) error
	DeleteArrayParameterTimeTime                                 func(id int) (int, error)
	DeleteArrayParameterTimeTimeOriginalPointer                  func(id int) (int, error)
	DeleteArrayParameterTimeTimePointerOriginal                  func(id int) (int, error)
	DeleteArrayParameterTimeTimePointerPointer                   func(id int) (int, error)
	DeleteArrayEntityTimeTime                                    func(id int) (int, error)
	DeleteArrayEntityTimeTimeOriginalPointer                     func(id int) (int, error)
	DeleteArrayEntityTimeTimePointerOriginal                     func(id int) (int, error)
	DeleteArrayEntityTimeTimePointerPointer                      func(id int) (int, error)
	DeleteParameterTimeWithTimezoneTime                          func(sid string) (int, error)
	DeleteParameterTimeWithTimezoneTimeOriginalPointer           func(sid string) (*int, error)
	DeleteParameterTimeWithTimezoneTimePointerOriginal           func(sid *string) (int, error)
	DeleteParameterTimeWithTimezoneTimePointerPointer            func(sid string) (*int, error)
	DeleteEntityTimeWithTimezoneTime                             func(item TimeWithTimezoneOriginal) error
	DeleteEntityTimeWithTimezoneTimeOriginalPointer              func(item TimeWithTimezoneOriginal) error
	DeleteEntityTimeWithTimezoneTimePointerOriginal              func(item *TimeWithTimezonePointer) error
	DeleteEntityTimeWithTimezoneTimePointerPointer               func(item *TimeWithTimezonePointer) error
	DeleteArrayParameterTimeWithTimezoneTime                     func(id int) (int, error)
	DeleteArrayParameterTimeWithTimezoneTimeOriginalPointer      func(id int) (int, error)
	DeleteArrayParameterTimeWithTimezoneTimePointerOriginal      func(id int) (int, error)
	DeleteArrayParameterTimeWithTimezoneTimePointerPointer       func(id int) (int, error)
	DeleteArrayEntityTimeWithTimezoneTime                        func(id int) (int, error)
	DeleteArrayEntityTimeWithTimezoneTimeOriginalPointer         func(id int) (int, error)
	DeleteArrayEntityTimeWithTimezoneTimePointerOriginal         func(id int) (int, error)
	DeleteArrayEntityTimeWithTimezoneTimePointerPointer          func(id int) (int, error)
	DeleteParameterTimetzTime                                    func(sid string) (int, error)
	DeleteParameterTimetzTimeOriginalPointer                     func(sid string) (*int, error)
	DeleteParameterTimetzTimePointerOriginal                     func(sid *string) (int, error)
	DeleteParameterTimetzTimePointerPointer                      func(sid string) (*int, error)
	DeleteEntityTimetzTime                                       func(item TimetzOriginal) error
	DeleteEntityTimetzTimeOriginalPointer                        func(item TimetzOriginal) error
	DeleteEntityTimetzTimePointerOriginal                        func(item *TimetzPointer) error
	DeleteEntityTimetzTimePointerPointer                         func(item *TimetzPointer) error
	DeleteArrayParameterTimetzTime                               func(id int) (int, error)
	DeleteArrayParameterTimetzTimeOriginalPointer                func(id int) (int, error)
	DeleteArrayParameterTimetzTimePointerOriginal                func(id int) (int, error)
	DeleteArrayParameterTimetzTimePointerPointer                 func(id int) (int, error)
	DeleteArrayEntityTimetzTime                                  func(id int) (int, error)
	DeleteArrayEntityTimetzTimeOriginalPointer                   func(id int) (int, error)
	DeleteArrayEntityTimetzTimePointerOriginal                   func(id int) (int, error)
	DeleteArrayEntityTimetzTimePointerPointer                    func(id int) (int, error)
	DeleteParameterTimestampTime                                 func(sid string) (int, error)
	DeleteParameterTimestampTimeOriginalPointer                  func(sid string) (*int, error)
	DeleteParameterTimestampTimePointerOriginal                  func(sid *string) (int, error)
	DeleteParameterTimestampTimePointerPointer                   func(sid string) (*int, error)
	DeleteEntityTimestampTime                                    func(item TimestampOriginal) error
	DeleteEntityTimestampTimeOriginalPointer                     func(item TimestampOriginal) error
	DeleteEntityTimestampTimePointerOriginal                     func(item *TimestampPointer) error
	DeleteEntityTimestampTimePointerPointer                      func(item *TimestampPointer) error
	DeleteArrayParameterTimestampTime                            func(id int) (int, error)
	DeleteArrayParameterTimestampTimeOriginalPointer             func(id int) (int, error)
	DeleteArrayParameterTimestampTimePointerOriginal             func(id int) (int, error)
	DeleteArrayParameterTimestampTimePointerPointer              func(id int) (int, error)
	DeleteArrayEntityTimestampTime                               func(id int) (int, error)
	DeleteArrayEntityTimestampTimeOriginalPointer                func(id int) (int, error)
	DeleteArrayEntityTimestampTimePointerOriginal                func(id int) (int, error)
	DeleteArrayEntityTimestampTimePointerPointer                 func(id int) (int, error)
	DeleteParameterTimestampWithTimezoneTime                     func(sid string) (int, error)
	DeleteParameterTimestampWithTimezoneTimeOriginalPointer      func(sid string) (*int, error)
	DeleteParameterTimestampWithTimezoneTimePointerOriginal      func(sid *string) (int, error)
	DeleteParameterTimestampWithTimezoneTimePointerPointer       func(sid string) (*int, error)
	DeleteEntityTimestampWithTimezoneTime                        func(item TimestampWithTimezoneOriginal) error
	DeleteEntityTimestampWithTimezoneTimeOriginalPointer         func(item TimestampWithTimezoneOriginal) error
	DeleteEntityTimestampWithTimezoneTimePointerOriginal         func(item *TimestampWithTimezonePointer) error
	DeleteEntityTimestampWithTimezoneTimePointerPointer          func(item *TimestampWithTimezonePointer) error
	DeleteArrayParameterTimestampWithTimezoneTime                func(id int) (int, error)
	DeleteArrayParameterTimestampWithTimezoneTimeOriginalPointer func(id int) (int, error)
	DeleteArrayParameterTimestampWithTimezoneTimePointerOriginal func(id int) (int, error)
	DeleteArrayParameterTimestampWithTimezoneTimePointerPointer  func(id int) (int, error)
	DeleteArrayEntityTimestampWithTimezoneTime                   func(id int) (int, error)
	DeleteArrayEntityTimestampWithTimezoneTimeOriginalPointer    func(id int) (int, error)
	DeleteArrayEntityTimestampWithTimezoneTimePointerOriginal    func(id int) (int, error)
	DeleteArrayEntityTimestampWithTimezoneTimePointerPointer     func(id int) (int, error)
	DeleteParameterTimestamptzTime                               func(sid string) (int, error)
	DeleteParameterTimestamptzTimeOriginalPointer                func(sid string) (*int, error)
	DeleteParameterTimestamptzTimePointerOriginal                func(sid *string) (int, error)
	DeleteParameterTimestamptzTimePointerPointer                 func(sid string) (*int, error)
	DeleteEntityTimestamptzTime                                  func(item TimestamptzOriginal) error
	DeleteEntityTimestamptzTimeOriginalPointer                   func(item TimestamptzOriginal) error
	DeleteEntityTimestamptzTimePointerOriginal                   func(item *TimestamptzPointer) error
	DeleteEntityTimestamptzTimePointerPointer                    func(item *TimestamptzPointer) error
	DeleteArrayParameterTimestamptzTime                          func(id int) (int, error)
	DeleteArrayParameterTimestamptzTimeOriginalPointer           func(id int) (int, error)
	DeleteArrayParameterTimestamptzTimePointerOriginal           func(id int) (int, error)
	DeleteArrayParameterTimestamptzTimePointerPointer            func(id int) (int, error)
	DeleteArrayEntityTimestamptzTime                             func(id int) (int, error)
	DeleteArrayEntityTimestamptzTimeOriginalPointer              func(id int) (int, error)
	DeleteArrayEntityTimestamptzTimePointerOriginal              func(id int) (int, error)
	DeleteArrayEntityTimestamptzTimePointerPointer               func(id int) (int, error)
}
