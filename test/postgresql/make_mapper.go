package postgresql

type MakeMapper struct {
	InsertParameterBigintInt64                                func(var_bigint int64) error
	InsertParameterBigintInt64PointerOriginal                 func(var_bigint *int64) error
	InsertArrayParameterBigintInt64                           func(items []int64) error
	InsertArrayParameterBigintInt64PointerOriginal            func(items []*int64) error
	InsertParameterCharacterString                            func(var_character string) error
	InsertParameterCharacterStringPointerOriginal             func(var_character *string) error
	InsertArrayParameterCharacterString                       func(items []string) error
	InsertArrayParameterCharacterStringPointerOriginal        func(items []*string) error
	InsertParameterCharacterVaryingString                     func(var_character_varying string) error
	InsertParameterCharacterVaryingStringPointerOriginal      func(var_character_varying *string) error
	InsertArrayParameterCharacterVaryingString                func(items []string) error
	InsertArrayParameterCharacterVaryingStringPointerOriginal func(items []*string) error
	SelectParameterBigintInt64                                func(id int) (int64, error)
	SelectParameterCharacterString                            func(id int) (string, error)
	SelectParameterCharacterVaryingString                     func(id int) (string, error)
}
