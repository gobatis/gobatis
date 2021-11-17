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
	InsertArrayParameterBigintInt64                           func(sid string, source string, items []int64) error
	InsertArrayParameterBigintInt64OriginalPointer            func(sid string, source string, items []int64) error
	InsertArrayParameterBigintInt64PointerOriginal            func(sid string, source string, items []*int64) error
	InsertArrayParameterBigintInt64PointerPointer             func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64                              func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64OriginalPointer               func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64PointerOriginal               func(sid string, source string, items []*int64) error
	InsertArrayEntityBigintInt64PointerPointer                func(sid string, source string, items []*int64) error
	InsertParameterInt8Int8                                   func(sid string, source string, var_int8 int8) (int, error)
	InsertParameterInt8Int8OriginalPointer                    func(sid string, source string, var_int8 int8) (*int, error)
	InsertParameterInt8Int8PointerOriginal                    func(sid *string, source *string, var_int8 *int8) (int, error)
	InsertParameterInt8Int8PointerPointer                     func(sid string, source string, var_int8 *int8) (*int, error)
	InsertEntityInt8Int8                                      func(item Int8Original) error
	InsertEntityInt8Int8OriginalPointer                       func(item Int8Original) error
	InsertEntityInt8Int8PointerOriginal                       func(item *Int8Pointer) error
	InsertEntityInt8Int8PointerPointer                        func(item *Int8Pointer) error
	InsertArrayParameterInt8Int8                              func(sid string, source string, items []int8) error
	InsertArrayParameterInt8Int8OriginalPointer               func(sid string, source string, items []int8) error
	InsertArrayParameterInt8Int8PointerOriginal               func(sid string, source string, items []*int8) error
	InsertArrayParameterInt8Int8PointerPointer                func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8                                 func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8OriginalPointer                  func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	InsertArrayEntityInt8Int8PointerPointer                   func(sid string, source string, items []*int8) error
	InsertParameterBooleanBool                                func(sid string, source string, var_boolean bool) (int, error)
	InsertParameterBooleanBoolOriginalPointer                 func(sid string, source string, var_boolean bool) (*int, error)
	InsertParameterBooleanBoolPointerOriginal                 func(sid *string, source *string, var_boolean *bool) (int, error)
	InsertParameterBooleanBoolPointerPointer                  func(sid string, source string, var_boolean *bool) (*int, error)
	InsertEntityBooleanBool                                   func(item BooleanOriginal) error
	InsertEntityBooleanBoolOriginalPointer                    func(item BooleanOriginal) error
	InsertEntityBooleanBoolPointerOriginal                    func(item *BooleanPointer) error
	InsertEntityBooleanBoolPointerPointer                     func(item *BooleanPointer) error
	InsertArrayParameterBooleanBool                           func(sid string, source string, items []bool) error
	InsertArrayParameterBooleanBoolOriginalPointer            func(sid string, source string, items []bool) error
	InsertArrayParameterBooleanBoolPointerOriginal            func(sid string, source string, items []*bool) error
	InsertArrayParameterBooleanBoolPointerPointer             func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBool                              func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBoolOriginalPointer               func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBoolPointerOriginal               func(sid string, source string, items []*bool) error
	InsertArrayEntityBooleanBoolPointerPointer                func(sid string, source string, items []*bool) error
	InsertParameterBoolBool                                   func(sid string, source string, var_bool bool) (int, error)
	InsertParameterBoolBoolOriginalPointer                    func(sid string, source string, var_bool bool) (*int, error)
	InsertParameterBoolBoolPointerOriginal                    func(sid *string, source *string, var_bool *bool) (int, error)
	InsertParameterBoolBoolPointerPointer                     func(sid string, source string, var_bool *bool) (*int, error)
	InsertEntityBoolBool                                      func(item BoolOriginal) error
	InsertEntityBoolBoolOriginalPointer                       func(item BoolOriginal) error
	InsertEntityBoolBoolPointerOriginal                       func(item *BoolPointer) error
	InsertEntityBoolBoolPointerPointer                        func(item *BoolPointer) error
	InsertArrayParameterBoolBool                              func(sid string, source string, items []bool) error
	InsertArrayParameterBoolBoolOriginalPointer               func(sid string, source string, items []bool) error
	InsertArrayParameterBoolBoolPointerOriginal               func(sid string, source string, items []*bool) error
	InsertArrayParameterBoolBoolPointerPointer                func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBool                                 func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBoolOriginalPointer                  func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBoolPointerOriginal                  func(sid string, source string, items []*bool) error
	InsertArrayEntityBoolBoolPointerPointer                   func(sid string, source string, items []*bool) error
	InsertParameterCharacterString                            func(sid string, source string, var_character string) (int, error)
	InsertParameterCharacterStringOriginalPointer             func(sid string, source string, var_character string) (*int, error)
	InsertParameterCharacterStringPointerOriginal             func(sid *string, source *string, var_character *string) (int, error)
	InsertParameterCharacterStringPointerPointer              func(sid string, source string, var_character *string) (*int, error)
	InsertEntityCharacterString                               func(item CharacterOriginal) error
	InsertEntityCharacterStringOriginalPointer                func(item CharacterOriginal) error
	InsertEntityCharacterStringPointerOriginal                func(item *CharacterPointer) error
	InsertEntityCharacterStringPointerPointer                 func(item *CharacterPointer) error
	InsertArrayParameterCharacterString                       func(sid string, source string, items []string) error
	InsertArrayParameterCharacterStringOriginalPointer        func(sid string, source string, items []string) error
	InsertArrayParameterCharacterStringPointerOriginal        func(sid string, source string, items []*string) error
	InsertArrayParameterCharacterStringPointerPointer         func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterString                          func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterStringOriginalPointer           func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterStringPointerOriginal           func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterStringPointerPointer            func(sid string, source string, items []*string) error
	InsertParameterCharString                                 func(sid string, source string, var_char string) (int, error)
	InsertParameterCharStringOriginalPointer                  func(sid string, source string, var_char string) (*int, error)
	InsertParameterCharStringPointerOriginal                  func(sid *string, source *string, var_char *string) (int, error)
	InsertParameterCharStringPointerPointer                   func(sid string, source string, var_char *string) (*int, error)
	InsertEntityCharString                                    func(item CharOriginal) error
	InsertEntityCharStringOriginalPointer                     func(item CharOriginal) error
	InsertEntityCharStringPointerOriginal                     func(item *CharPointer) error
	InsertEntityCharStringPointerPointer                      func(item *CharPointer) error
	InsertArrayParameterCharString                            func(sid string, source string, items []string) error
	InsertArrayParameterCharStringOriginalPointer             func(sid string, source string, items []string) error
	InsertArrayParameterCharStringPointerOriginal             func(sid string, source string, items []*string) error
	InsertArrayParameterCharStringPointerPointer              func(sid string, source string, items []*string) error
	InsertArrayEntityCharString                               func(sid string, source string, items []*string) error
	InsertArrayEntityCharStringOriginalPointer                func(sid string, source string, items []*string) error
	InsertArrayEntityCharStringPointerOriginal                func(sid string, source string, items []*string) error
	InsertArrayEntityCharStringPointerPointer                 func(sid string, source string, items []*string) error
	InsertParameterCharacterVaryingString                     func(sid string, source string, var_character_varying string) (int, error)
	InsertParameterCharacterVaryingStringOriginalPointer      func(sid string, source string, var_character_varying string) (*int, error)
	InsertParameterCharacterVaryingStringPointerOriginal      func(sid *string, source *string, var_character_varying *string) (int, error)
	InsertParameterCharacterVaryingStringPointerPointer       func(sid string, source string, var_character_varying *string) (*int, error)
	InsertEntityCharacterVaryingString                        func(item CharacterVaryingOriginal) error
	InsertEntityCharacterVaryingStringOriginalPointer         func(item CharacterVaryingOriginal) error
	InsertEntityCharacterVaryingStringPointerOriginal         func(item *CharacterVaryingPointer) error
	InsertEntityCharacterVaryingStringPointerPointer          func(item *CharacterVaryingPointer) error
	InsertArrayParameterCharacterVaryingString                func(sid string, source string, items []string) error
	InsertArrayParameterCharacterVaryingStringOriginalPointer func(sid string, source string, items []string) error
	InsertArrayParameterCharacterVaryingStringPointerOriginal func(sid string, source string, items []*string) error
	InsertArrayParameterCharacterVaryingStringPointerPointer  func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingString                   func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingStringOriginalPointer    func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingStringPointerOriginal    func(sid string, source string, items []*string) error
	InsertArrayEntityCharacterVaryingStringPointerPointer     func(sid string, source string, items []*string) error
	InsertParameterVarcharString                              func(sid string, source string, var_varchar string) (int, error)
	InsertParameterVarcharStringOriginalPointer               func(sid string, source string, var_varchar string) (*int, error)
	InsertParameterVarcharStringPointerOriginal               func(sid *string, source *string, var_varchar *string) (int, error)
	InsertParameterVarcharStringPointerPointer                func(sid string, source string, var_varchar *string) (*int, error)
	InsertEntityVarcharString                                 func(item VarcharOriginal) error
	InsertEntityVarcharStringOriginalPointer                  func(item VarcharOriginal) error
	InsertEntityVarcharStringPointerOriginal                  func(item *VarcharPointer) error
	InsertEntityVarcharStringPointerPointer                   func(item *VarcharPointer) error
	InsertArrayParameterVarcharString                         func(sid string, source string, items []string) error
	InsertArrayParameterVarcharStringOriginalPointer          func(sid string, source string, items []string) error
	InsertArrayParameterVarcharStringPointerOriginal          func(sid string, source string, items []*string) error
	InsertArrayParameterVarcharStringPointerPointer           func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharString                            func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharStringOriginalPointer             func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharStringPointerOriginal             func(sid string, source string, items []*string) error
	InsertArrayEntityVarcharStringPointerPointer              func(sid string, source string, items []*string) error
	InsertParameterIntegerInt8                                func(sid string, source string, var_integer int8) (int, error)
	InsertParameterIntegerInt8OriginalPointer                 func(sid string, source string, var_integer int8) (*int, error)
	InsertParameterIntegerInt8PointerOriginal                 func(sid *string, source *string, var_integer *int8) (int, error)
	InsertParameterIntegerInt8PointerPointer                  func(sid string, source string, var_integer *int8) (*int, error)
	InsertEntityIntegerInt8                                   func(item IntegerOriginal) error
	InsertEntityIntegerInt8OriginalPointer                    func(item IntegerOriginal) error
	InsertEntityIntegerInt8PointerOriginal                    func(item *IntegerPointer) error
	InsertEntityIntegerInt8PointerPointer                     func(item *IntegerPointer) error
	InsertArrayParameterIntegerInt8                           func(sid string, source string, items []int8) error
	InsertArrayParameterIntegerInt8OriginalPointer            func(sid string, source string, items []int8) error
	InsertArrayParameterIntegerInt8PointerOriginal            func(sid string, source string, items []*int8) error
	InsertArrayParameterIntegerInt8PointerPointer             func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8                              func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8OriginalPointer               func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8PointerOriginal               func(sid string, source string, items []*int8) error
	InsertArrayEntityIntegerInt8PointerPointer                func(sid string, source string, items []*int8) error
	InsertParameterIntInt8                                    func(sid string, source string, var_int int8) (int, error)
	InsertParameterIntInt8OriginalPointer                     func(sid string, source string, var_int int8) (*int, error)
	InsertParameterIntInt8PointerOriginal                     func(sid *string, source *string, var_int *int8) (int, error)
	InsertParameterIntInt8PointerPointer                      func(sid string, source string, var_int *int8) (*int, error)
	InsertEntityIntInt8                                       func(item IntOriginal) error
	InsertEntityIntInt8OriginalPointer                        func(item IntOriginal) error
	InsertEntityIntInt8PointerOriginal                        func(item *IntPointer) error
	InsertEntityIntInt8PointerPointer                         func(item *IntPointer) error
	InsertArrayParameterIntInt8                               func(sid string, source string, items []int8) error
	InsertArrayParameterIntInt8OriginalPointer                func(sid string, source string, items []int8) error
	InsertArrayParameterIntInt8PointerOriginal                func(sid string, source string, items []*int8) error
	InsertArrayParameterIntInt8PointerPointer                 func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8                                  func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8OriginalPointer                   func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8PointerOriginal                   func(sid string, source string, items []*int8) error
	InsertArrayEntityIntInt8PointerPointer                    func(sid string, source string, items []*int8) error
	InsertParameterInt4Int8                                   func(sid string, source string, var_int4 int8) (int, error)
	InsertParameterInt4Int8OriginalPointer                    func(sid string, source string, var_int4 int8) (*int, error)
	InsertParameterInt4Int8PointerOriginal                    func(sid *string, source *string, var_int4 *int8) (int, error)
	InsertParameterInt4Int8PointerPointer                     func(sid string, source string, var_int4 *int8) (*int, error)
	InsertEntityInt4Int8                                      func(item Int4Original) error
	InsertEntityInt4Int8OriginalPointer                       func(item Int4Original) error
	InsertEntityInt4Int8PointerOriginal                       func(item *Int4Pointer) error
	InsertEntityInt4Int8PointerPointer                        func(item *Int4Pointer) error
	InsertArrayParameterInt4Int8                              func(sid string, source string, items []int8) error
	InsertArrayParameterInt4Int8OriginalPointer               func(sid string, source string, items []int8) error
	InsertArrayParameterInt4Int8PointerOriginal               func(sid string, source string, items []*int8) error
	InsertArrayParameterInt4Int8PointerPointer                func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8                                 func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8OriginalPointer                  func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	InsertArrayEntityInt4Int8PointerPointer                   func(sid string, source string, items []*int8) error
	InsertParameterInt2Int8                                   func(sid string, source string, var_int2 int8) (int, error)
	InsertParameterInt2Int8OriginalPointer                    func(sid string, source string, var_int2 int8) (*int, error)
	InsertParameterInt2Int8PointerOriginal                    func(sid *string, source *string, var_int2 *int8) (int, error)
	InsertParameterInt2Int8PointerPointer                     func(sid string, source string, var_int2 *int8) (*int, error)
	InsertEntityInt2Int8                                      func(item Int2Original) error
	InsertEntityInt2Int8OriginalPointer                       func(item Int2Original) error
	InsertEntityInt2Int8PointerOriginal                       func(item *Int2Pointer) error
	InsertEntityInt2Int8PointerPointer                        func(item *Int2Pointer) error
	InsertArrayParameterInt2Int8                              func(sid string, source string, items []int8) error
	InsertArrayParameterInt2Int8OriginalPointer               func(sid string, source string, items []int8) error
	InsertArrayParameterInt2Int8PointerOriginal               func(sid string, source string, items []*int8) error
	InsertArrayParameterInt2Int8PointerPointer                func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8                                 func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8OriginalPointer                  func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8PointerOriginal                  func(sid string, source string, items []*int8) error
	InsertArrayEntityInt2Int8PointerPointer                   func(sid string, source string, items []*int8) error
	InsertParameterTextString                                 func(sid string, source string, var_text string) (int, error)
	InsertParameterTextStringOriginalPointer                  func(sid string, source string, var_text string) (*int, error)
	InsertParameterTextStringPointerOriginal                  func(sid *string, source *string, var_text *string) (int, error)
	InsertParameterTextStringPointerPointer                   func(sid string, source string, var_text *string) (*int, error)
	InsertEntityTextString                                    func(item TextOriginal) error
	InsertEntityTextStringOriginalPointer                     func(item TextOriginal) error
	InsertEntityTextStringPointerOriginal                     func(item *TextPointer) error
	InsertEntityTextStringPointerPointer                      func(item *TextPointer) error
	InsertArrayParameterTextString                            func(sid string, source string, items []string) error
	InsertArrayParameterTextStringOriginalPointer             func(sid string, source string, items []string) error
	InsertArrayParameterTextStringPointerOriginal             func(sid string, source string, items []*string) error
	InsertArrayParameterTextStringPointerPointer              func(sid string, source string, items []*string) error
	InsertArrayEntityTextString                               func(sid string, source string, items []*string) error
	InsertArrayEntityTextStringOriginalPointer                func(sid string, source string, items []*string) error
	InsertArrayEntityTextStringPointerOriginal                func(sid string, source string, items []*string) error
	InsertArrayEntityTextStringPointerPointer                 func(sid string, source string, items []*string) error
	SelectParameterBigintInt64                                func(sid string) (int64, error)
	SelectParameterBigintInt64OriginalPointer                 func(sid string) (*int64, error)
	SelectParameterBigintInt64PointerOriginal                 func(sid string) (int64, error)
	SelectParameterBigintInt64PointerPointer                  func(sid string) (*int64, error)
	SelectEntityBigintInt64                                   func(sid string) (*int64, error)
	SelectEntityBigintInt64OriginalPointer                    func(sid string) (*int64, error)
	SelectEntityBigintInt64PointerOriginal                    func(sid string) (*int64, error)
	SelectEntityBigintInt64PointerPointer                     func(sid string) (*int64, error)
	SelectArrayParameterBigintInt64                           func(id int) ([]int64, error)
	SelectArrayParameterBigintInt64OriginalPointer            func(id int) ([]int64, error)
	SelectArrayParameterBigintInt64PointerOriginal            func(id int) ([]*int64, error)
	SelectArrayParameterBigintInt64PointerPointer             func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64                              func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64OriginalPointer               func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64PointerOriginal               func(id int) ([]*int64, error)
	SelectArrayEntityBigintInt64PointerPointer                func(id int) ([]*int64, error)
	SelectParameterInt8Int8                                   func(sid string) (int8, error)
	SelectParameterInt8Int8OriginalPointer                    func(sid string) (*int8, error)
	SelectParameterInt8Int8PointerOriginal                    func(sid string) (int8, error)
	SelectParameterInt8Int8PointerPointer                     func(sid string) (*int8, error)
	SelectEntityInt8Int8                                      func(sid string) (*int8, error)
	SelectEntityInt8Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectEntityInt8Int8PointerOriginal                       func(sid string) (*int8, error)
	SelectEntityInt8Int8PointerPointer                        func(sid string) (*int8, error)
	SelectArrayParameterInt8Int8                              func(id int) ([]int8, error)
	SelectArrayParameterInt8Int8OriginalPointer               func(id int) ([]int8, error)
	SelectArrayParameterInt8Int8PointerOriginal               func(id int) ([]*int8, error)
	SelectArrayParameterInt8Int8PointerPointer                func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8                                 func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8OriginalPointer                  func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8PointerOriginal                  func(id int) ([]*int8, error)
	SelectArrayEntityInt8Int8PointerPointer                   func(id int) ([]*int8, error)
	SelectParameterBooleanBool                                func(sid string) (bool, error)
	SelectParameterBooleanBoolOriginalPointer                 func(sid string) (*bool, error)
	SelectParameterBooleanBoolPointerOriginal                 func(sid string) (bool, error)
	SelectParameterBooleanBoolPointerPointer                  func(sid string) (*bool, error)
	SelectEntityBooleanBool                                   func(sid string) (*bool, error)
	SelectEntityBooleanBoolOriginalPointer                    func(sid string) (*bool, error)
	SelectEntityBooleanBoolPointerOriginal                    func(sid string) (*bool, error)
	SelectEntityBooleanBoolPointerPointer                     func(sid string) (*bool, error)
	SelectArrayParameterBooleanBool                           func(id int) ([]bool, error)
	SelectArrayParameterBooleanBoolOriginalPointer            func(id int) ([]bool, error)
	SelectArrayParameterBooleanBoolPointerOriginal            func(id int) ([]*bool, error)
	SelectArrayParameterBooleanBoolPointerPointer             func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBool                              func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBoolOriginalPointer               func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBoolPointerOriginal               func(id int) ([]*bool, error)
	SelectArrayEntityBooleanBoolPointerPointer                func(id int) ([]*bool, error)
	SelectParameterBoolBool                                   func(sid string) (bool, error)
	SelectParameterBoolBoolOriginalPointer                    func(sid string) (*bool, error)
	SelectParameterBoolBoolPointerOriginal                    func(sid string) (bool, error)
	SelectParameterBoolBoolPointerPointer                     func(sid string) (*bool, error)
	SelectEntityBoolBool                                      func(sid string) (*bool, error)
	SelectEntityBoolBoolOriginalPointer                       func(sid string) (*bool, error)
	SelectEntityBoolBoolPointerOriginal                       func(sid string) (*bool, error)
	SelectEntityBoolBoolPointerPointer                        func(sid string) (*bool, error)
	SelectArrayParameterBoolBool                              func(id int) ([]bool, error)
	SelectArrayParameterBoolBoolOriginalPointer               func(id int) ([]bool, error)
	SelectArrayParameterBoolBoolPointerOriginal               func(id int) ([]*bool, error)
	SelectArrayParameterBoolBoolPointerPointer                func(id int) ([]*bool, error)
	SelectArrayEntityBoolBool                                 func(id int) ([]*bool, error)
	SelectArrayEntityBoolBoolOriginalPointer                  func(id int) ([]*bool, error)
	SelectArrayEntityBoolBoolPointerOriginal                  func(id int) ([]*bool, error)
	SelectArrayEntityBoolBoolPointerPointer                   func(id int) ([]*bool, error)
	SelectParameterCharacterString                            func(sid string) (string, error)
	SelectParameterCharacterStringOriginalPointer             func(sid string) (*string, error)
	SelectParameterCharacterStringPointerOriginal             func(sid string) (string, error)
	SelectParameterCharacterStringPointerPointer              func(sid string) (*string, error)
	SelectEntityCharacterString                               func(sid string) (*string, error)
	SelectEntityCharacterStringOriginalPointer                func(sid string) (*string, error)
	SelectEntityCharacterStringPointerOriginal                func(sid string) (*string, error)
	SelectEntityCharacterStringPointerPointer                 func(sid string) (*string, error)
	SelectArrayParameterCharacterString                       func(id int) ([]string, error)
	SelectArrayParameterCharacterStringOriginalPointer        func(id int) ([]string, error)
	SelectArrayParameterCharacterStringPointerOriginal        func(id int) ([]*string, error)
	SelectArrayParameterCharacterStringPointerPointer         func(id int) ([]*string, error)
	SelectArrayEntityCharacterString                          func(id int) ([]*string, error)
	SelectArrayEntityCharacterStringOriginalPointer           func(id int) ([]*string, error)
	SelectArrayEntityCharacterStringPointerOriginal           func(id int) ([]*string, error)
	SelectArrayEntityCharacterStringPointerPointer            func(id int) ([]*string, error)
	SelectParameterCharString                                 func(sid string) (string, error)
	SelectParameterCharStringOriginalPointer                  func(sid string) (*string, error)
	SelectParameterCharStringPointerOriginal                  func(sid string) (string, error)
	SelectParameterCharStringPointerPointer                   func(sid string) (*string, error)
	SelectEntityCharString                                    func(sid string) (*string, error)
	SelectEntityCharStringOriginalPointer                     func(sid string) (*string, error)
	SelectEntityCharStringPointerOriginal                     func(sid string) (*string, error)
	SelectEntityCharStringPointerPointer                      func(sid string) (*string, error)
	SelectArrayParameterCharString                            func(id int) ([]string, error)
	SelectArrayParameterCharStringOriginalPointer             func(id int) ([]string, error)
	SelectArrayParameterCharStringPointerOriginal             func(id int) ([]*string, error)
	SelectArrayParameterCharStringPointerPointer              func(id int) ([]*string, error)
	SelectArrayEntityCharString                               func(id int) ([]*string, error)
	SelectArrayEntityCharStringOriginalPointer                func(id int) ([]*string, error)
	SelectArrayEntityCharStringPointerOriginal                func(id int) ([]*string, error)
	SelectArrayEntityCharStringPointerPointer                 func(id int) ([]*string, error)
	SelectParameterCharacterVaryingString                     func(sid string) (string, error)
	SelectParameterCharacterVaryingStringOriginalPointer      func(sid string) (*string, error)
	SelectParameterCharacterVaryingStringPointerOriginal      func(sid string) (string, error)
	SelectParameterCharacterVaryingStringPointerPointer       func(sid string) (*string, error)
	SelectEntityCharacterVaryingString                        func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringOriginalPointer         func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringPointerOriginal         func(sid string) (*string, error)
	SelectEntityCharacterVaryingStringPointerPointer          func(sid string) (*string, error)
	SelectArrayParameterCharacterVaryingString                func(id int) ([]string, error)
	SelectArrayParameterCharacterVaryingStringOriginalPointer func(id int) ([]string, error)
	SelectArrayParameterCharacterVaryingStringPointerOriginal func(id int) ([]*string, error)
	SelectArrayParameterCharacterVaryingStringPointerPointer  func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingString                   func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringOriginalPointer    func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringPointerOriginal    func(id int) ([]*string, error)
	SelectArrayEntityCharacterVaryingStringPointerPointer     func(id int) ([]*string, error)
	SelectParameterVarcharString                              func(sid string) (string, error)
	SelectParameterVarcharStringOriginalPointer               func(sid string) (*string, error)
	SelectParameterVarcharStringPointerOriginal               func(sid string) (string, error)
	SelectParameterVarcharStringPointerPointer                func(sid string) (*string, error)
	SelectEntityVarcharString                                 func(sid string) (*string, error)
	SelectEntityVarcharStringOriginalPointer                  func(sid string) (*string, error)
	SelectEntityVarcharStringPointerOriginal                  func(sid string) (*string, error)
	SelectEntityVarcharStringPointerPointer                   func(sid string) (*string, error)
	SelectArrayParameterVarcharString                         func(id int) ([]string, error)
	SelectArrayParameterVarcharStringOriginalPointer          func(id int) ([]string, error)
	SelectArrayParameterVarcharStringPointerOriginal          func(id int) ([]*string, error)
	SelectArrayParameterVarcharStringPointerPointer           func(id int) ([]*string, error)
	SelectArrayEntityVarcharString                            func(id int) ([]*string, error)
	SelectArrayEntityVarcharStringOriginalPointer             func(id int) ([]*string, error)
	SelectArrayEntityVarcharStringPointerOriginal             func(id int) ([]*string, error)
	SelectArrayEntityVarcharStringPointerPointer              func(id int) ([]*string, error)
	SelectParameterIntegerInt8                                func(sid string) (int8, error)
	SelectParameterIntegerInt8OriginalPointer                 func(sid string) (*int8, error)
	SelectParameterIntegerInt8PointerOriginal                 func(sid string) (int8, error)
	SelectParameterIntegerInt8PointerPointer                  func(sid string) (*int8, error)
	SelectEntityIntegerInt8                                   func(sid string) (*int8, error)
	SelectEntityIntegerInt8OriginalPointer                    func(sid string) (*int8, error)
	SelectEntityIntegerInt8PointerOriginal                    func(sid string) (*int8, error)
	SelectEntityIntegerInt8PointerPointer                     func(sid string) (*int8, error)
	SelectArrayParameterIntegerInt8                           func(id int) ([]int8, error)
	SelectArrayParameterIntegerInt8OriginalPointer            func(id int) ([]int8, error)
	SelectArrayParameterIntegerInt8PointerOriginal            func(id int) ([]*int8, error)
	SelectArrayParameterIntegerInt8PointerPointer             func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8                              func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8OriginalPointer               func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8PointerOriginal               func(id int) ([]*int8, error)
	SelectArrayEntityIntegerInt8PointerPointer                func(id int) ([]*int8, error)
	SelectParameterIntInt8                                    func(sid string) (int8, error)
	SelectParameterIntInt8OriginalPointer                     func(sid string) (*int8, error)
	SelectParameterIntInt8PointerOriginal                     func(sid string) (int8, error)
	SelectParameterIntInt8PointerPointer                      func(sid string) (*int8, error)
	SelectEntityIntInt8                                       func(sid string) (*int8, error)
	SelectEntityIntInt8OriginalPointer                        func(sid string) (*int8, error)
	SelectEntityIntInt8PointerOriginal                        func(sid string) (*int8, error)
	SelectEntityIntInt8PointerPointer                         func(sid string) (*int8, error)
	SelectArrayParameterIntInt8                               func(id int) ([]int8, error)
	SelectArrayParameterIntInt8OriginalPointer                func(id int) ([]int8, error)
	SelectArrayParameterIntInt8PointerOriginal                func(id int) ([]*int8, error)
	SelectArrayParameterIntInt8PointerPointer                 func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8                                  func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8OriginalPointer                   func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8PointerOriginal                   func(id int) ([]*int8, error)
	SelectArrayEntityIntInt8PointerPointer                    func(id int) ([]*int8, error)
	SelectParameterInt4Int8                                   func(sid string) (int8, error)
	SelectParameterInt4Int8OriginalPointer                    func(sid string) (*int8, error)
	SelectParameterInt4Int8PointerOriginal                    func(sid string) (int8, error)
	SelectParameterInt4Int8PointerPointer                     func(sid string) (*int8, error)
	SelectEntityInt4Int8                                      func(sid string) (*int8, error)
	SelectEntityInt4Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectEntityInt4Int8PointerOriginal                       func(sid string) (*int8, error)
	SelectEntityInt4Int8PointerPointer                        func(sid string) (*int8, error)
	SelectArrayParameterInt4Int8                              func(id int) ([]int8, error)
	SelectArrayParameterInt4Int8OriginalPointer               func(id int) ([]int8, error)
	SelectArrayParameterInt4Int8PointerOriginal               func(id int) ([]*int8, error)
	SelectArrayParameterInt4Int8PointerPointer                func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8                                 func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8OriginalPointer                  func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8PointerOriginal                  func(id int) ([]*int8, error)
	SelectArrayEntityInt4Int8PointerPointer                   func(id int) ([]*int8, error)
	SelectParameterInt2Int8                                   func(sid string) (int8, error)
	SelectParameterInt2Int8OriginalPointer                    func(sid string) (*int8, error)
	SelectParameterInt2Int8PointerOriginal                    func(sid string) (int8, error)
	SelectParameterInt2Int8PointerPointer                     func(sid string) (*int8, error)
	SelectEntityInt2Int8                                      func(sid string) (*int8, error)
	SelectEntityInt2Int8OriginalPointer                       func(sid string) (*int8, error)
	SelectEntityInt2Int8PointerOriginal                       func(sid string) (*int8, error)
	SelectEntityInt2Int8PointerPointer                        func(sid string) (*int8, error)
	SelectArrayParameterInt2Int8                              func(id int) ([]int8, error)
	SelectArrayParameterInt2Int8OriginalPointer               func(id int) ([]int8, error)
	SelectArrayParameterInt2Int8PointerOriginal               func(id int) ([]*int8, error)
	SelectArrayParameterInt2Int8PointerPointer                func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8                                 func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8OriginalPointer                  func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8PointerOriginal                  func(id int) ([]*int8, error)
	SelectArrayEntityInt2Int8PointerPointer                   func(id int) ([]*int8, error)
	SelectParameterTextString                                 func(sid string) (string, error)
	SelectParameterTextStringOriginalPointer                  func(sid string) (*string, error)
	SelectParameterTextStringPointerOriginal                  func(sid string) (string, error)
	SelectParameterTextStringPointerPointer                   func(sid string) (*string, error)
	SelectEntityTextString                                    func(sid string) (*string, error)
	SelectEntityTextStringOriginalPointer                     func(sid string) (*string, error)
	SelectEntityTextStringPointerOriginal                     func(sid string) (*string, error)
	SelectEntityTextStringPointerPointer                      func(sid string) (*string, error)
	SelectArrayParameterTextString                            func(id int) ([]string, error)
	SelectArrayParameterTextStringOriginalPointer             func(id int) ([]string, error)
	SelectArrayParameterTextStringPointerOriginal             func(id int) ([]*string, error)
	SelectArrayParameterTextStringPointerPointer              func(id int) ([]*string, error)
	SelectArrayEntityTextString                               func(id int) ([]*string, error)
	SelectArrayEntityTextStringOriginalPointer                func(id int) ([]*string, error)
	SelectArrayEntityTextStringPointerOriginal                func(id int) ([]*string, error)
	SelectArrayEntityTextStringPointerPointer                 func(id int) ([]*string, error)
	UpdateParameterBigintInt64                                func(sid string, source string, var_bigint int64) (int, error)
	UpdateParameterBigintInt64OriginalPointer                 func(sid string, source string, var_bigint int64) (*int, error)
	UpdateParameterBigintInt64PointerOriginal                 func(sid *string, source *string, var_bigint *int64) (int, error)
	UpdateParameterBigintInt64PointerPointer                  func(sid string, source string, var_bigint *int64) (*int, error)
	UpdateEntityBigintInt64                                   func(item BigintOriginal) error
	UpdateEntityBigintInt64OriginalPointer                    func(item BigintOriginal) error
	UpdateEntityBigintInt64PointerOriginal                    func(item *BigintPointer) error
	UpdateEntityBigintInt64PointerPointer                     func(item *BigintPointer) error
	UpdateArrayParameterBigintInt64                           func(sid string, source string, items []int64) error
	UpdateArrayParameterBigintInt64OriginalPointer            func(sid string, source string, items []int64) error
	UpdateArrayParameterBigintInt64PointerOriginal            func(sid string, source string, items []*int64) error
	UpdateArrayParameterBigintInt64PointerPointer             func(sid string, source string, items []*int64) error
	UpdateArrayEntityBigintInt64                              func(item BigintOriginal) error
	UpdateArrayEntityBigintInt64OriginalPointer               func(item BigintPointer) error
	UpdateArrayEntityBigintInt64PointerOriginal               func(item BigintOriginal) error
	UpdateArrayEntityBigintInt64PointerPointer                func(item BigintPointer) error
	UpdateParameterInt8Int8                                   func(sid string, source string, var_int8 int8) (int, error)
	UpdateParameterInt8Int8OriginalPointer                    func(sid string, source string, var_int8 int8) (*int, error)
	UpdateParameterInt8Int8PointerOriginal                    func(sid *string, source *string, var_int8 *int8) (int, error)
	UpdateParameterInt8Int8PointerPointer                     func(sid string, source string, var_int8 *int8) (*int, error)
	UpdateEntityInt8Int8                                      func(item Int8Original) error
	UpdateEntityInt8Int8OriginalPointer                       func(item Int8Original) error
	UpdateEntityInt8Int8PointerOriginal                       func(item *Int8Pointer) error
	UpdateEntityInt8Int8PointerPointer                        func(item *Int8Pointer) error
	UpdateArrayParameterInt8Int8                              func(sid string, source string, items []int8) error
	UpdateArrayParameterInt8Int8OriginalPointer               func(sid string, source string, items []int8) error
	UpdateArrayParameterInt8Int8PointerOriginal               func(sid string, source string, items []*int8) error
	UpdateArrayParameterInt8Int8PointerPointer                func(sid string, source string, items []*int8) error
	UpdateArrayEntityInt8Int8                                 func(item Int8Original) error
	UpdateArrayEntityInt8Int8OriginalPointer                  func(item Int8Pointer) error
	UpdateArrayEntityInt8Int8PointerOriginal                  func(item Int8Original) error
	UpdateArrayEntityInt8Int8PointerPointer                   func(item Int8Pointer) error
	UpdateParameterBooleanBool                                func(sid string, source string, var_boolean bool) (int, error)
	UpdateParameterBooleanBoolOriginalPointer                 func(sid string, source string, var_boolean bool) (*int, error)
	UpdateParameterBooleanBoolPointerOriginal                 func(sid *string, source *string, var_boolean *bool) (int, error)
	UpdateParameterBooleanBoolPointerPointer                  func(sid string, source string, var_boolean *bool) (*int, error)
	UpdateEntityBooleanBool                                   func(item BooleanOriginal) error
	UpdateEntityBooleanBoolOriginalPointer                    func(item BooleanOriginal) error
	UpdateEntityBooleanBoolPointerOriginal                    func(item *BooleanPointer) error
	UpdateEntityBooleanBoolPointerPointer                     func(item *BooleanPointer) error
	UpdateArrayParameterBooleanBool                           func(sid string, source string, items []bool) error
	UpdateArrayParameterBooleanBoolOriginalPointer            func(sid string, source string, items []bool) error
	UpdateArrayParameterBooleanBoolPointerOriginal            func(sid string, source string, items []*bool) error
	UpdateArrayParameterBooleanBoolPointerPointer             func(sid string, source string, items []*bool) error
	UpdateArrayEntityBooleanBool                              func(item BooleanOriginal) error
	UpdateArrayEntityBooleanBoolOriginalPointer               func(item BooleanPointer) error
	UpdateArrayEntityBooleanBoolPointerOriginal               func(item BooleanOriginal) error
	UpdateArrayEntityBooleanBoolPointerPointer                func(item BooleanPointer) error
	UpdateParameterBoolBool                                   func(sid string, source string, var_bool bool) (int, error)
	UpdateParameterBoolBoolOriginalPointer                    func(sid string, source string, var_bool bool) (*int, error)
	UpdateParameterBoolBoolPointerOriginal                    func(sid *string, source *string, var_bool *bool) (int, error)
	UpdateParameterBoolBoolPointerPointer                     func(sid string, source string, var_bool *bool) (*int, error)
	UpdateEntityBoolBool                                      func(item BoolOriginal) error
	UpdateEntityBoolBoolOriginalPointer                       func(item BoolOriginal) error
	UpdateEntityBoolBoolPointerOriginal                       func(item *BoolPointer) error
	UpdateEntityBoolBoolPointerPointer                        func(item *BoolPointer) error
	UpdateArrayParameterBoolBool                              func(sid string, source string, items []bool) error
	UpdateArrayParameterBoolBoolOriginalPointer               func(sid string, source string, items []bool) error
	UpdateArrayParameterBoolBoolPointerOriginal               func(sid string, source string, items []*bool) error
	UpdateArrayParameterBoolBoolPointerPointer                func(sid string, source string, items []*bool) error
	UpdateArrayEntityBoolBool                                 func(item BoolOriginal) error
	UpdateArrayEntityBoolBoolOriginalPointer                  func(item BoolPointer) error
	UpdateArrayEntityBoolBoolPointerOriginal                  func(item BoolOriginal) error
	UpdateArrayEntityBoolBoolPointerPointer                   func(item BoolPointer) error
	UpdateParameterCharacterString                            func(sid string, source string, var_character string) (int, error)
	UpdateParameterCharacterStringOriginalPointer             func(sid string, source string, var_character string) (*int, error)
	UpdateParameterCharacterStringPointerOriginal             func(sid *string, source *string, var_character *string) (int, error)
	UpdateParameterCharacterStringPointerPointer              func(sid string, source string, var_character *string) (*int, error)
	UpdateEntityCharacterString                               func(item CharacterOriginal) error
	UpdateEntityCharacterStringOriginalPointer                func(item CharacterOriginal) error
	UpdateEntityCharacterStringPointerOriginal                func(item *CharacterPointer) error
	UpdateEntityCharacterStringPointerPointer                 func(item *CharacterPointer) error
	UpdateArrayParameterCharacterString                       func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterStringOriginalPointer        func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterStringPointerOriginal        func(sid string, source string, items []*string) error
	UpdateArrayParameterCharacterStringPointerPointer         func(sid string, source string, items []*string) error
	UpdateArrayEntityCharacterString                          func(item CharacterOriginal) error
	UpdateArrayEntityCharacterStringOriginalPointer           func(item CharacterPointer) error
	UpdateArrayEntityCharacterStringPointerOriginal           func(item CharacterOriginal) error
	UpdateArrayEntityCharacterStringPointerPointer            func(item CharacterPointer) error
	UpdateParameterCharString                                 func(sid string, source string, var_char string) (int, error)
	UpdateParameterCharStringOriginalPointer                  func(sid string, source string, var_char string) (*int, error)
	UpdateParameterCharStringPointerOriginal                  func(sid *string, source *string, var_char *string) (int, error)
	UpdateParameterCharStringPointerPointer                   func(sid string, source string, var_char *string) (*int, error)
	UpdateEntityCharString                                    func(item CharOriginal) error
	UpdateEntityCharStringOriginalPointer                     func(item CharOriginal) error
	UpdateEntityCharStringPointerOriginal                     func(item *CharPointer) error
	UpdateEntityCharStringPointerPointer                      func(item *CharPointer) error
	UpdateArrayParameterCharString                            func(sid string, source string, items []string) error
	UpdateArrayParameterCharStringOriginalPointer             func(sid string, source string, items []string) error
	UpdateArrayParameterCharStringPointerOriginal             func(sid string, source string, items []*string) error
	UpdateArrayParameterCharStringPointerPointer              func(sid string, source string, items []*string) error
	UpdateArrayEntityCharString                               func(item CharOriginal) error
	UpdateArrayEntityCharStringOriginalPointer                func(item CharPointer) error
	UpdateArrayEntityCharStringPointerOriginal                func(item CharOriginal) error
	UpdateArrayEntityCharStringPointerPointer                 func(item CharPointer) error
	UpdateParameterCharacterVaryingString                     func(sid string, source string, var_character_varying string) (int, error)
	UpdateParameterCharacterVaryingStringOriginalPointer      func(sid string, source string, var_character_varying string) (*int, error)
	UpdateParameterCharacterVaryingStringPointerOriginal      func(sid *string, source *string, var_character_varying *string) (int, error)
	UpdateParameterCharacterVaryingStringPointerPointer       func(sid string, source string, var_character_varying *string) (*int, error)
	UpdateEntityCharacterVaryingString                        func(item CharacterVaryingOriginal) error
	UpdateEntityCharacterVaryingStringOriginalPointer         func(item CharacterVaryingOriginal) error
	UpdateEntityCharacterVaryingStringPointerOriginal         func(item *CharacterVaryingPointer) error
	UpdateEntityCharacterVaryingStringPointerPointer          func(item *CharacterVaryingPointer) error
	UpdateArrayParameterCharacterVaryingString                func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterVaryingStringOriginalPointer func(sid string, source string, items []string) error
	UpdateArrayParameterCharacterVaryingStringPointerOriginal func(sid string, source string, items []*string) error
	UpdateArrayParameterCharacterVaryingStringPointerPointer  func(sid string, source string, items []*string) error
	UpdateArrayEntityCharacterVaryingString                   func(item CharacterVaryingOriginal) error
	UpdateArrayEntityCharacterVaryingStringOriginalPointer    func(item CharacterVaryingPointer) error
	UpdateArrayEntityCharacterVaryingStringPointerOriginal    func(item CharacterVaryingOriginal) error
	UpdateArrayEntityCharacterVaryingStringPointerPointer     func(item CharacterVaryingPointer) error
	UpdateParameterVarcharString                              func(sid string, source string, var_varchar string) (int, error)
	UpdateParameterVarcharStringOriginalPointer               func(sid string, source string, var_varchar string) (*int, error)
	UpdateParameterVarcharStringPointerOriginal               func(sid *string, source *string, var_varchar *string) (int, error)
	UpdateParameterVarcharStringPointerPointer                func(sid string, source string, var_varchar *string) (*int, error)
	UpdateEntityVarcharString                                 func(item VarcharOriginal) error
	UpdateEntityVarcharStringOriginalPointer                  func(item VarcharOriginal) error
	UpdateEntityVarcharStringPointerOriginal                  func(item *VarcharPointer) error
	UpdateEntityVarcharStringPointerPointer                   func(item *VarcharPointer) error
	UpdateArrayParameterVarcharString                         func(sid string, source string, items []string) error
	UpdateArrayParameterVarcharStringOriginalPointer          func(sid string, source string, items []string) error
	UpdateArrayParameterVarcharStringPointerOriginal          func(sid string, source string, items []*string) error
	UpdateArrayParameterVarcharStringPointerPointer           func(sid string, source string, items []*string) error
	UpdateArrayEntityVarcharString                            func(item VarcharOriginal) error
	UpdateArrayEntityVarcharStringOriginalPointer             func(item VarcharPointer) error
	UpdateArrayEntityVarcharStringPointerOriginal             func(item VarcharOriginal) error
	UpdateArrayEntityVarcharStringPointerPointer              func(item VarcharPointer) error
	UpdateParameterIntegerInt8                                func(sid string, source string, var_integer int8) (int, error)
	UpdateParameterIntegerInt8OriginalPointer                 func(sid string, source string, var_integer int8) (*int, error)
	UpdateParameterIntegerInt8PointerOriginal                 func(sid *string, source *string, var_integer *int8) (int, error)
	UpdateParameterIntegerInt8PointerPointer                  func(sid string, source string, var_integer *int8) (*int, error)
	UpdateEntityIntegerInt8                                   func(item IntegerOriginal) error
	UpdateEntityIntegerInt8OriginalPointer                    func(item IntegerOriginal) error
	UpdateEntityIntegerInt8PointerOriginal                    func(item *IntegerPointer) error
	UpdateEntityIntegerInt8PointerPointer                     func(item *IntegerPointer) error
	UpdateArrayParameterIntegerInt8                           func(sid string, source string, items []int8) error
	UpdateArrayParameterIntegerInt8OriginalPointer            func(sid string, source string, items []int8) error
	UpdateArrayParameterIntegerInt8PointerOriginal            func(sid string, source string, items []*int8) error
	UpdateArrayParameterIntegerInt8PointerPointer             func(sid string, source string, items []*int8) error
	UpdateArrayEntityIntegerInt8                              func(item IntegerOriginal) error
	UpdateArrayEntityIntegerInt8OriginalPointer               func(item IntegerPointer) error
	UpdateArrayEntityIntegerInt8PointerOriginal               func(item IntegerOriginal) error
	UpdateArrayEntityIntegerInt8PointerPointer                func(item IntegerPointer) error
	UpdateParameterIntInt8                                    func(sid string, source string, var_int int8) (int, error)
	UpdateParameterIntInt8OriginalPointer                     func(sid string, source string, var_int int8) (*int, error)
	UpdateParameterIntInt8PointerOriginal                     func(sid *string, source *string, var_int *int8) (int, error)
	UpdateParameterIntInt8PointerPointer                      func(sid string, source string, var_int *int8) (*int, error)
	UpdateEntityIntInt8                                       func(item IntOriginal) error
	UpdateEntityIntInt8OriginalPointer                        func(item IntOriginal) error
	UpdateEntityIntInt8PointerOriginal                        func(item *IntPointer) error
	UpdateEntityIntInt8PointerPointer                         func(item *IntPointer) error
	UpdateArrayParameterIntInt8                               func(sid string, source string, items []int8) error
	UpdateArrayParameterIntInt8OriginalPointer                func(sid string, source string, items []int8) error
	UpdateArrayParameterIntInt8PointerOriginal                func(sid string, source string, items []*int8) error
	UpdateArrayParameterIntInt8PointerPointer                 func(sid string, source string, items []*int8) error
	UpdateArrayEntityIntInt8                                  func(item IntOriginal) error
	UpdateArrayEntityIntInt8OriginalPointer                   func(item IntPointer) error
	UpdateArrayEntityIntInt8PointerOriginal                   func(item IntOriginal) error
	UpdateArrayEntityIntInt8PointerPointer                    func(item IntPointer) error
	UpdateParameterInt4Int8                                   func(sid string, source string, var_int4 int8) (int, error)
	UpdateParameterInt4Int8OriginalPointer                    func(sid string, source string, var_int4 int8) (*int, error)
	UpdateParameterInt4Int8PointerOriginal                    func(sid *string, source *string, var_int4 *int8) (int, error)
	UpdateParameterInt4Int8PointerPointer                     func(sid string, source string, var_int4 *int8) (*int, error)
	UpdateEntityInt4Int8                                      func(item Int4Original) error
	UpdateEntityInt4Int8OriginalPointer                       func(item Int4Original) error
	UpdateEntityInt4Int8PointerOriginal                       func(item *Int4Pointer) error
	UpdateEntityInt4Int8PointerPointer                        func(item *Int4Pointer) error
	UpdateArrayParameterInt4Int8                              func(sid string, source string, items []int8) error
	UpdateArrayParameterInt4Int8OriginalPointer               func(sid string, source string, items []int8) error
	UpdateArrayParameterInt4Int8PointerOriginal               func(sid string, source string, items []*int8) error
	UpdateArrayParameterInt4Int8PointerPointer                func(sid string, source string, items []*int8) error
	UpdateArrayEntityInt4Int8                                 func(item Int4Original) error
	UpdateArrayEntityInt4Int8OriginalPointer                  func(item Int4Pointer) error
	UpdateArrayEntityInt4Int8PointerOriginal                  func(item Int4Original) error
	UpdateArrayEntityInt4Int8PointerPointer                   func(item Int4Pointer) error
	UpdateParameterInt2Int8                                   func(sid string, source string, var_int2 int8) (int, error)
	UpdateParameterInt2Int8OriginalPointer                    func(sid string, source string, var_int2 int8) (*int, error)
	UpdateParameterInt2Int8PointerOriginal                    func(sid *string, source *string, var_int2 *int8) (int, error)
	UpdateParameterInt2Int8PointerPointer                     func(sid string, source string, var_int2 *int8) (*int, error)
	UpdateEntityInt2Int8                                      func(item Int2Original) error
	UpdateEntityInt2Int8OriginalPointer                       func(item Int2Original) error
	UpdateEntityInt2Int8PointerOriginal                       func(item *Int2Pointer) error
	UpdateEntityInt2Int8PointerPointer                        func(item *Int2Pointer) error
	UpdateArrayParameterInt2Int8                              func(sid string, source string, items []int8) error
	UpdateArrayParameterInt2Int8OriginalPointer               func(sid string, source string, items []int8) error
	UpdateArrayParameterInt2Int8PointerOriginal               func(sid string, source string, items []*int8) error
	UpdateArrayParameterInt2Int8PointerPointer                func(sid string, source string, items []*int8) error
	UpdateArrayEntityInt2Int8                                 func(item Int2Original) error
	UpdateArrayEntityInt2Int8OriginalPointer                  func(item Int2Pointer) error
	UpdateArrayEntityInt2Int8PointerOriginal                  func(item Int2Original) error
	UpdateArrayEntityInt2Int8PointerPointer                   func(item Int2Pointer) error
	UpdateParameterTextString                                 func(sid string, source string, var_text string) (int, error)
	UpdateParameterTextStringOriginalPointer                  func(sid string, source string, var_text string) (*int, error)
	UpdateParameterTextStringPointerOriginal                  func(sid *string, source *string, var_text *string) (int, error)
	UpdateParameterTextStringPointerPointer                   func(sid string, source string, var_text *string) (*int, error)
	UpdateEntityTextString                                    func(item TextOriginal) error
	UpdateEntityTextStringOriginalPointer                     func(item TextOriginal) error
	UpdateEntityTextStringPointerOriginal                     func(item *TextPointer) error
	UpdateEntityTextStringPointerPointer                      func(item *TextPointer) error
	UpdateArrayParameterTextString                            func(sid string, source string, items []string) error
	UpdateArrayParameterTextStringOriginalPointer             func(sid string, source string, items []string) error
	UpdateArrayParameterTextStringPointerOriginal             func(sid string, source string, items []*string) error
	UpdateArrayParameterTextStringPointerPointer              func(sid string, source string, items []*string) error
	UpdateArrayEntityTextString                               func(item TextOriginal) error
	UpdateArrayEntityTextStringOriginalPointer                func(item TextPointer) error
	UpdateArrayEntityTextStringPointerOriginal                func(item TextOriginal) error
	UpdateArrayEntityTextStringPointerPointer                 func(item TextPointer) error
	DeleteParameterBigintInt64                                func(sid string) (int, error)
	DeleteParameterBigintInt64OriginalPointer                 func(sid string) (*int, error)
	DeleteParameterBigintInt64PointerOriginal                 func(sid *string) (int, error)
	DeleteParameterBigintInt64PointerPointer                  func(sid string) (*int, error)
	DeleteEntityBigintInt64                                   func(item BigintOriginal) error
	DeleteEntityBigintInt64OriginalPointer                    func(item BigintOriginal) error
	DeleteEntityBigintInt64PointerOriginal                    func(item *BigintPointer) error
	DeleteEntityBigintInt64PointerPointer                     func(item *BigintPointer) error
	DeleteArrayParameterBigintInt64                           func(id int) (int, error)
	DeleteArrayParameterBigintInt64OriginalPointer            func(id int) (int, error)
	DeleteArrayParameterBigintInt64PointerOriginal            func(id int) (int, error)
	DeleteArrayParameterBigintInt64PointerPointer             func(id int) (int, error)
	DeleteArrayEntityBigintInt64                              func(id int) (int, error)
	DeleteArrayEntityBigintInt64OriginalPointer               func(id int) (int, error)
	DeleteArrayEntityBigintInt64PointerOriginal               func(id int) (int, error)
	DeleteArrayEntityBigintInt64PointerPointer                func(id int) (int, error)
	DeleteParameterInt8Int8                                   func(sid string) (int, error)
	DeleteParameterInt8Int8OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterInt8Int8PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterInt8Int8PointerPointer                     func(sid string) (*int, error)
	DeleteEntityInt8Int8                                      func(item Int8Original) error
	DeleteEntityInt8Int8OriginalPointer                       func(item Int8Original) error
	DeleteEntityInt8Int8PointerOriginal                       func(item *Int8Pointer) error
	DeleteEntityInt8Int8PointerPointer                        func(item *Int8Pointer) error
	DeleteArrayParameterInt8Int8                              func(id int) (int, error)
	DeleteArrayParameterInt8Int8OriginalPointer               func(id int) (int, error)
	DeleteArrayParameterInt8Int8PointerOriginal               func(id int) (int, error)
	DeleteArrayParameterInt8Int8PointerPointer                func(id int) (int, error)
	DeleteArrayEntityInt8Int8                                 func(id int) (int, error)
	DeleteArrayEntityInt8Int8OriginalPointer                  func(id int) (int, error)
	DeleteArrayEntityInt8Int8PointerOriginal                  func(id int) (int, error)
	DeleteArrayEntityInt8Int8PointerPointer                   func(id int) (int, error)
	DeleteParameterBooleanBool                                func(sid string) (int, error)
	DeleteParameterBooleanBoolOriginalPointer                 func(sid string) (*int, error)
	DeleteParameterBooleanBoolPointerOriginal                 func(sid *string) (int, error)
	DeleteParameterBooleanBoolPointerPointer                  func(sid string) (*int, error)
	DeleteEntityBooleanBool                                   func(item BooleanOriginal) error
	DeleteEntityBooleanBoolOriginalPointer                    func(item BooleanOriginal) error
	DeleteEntityBooleanBoolPointerOriginal                    func(item *BooleanPointer) error
	DeleteEntityBooleanBoolPointerPointer                     func(item *BooleanPointer) error
	DeleteArrayParameterBooleanBool                           func(id int) (int, error)
	DeleteArrayParameterBooleanBoolOriginalPointer            func(id int) (int, error)
	DeleteArrayParameterBooleanBoolPointerOriginal            func(id int) (int, error)
	DeleteArrayParameterBooleanBoolPointerPointer             func(id int) (int, error)
	DeleteArrayEntityBooleanBool                              func(id int) (int, error)
	DeleteArrayEntityBooleanBoolOriginalPointer               func(id int) (int, error)
	DeleteArrayEntityBooleanBoolPointerOriginal               func(id int) (int, error)
	DeleteArrayEntityBooleanBoolPointerPointer                func(id int) (int, error)
	DeleteParameterBoolBool                                   func(sid string) (int, error)
	DeleteParameterBoolBoolOriginalPointer                    func(sid string) (*int, error)
	DeleteParameterBoolBoolPointerOriginal                    func(sid *string) (int, error)
	DeleteParameterBoolBoolPointerPointer                     func(sid string) (*int, error)
	DeleteEntityBoolBool                                      func(item BoolOriginal) error
	DeleteEntityBoolBoolOriginalPointer                       func(item BoolOriginal) error
	DeleteEntityBoolBoolPointerOriginal                       func(item *BoolPointer) error
	DeleteEntityBoolBoolPointerPointer                        func(item *BoolPointer) error
	DeleteArrayParameterBoolBool                              func(id int) (int, error)
	DeleteArrayParameterBoolBoolOriginalPointer               func(id int) (int, error)
	DeleteArrayParameterBoolBoolPointerOriginal               func(id int) (int, error)
	DeleteArrayParameterBoolBoolPointerPointer                func(id int) (int, error)
	DeleteArrayEntityBoolBool                                 func(id int) (int, error)
	DeleteArrayEntityBoolBoolOriginalPointer                  func(id int) (int, error)
	DeleteArrayEntityBoolBoolPointerOriginal                  func(id int) (int, error)
	DeleteArrayEntityBoolBoolPointerPointer                   func(id int) (int, error)
	DeleteParameterCharacterString                            func(sid string) (int, error)
	DeleteParameterCharacterStringOriginalPointer             func(sid string) (*int, error)
	DeleteParameterCharacterStringPointerOriginal             func(sid *string) (int, error)
	DeleteParameterCharacterStringPointerPointer              func(sid string) (*int, error)
	DeleteEntityCharacterString                               func(item CharacterOriginal) error
	DeleteEntityCharacterStringOriginalPointer                func(item CharacterOriginal) error
	DeleteEntityCharacterStringPointerOriginal                func(item *CharacterPointer) error
	DeleteEntityCharacterStringPointerPointer                 func(item *CharacterPointer) error
	DeleteArrayParameterCharacterString                       func(id int) (int, error)
	DeleteArrayParameterCharacterStringOriginalPointer        func(id int) (int, error)
	DeleteArrayParameterCharacterStringPointerOriginal        func(id int) (int, error)
	DeleteArrayParameterCharacterStringPointerPointer         func(id int) (int, error)
	DeleteArrayEntityCharacterString                          func(id int) (int, error)
	DeleteArrayEntityCharacterStringOriginalPointer           func(id int) (int, error)
	DeleteArrayEntityCharacterStringPointerOriginal           func(id int) (int, error)
	DeleteArrayEntityCharacterStringPointerPointer            func(id int) (int, error)
	DeleteParameterCharString                                 func(sid string) (int, error)
	DeleteParameterCharStringOriginalPointer                  func(sid string) (*int, error)
	DeleteParameterCharStringPointerOriginal                  func(sid *string) (int, error)
	DeleteParameterCharStringPointerPointer                   func(sid string) (*int, error)
	DeleteEntityCharString                                    func(item CharOriginal) error
	DeleteEntityCharStringOriginalPointer                     func(item CharOriginal) error
	DeleteEntityCharStringPointerOriginal                     func(item *CharPointer) error
	DeleteEntityCharStringPointerPointer                      func(item *CharPointer) error
	DeleteArrayParameterCharString                            func(id int) (int, error)
	DeleteArrayParameterCharStringOriginalPointer             func(id int) (int, error)
	DeleteArrayParameterCharStringPointerOriginal             func(id int) (int, error)
	DeleteArrayParameterCharStringPointerPointer              func(id int) (int, error)
	DeleteArrayEntityCharString                               func(id int) (int, error)
	DeleteArrayEntityCharStringOriginalPointer                func(id int) (int, error)
	DeleteArrayEntityCharStringPointerOriginal                func(id int) (int, error)
	DeleteArrayEntityCharStringPointerPointer                 func(id int) (int, error)
	DeleteParameterCharacterVaryingString                     func(sid string) (int, error)
	DeleteParameterCharacterVaryingStringOriginalPointer      func(sid string) (*int, error)
	DeleteParameterCharacterVaryingStringPointerOriginal      func(sid *string) (int, error)
	DeleteParameterCharacterVaryingStringPointerPointer       func(sid string) (*int, error)
	DeleteEntityCharacterVaryingString                        func(item CharacterVaryingOriginal) error
	DeleteEntityCharacterVaryingStringOriginalPointer         func(item CharacterVaryingOriginal) error
	DeleteEntityCharacterVaryingStringPointerOriginal         func(item *CharacterVaryingPointer) error
	DeleteEntityCharacterVaryingStringPointerPointer          func(item *CharacterVaryingPointer) error
	DeleteArrayParameterCharacterVaryingString                func(id int) (int, error)
	DeleteArrayParameterCharacterVaryingStringOriginalPointer func(id int) (int, error)
	DeleteArrayParameterCharacterVaryingStringPointerOriginal func(id int) (int, error)
	DeleteArrayParameterCharacterVaryingStringPointerPointer  func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingString                   func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingStringOriginalPointer    func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingStringPointerOriginal    func(id int) (int, error)
	DeleteArrayEntityCharacterVaryingStringPointerPointer     func(id int) (int, error)
	DeleteParameterVarcharString                              func(sid string) (int, error)
	DeleteParameterVarcharStringOriginalPointer               func(sid string) (*int, error)
	DeleteParameterVarcharStringPointerOriginal               func(sid *string) (int, error)
	DeleteParameterVarcharStringPointerPointer                func(sid string) (*int, error)
	DeleteEntityVarcharString                                 func(item VarcharOriginal) error
	DeleteEntityVarcharStringOriginalPointer                  func(item VarcharOriginal) error
	DeleteEntityVarcharStringPointerOriginal                  func(item *VarcharPointer) error
	DeleteEntityVarcharStringPointerPointer                   func(item *VarcharPointer) error
	DeleteArrayParameterVarcharString                         func(id int) (int, error)
	DeleteArrayParameterVarcharStringOriginalPointer          func(id int) (int, error)
	DeleteArrayParameterVarcharStringPointerOriginal          func(id int) (int, error)
	DeleteArrayParameterVarcharStringPointerPointer           func(id int) (int, error)
	DeleteArrayEntityVarcharString                            func(id int) (int, error)
	DeleteArrayEntityVarcharStringOriginalPointer             func(id int) (int, error)
	DeleteArrayEntityVarcharStringPointerOriginal             func(id int) (int, error)
	DeleteArrayEntityVarcharStringPointerPointer              func(id int) (int, error)
	DeleteParameterIntegerInt8                                func(sid string) (int, error)
	DeleteParameterIntegerInt8OriginalPointer                 func(sid string) (*int, error)
	DeleteParameterIntegerInt8PointerOriginal                 func(sid *string) (int, error)
	DeleteParameterIntegerInt8PointerPointer                  func(sid string) (*int, error)
	DeleteEntityIntegerInt8                                   func(item IntegerOriginal) error
	DeleteEntityIntegerInt8OriginalPointer                    func(item IntegerOriginal) error
	DeleteEntityIntegerInt8PointerOriginal                    func(item *IntegerPointer) error
	DeleteEntityIntegerInt8PointerPointer                     func(item *IntegerPointer) error
	DeleteArrayParameterIntegerInt8                           func(id int) (int, error)
	DeleteArrayParameterIntegerInt8OriginalPointer            func(id int) (int, error)
	DeleteArrayParameterIntegerInt8PointerOriginal            func(id int) (int, error)
	DeleteArrayParameterIntegerInt8PointerPointer             func(id int) (int, error)
	DeleteArrayEntityIntegerInt8                              func(id int) (int, error)
	DeleteArrayEntityIntegerInt8OriginalPointer               func(id int) (int, error)
	DeleteArrayEntityIntegerInt8PointerOriginal               func(id int) (int, error)
	DeleteArrayEntityIntegerInt8PointerPointer                func(id int) (int, error)
	DeleteParameterIntInt8                                    func(sid string) (int, error)
	DeleteParameterIntInt8OriginalPointer                     func(sid string) (*int, error)
	DeleteParameterIntInt8PointerOriginal                     func(sid *string) (int, error)
	DeleteParameterIntInt8PointerPointer                      func(sid string) (*int, error)
	DeleteEntityIntInt8                                       func(item IntOriginal) error
	DeleteEntityIntInt8OriginalPointer                        func(item IntOriginal) error
	DeleteEntityIntInt8PointerOriginal                        func(item *IntPointer) error
	DeleteEntityIntInt8PointerPointer                         func(item *IntPointer) error
	DeleteArrayParameterIntInt8                               func(id int) (int, error)
	DeleteArrayParameterIntInt8OriginalPointer                func(id int) (int, error)
	DeleteArrayParameterIntInt8PointerOriginal                func(id int) (int, error)
	DeleteArrayParameterIntInt8PointerPointer                 func(id int) (int, error)
	DeleteArrayEntityIntInt8                                  func(id int) (int, error)
	DeleteArrayEntityIntInt8OriginalPointer                   func(id int) (int, error)
	DeleteArrayEntityIntInt8PointerOriginal                   func(id int) (int, error)
	DeleteArrayEntityIntInt8PointerPointer                    func(id int) (int, error)
	DeleteParameterInt4Int8                                   func(sid string) (int, error)
	DeleteParameterInt4Int8OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterInt4Int8PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterInt4Int8PointerPointer                     func(sid string) (*int, error)
	DeleteEntityInt4Int8                                      func(item Int4Original) error
	DeleteEntityInt4Int8OriginalPointer                       func(item Int4Original) error
	DeleteEntityInt4Int8PointerOriginal                       func(item *Int4Pointer) error
	DeleteEntityInt4Int8PointerPointer                        func(item *Int4Pointer) error
	DeleteArrayParameterInt4Int8                              func(id int) (int, error)
	DeleteArrayParameterInt4Int8OriginalPointer               func(id int) (int, error)
	DeleteArrayParameterInt4Int8PointerOriginal               func(id int) (int, error)
	DeleteArrayParameterInt4Int8PointerPointer                func(id int) (int, error)
	DeleteArrayEntityInt4Int8                                 func(id int) (int, error)
	DeleteArrayEntityInt4Int8OriginalPointer                  func(id int) (int, error)
	DeleteArrayEntityInt4Int8PointerOriginal                  func(id int) (int, error)
	DeleteArrayEntityInt4Int8PointerPointer                   func(id int) (int, error)
	DeleteParameterInt2Int8                                   func(sid string) (int, error)
	DeleteParameterInt2Int8OriginalPointer                    func(sid string) (*int, error)
	DeleteParameterInt2Int8PointerOriginal                    func(sid *string) (int, error)
	DeleteParameterInt2Int8PointerPointer                     func(sid string) (*int, error)
	DeleteEntityInt2Int8                                      func(item Int2Original) error
	DeleteEntityInt2Int8OriginalPointer                       func(item Int2Original) error
	DeleteEntityInt2Int8PointerOriginal                       func(item *Int2Pointer) error
	DeleteEntityInt2Int8PointerPointer                        func(item *Int2Pointer) error
	DeleteArrayParameterInt2Int8                              func(id int) (int, error)
	DeleteArrayParameterInt2Int8OriginalPointer               func(id int) (int, error)
	DeleteArrayParameterInt2Int8PointerOriginal               func(id int) (int, error)
	DeleteArrayParameterInt2Int8PointerPointer                func(id int) (int, error)
	DeleteArrayEntityInt2Int8                                 func(id int) (int, error)
	DeleteArrayEntityInt2Int8OriginalPointer                  func(id int) (int, error)
	DeleteArrayEntityInt2Int8PointerOriginal                  func(id int) (int, error)
	DeleteArrayEntityInt2Int8PointerPointer                   func(id int) (int, error)
	DeleteParameterTextString                                 func(sid string) (int, error)
	DeleteParameterTextStringOriginalPointer                  func(sid string) (*int, error)
	DeleteParameterTextStringPointerOriginal                  func(sid *string) (int, error)
	DeleteParameterTextStringPointerPointer                   func(sid string) (*int, error)
	DeleteEntityTextString                                    func(item TextOriginal) error
	DeleteEntityTextStringOriginalPointer                     func(item TextOriginal) error
	DeleteEntityTextStringPointerOriginal                     func(item *TextPointer) error
	DeleteEntityTextStringPointerPointer                      func(item *TextPointer) error
	DeleteArrayParameterTextString                            func(id int) (int, error)
	DeleteArrayParameterTextStringOriginalPointer             func(id int) (int, error)
	DeleteArrayParameterTextStringPointerOriginal             func(id int) (int, error)
	DeleteArrayParameterTextStringPointerPointer              func(id int) (int, error)
	DeleteArrayEntityTextString                               func(id int) (int, error)
	DeleteArrayEntityTextStringOriginalPointer                func(id int) (int, error)
	DeleteArrayEntityTextStringPointerOriginal                func(id int) (int, error)
	DeleteArrayEntityTextStringPointerPointer                 func(id int) (int, error)
}
