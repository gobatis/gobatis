package postgresql

import ()

type TypeOriginal struct {
	Sid               string `sql:"sid"`
	Source            string `sql:"source"`
	TBigint           int64  `sql:"t_bigint"`
	TInt8             int8   `sql:"t_int8"`
	TBool             bool   `sql:"t_bool"`
	TCharacter        string `sql:"t_character"`
	TChar             string `sql:"t_char"`
	TCharacterVarying string `sql:"t_character_varying"`
	TVarchar          string `sql:"t_varchar"`
	TInt              int8   `sql:"t_int"`
	TInt4             int8   `sql:"t_int4"`
}

type TypePointer struct {
	Sid               *string `sql:"sid"`
	Source            *string `sql:"source"`
	TBigint           *int64  `sql:"t_bigint"`
	TInt8             *int8   `sql:"t_int8"`
	TBool             *bool   `sql:"t_bool"`
	TCharacter        *string `sql:"t_character"`
	TChar             *string `sql:"t_char"`
	TCharacterVarying *string `sql:"t_character_varying"`
	TVarchar          *string `sql:"t_varchar"`
	TInt              *int8   `sql:"t_int"`
	TInt4             *int8   `sql:"t_int4"`
}

type ArrayTypeOriginal struct {
	Sid               string   `sql:"sid"`
	Source            string   `sql:"source"`
	TBigint           []int64  `sql:"t_bigint"`
	TInt8             []int8   `sql:"t_int8"`
	TBool             []bool   `sql:"t_bool"`
	TCharacter        []string `sql:"t_character"`
	TChar             []string `sql:"t_char"`
	TCharacterVarying []string `sql:"t_character_varying"`
	TVarchar          []string `sql:"t_varchar"`
	TInt              []int8   `sql:"t_int"`
	TInt4             []int8   `sql:"t_int4"`
}

type ArrayTypePointer struct {
	Sid               *string   `sql:"sid"`
	Source            *string   `sql:"source"`
	TBigint           []*int64  `sql:"t_bigint"`
	TInt8             []*int8   `sql:"t_int8"`
	TBool             []*bool   `sql:"t_bool"`
	TCharacter        []*string `sql:"t_character"`
	TChar             []*string `sql:"t_char"`
	TCharacterVarying []*string `sql:"t_character_varying"`
	TVarchar          []*string `sql:"t_varchar"`
	TInt              []*int8   `sql:"t_int"`
	TInt4             []*int8   `sql:"t_int4"`
}

type BigintOriginal struct {
	Sid     string `sql:"sid"`
	Source  string `sql:"source"`
	TBigint int64  `sql:"t_bigint"`
}

type BigintArrayOriginal struct {
	Sid     string  `sql:"sid"`
	Source  string  `sql:"source"`
	TBigint []int64 `sql:"t_bigint"`
}

type BigintPointer struct {
	Sid     *string `sql:"sid"`
	Source  *string `sql:"source"`
	TBigint *int64  `sql:"t_bigint"`
}

type BigintArrayPointer struct {
	Sid     *string  `sql:"sid"`
	Source  *string  `sql:"source"`
	TBigint []*int64 `sql:"t_bigint"`
}

type Int8Original struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt8  int8   `sql:"t_int8"`
}

type Int8ArrayOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt8  []int8 `sql:"t_int8"`
}

type Int8Pointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt8  *int8   `sql:"t_int8"`
}

type Int8ArrayPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt8  []*int8 `sql:"t_int8"`
}

type BoolOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TBool  bool   `sql:"t_bool"`
}

type BoolArrayOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TBool  []bool `sql:"t_bool"`
}

type BoolPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TBool  *bool   `sql:"t_bool"`
}

type BoolArrayPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TBool  []*bool `sql:"t_bool"`
}

type CharacterOriginal struct {
	Sid        string `sql:"sid"`
	Source     string `sql:"source"`
	TCharacter string `sql:"t_character"`
}

type CharacterArrayOriginal struct {
	Sid        string   `sql:"sid"`
	Source     string   `sql:"source"`
	TCharacter []string `sql:"t_character"`
}

type CharacterPointer struct {
	Sid        *string `sql:"sid"`
	Source     *string `sql:"source"`
	TCharacter *string `sql:"t_character"`
}

type CharacterArrayPointer struct {
	Sid        *string   `sql:"sid"`
	Source     *string   `sql:"source"`
	TCharacter []*string `sql:"t_character"`
}

type CharOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TChar  string `sql:"t_char"`
}

type CharArrayOriginal struct {
	Sid    string   `sql:"sid"`
	Source string   `sql:"source"`
	TChar  []string `sql:"t_char"`
}

type CharPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TChar  *string `sql:"t_char"`
}

type CharArrayPointer struct {
	Sid    *string   `sql:"sid"`
	Source *string   `sql:"source"`
	TChar  []*string `sql:"t_char"`
}

type CharacterVaryingOriginal struct {
	Sid               string `sql:"sid"`
	Source            string `sql:"source"`
	TCharacterVarying string `sql:"t_character_varying"`
}

type CharacterVaryingArrayOriginal struct {
	Sid               string   `sql:"sid"`
	Source            string   `sql:"source"`
	TCharacterVarying []string `sql:"t_character_varying"`
}

type CharacterVaryingPointer struct {
	Sid               *string `sql:"sid"`
	Source            *string `sql:"source"`
	TCharacterVarying *string `sql:"t_character_varying"`
}

type CharacterVaryingArrayPointer struct {
	Sid               *string   `sql:"sid"`
	Source            *string   `sql:"source"`
	TCharacterVarying []*string `sql:"t_character_varying"`
}

type VarcharOriginal struct {
	Sid      string `sql:"sid"`
	Source   string `sql:"source"`
	TVarchar string `sql:"t_varchar"`
}

type VarcharArrayOriginal struct {
	Sid      string   `sql:"sid"`
	Source   string   `sql:"source"`
	TVarchar []string `sql:"t_varchar"`
}

type VarcharPointer struct {
	Sid      *string `sql:"sid"`
	Source   *string `sql:"source"`
	TVarchar *string `sql:"t_varchar"`
}

type VarcharArrayPointer struct {
	Sid      *string   `sql:"sid"`
	Source   *string   `sql:"source"`
	TVarchar []*string `sql:"t_varchar"`
}

type IntOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt   int8   `sql:"t_int"`
}

type IntArrayOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt   []int8 `sql:"t_int"`
}

type IntPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt   *int8   `sql:"t_int"`
}

type IntArrayPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt   []*int8 `sql:"t_int"`
}

type Int4Original struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt4  int8   `sql:"t_int4"`
}

type Int4ArrayOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt4  []int8 `sql:"t_int4"`
}

type Int4Pointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt4  *int8   `sql:"t_int4"`
}

type Int4ArrayPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt4  []*int8 `sql:"t_int4"`
}
