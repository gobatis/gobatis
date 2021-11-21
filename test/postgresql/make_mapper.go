package postgresql

import ()

type MakeMapper struct {
	InsertParameterBigintInt64                                func(sid string, source string, var_bigint int64) (int, error)
	InsertParameterBigintInt64OriginalPointer                 func(sid string, source string, var_bigint int64) (*int, error)
	InsertParameterBigintInt64PointerOriginal                 func(sid *string, source *string, var_bigint *int64) (int, error)
	InsertParameterBigintInt64PointerPointer                  func(sid string, source string, var_bigint *int64) (*int, error)
	InsertEntityBigintInt64                                   func(item BigintOriginal) error
	InsertEntityBigintInt64OriginalPointer                    func(item BigintOriginal) error
	InsertEntityBigintInt64PointerOriginal                    func(item *BigintPointer) error
	InsertEntityBigintInt64PointerPointer                     func(item *BigintPointer) error
	InsertArrayParameterBigintInt64                           func(sid string, source string, items []int64) (int, error)
	InsertArrayParameterBigintInt64OriginalPointer            func(sid string, source string, items []int64) (*int, error)
	InsertArrayParameterBigintInt64PointerOriginal            func(sid string, source string, items []*int64) (int, error)
	InsertArrayParameterBigintInt64PointerPointer             func(sid string, source string, items []*int64) (*int, error)
	InsertArrayEntityBigintInt64                              func(sid string, source string, items []*int64) (int, error)
	InsertArrayEntityBigintInt64OriginalPointer               func(sid string, source string, items []*int64) (int, error)
	InsertArrayEntityBigintInt64PointerOriginal               func(sid string, source string, items []*int64) (int, error)
	InsertArrayEntityBigintInt64PointerPointer                func(sid string, source string, items []*int64) (int, error)
	InsertParameterInt8Int8                                   func(sid string, source string, var_int8 int8) (int, error)
	InsertParameterInt8Int8OriginalPointer                    func(sid string, source string, var_int8 int8) (*int, error)
	InsertParameterInt8Int8PointerOriginal                    func(sid *string, source *string, var_int8 *int8) (int, error)
	InsertParameterInt8Int8PointerPointer                     func(sid string, source string, var_int8 *int8) (*int, error)
	InsertEntityInt8Int8                                      func(item Int8Original) error
	InsertEntityInt8Int8OriginalPointer                       func(item Int8Original) error
	InsertEntityInt8Int8PointerOriginal                       func(item *Int8Pointer) error
	InsertEntityInt8Int8PointerPointer                        func(item *Int8Pointer) error
	InsertArrayParameterInt8Int8                              func(sid string, source string, items []int8) (int, error)
	InsertArrayParameterInt8Int8OriginalPointer               func(sid string, source string, items []int8) (*int, error)
	InsertArrayParameterInt8Int8PointerOriginal               func(sid string, source string, items []*int8) (int, error)
	InsertArrayParameterInt8Int8PointerPointer                func(sid string, source string, items []*int8) (*int, error)
	InsertArrayEntityInt8Int8                                 func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt8Int8OriginalPointer                  func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt8Int8PointerOriginal                  func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt8Int8PointerPointer                   func(sid string, source string, items []*int8) (int, error)
	InsertParameterBooleanBool                                func(sid string, source string, var_boolean bool) (int, error)
	InsertParameterBooleanBoolOriginalPointer                 func(sid string, source string, var_boolean bool) (*int, error)
	InsertParameterBooleanBoolPointerOriginal                 func(sid *string, source *string, var_boolean *bool) (int, error)
	InsertParameterBooleanBoolPointerPointer                  func(sid string, source string, var_boolean *bool) (*int, error)
	InsertEntityBooleanBool                                   func(item BooleanOriginal) error
	InsertEntityBooleanBoolOriginalPointer                    func(item BooleanOriginal) error
	InsertEntityBooleanBoolPointerOriginal                    func(item *BooleanPointer) error
	InsertEntityBooleanBoolPointerPointer                     func(item *BooleanPointer) error
	InsertArrayParameterBooleanBool                           func(sid string, source string, items []bool) (int, error)
	InsertArrayParameterBooleanBoolOriginalPointer            func(sid string, source string, items []bool) (*int, error)
	InsertArrayParameterBooleanBoolPointerOriginal            func(sid string, source string, items []*bool) (int, error)
	InsertArrayParameterBooleanBoolPointerPointer             func(sid string, source string, items []*bool) (*int, error)
	InsertArrayEntityBooleanBool                              func(sid string, source string, items []*bool) (int, error)
	InsertArrayEntityBooleanBoolOriginalPointer               func(sid string, source string, items []*bool) (int, error)
	InsertArrayEntityBooleanBoolPointerOriginal               func(sid string, source string, items []*bool) (int, error)
	InsertArrayEntityBooleanBoolPointerPointer                func(sid string, source string, items []*bool) (int, error)
	InsertParameterBoolBool                                   func(sid string, source string, var_bool bool) (int, error)
	InsertParameterBoolBoolOriginalPointer                    func(sid string, source string, var_bool bool) (*int, error)
	InsertParameterBoolBoolPointerOriginal                    func(sid *string, source *string, var_bool *bool) (int, error)
	InsertParameterBoolBoolPointerPointer                     func(sid string, source string, var_bool *bool) (*int, error)
	InsertEntityBoolBool                                      func(item BoolOriginal) error
	InsertEntityBoolBoolOriginalPointer                       func(item BoolOriginal) error
	InsertEntityBoolBoolPointerOriginal                       func(item *BoolPointer) error
	InsertEntityBoolBoolPointerPointer                        func(item *BoolPointer) error
	InsertArrayParameterBoolBool                              func(sid string, source string, items []bool) (int, error)
	InsertArrayParameterBoolBoolOriginalPointer               func(sid string, source string, items []bool) (*int, error)
	InsertArrayParameterBoolBoolPointerOriginal               func(sid string, source string, items []*bool) (int, error)
	InsertArrayParameterBoolBoolPointerPointer                func(sid string, source string, items []*bool) (*int, error)
	InsertArrayEntityBoolBool                                 func(sid string, source string, items []*bool) (int, error)
	InsertArrayEntityBoolBoolOriginalPointer                  func(sid string, source string, items []*bool) (int, error)
	InsertArrayEntityBoolBoolPointerOriginal                  func(sid string, source string, items []*bool) (int, error)
	InsertArrayEntityBoolBoolPointerPointer                   func(sid string, source string, items []*bool) (int, error)
	InsertParameterCharacterString                            func(sid string, source string, var_character string) (int, error)
	InsertParameterCharacterStringOriginalPointer             func(sid string, source string, var_character string) (*int, error)
	InsertParameterCharacterStringPointerOriginal             func(sid *string, source *string, var_character *string) (int, error)
	InsertParameterCharacterStringPointerPointer              func(sid string, source string, var_character *string) (*int, error)
	InsertEntityCharacterString                               func(item CharacterOriginal) error
	InsertEntityCharacterStringOriginalPointer                func(item CharacterOriginal) error
	InsertEntityCharacterStringPointerOriginal                func(item *CharacterPointer) error
	InsertEntityCharacterStringPointerPointer                 func(item *CharacterPointer) error
	InsertArrayParameterCharacterString                       func(sid string, source string, items []string) (int, error)
	InsertArrayParameterCharacterStringOriginalPointer        func(sid string, source string, items []string) (*int, error)
	InsertArrayParameterCharacterStringPointerOriginal        func(sid string, source string, items []*string) (int, error)
	InsertArrayParameterCharacterStringPointerPointer         func(sid string, source string, items []*string) (*int, error)
	InsertArrayEntityCharacterString                          func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharacterStringOriginalPointer           func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharacterStringPointerOriginal           func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharacterStringPointerPointer            func(sid string, source string, items []*string) (int, error)
	InsertParameterCharString                                 func(sid string, source string, var_char string) (int, error)
	InsertParameterCharStringOriginalPointer                  func(sid string, source string, var_char string) (*int, error)
	InsertParameterCharStringPointerOriginal                  func(sid *string, source *string, var_char *string) (int, error)
	InsertParameterCharStringPointerPointer                   func(sid string, source string, var_char *string) (*int, error)
	InsertEntityCharString                                    func(item CharOriginal) error
	InsertEntityCharStringOriginalPointer                     func(item CharOriginal) error
	InsertEntityCharStringPointerOriginal                     func(item *CharPointer) error
	InsertEntityCharStringPointerPointer                      func(item *CharPointer) error
	InsertArrayParameterCharString                            func(sid string, source string, items []string) (int, error)
	InsertArrayParameterCharStringOriginalPointer             func(sid string, source string, items []string) (*int, error)
	InsertArrayParameterCharStringPointerOriginal             func(sid string, source string, items []*string) (int, error)
	InsertArrayParameterCharStringPointerPointer              func(sid string, source string, items []*string) (*int, error)
	InsertArrayEntityCharString                               func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharStringOriginalPointer                func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharStringPointerOriginal                func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharStringPointerPointer                 func(sid string, source string, items []*string) (int, error)
	InsertParameterCharacterVaryingString                     func(sid string, source string, var_character_varying string) (int, error)
	InsertParameterCharacterVaryingStringOriginalPointer      func(sid string, source string, var_character_varying string) (*int, error)
	InsertParameterCharacterVaryingStringPointerOriginal      func(sid *string, source *string, var_character_varying *string) (int, error)
	InsertParameterCharacterVaryingStringPointerPointer       func(sid string, source string, var_character_varying *string) (*int, error)
	InsertEntityCharacterVaryingString                        func(item CharacterVaryingOriginal) error
	InsertEntityCharacterVaryingStringOriginalPointer         func(item CharacterVaryingOriginal) error
	InsertEntityCharacterVaryingStringPointerOriginal         func(item *CharacterVaryingPointer) error
	InsertEntityCharacterVaryingStringPointerPointer          func(item *CharacterVaryingPointer) error
	InsertArrayParameterCharacterVaryingString                func(sid string, source string, items []string) (int, error)
	InsertArrayParameterCharacterVaryingStringOriginalPointer func(sid string, source string, items []string) (*int, error)
	InsertArrayParameterCharacterVaryingStringPointerOriginal func(sid string, source string, items []*string) (int, error)
	InsertArrayParameterCharacterVaryingStringPointerPointer  func(sid string, source string, items []*string) (*int, error)
	InsertArrayEntityCharacterVaryingString                   func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharacterVaryingStringOriginalPointer    func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharacterVaryingStringPointerOriginal    func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityCharacterVaryingStringPointerPointer     func(sid string, source string, items []*string) (int, error)
	InsertParameterVarcharString                              func(sid string, source string, var_varchar string) (int, error)
	InsertParameterVarcharStringOriginalPointer               func(sid string, source string, var_varchar string) (*int, error)
	InsertParameterVarcharStringPointerOriginal               func(sid *string, source *string, var_varchar *string) (int, error)
	InsertParameterVarcharStringPointerPointer                func(sid string, source string, var_varchar *string) (*int, error)
	InsertEntityVarcharString                                 func(item VarcharOriginal) error
	InsertEntityVarcharStringOriginalPointer                  func(item VarcharOriginal) error
	InsertEntityVarcharStringPointerOriginal                  func(item *VarcharPointer) error
	InsertEntityVarcharStringPointerPointer                   func(item *VarcharPointer) error
	InsertArrayParameterVarcharString                         func(sid string, source string, items []string) (int, error)
	InsertArrayParameterVarcharStringOriginalPointer          func(sid string, source string, items []string) (*int, error)
	InsertArrayParameterVarcharStringPointerOriginal          func(sid string, source string, items []*string) (int, error)
	InsertArrayParameterVarcharStringPointerPointer           func(sid string, source string, items []*string) (*int, error)
	InsertArrayEntityVarcharString                            func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityVarcharStringOriginalPointer             func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityVarcharStringPointerOriginal             func(sid string, source string, items []*string) (int, error)
	InsertArrayEntityVarcharStringPointerPointer              func(sid string, source string, items []*string) (int, error)
	InsertParameterFloat8Float32                              func(sid string, source string, var_float8 float32) (int, error)
	InsertParameterFloat8Float32OriginalPointer               func(sid string, source string, var_float8 float32) (*int, error)
	InsertParameterFloat8Float32PointerOriginal               func(sid *string, source *string, var_float8 *float32) (int, error)
	InsertParameterFloat8Float32PointerPointer                func(sid string, source string, var_float8 *float32) (*int, error)
	InsertEntityFloat8Float32                                 func(item Float8Original) error
	InsertEntityFloat8Float32OriginalPointer                  func(item Float8Original) error
	InsertEntityFloat8Float32PointerOriginal                  func(item *Float8Pointer) error
	InsertEntityFloat8Float32PointerPointer                   func(item *Float8Pointer) error
	InsertArrayParameterFloat8Float32                         func(sid string, source string, items []float32) (int, error)
	InsertArrayParameterFloat8Float32OriginalPointer          func(sid string, source string, items []float32) (*int, error)
	InsertArrayParameterFloat8Float32PointerOriginal          func(sid string, source string, items []*float32) (int, error)
	InsertArrayParameterFloat8Float32PointerPointer           func(sid string, source string, items []*float32) (*int, error)
	InsertArrayEntityFloat8Float32                            func(sid string, source string, items []*float32) (int, error)
	InsertArrayEntityFloat8Float32OriginalPointer             func(sid string, source string, items []*float32) (int, error)
	InsertArrayEntityFloat8Float32PointerOriginal             func(sid string, source string, items []*float32) (int, error)
	InsertArrayEntityFloat8Float32PointerPointer              func(sid string, source string, items []*float32) (int, error)
	InsertParameterIntegerInt8                                func(sid string, source string, var_integer int8) (int, error)
	InsertParameterIntegerInt8OriginalPointer                 func(sid string, source string, var_integer int8) (*int, error)
	InsertParameterIntegerInt8PointerOriginal                 func(sid *string, source *string, var_integer *int8) (int, error)
	InsertParameterIntegerInt8PointerPointer                  func(sid string, source string, var_integer *int8) (*int, error)
	InsertEntityIntegerInt8                                   func(item IntegerOriginal) error
	InsertEntityIntegerInt8OriginalPointer                    func(item IntegerOriginal) error
	InsertEntityIntegerInt8PointerOriginal                    func(item *IntegerPointer) error
	InsertEntityIntegerInt8PointerPointer                     func(item *IntegerPointer) error
	InsertArrayParameterIntegerInt8                           func(sid string, source string, items []int8) (int, error)
	InsertArrayParameterIntegerInt8OriginalPointer            func(sid string, source string, items []int8) (*int, error)
	InsertArrayParameterIntegerInt8PointerOriginal            func(sid string, source string, items []*int8) (int, error)
	InsertArrayParameterIntegerInt8PointerPointer             func(sid string, source string, items []*int8) (*int, error)
	InsertArrayEntityIntegerInt8                              func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityIntegerInt8OriginalPointer               func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityIntegerInt8PointerOriginal               func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityIntegerInt8PointerPointer                func(sid string, source string, items []*int8) (int, error)
	InsertParameterIntInt8                                    func(sid string, source string, var_int int8) (int, error)
	InsertParameterIntInt8OriginalPointer                     func(sid string, source string, var_int int8) (*int, error)
	InsertParameterIntInt8PointerOriginal                     func(sid *string, source *string, var_int *int8) (int, error)
	InsertParameterIntInt8PointerPointer                      func(sid string, source string, var_int *int8) (*int, error)
	InsertEntityIntInt8                                       func(item IntOriginal) error
	InsertEntityIntInt8OriginalPointer                        func(item IntOriginal) error
	InsertEntityIntInt8PointerOriginal                        func(item *IntPointer) error
	InsertEntityIntInt8PointerPointer                         func(item *IntPointer) error
	InsertArrayParameterIntInt8                               func(sid string, source string, items []int8) (int, error)
	InsertArrayParameterIntInt8OriginalPointer                func(sid string, source string, items []int8) (*int, error)
	InsertArrayParameterIntInt8PointerOriginal                func(sid string, source string, items []*int8) (int, error)
	InsertArrayParameterIntInt8PointerPointer                 func(sid string, source string, items []*int8) (*int, error)
	InsertArrayEntityIntInt8                                  func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityIntInt8OriginalPointer                   func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityIntInt8PointerOriginal                   func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityIntInt8PointerPointer                    func(sid string, source string, items []*int8) (int, error)
	InsertParameterInt4Int8                                   func(sid string, source string, var_int4 int8) (int, error)
	InsertParameterInt4Int8OriginalPointer                    func(sid string, source string, var_int4 int8) (*int, error)
	InsertParameterInt4Int8PointerOriginal                    func(sid *string, source *string, var_int4 *int8) (int, error)
	InsertParameterInt4Int8PointerPointer                     func(sid string, source string, var_int4 *int8) (*int, error)
	InsertEntityInt4Int8                                      func(item Int4Original) error
	InsertEntityInt4Int8OriginalPointer                       func(item Int4Original) error
	InsertEntityInt4Int8PointerOriginal                       func(item *Int4Pointer) error
	InsertEntityInt4Int8PointerPointer                        func(item *Int4Pointer) error
	InsertArrayParameterInt4Int8                              func(sid string, source string, items []int8) (int, error)
	InsertArrayParameterInt4Int8OriginalPointer               func(sid string, source string, items []int8) (*int, error)
	InsertArrayParameterInt4Int8PointerOriginal               func(sid string, source string, items []*int8) (int, error)
	InsertArrayParameterInt4Int8PointerPointer                func(sid string, source string, items []*int8) (*int, error)
	InsertArrayEntityInt4Int8                                 func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt4Int8OriginalPointer                  func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt4Int8PointerOriginal                  func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt4Int8PointerPointer                   func(sid string, source string, items []*int8) (int, error)
	InsertParameterNumericDecimal                             func(sid string, source string, var_numeric decimal) (int, error)
	InsertParameterNumericDecimalOriginalPointer              func(sid string, source string, var_numeric decimal) (*int, error)
	InsertParameterNumericDecimalPointerOriginal              func(sid *string, source *string, var_numeric *decimal) (int, error)
	InsertParameterNumericDecimalPointerPointer               func(sid string, source string, var_numeric *decimal) (*int, error)
	InsertEntityNumericDecimal                                func(item NumericOriginal) error
	InsertEntityNumericDecimalOriginalPointer                 func(item NumericOriginal) error
	InsertEntityNumericDecimalPointerOriginal                 func(item *NumericPointer) error
	InsertEntityNumericDecimalPointerPointer                  func(item *NumericPointer) error
	InsertArrayParameterNumericDecimal                        func(sid string, source string, items []decimal) (int, error)
	InsertArrayParameterNumericDecimalOriginalPointer         func(sid string, source string, items []decimal) (*int, error)
	InsertArrayParameterNumericDecimalPointerOriginal         func(sid string, source string, items []*decimal) (int, error)
	InsertArrayParameterNumericDecimalPointerPointer          func(sid string, source string, items []*decimal) (*int, error)
	InsertArrayEntityNumericDecimal                           func(sid string, source string, items []*decimal) (int, error)
	InsertArrayEntityNumericDecimalOriginalPointer            func(sid string, source string, items []*decimal) (int, error)
	InsertArrayEntityNumericDecimalPointerOriginal            func(sid string, source string, items []*decimal) (int, error)
	InsertArrayEntityNumericDecimalPointerPointer             func(sid string, source string, items []*decimal) (int, error)
	InsertParameterDecimalDecimal                             func(sid string, source string, var_decimal decimal) (int, error)
	InsertParameterDecimalDecimalOriginalPointer              func(sid string, source string, var_decimal decimal) (*int, error)
	InsertParameterDecimalDecimalPointerOriginal              func(sid *string, source *string, var_decimal *decimal) (int, error)
	InsertParameterDecimalDecimalPointerPointer               func(sid string, source string, var_decimal *decimal) (*int, error)
	InsertEntityDecimalDecimal                                func(item DecimalOriginal) error
	InsertEntityDecimalDecimalOriginalPointer                 func(item DecimalOriginal) error
	InsertEntityDecimalDecimalPointerOriginal                 func(item *DecimalPointer) error
	InsertEntityDecimalDecimalPointerPointer                  func(item *DecimalPointer) error
	InsertArrayParameterDecimalDecimal                        func(sid string, source string, items []decimal) (int, error)
	InsertArrayParameterDecimalDecimalOriginalPointer         func(sid string, source string, items []decimal) (*int, error)
	InsertArrayParameterDecimalDecimalPointerOriginal         func(sid string, source string, items []*decimal) (int, error)
	InsertArrayParameterDecimalDecimalPointerPointer          func(sid string, source string, items []*decimal) (*int, error)
	InsertArrayEntityDecimalDecimal                           func(sid string, source string, items []*decimal) (int, error)
	InsertArrayEntityDecimalDecimalOriginalPointer            func(sid string, source string, items []*decimal) (int, error)
	InsertArrayEntityDecimalDecimalPointerOriginal            func(sid string, source string, items []*decimal) (int, error)
	InsertArrayEntityDecimalDecimalPointerPointer             func(sid string, source string, items []*decimal) (int, error)
	InsertParameterFloat4Float32                              func(sid string, source string, var_float4 float32) (int, error)
	InsertParameterFloat4Float32OriginalPointer               func(sid string, source string, var_float4 float32) (*int, error)
	InsertParameterFloat4Float32PointerOriginal               func(sid *string, source *string, var_float4 *float32) (int, error)
	InsertParameterFloat4Float32PointerPointer                func(sid string, source string, var_float4 *float32) (*int, error)
	InsertEntityFloat4Float32                                 func(item Float4Original) error
	InsertEntityFloat4Float32OriginalPointer                  func(item Float4Original) error
	InsertEntityFloat4Float32PointerOriginal                  func(item *Float4Pointer) error
	InsertEntityFloat4Float32PointerPointer                   func(item *Float4Pointer) error
	InsertArrayParameterFloat4Float32                         func(sid string, source string, items []float32) (int, error)
	InsertArrayParameterFloat4Float32OriginalPointer          func(sid string, source string, items []float32) (*int, error)
	InsertArrayParameterFloat4Float32PointerOriginal          func(sid string, source string, items []*float32) (int, error)
	InsertArrayParameterFloat4Float32PointerPointer           func(sid string, source string, items []*float32) (*int, error)
	InsertArrayEntityFloat4Float32                            func(sid string, source string, items []*float32) (int, error)
	InsertArrayEntityFloat4Float32OriginalPointer             func(sid string, source string, items []*float32) (int, error)
	InsertArrayEntityFloat4Float32PointerOriginal             func(sid string, source string, items []*float32) (int, error)
	InsertArrayEntityFloat4Float32PointerPointer              func(sid string, source string, items []*float32) (int, error)
	InsertParameterSmallintInt8                               func(sid string, source string, var_smallint int8) (int, error)
	InsertParameterSmallintInt8OriginalPointer                func(sid string, source string, var_smallint int8) (*int, error)
	InsertParameterSmallintInt8PointerOriginal                func(sid *string, source *string, var_smallint *int8) (int, error)
	InsertParameterSmallintInt8PointerPointer                 func(sid string, source string, var_smallint *int8) (*int, error)
	InsertEntitySmallintInt8                                  func(item SmallintOriginal) error
	InsertEntitySmallintInt8OriginalPointer                   func(item SmallintOriginal) error
	InsertEntitySmallintInt8PointerOriginal                   func(item *SmallintPointer) error
	InsertEntitySmallintInt8PointerPointer                    func(item *SmallintPointer) error
	InsertArrayParameterSmallintInt8                          func(sid string, source string, items []int8) (int, error)
	InsertArrayParameterSmallintInt8OriginalPointer           func(sid string, source string, items []int8) (*int, error)
	InsertArrayParameterSmallintInt8PointerOriginal           func(sid string, source string, items []*int8) (int, error)
	InsertArrayParameterSmallintInt8PointerPointer            func(sid string, source string, items []*int8) (*int, error)
	InsertArrayEntitySmallintInt8                             func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntitySmallintInt8OriginalPointer              func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntitySmallintInt8PointerOriginal              func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntitySmallintInt8PointerPointer               func(sid string, source string, items []*int8) (int, error)
	InsertParameterInt2Int8                                   func(sid string, source string, var_int2 int8) (int, error)
	InsertParameterInt2Int8OriginalPointer                    func(sid string, source string, var_int2 int8) (*int, error)
	InsertParameterInt2Int8PointerOriginal                    func(sid *string, source *string, var_int2 *int8) (int, error)
	InsertParameterInt2Int8PointerPointer                     func(sid string, source string, var_int2 *int8) (*int, error)
	InsertEntityInt2Int8                                      func(item Int2Original) error
	InsertEntityInt2Int8OriginalPointer                       func(item Int2Original) error
	InsertEntityInt2Int8PointerOriginal                       func(item *Int2Pointer) error
	InsertEntityInt2Int8PointerPointer                        func(item *Int2Pointer) error
	InsertArrayParameterInt2Int8                              func(sid string, source string, items []int8) (int, error)
	InsertArrayParameterInt2Int8OriginalPointer               func(sid string, source string, items []int8) (*int, error)
	InsertArrayParameterInt2Int8PointerOriginal               func(sid string, source string, items []*int8) (int, error)
	InsertArrayParameterInt2Int8PointerPointer                func(sid string, source string, items []*int8) (*int, error)
	InsertArrayEntityInt2Int8                                 func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt2Int8OriginalPointer                  func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt2Int8PointerOriginal                  func(sid string, source string, items []*int8) (int, error)
	InsertArrayEntityInt2Int8PointerPointer                   func(sid string, source string, items []*int8) (int, error)
	InsertParameterTextString                                 func(sid string, source string, var_text string) (int, error)
	InsertParameterTextStringOriginalPointer                  func(sid string, source string, var_text string) (*int, error)
	InsertParameterTextStringPointerOriginal                  func(sid *string, source *string, var_text *string) (int, error)
	InsertParameterTextStringPointerPointer                   func(sid string, source string, var_text *string) (*int, error)
	InsertEntityTextString                                    func(item TextOriginal) error
	InsertEntityTextStringOriginalPointer                     func(item TextOriginal) error
	InsertEntityTextStringPointerOriginal                     func(item *TextPointer) error
	InsertEntityTextStringPointerPointer                      func(item *TextPointer) error
	InsertParameterTimeTime                                   func(sid string, source string, var_time time) (int, error)
	InsertParameterTimeTimeOriginalPointer                    func(sid string, source string, var_time time) (*int, error)
	InsertParameterTimeTimePointerOriginal                    func(sid *string, source *string, var_time *time) (int, error)
	InsertParameterTimeTimePointerPointer                     func(sid string, source string, var_time *time) (*int, error)
	InsertEntityTimeTime                                      func(item TimeOriginal) error
	InsertEntityTimeTimeOriginalPointer                       func(item TimeOriginal) error
	InsertEntityTimeTimePointerOriginal                       func(item *TimePointer) error
	InsertEntityTimeTimePointerPointer                        func(item *TimePointer) error
	InsertParameterTimeWithTimezoneTime                       func(sid string, source string, var_time_with_timezone time) (int, error)
	InsertParameterTimeWithTimezoneTimeOriginalPointer        func(sid string, source string, var_time_with_timezone time) (*int, error)
	InsertParameterTimeWithTimezoneTimePointerOriginal        func(sid *string, source *string, var_time_with_timezone *time) (int, error)
	InsertParameterTimeWithTimezoneTimePointerPointer         func(sid string, source string, var_time_with_timezone *time) (*int, error)
	InsertEntityTimeWithTimezoneTime                          func(item TimeWithTimezoneOriginal) error
	InsertEntityTimeWithTimezoneTimeOriginalPointer           func(item TimeWithTimezoneOriginal) error
	InsertEntityTimeWithTimezoneTimePointerOriginal           func(item *TimeWithTimezonePointer) error
	InsertEntityTimeWithTimezoneTimePointerPointer            func(item *TimeWithTimezonePointer) error
	InsertParameterTimetzTime                                 func(sid string, source string, var_timetz time) (int, error)
	InsertParameterTimetzTimeOriginalPointer                  func(sid string, source string, var_timetz time) (*int, error)
	InsertParameterTimetzTimePointerOriginal                  func(sid *string, source *string, var_timetz *time) (int, error)
	InsertParameterTimetzTimePointerPointer                   func(sid string, source string, var_timetz *time) (*int, error)
	InsertEntityTimetzTime                                    func(item TimetzOriginal) error
	InsertEntityTimetzTimeOriginalPointer                     func(item TimetzOriginal) error
	InsertEntityTimetzTimePointerOriginal                     func(item *TimetzPointer) error
	InsertEntityTimetzTimePointerPointer                      func(item *TimetzPointer) error
	InsertParameterTimestampTime                              func(sid string, source string, var_timestamp time) (int, error)
	InsertParameterTimestampTimeOriginalPointer               func(sid string, source string, var_timestamp time) (*int, error)
	InsertParameterTimestampTimePointerOriginal               func(sid *string, source *string, var_timestamp *time) (int, error)
	InsertParameterTimestampTimePointerPointer                func(sid string, source string, var_timestamp *time) (*int, error)
	InsertEntityTimestampTime                                 func(item TimestampOriginal) error
	InsertEntityTimestampTimeOriginalPointer                  func(item TimestampOriginal) error
	InsertEntityTimestampTimePointerOriginal                  func(item *TimestampPointer) error
	InsertEntityTimestampTimePointerPointer                   func(item *TimestampPointer) error
	InsertParameterTimestampWithTimezoneTime                  func(sid string, source string, var_timestamp_with_timezone time) (int, error)
	InsertParameterTimestampWithTimezoneTimeOriginalPointer   func(sid string, source string, var_timestamp_with_timezone time) (*int, error)
	InsertParameterTimestampWithTimezoneTimePointerOriginal   func(sid *string, source *string, var_timestamp_with_timezone *time) (int, error)
	InsertParameterTimestampWithTimezoneTimePointerPointer    func(sid string, source string, var_timestamp_with_timezone *time) (*int, error)
	InsertEntityTimestampWithTimezoneTime                     func(item TimestampWithTimezoneOriginal) error
	InsertEntityTimestampWithTimezoneTimeOriginalPointer      func(item TimestampWithTimezoneOriginal) error
	InsertEntityTimestampWithTimezoneTimePointerOriginal      func(item *TimestampWithTimezonePointer) error
	InsertEntityTimestampWithTimezoneTimePointerPointer       func(item *TimestampWithTimezonePointer) error
	InsertParameterTimestamptzTime                            func(sid string, source string, var_timestamptz time) (int, error)
	InsertParameterTimestamptzTimeOriginalPointer             func(sid string, source string, var_timestamptz time) (*int, error)
	InsertParameterTimestamptzTimePointerOriginal             func(sid *string, source *string, var_timestamptz *time) (int, error)
	InsertParameterTimestamptzTimePointerPointer              func(sid string, source string, var_timestamptz *time) (*int, error)
	InsertEntityTimestamptzTime                               func(item TimestamptzOriginal) error
	InsertEntityTimestamptzTimeOriginalPointer                func(item TimestamptzOriginal) error
	InsertEntityTimestamptzTimePointerOriginal                func(item *TimestamptzPointer) error
	InsertEntityTimestamptzTimePointerPointer                 func(item *TimestamptzPointer) error
	SelectParameterBigintInt64                                func(sid string) (int64, error)
	SelectParameterBigintInt64OriginalPointer                 func(sid string) (*int64, error)
	SelectParameterBigintInt64PointerOriginal                 func(sid string) (int64, error)
	SelectParameterBigintInt64PointerPointer                  func(sid string) (*int64, error)
	SelectEntityBigintInt64                                   func(sid string) (*int64, error)
	SelectEntityBigintInt64OriginalPointer                    func(sid string) (*int64, error)
	SelectEntityBigintInt64PointerOriginal                    func(sid string) (*int64, error)
	SelectEntityBigintInt64PointerPointer                     func(sid string) (*int64, error)
	SelectArrayParameterBigintInt64                           func(sid string) ([]int64, error)
	SelectArrayParameterBigintInt64OriginalPointer            func(sid string) ([]*int64, error)
	SelectArrayParameterBigintInt64PointerOriginal            func(sid string) ([]int64, error)
	SelectArrayParameterBigintInt64PointerPointer             func(sid string) ([]*int64, error)
	SelectArrayEntityBigintInt64                              func(sid string) ([]*int64, error)
	SelectArrayEntityBigintInt64OriginalPointer               func(sid string) ([]*int64, error)
	SelectArrayEntityBigintInt64PointerOriginal               func(sid string) ([]*int64, error)
	SelectArrayEntityBigintInt64PointerPointer                func(sid string) ([]*int64, error)
	SelectParameterInt8Int8                                   func(sid string) (int8, error)
	SelectParameterInt8Int8OriginalPointer                    func(sid string) (*int8, error)
	SelectParameterInt8Int8PointerOriginal                    func(sid string) (int8, error)
	SelectParameterInt8Int8PointerPointer                     func(sid string) (*int8, error)
	SelectEntityInt8Int8                                      func(sid string) (*int8, error)
	SelectEntityInt8Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectEntityInt8Int8PointerOriginal                       func(sid string) (*int8, error)
	SelectEntityInt8Int8PointerPointer                        func(sid string) (*int8, error)
	SelectArrayParameterInt8Int8                              func(sid string) ([]int8, error)
	SelectArrayParameterInt8Int8OriginalPointer               func(sid string) ([]*int8, error)
	SelectArrayParameterInt8Int8PointerOriginal               func(sid string) ([]int8, error)
	SelectArrayParameterInt8Int8PointerPointer                func(sid string) ([]*int8, error)
	SelectArrayEntityInt8Int8                                 func(sid string) ([]*int8, error)
	SelectArrayEntityInt8Int8OriginalPointer                  func(sid string) ([]*int8, error)
	SelectArrayEntityInt8Int8PointerOriginal                  func(sid string) ([]*int8, error)
	SelectArrayEntityInt8Int8PointerPointer                   func(sid string) ([]*int8, error)
	SelectParameterBooleanBool                                func(sid string) (bool, error)
	SelectParameterBooleanBoolOriginalPointer                 func(sid string) (*bool, error)
	SelectParameterBooleanBoolPointerOriginal                 func(sid string) (bool, error)
	SelectParameterBooleanBoolPointerPointer                  func(sid string) (*bool, error)
	SelectEntityBooleanBool                                   func(sid string) (*bool, error)
	SelectEntityBooleanBoolOriginalPointer                    func(sid string) (*bool, error)
	SelectEntityBooleanBoolPointerOriginal                    func(sid string) (*bool, error)
	SelectEntityBooleanBoolPointerPointer                     func(sid string) (*bool, error)
	SelectArrayParameterBooleanBool                           func(sid string) ([]bool, error)
	SelectArrayParameterBooleanBoolOriginalPointer            func(sid string) ([]*bool, error)
	SelectArrayParameterBooleanBoolPointerOriginal            func(sid string) ([]bool, error)
	SelectArrayParameterBooleanBoolPointerPointer             func(sid string) ([]*bool, error)
	SelectArrayEntityBooleanBool                              func(sid string) ([]*bool, error)
	SelectArrayEntityBooleanBoolOriginalPointer               func(sid string) ([]*bool, error)
	SelectArrayEntityBooleanBoolPointerOriginal               func(sid string) ([]*bool, error)
	SelectArrayEntityBooleanBoolPointerPointer                func(sid string) ([]*bool, error)
	SelectParameterBoolBool                                   func(sid string) (bool, error)
	SelectParameterBoolBoolOriginalPointer                    func(sid string) (*bool, error)
	SelectParameterBoolBoolPointerOriginal                    func(sid string) (bool, error)
	SelectParameterBoolBoolPointerPointer                     func(sid string) (*bool, error)
	SelectEntityBoolBool                                      func(sid string) (*bool, error)
	SelectEntityBoolBoolOriginalPointer                       func(sid string) (*bool, error)
	SelectEntityBoolBoolPointerOriginal                       func(sid string) (*bool, error)
	SelectEntityBoolBoolPointerPointer                        func(sid string) (*bool, error)
	SelectArrayParameterBoolBool                              func(sid string) ([]bool, error)
	SelectArrayParameterBoolBoolOriginalPointer               func(sid string) ([]*bool, error)
	SelectArrayParameterBoolBoolPointerOriginal               func(sid string) ([]bool, error)
	SelectArrayParameterBoolBoolPointerPointer                func(sid string) ([]*bool, error)
	SelectArrayEntityBoolBool                                 func(sid string) ([]*bool, error)
	SelectArrayEntityBoolBoolOriginalPointer                  func(sid string) ([]*bool, error)
	SelectArrayEntityBoolBoolPointerOriginal                  func(sid string) ([]*bool, error)
	SelectArrayEntityBoolBoolPointerPointer                   func(sid string) ([]*bool, error)
	SelectParameterCharacterString                            func(sid string) (string, error)
	SelectParameterCharacterStringOriginalPointer             func(sid string) (*string, error)
	SelectParameterCharacterStringPointerOriginal             func(sid string) (string, error)
	SelectParameterCharacterStringPointerPointer              func(sid string) (*string, error)
	SelectEntityCharacterString                               func(sid string) (*string, error)
	SelectEntityCharacterStringOriginalPointer                func(sid string) (*string, error)
	SelectEntityCharacterStringPointerOriginal                func(sid string) (*string, error)
	SelectEntityCharacterStringPointerPointer                 func(sid string) (*string, error)
	SelectArrayParameterCharacterString                       func(sid string) ([]string, error)
	SelectArrayParameterCharacterStringOriginalPointer        func(sid string) ([]*string, error)
	SelectArrayParameterCharacterStringPointerOriginal        func(sid string) ([]string, error)
	SelectArrayParameterCharacterStringPointerPointer         func(sid string) ([]*string, error)
	SelectArrayEntityCharacterString                          func(sid string) ([]*string, error)
	SelectArrayEntityCharacterStringOriginalPointer           func(sid string) ([]*string, error)
	SelectArrayEntityCharacterStringPointerOriginal           func(sid string) ([]*string, error)
	SelectArrayEntityCharacterStringPointerPointer            func(sid string) ([]*string, error)
	SelectParameterCharString                                 func(sid string) (string, error)
	SelectParameterCharStringOriginalPointer                  func(sid string) (*string, error)
	SelectParameterCharStringPointerOriginal                  func(sid string) (string, error)
	SelectParameterCharStringPointerPointer                   func(sid string) (*string, error)
	SelectEntityCharString                                    func(sid string) (*string, error)
	SelectEntityCharStringOriginalPointer                     func(sid string) (*string, error)
	SelectEntityCharStringPointerOriginal                     func(sid string) (*string, error)
	SelectEntityCharStringPointerPointer                      func(sid string) (*string, error)
	SelectArrayParameterCharString                            func(sid string) ([]string, error)
	SelectArrayParameterCharStringOriginalPointer             func(sid string) ([]*string, error)
	SelectArrayParameterCharStringPointerOriginal             func(sid string) ([]string, error)
	SelectArrayParameterCharStringPointerPointer              func(sid string) ([]*string, error)
	SelectArrayEntityCharString                               func(sid string) ([]*string, error)
	SelectArrayEntityCharStringOriginalPointer                func(sid string) ([]*string, error)
	SelectArrayEntityCharStringPointerOriginal                func(sid string) ([]*string, error)
	SelectArrayEntityCharStringPointerPointer                 func(sid string) ([]*string, error)
	SelectParameterCharacterVaryingString                     func(sid string) (string, error)
	SelectParameterCharacterVaryingStringOriginalPointer      func(sid string) (*string, error)
	SelectParameterCharacterVaryingStringPointerOriginal      func(sid string) (string, error)
	SelectParameterCharacterVaryingStringPointerPointer       func(sid string) (*string, error)
	SelectEntityCharacterVaryingString                        func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringOriginalPointer         func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringPointerOriginal         func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringPointerPointer          func(sid string) (*string, error)
	SelectArrayParameterCharacterVaryingString                func(sid string) ([]string, error)
	SelectArrayParameterCharacterVaryingStringOriginalPointer func(sid string) ([]*string, error)
	SelectArrayParameterCharacterVaryingStringPointerOriginal func(sid string) ([]string, error)
	SelectArrayParameterCharacterVaryingStringPointerPointer  func(sid string) ([]*string, error)
	SelectArrayEntityCharacterVaryingString                   func(sid string) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringOriginalPointer    func(sid string) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringPointerOriginal    func(sid string) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringPointerPointer     func(sid string) ([]*string, error)
	SelectParameterVarcharString                              func(sid string) (string, error)
	SelectParameterVarcharStringOriginalPointer               func(sid string) (*string, error)
	SelectParameterVarcharStringPointerOriginal               func(sid string) (string, error)
	SelectParameterVarcharStringPointerPointer                func(sid string) (*string, error)
	SelectEntityVarcharString                                 func(sid string) (*string, error)
	SelectEntityVarcharStringOriginalPointer                  func(sid string) (*string, error)
	SelectEntityVarcharStringPointerOriginal                  func(sid string) (*string, error)
	SelectEntityVarcharStringPointerPointer                   func(sid string) (*string, error)
	SelectArrayParameterVarcharString                         func(sid string) ([]string, error)
	SelectArrayParameterVarcharStringOriginalPointer          func(sid string) ([]*string, error)
	SelectArrayParameterVarcharStringPointerOriginal          func(sid string) ([]string, error)
	SelectArrayParameterVarcharStringPointerPointer           func(sid string) ([]*string, error)
	SelectArrayEntityVarcharString                            func(sid string) ([]*string, error)
	SelectArrayEntityVarcharStringOriginalPointer             func(sid string) ([]*string, error)
	SelectArrayEntityVarcharStringPointerOriginal             func(sid string) ([]*string, error)
	SelectArrayEntityVarcharStringPointerPointer              func(sid string) ([]*string, error)
	SelectParameterFloat8Float32                              func(sid string) (float32, error)
	SelectParameterFloat8Float32OriginalPointer               func(sid string) (*float32, error)
	SelectParameterFloat8Float32PointerOriginal               func(sid string) (float32, error)
	SelectParameterFloat8Float32PointerPointer                func(sid string) (*float32, error)
	SelectEntityFloat8Float32                                 func(sid string) (*float32, error)
	SelectEntityFloat8Float32OriginalPointer                  func(sid string) (*float32, error)
	SelectEntityFloat8Float32PointerOriginal                  func(sid string) (*float32, error)
	SelectEntityFloat8Float32PointerPointer                   func(sid string) (*float32, error)
	SelectArrayParameterFloat8Float32                         func(sid string) ([]float32, error)
	SelectArrayParameterFloat8Float32OriginalPointer          func(sid string) ([]*float32, error)
	SelectArrayParameterFloat8Float32PointerOriginal          func(sid string) ([]float32, error)
	SelectArrayParameterFloat8Float32PointerPointer           func(sid string) ([]*float32, error)
	SelectArrayEntityFloat8Float32                            func(sid string) ([]*float32, error)
	SelectArrayEntityFloat8Float32OriginalPointer             func(sid string) ([]*float32, error)
	SelectArrayEntityFloat8Float32PointerOriginal             func(sid string) ([]*float32, error)
	SelectArrayEntityFloat8Float32PointerPointer              func(sid string) ([]*float32, error)
	SelectParameterIntegerInt8                                func(sid string) (int8, error)
	SelectParameterIntegerInt8OriginalPointer                 func(sid string) (*int8, error)
	SelectParameterIntegerInt8PointerOriginal                 func(sid string) (int8, error)
	SelectParameterIntegerInt8PointerPointer                  func(sid string) (*int8, error)
	SelectEntityIntegerInt8                                   func(sid string) (*int8, error)
	SelectEntityIntegerInt8OriginalPointer                    func(sid string) (*int8, error)
	SelectEntityIntegerInt8PointerOriginal                    func(sid string) (*int8, error)
	SelectEntityIntegerInt8PointerPointer                     func(sid string) (*int8, error)
	SelectArrayParameterIntegerInt8                           func(sid string) ([]int8, error)
	SelectArrayParameterIntegerInt8OriginalPointer            func(sid string) ([]*int8, error)
	SelectArrayParameterIntegerInt8PointerOriginal            func(sid string) ([]int8, error)
	SelectArrayParameterIntegerInt8PointerPointer             func(sid string) ([]*int8, error)
	SelectArrayEntityIntegerInt8                              func(sid string) ([]*int8, error)
	SelectArrayEntityIntegerInt8OriginalPointer               func(sid string) ([]*int8, error)
	SelectArrayEntityIntegerInt8PointerOriginal               func(sid string) ([]*int8, error)
	SelectArrayEntityIntegerInt8PointerPointer                func(sid string) ([]*int8, error)
	SelectParameterIntInt8                                    func(sid string) (int8, error)
	SelectParameterIntInt8OriginalPointer                     func(sid string) (*int8, error)
	SelectParameterIntInt8PointerOriginal                     func(sid string) (int8, error)
	SelectParameterIntInt8PointerPointer                      func(sid string) (*int8, error)
	SelectEntityIntInt8                                       func(sid string) (*int8, error)
	SelectEntityIntInt8OriginalPointer                        func(sid string) (*int8, error)
	SelectEntityIntInt8PointerOriginal                        func(sid string) (*int8, error)
	SelectEntityIntInt8PointerPointer                         func(sid string) (*int8, error)
	SelectArrayParameterIntInt8                               func(sid string) ([]int8, error)
	SelectArrayParameterIntInt8OriginalPointer                func(sid string) ([]*int8, error)
	SelectArrayParameterIntInt8PointerOriginal                func(sid string) ([]int8, error)
	SelectArrayParameterIntInt8PointerPointer                 func(sid string) ([]*int8, error)
	SelectArrayEntityIntInt8                                  func(sid string) ([]*int8, error)
	SelectArrayEntityIntInt8OriginalPointer                   func(sid string) ([]*int8, error)
	SelectArrayEntityIntInt8PointerOriginal                   func(sid string) ([]*int8, error)
	SelectArrayEntityIntInt8PointerPointer                    func(sid string) ([]*int8, error)
	SelectParameterInt4Int8                                   func(sid string) (int8, error)
	SelectParameterInt4Int8OriginalPointer                    func(sid string) (*int8, error)
	SelectParameterInt4Int8PointerOriginal                    func(sid string) (int8, error)
	SelectParameterInt4Int8PointerPointer                     func(sid string) (*int8, error)
	SelectEntityInt4Int8                                      func(sid string) (*int8, error)
	SelectEntityInt4Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectEntityInt4Int8PointerOriginal                       func(sid string) (*int8, error)
	SelectEntityInt4Int8PointerPointer                        func(sid string) (*int8, error)
	SelectArrayParameterInt4Int8                              func(sid string) ([]int8, error)
	SelectArrayParameterInt4Int8OriginalPointer               func(sid string) ([]*int8, error)
	SelectArrayParameterInt4Int8PointerOriginal               func(sid string) ([]int8, error)
	SelectArrayParameterInt4Int8PointerPointer                func(sid string) ([]*int8, error)
	SelectArrayEntityInt4Int8                                 func(sid string) ([]*int8, error)
	SelectArrayEntityInt4Int8OriginalPointer                  func(sid string) ([]*int8, error)
	SelectArrayEntityInt4Int8PointerOriginal                  func(sid string) ([]*int8, error)
	SelectArrayEntityInt4Int8PointerPointer                   func(sid string) ([]*int8, error)
	SelectParameterNumericDecimal                             func(sid string) (decimal, error)
	SelectParameterNumericDecimalOriginalPointer              func(sid string) (*decimal, error)
	SelectParameterNumericDecimalPointerOriginal              func(sid string) (decimal, error)
	SelectParameterNumericDecimalPointerPointer               func(sid string) (*decimal, error)
	SelectEntityNumericDecimal                                func(sid string) (*decimal, error)
	SelectEntityNumericDecimalOriginalPointer                 func(sid string) (*decimal, error)
	SelectEntityNumericDecimalPointerOriginal                 func(sid string) (*decimal, error)
	SelectEntityNumericDecimalPointerPointer                  func(sid string) (*decimal, error)
	SelectArrayParameterNumericDecimal                        func(sid string) ([]decimal, error)
	SelectArrayParameterNumericDecimalOriginalPointer         func(sid string) ([]*decimal, error)
	SelectArrayParameterNumericDecimalPointerOriginal         func(sid string) ([]decimal, error)
	SelectArrayParameterNumericDecimalPointerPointer          func(sid string) ([]*decimal, error)
	SelectArrayEntityNumericDecimal                           func(sid string) ([]*decimal, error)
	SelectArrayEntityNumericDecimalOriginalPointer            func(sid string) ([]*decimal, error)
	SelectArrayEntityNumericDecimalPointerOriginal            func(sid string) ([]*decimal, error)
	SelectArrayEntityNumericDecimalPointerPointer             func(sid string) ([]*decimal, error)
	SelectParameterDecimalDecimal                             func(sid string) (decimal, error)
	SelectParameterDecimalDecimalOriginalPointer              func(sid string) (*decimal, error)
	SelectParameterDecimalDecimalPointerOriginal              func(sid string) (decimal, error)
	SelectParameterDecimalDecimalPointerPointer               func(sid string) (*decimal, error)
	SelectEntityDecimalDecimal                                func(sid string) (*decimal, error)
	SelectEntityDecimalDecimalOriginalPointer                 func(sid string) (*decimal, error)
	SelectEntityDecimalDecimalPointerOriginal                 func(sid string) (*decimal, error)
	SelectEntityDecimalDecimalPointerPointer                  func(sid string) (*decimal, error)
	SelectArrayParameterDecimalDecimal                        func(sid string) ([]decimal, error)
	SelectArrayParameterDecimalDecimalOriginalPointer         func(sid string) ([]*decimal, error)
	SelectArrayParameterDecimalDecimalPointerOriginal         func(sid string) ([]decimal, error)
	SelectArrayParameterDecimalDecimalPointerPointer          func(sid string) ([]*decimal, error)
	SelectArrayEntityDecimalDecimal                           func(sid string) ([]*decimal, error)
	SelectArrayEntityDecimalDecimalOriginalPointer            func(sid string) ([]*decimal, error)
	SelectArrayEntityDecimalDecimalPointerOriginal            func(sid string) ([]*decimal, error)
	SelectArrayEntityDecimalDecimalPointerPointer             func(sid string) ([]*decimal, error)
	SelectParameterFloat4Float32                              func(sid string) (float32, error)
	SelectParameterFloat4Float32OriginalPointer               func(sid string) (*float32, error)
	SelectParameterFloat4Float32PointerOriginal               func(sid string) (float32, error)
	SelectParameterFloat4Float32PointerPointer                func(sid string) (*float32, error)
	SelectEntityFloat4Float32                                 func(sid string) (*float32, error)
	SelectEntityFloat4Float32OriginalPointer                  func(sid string) (*float32, error)
	SelectEntityFloat4Float32PointerOriginal                  func(sid string) (*float32, error)
	SelectEntityFloat4Float32PointerPointer                   func(sid string) (*float32, error)
	SelectArrayParameterFloat4Float32                         func(sid string) ([]float32, error)
	SelectArrayParameterFloat4Float32OriginalPointer          func(sid string) ([]*float32, error)
	SelectArrayParameterFloat4Float32PointerOriginal          func(sid string) ([]float32, error)
	SelectArrayParameterFloat4Float32PointerPointer           func(sid string) ([]*float32, error)
	SelectArrayEntityFloat4Float32                            func(sid string) ([]*float32, error)
	SelectArrayEntityFloat4Float32OriginalPointer             func(sid string) ([]*float32, error)
	SelectArrayEntityFloat4Float32PointerOriginal             func(sid string) ([]*float32, error)
	SelectArrayEntityFloat4Float32PointerPointer              func(sid string) ([]*float32, error)
	SelectParameterSmallintInt8                               func(sid string) (int8, error)
	SelectParameterSmallintInt8OriginalPointer                func(sid string) (*int8, error)
	SelectParameterSmallintInt8PointerOriginal                func(sid string) (int8, error)
	SelectParameterSmallintInt8PointerPointer                 func(sid string) (*int8, error)
	SelectEntitySmallintInt8                                  func(sid string) (*int8, error)
	SelectEntitySmallintInt8OriginalPointer                   func(sid string) (*int8, error)
	SelectEntitySmallintInt8PointerOriginal                   func(sid string) (*int8, error)
	SelectEntitySmallintInt8PointerPointer                    func(sid string) (*int8, error)
	SelectArrayParameterSmallintInt8                          func(sid string) ([]int8, error)
	SelectArrayParameterSmallintInt8OriginalPointer           func(sid string) ([]*int8, error)
	SelectArrayParameterSmallintInt8PointerOriginal           func(sid string) ([]int8, error)
	SelectArrayParameterSmallintInt8PointerPointer            func(sid string) ([]*int8, error)
	SelectArrayEntitySmallintInt8                             func(sid string) ([]*int8, error)
	SelectArrayEntitySmallintInt8OriginalPointer              func(sid string) ([]*int8, error)
	SelectArrayEntitySmallintInt8PointerOriginal              func(sid string) ([]*int8, error)
	SelectArrayEntitySmallintInt8PointerPointer               func(sid string) ([]*int8, error)
	SelectParameterInt2Int8                                   func(sid string) (int8, error)
	SelectParameterInt2Int8OriginalPointer                    func(sid string) (*int8, error)
	SelectParameterInt2Int8PointerOriginal                    func(sid string) (int8, error)
	SelectParameterInt2Int8PointerPointer                     func(sid string) (*int8, error)
	SelectEntityInt2Int8                                      func(sid string) (*int8, error)
	SelectEntityInt2Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectEntityInt2Int8PointerOriginal                       func(sid string) (*int8, error)
	SelectEntityInt2Int8PointerPointer                        func(sid string) (*int8, error)
	SelectArrayParameterInt2Int8                              func(sid string) ([]int8, error)
	SelectArrayParameterInt2Int8OriginalPointer               func(sid string) ([]*int8, error)
	SelectArrayParameterInt2Int8PointerOriginal               func(sid string) ([]int8, error)
	SelectArrayParameterInt2Int8PointerPointer                func(sid string) ([]*int8, error)
	SelectArrayEntityInt2Int8                                 func(sid string) ([]*int8, error)
	SelectArrayEntityInt2Int8OriginalPointer                  func(sid string) ([]*int8, error)
	SelectArrayEntityInt2Int8PointerOriginal                  func(sid string) ([]*int8, error)
	SelectArrayEntityInt2Int8PointerPointer                   func(sid string) ([]*int8, error)
	SelectParameterTextString                                 func(sid string) (string, error)
	SelectParameterTextStringOriginalPointer                  func(sid string) (*string, error)
	SelectParameterTextStringPointerOriginal                  func(sid string) (string, error)
	SelectParameterTextStringPointerPointer                   func(sid string) (*string, error)
	SelectEntityTextString                                    func(sid string) (*string, error)
	SelectEntityTextStringOriginalPointer                     func(sid string) (*string, error)
	SelectEntityTextStringPointerOriginal                     func(sid string) (*string, error)
	SelectEntityTextStringPointerPointer                      func(sid string) (*string, error)
	SelectParameterTimeTime                                   func(sid string) (time, error)
	SelectParameterTimeTimeOriginalPointer                    func(sid string) (*time, error)
	SelectParameterTimeTimePointerOriginal                    func(sid string) (time, error)
	SelectParameterTimeTimePointerPointer                     func(sid string) (*time, error)
	SelectEntityTimeTime                                      func(sid string) (*time, error)
	SelectEntityTimeTimeOriginalPointer                       func(sid string) (*time, error)
	SelectEntityTimeTimePointerOriginal                       func(sid string) (*time, error)
	SelectEntityTimeTimePointerPointer                        func(sid string) (*time, error)
	SelectParameterTimeWithTimezoneTime                       func(sid string) (time, error)
	SelectParameterTimeWithTimezoneTimeOriginalPointer        func(sid string) (*time, error)
	SelectParameterTimeWithTimezoneTimePointerOriginal        func(sid string) (time, error)
	SelectParameterTimeWithTimezoneTimePointerPointer         func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTime                          func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTimeOriginalPointer           func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTimePointerOriginal           func(sid string) (*time, error)
	SelectEntityTimeWithTimezoneTimePointerPointer            func(sid string) (*time, error)
	SelectParameterTimetzTime                                 func(sid string) (time, error)
	SelectParameterTimetzTimeOriginalPointer                  func(sid string) (*time, error)
	SelectParameterTimetzTimePointerOriginal                  func(sid string) (time, error)
	SelectParameterTimetzTimePointerPointer                   func(sid string) (*time, error)
	SelectEntityTimetzTime                                    func(sid string) (*time, error)
	SelectEntityTimetzTimeOriginalPointer                     func(sid string) (*time, error)
	SelectEntityTimetzTimePointerOriginal                     func(sid string) (*time, error)
	SelectEntityTimetzTimePointerPointer                      func(sid string) (*time, error)
	SelectParameterTimestampTime                              func(sid string) (time, error)
	SelectParameterTimestampTimeOriginalPointer               func(sid string) (*time, error)
	SelectParameterTimestampTimePointerOriginal               func(sid string) (time, error)
	SelectParameterTimestampTimePointerPointer                func(sid string) (*time, error)
	SelectEntityTimestampTime                                 func(sid string) (*time, error)
	SelectEntityTimestampTimeOriginalPointer                  func(sid string) (*time, error)
	SelectEntityTimestampTimePointerOriginal                  func(sid string) (*time, error)
	SelectEntityTimestampTimePointerPointer                   func(sid string) (*time, error)
	SelectParameterTimestampWithTimezoneTime                  func(sid string) (time, error)
	SelectParameterTimestampWithTimezoneTimeOriginalPointer   func(sid string) (*time, error)
	SelectParameterTimestampWithTimezoneTimePointerOriginal   func(sid string) (time, error)
	SelectParameterTimestampWithTimezoneTimePointerPointer    func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTime                     func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTimeOriginalPointer      func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTimePointerOriginal      func(sid string) (*time, error)
	SelectEntityTimestampWithTimezoneTimePointerPointer       func(sid string) (*time, error)
	SelectParameterTimestamptzTime                            func(sid string) (time, error)
	SelectParameterTimestamptzTimeOriginalPointer             func(sid string) (*time, error)
	SelectParameterTimestamptzTimePointerOriginal             func(sid string) (time, error)
	SelectParameterTimestamptzTimePointerPointer              func(sid string) (*time, error)
	SelectEntityTimestamptzTime                               func(sid string) (*time, error)
	SelectEntityTimestamptzTimeOriginalPointer                func(sid string) (*time, error)
	SelectEntityTimestamptzTimePointerOriginal                func(sid string) (*time, error)
	SelectEntityTimestamptzTimePointerPointer                 func(sid string) (*time, error)
	UpdateParameterBigintInt64                                func(sid string, source string, var_bigint int64) (int, error)
	UpdateParameterBigintInt64OriginalPointer                 func(sid string, source string, var_bigint int64) (*int, error)
	UpdateParameterBigintInt64PointerOriginal                 func(sid *string, source *string, var_bigint *int64) (int, error)
	UpdateParameterBigintInt64PointerPointer                  func(sid string, source string, var_bigint *int64) (*int, error)
	UpdateEntityBigintInt64                                   func(item BigintOriginal) error
	UpdateEntityBigintInt64OriginalPointer                    func(item BigintOriginal) error
	UpdateEntityBigintInt64PointerOriginal                    func(item *BigintPointer) error
	UpdateEntityBigintInt64PointerPointer                     func(item *BigintPointer) error
	UpdateArrayParameterBigintInt64                           func(sid string, source string, items []int64) (int, error)
	UpdateArrayParameterBigintInt64OriginalPointer            func(sid string, source string, items []int64) (*int, error)
	UpdateArrayParameterBigintInt64PointerOriginal            func(sid string, source string, items []*int64) (int, error)
	UpdateArrayParameterBigintInt64PointerPointer             func(sid string, source string, items []*int64) (*int, error)
	UpdateArrayEntityBigintInt64                              func(item BigintOriginal) (int, error)
	UpdateArrayEntityBigintInt64OriginalPointer               func(item BigintPointer) (int, error)
	UpdateArrayEntityBigintInt64PointerOriginal               func(item BigintOriginal) (int, error)
	UpdateArrayEntityBigintInt64PointerPointer                func(item BigintPointer) (int, error)
	UpdateParameterInt8Int8                                   func(sid string, source string, var_int8 int8) (int, error)
	UpdateParameterInt8Int8OriginalPointer                    func(sid string, source string, var_int8 int8) (*int, error)
	UpdateParameterInt8Int8PointerOriginal                    func(sid *string, source *string, var_int8 *int8) (int, error)
	UpdateParameterInt8Int8PointerPointer                     func(sid string, source string, var_int8 *int8) (*int, error)
	UpdateEntityInt8Int8                                      func(item Int8Original) error
	UpdateEntityInt8Int8OriginalPointer                       func(item Int8Original) error
	UpdateEntityInt8Int8PointerOriginal                       func(item *Int8Pointer) error
	UpdateEntityInt8Int8PointerPointer                        func(item *Int8Pointer) error
	UpdateArrayParameterInt8Int8                              func(sid string, source string, items []int8) (int, error)
	UpdateArrayParameterInt8Int8OriginalPointer               func(sid string, source string, items []int8) (*int, error)
	UpdateArrayParameterInt8Int8PointerOriginal               func(sid string, source string, items []*int8) (int, error)
	UpdateArrayParameterInt8Int8PointerPointer                func(sid string, source string, items []*int8) (*int, error)
	UpdateArrayEntityInt8Int8                                 func(item Int8Original) (int, error)
	UpdateArrayEntityInt8Int8OriginalPointer                  func(item Int8Pointer) (int, error)
	UpdateArrayEntityInt8Int8PointerOriginal                  func(item Int8Original) (int, error)
	UpdateArrayEntityInt8Int8PointerPointer                   func(item Int8Pointer) (int, error)
	UpdateParameterBooleanBool                                func(sid string, source string, var_boolean bool) (int, error)
	UpdateParameterBooleanBoolOriginalPointer                 func(sid string, source string, var_boolean bool) (*int, error)
	UpdateParameterBooleanBoolPointerOriginal                 func(sid *string, source *string, var_boolean *bool) (int, error)
	UpdateParameterBooleanBoolPointerPointer                  func(sid string, source string, var_boolean *bool) (*int, error)
	UpdateEntityBooleanBool                                   func(item BooleanOriginal) error
	UpdateEntityBooleanBoolOriginalPointer                    func(item BooleanOriginal) error
	UpdateEntityBooleanBoolPointerOriginal                    func(item *BooleanPointer) error
	UpdateEntityBooleanBoolPointerPointer                     func(item *BooleanPointer) error
	UpdateArrayParameterBooleanBool                           func(sid string, source string, items []bool) (int, error)
	UpdateArrayParameterBooleanBoolOriginalPointer            func(sid string, source string, items []bool) (*int, error)
	UpdateArrayParameterBooleanBoolPointerOriginal            func(sid string, source string, items []*bool) (int, error)
	UpdateArrayParameterBooleanBoolPointerPointer             func(sid string, source string, items []*bool) (*int, error)
	UpdateArrayEntityBooleanBool                              func(item BooleanOriginal) (int, error)
	UpdateArrayEntityBooleanBoolOriginalPointer               func(item BooleanPointer) (int, error)
	UpdateArrayEntityBooleanBoolPointerOriginal               func(item BooleanOriginal) (int, error)
	UpdateArrayEntityBooleanBoolPointerPointer                func(item BooleanPointer) (int, error)
	UpdateParameterBoolBool                                   func(sid string, source string, var_bool bool) (int, error)
	UpdateParameterBoolBoolOriginalPointer                    func(sid string, source string, var_bool bool) (*int, error)
	UpdateParameterBoolBoolPointerOriginal                    func(sid *string, source *string, var_bool *bool) (int, error)
	UpdateParameterBoolBoolPointerPointer                     func(sid string, source string, var_bool *bool) (*int, error)
	UpdateEntityBoolBool                                      func(item BoolOriginal) error
	UpdateEntityBoolBoolOriginalPointer                       func(item BoolOriginal) error
	UpdateEntityBoolBoolPointerOriginal                       func(item *BoolPointer) error
	UpdateEntityBoolBoolPointerPointer                        func(item *BoolPointer) error
	UpdateArrayParameterBoolBool                              func(sid string, source string, items []bool) (int, error)
	UpdateArrayParameterBoolBoolOriginalPointer               func(sid string, source string, items []bool) (*int, error)
	UpdateArrayParameterBoolBoolPointerOriginal               func(sid string, source string, items []*bool) (int, error)
	UpdateArrayParameterBoolBoolPointerPointer                func(sid string, source string, items []*bool) (*int, error)
	UpdateArrayEntityBoolBool                                 func(item BoolOriginal) (int, error)
	UpdateArrayEntityBoolBoolOriginalPointer                  func(item BoolPointer) (int, error)
	UpdateArrayEntityBoolBoolPointerOriginal                  func(item BoolOriginal) (int, error)
	UpdateArrayEntityBoolBoolPointerPointer                   func(item BoolPointer) (int, error)
	UpdateParameterCharacterString                            func(sid string, source string, var_character string) (int, error)
	UpdateParameterCharacterStringOriginalPointer             func(sid string, source string, var_character string) (*int, error)
	UpdateParameterCharacterStringPointerOriginal             func(sid *string, source *string, var_character *string) (int, error)
	UpdateParameterCharacterStringPointerPointer              func(sid string, source string, var_character *string) (*int, error)
	UpdateEntityCharacterString                               func(item CharacterOriginal) error
	UpdateEntityCharacterStringOriginalPointer                func(item CharacterOriginal) error
	UpdateEntityCharacterStringPointerOriginal                func(item *CharacterPointer) error
	UpdateEntityCharacterStringPointerPointer                 func(item *CharacterPointer) error
	UpdateArrayParameterCharacterString                       func(sid string, source string, items []string) (int, error)
	UpdateArrayParameterCharacterStringOriginalPointer        func(sid string, source string, items []string) (*int, error)
	UpdateArrayParameterCharacterStringPointerOriginal        func(sid string, source string, items []*string) (int, error)
	UpdateArrayParameterCharacterStringPointerPointer         func(sid string, source string, items []*string) (*int, error)
	UpdateArrayEntityCharacterString                          func(item CharacterOriginal) (int, error)
	UpdateArrayEntityCharacterStringOriginalPointer           func(item CharacterPointer) (int, error)
	UpdateArrayEntityCharacterStringPointerOriginal           func(item CharacterOriginal) (int, error)
	UpdateArrayEntityCharacterStringPointerPointer            func(item CharacterPointer) (int, error)
	UpdateParameterCharString                                 func(sid string, source string, var_char string) (int, error)
	UpdateParameterCharStringOriginalPointer                  func(sid string, source string, var_char string) (*int, error)
	UpdateParameterCharStringPointerOriginal                  func(sid *string, source *string, var_char *string) (int, error)
	UpdateParameterCharStringPointerPointer                   func(sid string, source string, var_char *string) (*int, error)
	UpdateEntityCharString                                    func(item CharOriginal) error
	UpdateEntityCharStringOriginalPointer                     func(item CharOriginal) error
	UpdateEntityCharStringPointerOriginal                     func(item *CharPointer) error
	UpdateEntityCharStringPointerPointer                      func(item *CharPointer) error
	UpdateArrayParameterCharString                            func(sid string, source string, items []string) (int, error)
	UpdateArrayParameterCharStringOriginalPointer             func(sid string, source string, items []string) (*int, error)
	UpdateArrayParameterCharStringPointerOriginal             func(sid string, source string, items []*string) (int, error)
	UpdateArrayParameterCharStringPointerPointer              func(sid string, source string, items []*string) (*int, error)
	UpdateArrayEntityCharString                               func(item CharOriginal) (int, error)
	UpdateArrayEntityCharStringOriginalPointer                func(item CharPointer) (int, error)
	UpdateArrayEntityCharStringPointerOriginal                func(item CharOriginal) (int, error)
	UpdateArrayEntityCharStringPointerPointer                 func(item CharPointer) (int, error)
	UpdateParameterCharacterVaryingString                     func(sid string, source string, var_character_varying string) (int, error)
	UpdateParameterCharacterVaryingStringOriginalPointer      func(sid string, source string, var_character_varying string) (*int, error)
	UpdateParameterCharacterVaryingStringPointerOriginal      func(sid *string, source *string, var_character_varying *string) (int, error)
	UpdateParameterCharacterVaryingStringPointerPointer       func(sid string, source string, var_character_varying *string) (*int, error)
	UpdateEntityCharacterVaryingString                        func(item CharacterVaryingOriginal) error
	UpdateEntityCharacterVaryingStringOriginalPointer         func(item CharacterVaryingOriginal) error
	UpdateEntityCharacterVaryingStringPointerOriginal         func(item *CharacterVaryingPointer) error
	UpdateEntityCharacterVaryingStringPointerPointer          func(item *CharacterVaryingPointer) error
	UpdateArrayParameterCharacterVaryingString                func(sid string, source string, items []string) (int, error)
	UpdateArrayParameterCharacterVaryingStringOriginalPointer func(sid string, source string, items []string) (*int, error)
	UpdateArrayParameterCharacterVaryingStringPointerOriginal func(sid string, source string, items []*string) (int, error)
	UpdateArrayParameterCharacterVaryingStringPointerPointer  func(sid string, source string, items []*string) (*int, error)
	UpdateArrayEntityCharacterVaryingString                   func(item CharacterVaryingOriginal) (int, error)
	UpdateArrayEntityCharacterVaryingStringOriginalPointer    func(item CharacterVaryingPointer) (int, error)
	UpdateArrayEntityCharacterVaryingStringPointerOriginal    func(item CharacterVaryingOriginal) (int, error)
	UpdateArrayEntityCharacterVaryingStringPointerPointer     func(item CharacterVaryingPointer) (int, error)
	UpdateParameterVarcharString                              func(sid string, source string, var_varchar string) (int, error)
	UpdateParameterVarcharStringOriginalPointer               func(sid string, source string, var_varchar string) (*int, error)
	UpdateParameterVarcharStringPointerOriginal               func(sid *string, source *string, var_varchar *string) (int, error)
	UpdateParameterVarcharStringPointerPointer                func(sid string, source string, var_varchar *string) (*int, error)
	UpdateEntityVarcharString                                 func(item VarcharOriginal) error
	UpdateEntityVarcharStringOriginalPointer                  func(item VarcharOriginal) error
	UpdateEntityVarcharStringPointerOriginal                  func(item *VarcharPointer) error
	UpdateEntityVarcharStringPointerPointer                   func(item *VarcharPointer) error
	UpdateArrayParameterVarcharString                         func(sid string, source string, items []string) (int, error)
	UpdateArrayParameterVarcharStringOriginalPointer          func(sid string, source string, items []string) (*int, error)
	UpdateArrayParameterVarcharStringPointerOriginal          func(sid string, source string, items []*string) (int, error)
	UpdateArrayParameterVarcharStringPointerPointer           func(sid string, source string, items []*string) (*int, error)
	UpdateArrayEntityVarcharString                            func(item VarcharOriginal) (int, error)
	UpdateArrayEntityVarcharStringOriginalPointer             func(item VarcharPointer) (int, error)
	UpdateArrayEntityVarcharStringPointerOriginal             func(item VarcharOriginal) (int, error)
	UpdateArrayEntityVarcharStringPointerPointer              func(item VarcharPointer) (int, error)
	UpdateParameterFloat8Float32                              func(sid string, source string, var_float8 float32) (int, error)
	UpdateParameterFloat8Float32OriginalPointer               func(sid string, source string, var_float8 float32) (*int, error)
	UpdateParameterFloat8Float32PointerOriginal               func(sid *string, source *string, var_float8 *float32) (int, error)
	UpdateParameterFloat8Float32PointerPointer                func(sid string, source string, var_float8 *float32) (*int, error)
	UpdateEntityFloat8Float32                                 func(item Float8Original) error
	UpdateEntityFloat8Float32OriginalPointer                  func(item Float8Original) error
	UpdateEntityFloat8Float32PointerOriginal                  func(item *Float8Pointer) error
	UpdateEntityFloat8Float32PointerPointer                   func(item *Float8Pointer) error
	UpdateArrayParameterFloat8Float32                         func(sid string, source string, items []float32) (int, error)
	UpdateArrayParameterFloat8Float32OriginalPointer          func(sid string, source string, items []float32) (*int, error)
	UpdateArrayParameterFloat8Float32PointerOriginal          func(sid string, source string, items []*float32) (int, error)
	UpdateArrayParameterFloat8Float32PointerPointer           func(sid string, source string, items []*float32) (*int, error)
	UpdateArrayEntityFloat8Float32                            func(item Float8Original) (int, error)
	UpdateArrayEntityFloat8Float32OriginalPointer             func(item Float8Pointer) (int, error)
	UpdateArrayEntityFloat8Float32PointerOriginal             func(item Float8Original) (int, error)
	UpdateArrayEntityFloat8Float32PointerPointer              func(item Float8Pointer) (int, error)
	UpdateParameterIntegerInt8                                func(sid string, source string, var_integer int8) (int, error)
	UpdateParameterIntegerInt8OriginalPointer                 func(sid string, source string, var_integer int8) (*int, error)
	UpdateParameterIntegerInt8PointerOriginal                 func(sid *string, source *string, var_integer *int8) (int, error)
	UpdateParameterIntegerInt8PointerPointer                  func(sid string, source string, var_integer *int8) (*int, error)
	UpdateEntityIntegerInt8                                   func(item IntegerOriginal) error
	UpdateEntityIntegerInt8OriginalPointer                    func(item IntegerOriginal) error
	UpdateEntityIntegerInt8PointerOriginal                    func(item *IntegerPointer) error
	UpdateEntityIntegerInt8PointerPointer                     func(item *IntegerPointer) error
	UpdateArrayParameterIntegerInt8                           func(sid string, source string, items []int8) (int, error)
	UpdateArrayParameterIntegerInt8OriginalPointer            func(sid string, source string, items []int8) (*int, error)
	UpdateArrayParameterIntegerInt8PointerOriginal            func(sid string, source string, items []*int8) (int, error)
	UpdateArrayParameterIntegerInt8PointerPointer             func(sid string, source string, items []*int8) (*int, error)
	UpdateArrayEntityIntegerInt8                              func(item IntegerOriginal) (int, error)
	UpdateArrayEntityIntegerInt8OriginalPointer               func(item IntegerPointer) (int, error)
	UpdateArrayEntityIntegerInt8PointerOriginal               func(item IntegerOriginal) (int, error)
	UpdateArrayEntityIntegerInt8PointerPointer                func(item IntegerPointer) (int, error)
	UpdateParameterIntInt8                                    func(sid string, source string, var_int int8) (int, error)
	UpdateParameterIntInt8OriginalPointer                     func(sid string, source string, var_int int8) (*int, error)
	UpdateParameterIntInt8PointerOriginal                     func(sid *string, source *string, var_int *int8) (int, error)
	UpdateParameterIntInt8PointerPointer                      func(sid string, source string, var_int *int8) (*int, error)
	UpdateEntityIntInt8                                       func(item IntOriginal) error
	UpdateEntityIntInt8OriginalPointer                        func(item IntOriginal) error
	UpdateEntityIntInt8PointerOriginal                        func(item *IntPointer) error
	UpdateEntityIntInt8PointerPointer                         func(item *IntPointer) error
	UpdateArrayParameterIntInt8                               func(sid string, source string, items []int8) (int, error)
	UpdateArrayParameterIntInt8OriginalPointer                func(sid string, source string, items []int8) (*int, error)
	UpdateArrayParameterIntInt8PointerOriginal                func(sid string, source string, items []*int8) (int, error)
	UpdateArrayParameterIntInt8PointerPointer                 func(sid string, source string, items []*int8) (*int, error)
	UpdateArrayEntityIntInt8                                  func(item IntOriginal) (int, error)
	UpdateArrayEntityIntInt8OriginalPointer                   func(item IntPointer) (int, error)
	UpdateArrayEntityIntInt8PointerOriginal                   func(item IntOriginal) (int, error)
	UpdateArrayEntityIntInt8PointerPointer                    func(item IntPointer) (int, error)
	UpdateParameterInt4Int8                                   func(sid string, source string, var_int4 int8) (int, error)
	UpdateParameterInt4Int8OriginalPointer                    func(sid string, source string, var_int4 int8) (*int, error)
	UpdateParameterInt4Int8PointerOriginal                    func(sid *string, source *string, var_int4 *int8) (int, error)
	UpdateParameterInt4Int8PointerPointer                     func(sid string, source string, var_int4 *int8) (*int, error)
	UpdateEntityInt4Int8                                      func(item Int4Original) error
	UpdateEntityInt4Int8OriginalPointer                       func(item Int4Original) error
	UpdateEntityInt4Int8PointerOriginal                       func(item *Int4Pointer) error
	UpdateEntityInt4Int8PointerPointer                        func(item *Int4Pointer) error
	UpdateArrayParameterInt4Int8                              func(sid string, source string, items []int8) (int, error)
	UpdateArrayParameterInt4Int8OriginalPointer               func(sid string, source string, items []int8) (*int, error)
	UpdateArrayParameterInt4Int8PointerOriginal               func(sid string, source string, items []*int8) (int, error)
	UpdateArrayParameterInt4Int8PointerPointer                func(sid string, source string, items []*int8) (*int, error)
	UpdateArrayEntityInt4Int8                                 func(item Int4Original) (int, error)
	UpdateArrayEntityInt4Int8OriginalPointer                  func(item Int4Pointer) (int, error)
	UpdateArrayEntityInt4Int8PointerOriginal                  func(item Int4Original) (int, error)
	UpdateArrayEntityInt4Int8PointerPointer                   func(item Int4Pointer) (int, error)
	UpdateParameterNumericDecimal                             func(sid string, source string, var_numeric decimal) (int, error)
	UpdateParameterNumericDecimalOriginalPointer              func(sid string, source string, var_numeric decimal) (*int, error)
	UpdateParameterNumericDecimalPointerOriginal              func(sid *string, source *string, var_numeric *decimal) (int, error)
	UpdateParameterNumericDecimalPointerPointer               func(sid string, source string, var_numeric *decimal) (*int, error)
	UpdateEntityNumericDecimal                                func(item NumericOriginal) error
	UpdateEntityNumericDecimalOriginalPointer                 func(item NumericOriginal) error
	UpdateEntityNumericDecimalPointerOriginal                 func(item *NumericPointer) error
	UpdateEntityNumericDecimalPointerPointer                  func(item *NumericPointer) error
	UpdateArrayParameterNumericDecimal                        func(sid string, source string, items []decimal) (int, error)
	UpdateArrayParameterNumericDecimalOriginalPointer         func(sid string, source string, items []decimal) (*int, error)
	UpdateArrayParameterNumericDecimalPointerOriginal         func(sid string, source string, items []*decimal) (int, error)
	UpdateArrayParameterNumericDecimalPointerPointer          func(sid string, source string, items []*decimal) (*int, error)
	UpdateArrayEntityNumericDecimal                           func(item NumericOriginal) (int, error)
	UpdateArrayEntityNumericDecimalOriginalPointer            func(item NumericPointer) (int, error)
	UpdateArrayEntityNumericDecimalPointerOriginal            func(item NumericOriginal) (int, error)
	UpdateArrayEntityNumericDecimalPointerPointer             func(item NumericPointer) (int, error)
	UpdateParameterDecimalDecimal                             func(sid string, source string, var_decimal decimal) (int, error)
	UpdateParameterDecimalDecimalOriginalPointer              func(sid string, source string, var_decimal decimal) (*int, error)
	UpdateParameterDecimalDecimalPointerOriginal              func(sid *string, source *string, var_decimal *decimal) (int, error)
	UpdateParameterDecimalDecimalPointerPointer               func(sid string, source string, var_decimal *decimal) (*int, error)
	UpdateEntityDecimalDecimal                                func(item DecimalOriginal) error
	UpdateEntityDecimalDecimalOriginalPointer                 func(item DecimalOriginal) error
	UpdateEntityDecimalDecimalPointerOriginal                 func(item *DecimalPointer) error
	UpdateEntityDecimalDecimalPointerPointer                  func(item *DecimalPointer) error
	UpdateArrayParameterDecimalDecimal                        func(sid string, source string, items []decimal) (int, error)
	UpdateArrayParameterDecimalDecimalOriginalPointer         func(sid string, source string, items []decimal) (*int, error)
	UpdateArrayParameterDecimalDecimalPointerOriginal         func(sid string, source string, items []*decimal) (int, error)
	UpdateArrayParameterDecimalDecimalPointerPointer          func(sid string, source string, items []*decimal) (*int, error)
	UpdateArrayEntityDecimalDecimal                           func(item DecimalOriginal) (int, error)
	UpdateArrayEntityDecimalDecimalOriginalPointer            func(item DecimalPointer) (int, error)
	UpdateArrayEntityDecimalDecimalPointerOriginal            func(item DecimalOriginal) (int, error)
	UpdateArrayEntityDecimalDecimalPointerPointer             func(item DecimalPointer) (int, error)
	UpdateParameterFloat4Float32                              func(sid string, source string, var_float4 float32) (int, error)
	UpdateParameterFloat4Float32OriginalPointer               func(sid string, source string, var_float4 float32) (*int, error)
	UpdateParameterFloat4Float32PointerOriginal               func(sid *string, source *string, var_float4 *float32) (int, error)
	UpdateParameterFloat4Float32PointerPointer                func(sid string, source string, var_float4 *float32) (*int, error)
	UpdateEntityFloat4Float32                                 func(item Float4Original) error
	UpdateEntityFloat4Float32OriginalPointer                  func(item Float4Original) error
	UpdateEntityFloat4Float32PointerOriginal                  func(item *Float4Pointer) error
	UpdateEntityFloat4Float32PointerPointer                   func(item *Float4Pointer) error
	UpdateArrayParameterFloat4Float32                         func(sid string, source string, items []float32) (int, error)
	UpdateArrayParameterFloat4Float32OriginalPointer          func(sid string, source string, items []float32) (*int, error)
	UpdateArrayParameterFloat4Float32PointerOriginal          func(sid string, source string, items []*float32) (int, error)
	UpdateArrayParameterFloat4Float32PointerPointer           func(sid string, source string, items []*float32) (*int, error)
	UpdateArrayEntityFloat4Float32                            func(item Float4Original) (int, error)
	UpdateArrayEntityFloat4Float32OriginalPointer             func(item Float4Pointer) (int, error)
	UpdateArrayEntityFloat4Float32PointerOriginal             func(item Float4Original) (int, error)
	UpdateArrayEntityFloat4Float32PointerPointer              func(item Float4Pointer) (int, error)
	UpdateParameterSmallintInt8                               func(sid string, source string, var_smallint int8) (int, error)
	UpdateParameterSmallintInt8OriginalPointer                func(sid string, source string, var_smallint int8) (*int, error)
	UpdateParameterSmallintInt8PointerOriginal                func(sid *string, source *string, var_smallint *int8) (int, error)
	UpdateParameterSmallintInt8PointerPointer                 func(sid string, source string, var_smallint *int8) (*int, error)
	UpdateEntitySmallintInt8                                  func(item SmallintOriginal) error
	UpdateEntitySmallintInt8OriginalPointer                   func(item SmallintOriginal) error
	UpdateEntitySmallintInt8PointerOriginal                   func(item *SmallintPointer) error
	UpdateEntitySmallintInt8PointerPointer                    func(item *SmallintPointer) error
	UpdateArrayParameterSmallintInt8                          func(sid string, source string, items []int8) (int, error)
	UpdateArrayParameterSmallintInt8OriginalPointer           func(sid string, source string, items []int8) (*int, error)
	UpdateArrayParameterSmallintInt8PointerOriginal           func(sid string, source string, items []*int8) (int, error)
	UpdateArrayParameterSmallintInt8PointerPointer            func(sid string, source string, items []*int8) (*int, error)
	UpdateArrayEntitySmallintInt8                             func(item SmallintOriginal) (int, error)
	UpdateArrayEntitySmallintInt8OriginalPointer              func(item SmallintPointer) (int, error)
	UpdateArrayEntitySmallintInt8PointerOriginal              func(item SmallintOriginal) (int, error)
	UpdateArrayEntitySmallintInt8PointerPointer               func(item SmallintPointer) (int, error)
	UpdateParameterInt2Int8                                   func(sid string, source string, var_int2 int8) (int, error)
	UpdateParameterInt2Int8OriginalPointer                    func(sid string, source string, var_int2 int8) (*int, error)
	UpdateParameterInt2Int8PointerOriginal                    func(sid *string, source *string, var_int2 *int8) (int, error)
	UpdateParameterInt2Int8PointerPointer                     func(sid string, source string, var_int2 *int8) (*int, error)
	UpdateEntityInt2Int8                                      func(item Int2Original) error
	UpdateEntityInt2Int8OriginalPointer                       func(item Int2Original) error
	UpdateEntityInt2Int8PointerOriginal                       func(item *Int2Pointer) error
	UpdateEntityInt2Int8PointerPointer                        func(item *Int2Pointer) error
	UpdateArrayParameterInt2Int8                              func(sid string, source string, items []int8) (int, error)
	UpdateArrayParameterInt2Int8OriginalPointer               func(sid string, source string, items []int8) (*int, error)
	UpdateArrayParameterInt2Int8PointerOriginal               func(sid string, source string, items []*int8) (int, error)
	UpdateArrayParameterInt2Int8PointerPointer                func(sid string, source string, items []*int8) (*int, error)
	UpdateArrayEntityInt2Int8                                 func(item Int2Original) (int, error)
	UpdateArrayEntityInt2Int8OriginalPointer                  func(item Int2Pointer) (int, error)
	UpdateArrayEntityInt2Int8PointerOriginal                  func(item Int2Original) (int, error)
	UpdateArrayEntityInt2Int8PointerPointer                   func(item Int2Pointer) (int, error)
	UpdateParameterTextString                                 func(sid string, source string, var_text string) (int, error)
	UpdateParameterTextStringOriginalPointer                  func(sid string, source string, var_text string) (*int, error)
	UpdateParameterTextStringPointerOriginal                  func(sid *string, source *string, var_text *string) (int, error)
	UpdateParameterTextStringPointerPointer                   func(sid string, source string, var_text *string) (*int, error)
	UpdateEntityTextString                                    func(item TextOriginal) error
	UpdateEntityTextStringOriginalPointer                     func(item TextOriginal) error
	UpdateEntityTextStringPointerOriginal                     func(item *TextPointer) error
	UpdateEntityTextStringPointerPointer                      func(item *TextPointer) error
	UpdateParameterTimeTime                                   func(sid string, source string, var_time time) (int, error)
	UpdateParameterTimeTimeOriginalPointer                    func(sid string, source string, var_time time) (*int, error)
	UpdateParameterTimeTimePointerOriginal                    func(sid *string, source *string, var_time *time) (int, error)
	UpdateParameterTimeTimePointerPointer                     func(sid string, source string, var_time *time) (*int, error)
	UpdateEntityTimeTime                                      func(item TimeOriginal) error
	UpdateEntityTimeTimeOriginalPointer                       func(item TimeOriginal) error
	UpdateEntityTimeTimePointerOriginal                       func(item *TimePointer) error
	UpdateEntityTimeTimePointerPointer                        func(item *TimePointer) error
	UpdateParameterTimeWithTimezoneTime                       func(sid string, source string, var_time_with_timezone time) (int, error)
	UpdateParameterTimeWithTimezoneTimeOriginalPointer        func(sid string, source string, var_time_with_timezone time) (*int, error)
	UpdateParameterTimeWithTimezoneTimePointerOriginal        func(sid *string, source *string, var_time_with_timezone *time) (int, error)
	UpdateParameterTimeWithTimezoneTimePointerPointer         func(sid string, source string, var_time_with_timezone *time) (*int, error)
	UpdateEntityTimeWithTimezoneTime                          func(item TimeWithTimezoneOriginal) error
	UpdateEntityTimeWithTimezoneTimeOriginalPointer           func(item TimeWithTimezoneOriginal) error
	UpdateEntityTimeWithTimezoneTimePointerOriginal           func(item *TimeWithTimezonePointer) error
	UpdateEntityTimeWithTimezoneTimePointerPointer            func(item *TimeWithTimezonePointer) error
	UpdateParameterTimetzTime                                 func(sid string, source string, var_timetz time) (int, error)
	UpdateParameterTimetzTimeOriginalPointer                  func(sid string, source string, var_timetz time) (*int, error)
	UpdateParameterTimetzTimePointerOriginal                  func(sid *string, source *string, var_timetz *time) (int, error)
	UpdateParameterTimetzTimePointerPointer                   func(sid string, source string, var_timetz *time) (*int, error)
	UpdateEntityTimetzTime                                    func(item TimetzOriginal) error
	UpdateEntityTimetzTimeOriginalPointer                     func(item TimetzOriginal) error
	UpdateEntityTimetzTimePointerOriginal                     func(item *TimetzPointer) error
	UpdateEntityTimetzTimePointerPointer                      func(item *TimetzPointer) error
	UpdateParameterTimestampTime                              func(sid string, source string, var_timestamp time) (int, error)
	UpdateParameterTimestampTimeOriginalPointer               func(sid string, source string, var_timestamp time) (*int, error)
	UpdateParameterTimestampTimePointerOriginal               func(sid *string, source *string, var_timestamp *time) (int, error)
	UpdateParameterTimestampTimePointerPointer                func(sid string, source string, var_timestamp *time) (*int, error)
	UpdateEntityTimestampTime                                 func(item TimestampOriginal) error
	UpdateEntityTimestampTimeOriginalPointer                  func(item TimestampOriginal) error
	UpdateEntityTimestampTimePointerOriginal                  func(item *TimestampPointer) error
	UpdateEntityTimestampTimePointerPointer                   func(item *TimestampPointer) error
	UpdateParameterTimestampWithTimezoneTime                  func(sid string, source string, var_timestamp_with_timezone time) (int, error)
	UpdateParameterTimestampWithTimezoneTimeOriginalPointer   func(sid string, source string, var_timestamp_with_timezone time) (*int, error)
	UpdateParameterTimestampWithTimezoneTimePointerOriginal   func(sid *string, source *string, var_timestamp_with_timezone *time) (int, error)
	UpdateParameterTimestampWithTimezoneTimePointerPointer    func(sid string, source string, var_timestamp_with_timezone *time) (*int, error)
	UpdateEntityTimestampWithTimezoneTime                     func(item TimestampWithTimezoneOriginal) error
	UpdateEntityTimestampWithTimezoneTimeOriginalPointer      func(item TimestampWithTimezoneOriginal) error
	UpdateEntityTimestampWithTimezoneTimePointerOriginal      func(item *TimestampWithTimezonePointer) error
	UpdateEntityTimestampWithTimezoneTimePointerPointer       func(item *TimestampWithTimezonePointer) error
	UpdateParameterTimestamptzTime                            func(sid string, source string, var_timestamptz time) (int, error)
	UpdateParameterTimestamptzTimeOriginalPointer             func(sid string, source string, var_timestamptz time) (*int, error)
	UpdateParameterTimestamptzTimePointerOriginal             func(sid *string, source *string, var_timestamptz *time) (int, error)
	UpdateParameterTimestamptzTimePointerPointer              func(sid string, source string, var_timestamptz *time) (*int, error)
	UpdateEntityTimestamptzTime                               func(item TimestamptzOriginal) error
	UpdateEntityTimestamptzTimeOriginalPointer                func(item TimestamptzOriginal) error
	UpdateEntityTimestamptzTimePointerOriginal                func(item *TimestamptzPointer) error
	UpdateEntityTimestamptzTimePointerPointer                 func(item *TimestamptzPointer) error
	DeleteParameterBigintInt64                                func(sid string) (int, error)
	DeleteParameterBigintInt64OriginalPointer                 func(sid string) (*int, error)
	DeleteParameterBigintInt64PointerOriginal                 func(sid *string) (int, error)
	DeleteParameterBigintInt64PointerPointer                  func(sid string) (*int, error)
	DeleteEntityBigintInt64                                   func(item BigintOriginal) error
	DeleteEntityBigintInt64OriginalPointer                    func(item BigintOriginal) error
	DeleteEntityBigintInt64PointerOriginal                    func(item *BigintPointer) error
	DeleteEntityBigintInt64PointerPointer                     func(item *BigintPointer) error
	DeleteArrayParameterBigintInt64                           func(sid string) (int, error)
	DeleteArrayParameterBigintInt64OriginalPointer            func(sid string) (*int, error)
	DeleteArrayParameterBigintInt64PointerOriginal            func(sid string) (int, error)
	DeleteArrayParameterBigintInt64PointerPointer             func(sid string) (*int, error)
	DeleteArrayEntityBigintInt64                              func(sid string) (int, error)
	DeleteArrayEntityBigintInt64OriginalPointer               func(sid string) (int, error)
	DeleteArrayEntityBigintInt64PointerOriginal               func(sid string) (int, error)
	DeleteArrayEntityBigintInt64PointerPointer                func(sid string) (int, error)
	DeleteParameterInt8Int8                                   func(sid string) (int, error)
	DeleteParameterInt8Int8OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterInt8Int8PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterInt8Int8PointerPointer                     func(sid string) (*int, error)
	DeleteEntityInt8Int8                                      func(item Int8Original) error
	DeleteEntityInt8Int8OriginalPointer                       func(item Int8Original) error
	DeleteEntityInt8Int8PointerOriginal                       func(item *Int8Pointer) error
	DeleteEntityInt8Int8PointerPointer                        func(item *Int8Pointer) error
	DeleteArrayParameterInt8Int8                              func(sid string) (int, error)
	DeleteArrayParameterInt8Int8OriginalPointer               func(sid string) (*int, error)
	DeleteArrayParameterInt8Int8PointerOriginal               func(sid string) (int, error)
	DeleteArrayParameterInt8Int8PointerPointer                func(sid string) (*int, error)
	DeleteArrayEntityInt8Int8                                 func(sid string) (int, error)
	DeleteArrayEntityInt8Int8OriginalPointer                  func(sid string) (int, error)
	DeleteArrayEntityInt8Int8PointerOriginal                  func(sid string) (int, error)
	DeleteArrayEntityInt8Int8PointerPointer                   func(sid string) (int, error)
	DeleteParameterBooleanBool                                func(sid string) (int, error)
	DeleteParameterBooleanBoolOriginalPointer                 func(sid string) (*int, error)
	DeleteParameterBooleanBoolPointerOriginal                 func(sid *string) (int, error)
	DeleteParameterBooleanBoolPointerPointer                  func(sid string) (*int, error)
	DeleteEntityBooleanBool                                   func(item BooleanOriginal) error
	DeleteEntityBooleanBoolOriginalPointer                    func(item BooleanOriginal) error
	DeleteEntityBooleanBoolPointerOriginal                    func(item *BooleanPointer) error
	DeleteEntityBooleanBoolPointerPointer                     func(item *BooleanPointer) error
	DeleteArrayParameterBooleanBool                           func(sid string) (int, error)
	DeleteArrayParameterBooleanBoolOriginalPointer            func(sid string) (*int, error)
	DeleteArrayParameterBooleanBoolPointerOriginal            func(sid string) (int, error)
	DeleteArrayParameterBooleanBoolPointerPointer             func(sid string) (*int, error)
	DeleteArrayEntityBooleanBool                              func(sid string) (int, error)
	DeleteArrayEntityBooleanBoolOriginalPointer               func(sid string) (int, error)
	DeleteArrayEntityBooleanBoolPointerOriginal               func(sid string) (int, error)
	DeleteArrayEntityBooleanBoolPointerPointer                func(sid string) (int, error)
	DeleteParameterBoolBool                                   func(sid string) (int, error)
	DeleteParameterBoolBoolOriginalPointer                    func(sid string) (*int, error)
	DeleteParameterBoolBoolPointerOriginal                    func(sid *string) (int, error)
	DeleteParameterBoolBoolPointerPointer                     func(sid string) (*int, error)
	DeleteEntityBoolBool                                      func(item BoolOriginal) error
	DeleteEntityBoolBoolOriginalPointer                       func(item BoolOriginal) error
	DeleteEntityBoolBoolPointerOriginal                       func(item *BoolPointer) error
	DeleteEntityBoolBoolPointerPointer                        func(item *BoolPointer) error
	DeleteArrayParameterBoolBool                              func(sid string) (int, error)
	DeleteArrayParameterBoolBoolOriginalPointer               func(sid string) (*int, error)
	DeleteArrayParameterBoolBoolPointerOriginal               func(sid string) (int, error)
	DeleteArrayParameterBoolBoolPointerPointer                func(sid string) (*int, error)
	DeleteArrayEntityBoolBool                                 func(sid string) (int, error)
	DeleteArrayEntityBoolBoolOriginalPointer                  func(sid string) (int, error)
	DeleteArrayEntityBoolBoolPointerOriginal                  func(sid string) (int, error)
	DeleteArrayEntityBoolBoolPointerPointer                   func(sid string) (int, error)
	DeleteParameterCharacterString                            func(sid string) (int, error)
	DeleteParameterCharacterStringOriginalPointer             func(sid string) (*int, error)
	DeleteParameterCharacterStringPointerOriginal             func(sid *string) (int, error)
	DeleteParameterCharacterStringPointerPointer              func(sid string) (*int, error)
	DeleteEntityCharacterString                               func(item CharacterOriginal) error
	DeleteEntityCharacterStringOriginalPointer                func(item CharacterOriginal) error
	DeleteEntityCharacterStringPointerOriginal                func(item *CharacterPointer) error
	DeleteEntityCharacterStringPointerPointer                 func(item *CharacterPointer) error
	DeleteArrayParameterCharacterString                       func(sid string) (int, error)
	DeleteArrayParameterCharacterStringOriginalPointer        func(sid string) (*int, error)
	DeleteArrayParameterCharacterStringPointerOriginal        func(sid string) (int, error)
	DeleteArrayParameterCharacterStringPointerPointer         func(sid string) (*int, error)
	DeleteArrayEntityCharacterString                          func(sid string) (int, error)
	DeleteArrayEntityCharacterStringOriginalPointer           func(sid string) (int, error)
	DeleteArrayEntityCharacterStringPointerOriginal           func(sid string) (int, error)
	DeleteArrayEntityCharacterStringPointerPointer            func(sid string) (int, error)
	DeleteParameterCharString                                 func(sid string) (int, error)
	DeleteParameterCharStringOriginalPointer                  func(sid string) (*int, error)
	DeleteParameterCharStringPointerOriginal                  func(sid *string) (int, error)
	DeleteParameterCharStringPointerPointer                   func(sid string) (*int, error)
	DeleteEntityCharString                                    func(item CharOriginal) error
	DeleteEntityCharStringOriginalPointer                     func(item CharOriginal) error
	DeleteEntityCharStringPointerOriginal                     func(item *CharPointer) error
	DeleteEntityCharStringPointerPointer                      func(item *CharPointer) error
	DeleteArrayParameterCharString                            func(sid string) (int, error)
	DeleteArrayParameterCharStringOriginalPointer             func(sid string) (*int, error)
	DeleteArrayParameterCharStringPointerOriginal             func(sid string) (int, error)
	DeleteArrayParameterCharStringPointerPointer              func(sid string) (*int, error)
	DeleteArrayEntityCharString                               func(sid string) (int, error)
	DeleteArrayEntityCharStringOriginalPointer                func(sid string) (int, error)
	DeleteArrayEntityCharStringPointerOriginal                func(sid string) (int, error)
	DeleteArrayEntityCharStringPointerPointer                 func(sid string) (int, error)
	DeleteParameterCharacterVaryingString                     func(sid string) (int, error)
	DeleteParameterCharacterVaryingStringOriginalPointer      func(sid string) (*int, error)
	DeleteParameterCharacterVaryingStringPointerOriginal      func(sid *string) (int, error)
	DeleteParameterCharacterVaryingStringPointerPointer       func(sid string) (*int, error)
	DeleteEntityCharacterVaryingString                        func(item CharacterVaryingOriginal) error
	DeleteEntityCharacterVaryingStringOriginalPointer         func(item CharacterVaryingOriginal) error
	DeleteEntityCharacterVaryingStringPointerOriginal         func(item *CharacterVaryingPointer) error
	DeleteEntityCharacterVaryingStringPointerPointer          func(item *CharacterVaryingPointer) error
	DeleteArrayParameterCharacterVaryingString                func(sid string) (int, error)
	DeleteArrayParameterCharacterVaryingStringOriginalPointer func(sid string) (*int, error)
	DeleteArrayParameterCharacterVaryingStringPointerOriginal func(sid string) (int, error)
	DeleteArrayParameterCharacterVaryingStringPointerPointer  func(sid string) (*int, error)
	DeleteArrayEntityCharacterVaryingString                   func(sid string) (int, error)
	DeleteArrayEntityCharacterVaryingStringOriginalPointer    func(sid string) (int, error)
	DeleteArrayEntityCharacterVaryingStringPointerOriginal    func(sid string) (int, error)
	DeleteArrayEntityCharacterVaryingStringPointerPointer     func(sid string) (int, error)
	DeleteParameterVarcharString                              func(sid string) (int, error)
	DeleteParameterVarcharStringOriginalPointer               func(sid string) (*int, error)
	DeleteParameterVarcharStringPointerOriginal               func(sid *string) (int, error)
	DeleteParameterVarcharStringPointerPointer                func(sid string) (*int, error)
	DeleteEntityVarcharString                                 func(item VarcharOriginal) error
	DeleteEntityVarcharStringOriginalPointer                  func(item VarcharOriginal) error
	DeleteEntityVarcharStringPointerOriginal                  func(item *VarcharPointer) error
	DeleteEntityVarcharStringPointerPointer                   func(item *VarcharPointer) error
	DeleteArrayParameterVarcharString                         func(sid string) (int, error)
	DeleteArrayParameterVarcharStringOriginalPointer          func(sid string) (*int, error)
	DeleteArrayParameterVarcharStringPointerOriginal          func(sid string) (int, error)
	DeleteArrayParameterVarcharStringPointerPointer           func(sid string) (*int, error)
	DeleteArrayEntityVarcharString                            func(sid string) (int, error)
	DeleteArrayEntityVarcharStringOriginalPointer             func(sid string) (int, error)
	DeleteArrayEntityVarcharStringPointerOriginal             func(sid string) (int, error)
	DeleteArrayEntityVarcharStringPointerPointer              func(sid string) (int, error)
	DeleteParameterFloat8Float32                              func(sid string) (int, error)
	DeleteParameterFloat8Float32OriginalPointer               func(sid string) (*int, error)
	DeleteParameterFloat8Float32PointerOriginal               func(sid *string) (int, error)
	DeleteParameterFloat8Float32PointerPointer                func(sid string) (*int, error)
	DeleteEntityFloat8Float32                                 func(item Float8Original) error
	DeleteEntityFloat8Float32OriginalPointer                  func(item Float8Original) error
	DeleteEntityFloat8Float32PointerOriginal                  func(item *Float8Pointer) error
	DeleteEntityFloat8Float32PointerPointer                   func(item *Float8Pointer) error
	DeleteArrayParameterFloat8Float32                         func(sid string) (int, error)
	DeleteArrayParameterFloat8Float32OriginalPointer          func(sid string) (*int, error)
	DeleteArrayParameterFloat8Float32PointerOriginal          func(sid string) (int, error)
	DeleteArrayParameterFloat8Float32PointerPointer           func(sid string) (*int, error)
	DeleteArrayEntityFloat8Float32                            func(sid string) (int, error)
	DeleteArrayEntityFloat8Float32OriginalPointer             func(sid string) (int, error)
	DeleteArrayEntityFloat8Float32PointerOriginal             func(sid string) (int, error)
	DeleteArrayEntityFloat8Float32PointerPointer              func(sid string) (int, error)
	DeleteParameterIntegerInt8                                func(sid string) (int, error)
	DeleteParameterIntegerInt8OriginalPointer                 func(sid string) (*int, error)
	DeleteParameterIntegerInt8PointerOriginal                 func(sid *string) (int, error)
	DeleteParameterIntegerInt8PointerPointer                  func(sid string) (*int, error)
	DeleteEntityIntegerInt8                                   func(item IntegerOriginal) error
	DeleteEntityIntegerInt8OriginalPointer                    func(item IntegerOriginal) error
	DeleteEntityIntegerInt8PointerOriginal                    func(item *IntegerPointer) error
	DeleteEntityIntegerInt8PointerPointer                     func(item *IntegerPointer) error
	DeleteArrayParameterIntegerInt8                           func(sid string) (int, error)
	DeleteArrayParameterIntegerInt8OriginalPointer            func(sid string) (*int, error)
	DeleteArrayParameterIntegerInt8PointerOriginal            func(sid string) (int, error)
	DeleteArrayParameterIntegerInt8PointerPointer             func(sid string) (*int, error)
	DeleteArrayEntityIntegerInt8                              func(sid string) (int, error)
	DeleteArrayEntityIntegerInt8OriginalPointer               func(sid string) (int, error)
	DeleteArrayEntityIntegerInt8PointerOriginal               func(sid string) (int, error)
	DeleteArrayEntityIntegerInt8PointerPointer                func(sid string) (int, error)
	DeleteParameterIntInt8                                    func(sid string) (int, error)
	DeleteParameterIntInt8OriginalPointer                     func(sid string) (*int, error)
	DeleteParameterIntInt8PointerOriginal                     func(sid *string) (int, error)
	DeleteParameterIntInt8PointerPointer                      func(sid string) (*int, error)
	DeleteEntityIntInt8                                       func(item IntOriginal) error
	DeleteEntityIntInt8OriginalPointer                        func(item IntOriginal) error
	DeleteEntityIntInt8PointerOriginal                        func(item *IntPointer) error
	DeleteEntityIntInt8PointerPointer                         func(item *IntPointer) error
	DeleteArrayParameterIntInt8                               func(sid string) (int, error)
	DeleteArrayParameterIntInt8OriginalPointer                func(sid string) (*int, error)
	DeleteArrayParameterIntInt8PointerOriginal                func(sid string) (int, error)
	DeleteArrayParameterIntInt8PointerPointer                 func(sid string) (*int, error)
	DeleteArrayEntityIntInt8                                  func(sid string) (int, error)
	DeleteArrayEntityIntInt8OriginalPointer                   func(sid string) (int, error)
	DeleteArrayEntityIntInt8PointerOriginal                   func(sid string) (int, error)
	DeleteArrayEntityIntInt8PointerPointer                    func(sid string) (int, error)
	DeleteParameterInt4Int8                                   func(sid string) (int, error)
	DeleteParameterInt4Int8OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterInt4Int8PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterInt4Int8PointerPointer                     func(sid string) (*int, error)
	DeleteEntityInt4Int8                                      func(item Int4Original) error
	DeleteEntityInt4Int8OriginalPointer                       func(item Int4Original) error
	DeleteEntityInt4Int8PointerOriginal                       func(item *Int4Pointer) error
	DeleteEntityInt4Int8PointerPointer                        func(item *Int4Pointer) error
	DeleteArrayParameterInt4Int8                              func(sid string) (int, error)
	DeleteArrayParameterInt4Int8OriginalPointer               func(sid string) (*int, error)
	DeleteArrayParameterInt4Int8PointerOriginal               func(sid string) (int, error)
	DeleteArrayParameterInt4Int8PointerPointer                func(sid string) (*int, error)
	DeleteArrayEntityInt4Int8                                 func(sid string) (int, error)
	DeleteArrayEntityInt4Int8OriginalPointer                  func(sid string) (int, error)
	DeleteArrayEntityInt4Int8PointerOriginal                  func(sid string) (int, error)
	DeleteArrayEntityInt4Int8PointerPointer                   func(sid string) (int, error)
	DeleteParameterNumericDecimal                             func(sid string) (int, error)
	DeleteParameterNumericDecimalOriginalPointer              func(sid string) (*int, error)
	DeleteParameterNumericDecimalPointerOriginal              func(sid *string) (int, error)
	DeleteParameterNumericDecimalPointerPointer               func(sid string) (*int, error)
	DeleteEntityNumericDecimal                                func(item NumericOriginal) error
	DeleteEntityNumericDecimalOriginalPointer                 func(item NumericOriginal) error
	DeleteEntityNumericDecimalPointerOriginal                 func(item *NumericPointer) error
	DeleteEntityNumericDecimalPointerPointer                  func(item *NumericPointer) error
	DeleteArrayParameterNumericDecimal                        func(sid string) (int, error)
	DeleteArrayParameterNumericDecimalOriginalPointer         func(sid string) (*int, error)
	DeleteArrayParameterNumericDecimalPointerOriginal         func(sid string) (int, error)
	DeleteArrayParameterNumericDecimalPointerPointer          func(sid string) (*int, error)
	DeleteArrayEntityNumericDecimal                           func(sid string) (int, error)
	DeleteArrayEntityNumericDecimalOriginalPointer            func(sid string) (int, error)
	DeleteArrayEntityNumericDecimalPointerOriginal            func(sid string) (int, error)
	DeleteArrayEntityNumericDecimalPointerPointer             func(sid string) (int, error)
	DeleteParameterDecimalDecimal                             func(sid string) (int, error)
	DeleteParameterDecimalDecimalOriginalPointer              func(sid string) (*int, error)
	DeleteParameterDecimalDecimalPointerOriginal              func(sid *string) (int, error)
	DeleteParameterDecimalDecimalPointerPointer               func(sid string) (*int, error)
	DeleteEntityDecimalDecimal                                func(item DecimalOriginal) error
	DeleteEntityDecimalDecimalOriginalPointer                 func(item DecimalOriginal) error
	DeleteEntityDecimalDecimalPointerOriginal                 func(item *DecimalPointer) error
	DeleteEntityDecimalDecimalPointerPointer                  func(item *DecimalPointer) error
	DeleteArrayParameterDecimalDecimal                        func(sid string) (int, error)
	DeleteArrayParameterDecimalDecimalOriginalPointer         func(sid string) (*int, error)
	DeleteArrayParameterDecimalDecimalPointerOriginal         func(sid string) (int, error)
	DeleteArrayParameterDecimalDecimalPointerPointer          func(sid string) (*int, error)
	DeleteArrayEntityDecimalDecimal                           func(sid string) (int, error)
	DeleteArrayEntityDecimalDecimalOriginalPointer            func(sid string) (int, error)
	DeleteArrayEntityDecimalDecimalPointerOriginal            func(sid string) (int, error)
	DeleteArrayEntityDecimalDecimalPointerPointer             func(sid string) (int, error)
	DeleteParameterFloat4Float32                              func(sid string) (int, error)
	DeleteParameterFloat4Float32OriginalPointer               func(sid string) (*int, error)
	DeleteParameterFloat4Float32PointerOriginal               func(sid *string) (int, error)
	DeleteParameterFloat4Float32PointerPointer                func(sid string) (*int, error)
	DeleteEntityFloat4Float32                                 func(item Float4Original) error
	DeleteEntityFloat4Float32OriginalPointer                  func(item Float4Original) error
	DeleteEntityFloat4Float32PointerOriginal                  func(item *Float4Pointer) error
	DeleteEntityFloat4Float32PointerPointer                   func(item *Float4Pointer) error
	DeleteArrayParameterFloat4Float32                         func(sid string) (int, error)
	DeleteArrayParameterFloat4Float32OriginalPointer          func(sid string) (*int, error)
	DeleteArrayParameterFloat4Float32PointerOriginal          func(sid string) (int, error)
	DeleteArrayParameterFloat4Float32PointerPointer           func(sid string) (*int, error)
	DeleteArrayEntityFloat4Float32                            func(sid string) (int, error)
	DeleteArrayEntityFloat4Float32OriginalPointer             func(sid string) (int, error)
	DeleteArrayEntityFloat4Float32PointerOriginal             func(sid string) (int, error)
	DeleteArrayEntityFloat4Float32PointerPointer              func(sid string) (int, error)
	DeleteParameterSmallintInt8                               func(sid string) (int, error)
	DeleteParameterSmallintInt8OriginalPointer                func(sid string) (*int, error)
	DeleteParameterSmallintInt8PointerOriginal                func(sid *string) (int, error)
	DeleteParameterSmallintInt8PointerPointer                 func(sid string) (*int, error)
	DeleteEntitySmallintInt8                                  func(item SmallintOriginal) error
	DeleteEntitySmallintInt8OriginalPointer                   func(item SmallintOriginal) error
	DeleteEntitySmallintInt8PointerOriginal                   func(item *SmallintPointer) error
	DeleteEntitySmallintInt8PointerPointer                    func(item *SmallintPointer) error
	DeleteArrayParameterSmallintInt8                          func(sid string) (int, error)
	DeleteArrayParameterSmallintInt8OriginalPointer           func(sid string) (*int, error)
	DeleteArrayParameterSmallintInt8PointerOriginal           func(sid string) (int, error)
	DeleteArrayParameterSmallintInt8PointerPointer            func(sid string) (*int, error)
	DeleteArrayEntitySmallintInt8                             func(sid string) (int, error)
	DeleteArrayEntitySmallintInt8OriginalPointer              func(sid string) (int, error)
	DeleteArrayEntitySmallintInt8PointerOriginal              func(sid string) (int, error)
	DeleteArrayEntitySmallintInt8PointerPointer               func(sid string) (int, error)
	DeleteParameterInt2Int8                                   func(sid string) (int, error)
	DeleteParameterInt2Int8OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterInt2Int8PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterInt2Int8PointerPointer                     func(sid string) (*int, error)
	DeleteEntityInt2Int8                                      func(item Int2Original) error
	DeleteEntityInt2Int8OriginalPointer                       func(item Int2Original) error
	DeleteEntityInt2Int8PointerOriginal                       func(item *Int2Pointer) error
	DeleteEntityInt2Int8PointerPointer                        func(item *Int2Pointer) error
	DeleteArrayParameterInt2Int8                              func(sid string) (int, error)
	DeleteArrayParameterInt2Int8OriginalPointer               func(sid string) (*int, error)
	DeleteArrayParameterInt2Int8PointerOriginal               func(sid string) (int, error)
	DeleteArrayParameterInt2Int8PointerPointer                func(sid string) (*int, error)
	DeleteArrayEntityInt2Int8                                 func(sid string) (int, error)
	DeleteArrayEntityInt2Int8OriginalPointer                  func(sid string) (int, error)
	DeleteArrayEntityInt2Int8PointerOriginal                  func(sid string) (int, error)
	DeleteArrayEntityInt2Int8PointerPointer                   func(sid string) (int, error)
	DeleteParameterTextString                                 func(sid string) (int, error)
	DeleteParameterTextStringOriginalPointer                  func(sid string) (*int, error)
	DeleteParameterTextStringPointerOriginal                  func(sid *string) (int, error)
	DeleteParameterTextStringPointerPointer                   func(sid string) (*int, error)
	DeleteEntityTextString                                    func(item TextOriginal) error
	DeleteEntityTextStringOriginalPointer                     func(item TextOriginal) error
	DeleteEntityTextStringPointerOriginal                     func(item *TextPointer) error
	DeleteEntityTextStringPointerPointer                      func(item *TextPointer) error
	DeleteParameterTimeTime                                   func(sid string) (int, error)
	DeleteParameterTimeTimeOriginalPointer                    func(sid string) (*int, error)
	DeleteParameterTimeTimePointerOriginal                    func(sid *string) (int, error)
	DeleteParameterTimeTimePointerPointer                     func(sid string) (*int, error)
	DeleteEntityTimeTime                                      func(item TimeOriginal) error
	DeleteEntityTimeTimeOriginalPointer                       func(item TimeOriginal) error
	DeleteEntityTimeTimePointerOriginal                       func(item *TimePointer) error
	DeleteEntityTimeTimePointerPointer                        func(item *TimePointer) error
	DeleteParameterTimeWithTimezoneTime                       func(sid string) (int, error)
	DeleteParameterTimeWithTimezoneTimeOriginalPointer        func(sid string) (*int, error)
	DeleteParameterTimeWithTimezoneTimePointerOriginal        func(sid *string) (int, error)
	DeleteParameterTimeWithTimezoneTimePointerPointer         func(sid string) (*int, error)
	DeleteEntityTimeWithTimezoneTime                          func(item TimeWithTimezoneOriginal) error
	DeleteEntityTimeWithTimezoneTimeOriginalPointer           func(item TimeWithTimezoneOriginal) error
	DeleteEntityTimeWithTimezoneTimePointerOriginal           func(item *TimeWithTimezonePointer) error
	DeleteEntityTimeWithTimezoneTimePointerPointer            func(item *TimeWithTimezonePointer) error
	DeleteParameterTimetzTime                                 func(sid string) (int, error)
	DeleteParameterTimetzTimeOriginalPointer                  func(sid string) (*int, error)
	DeleteParameterTimetzTimePointerOriginal                  func(sid *string) (int, error)
	DeleteParameterTimetzTimePointerPointer                   func(sid string) (*int, error)
	DeleteEntityTimetzTime                                    func(item TimetzOriginal) error
	DeleteEntityTimetzTimeOriginalPointer                     func(item TimetzOriginal) error
	DeleteEntityTimetzTimePointerOriginal                     func(item *TimetzPointer) error
	DeleteEntityTimetzTimePointerPointer                      func(item *TimetzPointer) error
	DeleteParameterTimestampTime                              func(sid string) (int, error)
	DeleteParameterTimestampTimeOriginalPointer               func(sid string) (*int, error)
	DeleteParameterTimestampTimePointerOriginal               func(sid *string) (int, error)
	DeleteParameterTimestampTimePointerPointer                func(sid string) (*int, error)
	DeleteEntityTimestampTime                                 func(item TimestampOriginal) error
	DeleteEntityTimestampTimeOriginalPointer                  func(item TimestampOriginal) error
	DeleteEntityTimestampTimePointerOriginal                  func(item *TimestampPointer) error
	DeleteEntityTimestampTimePointerPointer                   func(item *TimestampPointer) error
	DeleteParameterTimestampWithTimezoneTime                  func(sid string) (int, error)
	DeleteParameterTimestampWithTimezoneTimeOriginalPointer   func(sid string) (*int, error)
	DeleteParameterTimestampWithTimezoneTimePointerOriginal   func(sid *string) (int, error)
	DeleteParameterTimestampWithTimezoneTimePointerPointer    func(sid string) (*int, error)
	DeleteEntityTimestampWithTimezoneTime                     func(item TimestampWithTimezoneOriginal) error
	DeleteEntityTimestampWithTimezoneTimeOriginalPointer      func(item TimestampWithTimezoneOriginal) error
	DeleteEntityTimestampWithTimezoneTimePointerOriginal      func(item *TimestampWithTimezonePointer) error
	DeleteEntityTimestampWithTimezoneTimePointerPointer       func(item *TimestampWithTimezonePointer) error
	DeleteParameterTimestamptzTime                            func(sid string) (int, error)
	DeleteParameterTimestamptzTimeOriginalPointer             func(sid string) (*int, error)
	DeleteParameterTimestamptzTimePointerOriginal             func(sid *string) (int, error)
	DeleteParameterTimestamptzTimePointerPointer              func(sid string) (*int, error)
	DeleteEntityTimestamptzTime                               func(item TimestamptzOriginal) error
	DeleteEntityTimestamptzTimeOriginalPointer                func(item TimestamptzOriginal) error
	DeleteEntityTimestamptzTimePointerOriginal                func(item *TimestamptzPointer) error
	DeleteEntityTimestamptzTimePointerPointer                 func(item *TimestamptzPointer) error
}
