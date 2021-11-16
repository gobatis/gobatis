package postgresql

import ()

type TypeOriginal struct {
	TBigint           *int64  `sql:"t_bigint"`
	TCharacter        *string `sql:"t_character"`
	TCharacterVarying *string `sql:"t_character_varying"`
}

type TypePointer struct {
	TBigint           *int64  `sql:"t_bigint"`
	TCharacter        *string `sql:"t_character"`
	TCharacterVarying *string `sql:"t_character_varying"`
}

type ArrayTypeOriginal struct {
}

type ArrayTypePointer struct {
}

type BigintOriginal struct {
	TBigint *int64 `sql:"t_bigint"`
}

type BigintPointer struct {
	TBigint *int64 `sql:"t_bigint"`
}

type CharacterOriginal struct {
	TCharacter *string `sql:"t_character"`
}

type CharacterPointer struct {
	TCharacter *string `sql:"t_character"`
}

type CharacterVaryingOriginal struct {
	TCharacterVarying *string `sql:"t_character_varying"`
}

type CharacterVaryingPointer struct {
	TCharacterVarying *string `sql:"t_character_varying"`
}
