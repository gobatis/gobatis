package postgresql

import ()

type TypeOriginal struct {
	Sid               string  `sql:"sid"`
	Source            string  `sql:"source"`
	TBigint           *int64  `sql:"t_bigint"`
	TCharacter        *string `sql:"t_character"`
	TCharacterVarying *string `sql:"t_character_varying"`
}

type TypePointer struct {
	Sid               *string `sql:"sid"`
	Source            *string `sql:"source"`
	TBigint           *int64  `sql:"t_bigint"`
	TCharacter        *string `sql:"t_character"`
	TCharacterVarying *string `sql:"t_character_varying"`
}

type ArrayTypeOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
}

type ArrayTypePointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
}

type BigintOriginal struct {
	Sid     string `sql:"sid"`
	Source  string `sql:"source"`
	TBigint *int64 `sql:"t_bigint"`
}

type BigintPointer struct {
	Sid     *string `sql:"sid"`
	Source  *string `sql:"source"`
	TBigint *int64  `sql:"t_bigint"`
}

type CharacterOriginal struct {
	Sid        string  `sql:"sid"`
	Source     string  `sql:"source"`
	TCharacter *string `sql:"t_character"`
}

type CharacterPointer struct {
	Sid        *string `sql:"sid"`
	Source     *string `sql:"source"`
	TCharacter *string `sql:"t_character"`
}

type CharacterVaryingOriginal struct {
	Sid               string  `sql:"sid"`
	Source            string  `sql:"source"`
	TCharacterVarying *string `sql:"t_character_varying"`
}

type CharacterVaryingPointer struct {
	Sid               *string `sql:"sid"`
	Source            *string `sql:"source"`
	TCharacterVarying *string `sql:"t_character_varying"`
}
