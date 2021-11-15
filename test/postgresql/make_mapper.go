package postgresql

type MakeMapper struct {
	InsertParameterBigintInt64                           func(var_bigint int64) error
	InsertParameterBigintInt64PointerOriginal            func(var_bigint *int64) error
	InsertParameterCharacterString                       func(var_character string) error
	InsertParameterCharacterStringPointerOriginal        func(var_character *string) error
	InsertParameterCharacterVaryingString                func(var_character_varying string) error
	InsertParameterCharacterVaryingStringPointerOriginal func(var_character_varying *string) error
}
