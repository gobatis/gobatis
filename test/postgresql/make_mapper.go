package postgresql

import (
	"context"
	"database/sql"
	"github.com/gobatis/gobatis"
)

type MakeMapper struct {
	InsertParameterBigintInt64                                func(id string, source string, var_bigint int64) error
	InsertParameterBigintInt64OriginalPointer                 func(id string, source string, var_bigint int64) error
	InsertParameterBigintInt64PointerOriginal                 func(id string, source string, var_bigint *int64) error
	InsertParameterBigintInt64PointerPointer                  func(id string, source string, var_bigint *int64) error
	InsertParameterBigintInt64Must                            func(id string, source string, var_bigint int64) error
	InsertParameterBigintInt64Rows                            func(id string, source string, var_bigint int64) (int, error)
	InsertParameterBigintInt64Context                         func(ctx context.Context, id string, source string, var_bigint int64) error
	InsertParameterBigintInt64Tx                              func(ctx *sql.Tx, id string, source string, var_bigint int64) error
	InsertParameterBigintInt64Stmt                            func(id string, source string, var_bigint int64) (*gobatis.Stmt, error)
	InsertParameterBigintInt64Embed                           func(id string, source string, var_bigint int64) error
	InsertEntityBigintInt64                                   func(item BigintOriginal) error
	InsertEntityBigintInt64PointerOriginal                    func(item *BigintPointer) error
	InsertArrayParameterBigintInt64                           func(id string, source string, items []int64) error
	InsertArrayParameterBigintInt64PointerOriginal            func(id string, source string, items []*int64) error
	InsertParameterCharacterString                            func(id string, source string, var_character string) error
	InsertParameterCharacterStringOriginalPointer             func(id string, source string, var_character string) error
	InsertParameterCharacterStringPointerOriginal             func(id string, source string, var_character *string) error
	InsertParameterCharacterStringPointerPointer              func(id string, source string, var_character *string) error
	InsertParameterCharacterStringMust                        func(id string, source string, var_character string) error
	InsertParameterCharacterStringRows                        func(id string, source string, var_character string) (int, error)
	InsertParameterCharacterStringContext                     func(ctx context.Context, id string, source string, var_character string) error
	InsertParameterCharacterStringTx                          func(ctx *sql.Tx, id string, source string, var_character string) error
	InsertParameterCharacterStringStmt                        func(id string, source string, var_character string) (*gobatis.Stmt, error)
	InsertParameterCharacterStringEmbed                       func(id string, source string, var_character string) error
	InsertEntityCharacterString                               func(item CharacterOriginal) error
	InsertEntityCharacterStringPointerOriginal                func(item *CharacterPointer) error
	InsertArrayParameterCharacterString                       func(id string, source string, items []string) error
	InsertArrayParameterCharacterStringPointerOriginal        func(id string, source string, items []*string) error
	InsertParameterCharacterVaryingString                     func(id string, source string, var_character_varying string) error
	InsertParameterCharacterVaryingStringOriginalPointer      func(id string, source string, var_character_varying string) error
	InsertParameterCharacterVaryingStringPointerOriginal      func(id string, source string, var_character_varying *string) error
	InsertParameterCharacterVaryingStringPointerPointer       func(id string, source string, var_character_varying *string) error
	InsertParameterCharacterVaryingStringMust                 func(id string, source string, var_character_varying string) error
	InsertParameterCharacterVaryingStringRows                 func(id string, source string, var_character_varying string) (int, error)
	InsertParameterCharacterVaryingStringContext              func(ctx context.Context, id string, source string, var_character_varying string) error
	InsertParameterCharacterVaryingStringTx                   func(ctx *sql.Tx, id string, source string, var_character_varying string) error
	InsertParameterCharacterVaryingStringStmt                 func(id string, source string, var_character_varying string) (*gobatis.Stmt, error)
	InsertParameterCharacterVaryingStringEmbed                func(id string, source string, var_character_varying string) error
	InsertEntityCharacterVaryingString                        func(item CharacterVaryingOriginal) error
	InsertEntityCharacterVaryingStringPointerOriginal         func(item *CharacterVaryingPointer) error
	InsertArrayParameterCharacterVaryingString                func(id string, source string, items []string) error
	InsertArrayParameterCharacterVaryingStringPointerOriginal func(id string, source string, items []*string) error
	SelectParameterBigintInt64                                func(id string) (int64, error)
	SelectParameterBigintInt64OriginalPointer                 func(id string) (*int64, error)
	SelectArrayParameterBigintInt64                           func(id int) ([]int64, error)
	SelectArrayParameterBigintInt64OriginalPointer            func(id int) ([]*int64, error)
	SelectParameterCharacterString                            func(id string) (string, error)
	SelectParameterCharacterStringOriginalPointer             func(id string) (*string, error)
	SelectArrayParameterCharacterString                       func(id int) ([]string, error)
	SelectArrayParameterCharacterStringOriginalPointer        func(id int) ([]*string, error)
	SelectParameterCharacterVaryingString                     func(id string) (string, error)
	SelectParameterCharacterVaryingStringOriginalPointer      func(id string) (*string, error)
	SelectArrayParameterCharacterVaryingString                func(id int) ([]string, error)
	SelectArrayParameterCharacterVaryingStringOriginalPointer func(id int) ([]*string, error)
}
