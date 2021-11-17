package postgresql

import ()

type TypeOriginal struct {
	Sid               string `sql:"sid"`
	Source            string `sql:"source"`
	TBigint           int64  `sql:"t_bigint"`
	TInt8             int8   `sql:"t_int8"`
	TBoolean          bool   `sql:"t_boolean"`
	TBool             bool   `sql:"t_bool"`
	TCharacter        string `sql:"t_character"`
	TChar             string `sql:"t_char"`
	TCharacterVarying string `sql:"t_character_varying"`
	TVarchar          string `sql:"t_varchar"`
	TInteger          int8   `sql:"t_integer"`
	TInt              int8   `sql:"t_int"`
	TInt4             int8   `sql:"t_int4"`
	TInt2             int8   `sql:"t_int2"`
	TText             string `sql:"t_text"`
}

type TypePointer struct {
	Sid               *string `sql:"sid"`
	Source            *string `sql:"source"`
	TBigint           *int64  `sql:"t_bigint"`
	TInt8             *int8   `sql:"t_int8"`
	TBoolean          *bool   `sql:"t_boolean"`
	TBool             *bool   `sql:"t_bool"`
	TCharacter        *string `sql:"t_character"`
	TChar             *string `sql:"t_char"`
	TCharacterVarying *string `sql:"t_character_varying"`
	TVarchar          *string `sql:"t_varchar"`
	TInteger          *int8   `sql:"t_integer"`
	TInt              *int8   `sql:"t_int"`
	TInt4             *int8   `sql:"t_int4"`
	TInt2             *int8   `sql:"t_int2"`
	TText             *string `sql:"t_text"`
}

type ArrayTypeOriginal struct {
	Sid               string   `sql:"sid"`
	Source            string   `sql:"source"`
	TBigint           []int64  `sql:"t_bigint"`
	TInt8             []int8   `sql:"t_int8"`
	TBoolean          []bool   `sql:"t_boolean"`
	TBool             []bool   `sql:"t_bool"`
	TCharacter        []string `sql:"t_character"`
	TChar             []string `sql:"t_char"`
	TCharacterVarying []string `sql:"t_character_varying"`
	TVarchar          []string `sql:"t_varchar"`
	TInteger          []int8   `sql:"t_integer"`
	TInt              []int8   `sql:"t_int"`
	TInt4             []int8   `sql:"t_int4"`
	TInt2             []int8   `sql:"t_int2"`
	TText             []string `sql:"t_text"`
}

type ArrayTypePointer struct {
	Sid               *string   `sql:"sid"`
	Source            *string   `sql:"source"`
	TBigint           []*int64  `sql:"t_bigint"`
	TInt8             []*int8   `sql:"t_int8"`
	TBoolean          []*bool   `sql:"t_boolean"`
	TBool             []*bool   `sql:"t_bool"`
	TCharacter        []*string `sql:"t_character"`
	TChar             []*string `sql:"t_char"`
	TCharacterVarying []*string `sql:"t_character_varying"`
	TVarchar          []*string `sql:"t_varchar"`
	TInteger          []*int8   `sql:"t_integer"`
	TInt              []*int8   `sql:"t_int"`
	TInt4             []*int8   `sql:"t_int4"`
	TInt2             []*int8   `sql:"t_int2"`
	TText             []*string `sql:"t_text"`
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

type BooleanOriginal struct {
	Sid      string `sql:"sid"`
	Source   string `sql:"source"`
	TBoolean bool   `sql:"t_boolean"`
}

type BooleanArrayOriginal struct {
	Sid      string `sql:"sid"`
	Source   string `sql:"source"`
	TBoolean []bool `sql:"t_boolean"`
}

type BooleanPointer struct {
	Sid      *string `sql:"sid"`
	Source   *string `sql:"source"`
	TBoolean *bool   `sql:"t_boolean"`
}

type BooleanArrayPointer struct {
	Sid      *string `sql:"sid"`
	Source   *string `sql:"source"`
	TBoolean []*bool `sql:"t_boolean"`
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

type IntegerOriginal struct {
	Sid      string `sql:"sid"`
	Source   string `sql:"source"`
	TInteger int8   `sql:"t_integer"`
}

type IntegerArrayOriginal struct {
	Sid      string `sql:"sid"`
	Source   string `sql:"source"`
	TInteger []int8 `sql:"t_integer"`
}

type IntegerPointer struct {
	Sid      *string `sql:"sid"`
	Source   *string `sql:"source"`
	TInteger *int8   `sql:"t_integer"`
}

type IntegerArrayPointer struct {
	Sid      *string `sql:"sid"`
	Source   *string `sql:"source"`
	TInteger []*int8 `sql:"t_integer"`
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

type Int2Original struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt2  int8   `sql:"t_int2"`
}

type Int2ArrayOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TInt2  []int8 `sql:"t_int2"`
}

type Int2Pointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt2  *int8   `sql:"t_int2"`
}

type Int2ArrayPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TInt2  []*int8 `sql:"t_int2"`
}

type TextOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TText  string `sql:"t_text"`
}

type TextArrayOriginal struct {
	Sid    string   `sql:"sid"`
	Source string   `sql:"source"`
	TText  []string `sql:"t_text"`
}

type TextPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TText  *string `sql:"t_text"`
}

type TextArrayPointer struct {
	Sid    *string   `sql:"sid"`
	Source *string   `sql:"source"`
	TText  []*string `sql:"t_text"`
}
