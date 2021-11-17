package postgresql

import ()

type TypeOriginal struct {
	Sid                    string  `sql:"sid"`
	Source                 string  `sql:"source"`
	TBigint                int64   `sql:"t_bigint"`
	TInt8                  int8    `sql:"t_int8"`
	TBoolean               bool    `sql:"t_boolean"`
	TBool                  bool    `sql:"t_bool"`
	TCharacter             string  `sql:"t_character"`
	TChar                  string  `sql:"t_char"`
	TCharacterVarying      string  `sql:"t_character_varying"`
	TVarchar               string  `sql:"t_varchar"`
	TFloat8                float32 `sql:"t_float8"`
	TInteger               int8    `sql:"t_integer"`
	TInt                   int8    `sql:"t_int"`
	TInt4                  int8    `sql:"t_int4"`
	TNumeric               decimal `sql:"t_numeric"`
	TDecimal               decimal `sql:"t_decimal"`
	TFloat4                float32 `sql:"t_float4"`
	TSmallint              int8    `sql:"t_smallint"`
	TInt2                  int8    `sql:"t_int2"`
	TText                  string  `sql:"t_text"`
	TTime                  time    `sql:"t_time"`
	TTimeWithTimezone      time    `sql:"t_time_with_timezone"`
	TTimetz                time    `sql:"t_timetz"`
	TTimestamp             time    `sql:"t_timestamp"`
	TTimestampWithTimezone time    `sql:"t_timestamp_with_timezone"`
	TTimestamptz           time    `sql:"t_timestamptz"`
}

type TypePointer struct {
	Sid                    *string  `sql:"sid"`
	Source                 *string  `sql:"source"`
	TBigint                *int64   `sql:"t_bigint"`
	TInt8                  *int8    `sql:"t_int8"`
	TBoolean               *bool    `sql:"t_boolean"`
	TBool                  *bool    `sql:"t_bool"`
	TCharacter             *string  `sql:"t_character"`
	TChar                  *string  `sql:"t_char"`
	TCharacterVarying      *string  `sql:"t_character_varying"`
	TVarchar               *string  `sql:"t_varchar"`
	TFloat8                *float32 `sql:"t_float8"`
	TInteger               *int8    `sql:"t_integer"`
	TInt                   *int8    `sql:"t_int"`
	TInt4                  *int8    `sql:"t_int4"`
	TNumeric               *decimal `sql:"t_numeric"`
	TDecimal               *decimal `sql:"t_decimal"`
	TFloat4                *float32 `sql:"t_float4"`
	TSmallint              *int8    `sql:"t_smallint"`
	TInt2                  *int8    `sql:"t_int2"`
	TText                  *string  `sql:"t_text"`
	TTime                  *time    `sql:"t_time"`
	TTimeWithTimezone      *time    `sql:"t_time_with_timezone"`
	TTimetz                *time    `sql:"t_timetz"`
	TTimestamp             *time    `sql:"t_timestamp"`
	TTimestampWithTimezone *time    `sql:"t_timestamp_with_timezone"`
	TTimestamptz           *time    `sql:"t_timestamptz"`
}

type ArrayTypeOriginal struct {
	Sid                    string    `sql:"sid"`
	Source                 string    `sql:"source"`
	TBigint                []int64   `sql:"t_bigint"`
	TInt8                  []int8    `sql:"t_int8"`
	TBoolean               []bool    `sql:"t_boolean"`
	TBool                  []bool    `sql:"t_bool"`
	TCharacter             []string  `sql:"t_character"`
	TChar                  []string  `sql:"t_char"`
	TCharacterVarying      []string  `sql:"t_character_varying"`
	TVarchar               []string  `sql:"t_varchar"`
	TFloat8                []float32 `sql:"t_float8"`
	TInteger               []int8    `sql:"t_integer"`
	TInt                   []int8    `sql:"t_int"`
	TInt4                  []int8    `sql:"t_int4"`
	TNumeric               []decimal `sql:"t_numeric"`
	TDecimal               []decimal `sql:"t_decimal"`
	TFloat4                []float32 `sql:"t_float4"`
	TSmallint              []int8    `sql:"t_smallint"`
	TInt2                  []int8    `sql:"t_int2"`
	TText                  []string  `sql:"t_text"`
	TTime                  []time    `sql:"t_time"`
	TTimeWithTimezone      []time    `sql:"t_time_with_timezone"`
	TTimetz                []time    `sql:"t_timetz"`
	TTimestamp             []time    `sql:"t_timestamp"`
	TTimestampWithTimezone []time    `sql:"t_timestamp_with_timezone"`
	TTimestamptz           []time    `sql:"t_timestamptz"`
}

type ArrayTypePointer struct {
	Sid                    *string    `sql:"sid"`
	Source                 *string    `sql:"source"`
	TBigint                []*int64   `sql:"t_bigint"`
	TInt8                  []*int8    `sql:"t_int8"`
	TBoolean               []*bool    `sql:"t_boolean"`
	TBool                  []*bool    `sql:"t_bool"`
	TCharacter             []*string  `sql:"t_character"`
	TChar                  []*string  `sql:"t_char"`
	TCharacterVarying      []*string  `sql:"t_character_varying"`
	TVarchar               []*string  `sql:"t_varchar"`
	TFloat8                []*float32 `sql:"t_float8"`
	TInteger               []*int8    `sql:"t_integer"`
	TInt                   []*int8    `sql:"t_int"`
	TInt4                  []*int8    `sql:"t_int4"`
	TNumeric               []*decimal `sql:"t_numeric"`
	TDecimal               []*decimal `sql:"t_decimal"`
	TFloat4                []*float32 `sql:"t_float4"`
	TSmallint              []*int8    `sql:"t_smallint"`
	TInt2                  []*int8    `sql:"t_int2"`
	TText                  []*string  `sql:"t_text"`
	TTime                  []*time    `sql:"t_time"`
	TTimeWithTimezone      []*time    `sql:"t_time_with_timezone"`
	TTimetz                []*time    `sql:"t_timetz"`
	TTimestamp             []*time    `sql:"t_timestamp"`
	TTimestampWithTimezone []*time    `sql:"t_timestamp_with_timezone"`
	TTimestamptz           []*time    `sql:"t_timestamptz"`
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

type Float8Original struct {
	Sid     string  `sql:"sid"`
	Source  string  `sql:"source"`
	TFloat8 float32 `sql:"t_float8"`
}

type Float8ArrayOriginal struct {
	Sid     string    `sql:"sid"`
	Source  string    `sql:"source"`
	TFloat8 []float32 `sql:"t_float8"`
}

type Float8Pointer struct {
	Sid     *string  `sql:"sid"`
	Source  *string  `sql:"source"`
	TFloat8 *float32 `sql:"t_float8"`
}

type Float8ArrayPointer struct {
	Sid     *string    `sql:"sid"`
	Source  *string    `sql:"source"`
	TFloat8 []*float32 `sql:"t_float8"`
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

type NumericOriginal struct {
	Sid      string  `sql:"sid"`
	Source   string  `sql:"source"`
	TNumeric decimal `sql:"t_numeric"`
}

type NumericArrayOriginal struct {
	Sid      string    `sql:"sid"`
	Source   string    `sql:"source"`
	TNumeric []decimal `sql:"t_numeric"`
}

type NumericPointer struct {
	Sid      *string  `sql:"sid"`
	Source   *string  `sql:"source"`
	TNumeric *decimal `sql:"t_numeric"`
}

type NumericArrayPointer struct {
	Sid      *string    `sql:"sid"`
	Source   *string    `sql:"source"`
	TNumeric []*decimal `sql:"t_numeric"`
}

type DecimalOriginal struct {
	Sid      string  `sql:"sid"`
	Source   string  `sql:"source"`
	TDecimal decimal `sql:"t_decimal"`
}

type DecimalArrayOriginal struct {
	Sid      string    `sql:"sid"`
	Source   string    `sql:"source"`
	TDecimal []decimal `sql:"t_decimal"`
}

type DecimalPointer struct {
	Sid      *string  `sql:"sid"`
	Source   *string  `sql:"source"`
	TDecimal *decimal `sql:"t_decimal"`
}

type DecimalArrayPointer struct {
	Sid      *string    `sql:"sid"`
	Source   *string    `sql:"source"`
	TDecimal []*decimal `sql:"t_decimal"`
}

type Float4Original struct {
	Sid     string  `sql:"sid"`
	Source  string  `sql:"source"`
	TFloat4 float32 `sql:"t_float4"`
}

type Float4ArrayOriginal struct {
	Sid     string    `sql:"sid"`
	Source  string    `sql:"source"`
	TFloat4 []float32 `sql:"t_float4"`
}

type Float4Pointer struct {
	Sid     *string  `sql:"sid"`
	Source  *string  `sql:"source"`
	TFloat4 *float32 `sql:"t_float4"`
}

type Float4ArrayPointer struct {
	Sid     *string    `sql:"sid"`
	Source  *string    `sql:"source"`
	TFloat4 []*float32 `sql:"t_float4"`
}

type SmallintOriginal struct {
	Sid       string `sql:"sid"`
	Source    string `sql:"source"`
	TSmallint int8   `sql:"t_smallint"`
}

type SmallintArrayOriginal struct {
	Sid       string `sql:"sid"`
	Source    string `sql:"source"`
	TSmallint []int8 `sql:"t_smallint"`
}

type SmallintPointer struct {
	Sid       *string `sql:"sid"`
	Source    *string `sql:"source"`
	TSmallint *int8   `sql:"t_smallint"`
}

type SmallintArrayPointer struct {
	Sid       *string `sql:"sid"`
	Source    *string `sql:"source"`
	TSmallint []*int8 `sql:"t_smallint"`
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

type TimeOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TTime  time   `sql:"t_time"`
}

type TimeArrayOriginal struct {
	Sid    string `sql:"sid"`
	Source string `sql:"source"`
	TTime  []time `sql:"t_time"`
}

type TimePointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TTime  *time   `sql:"t_time"`
}

type TimeArrayPointer struct {
	Sid    *string `sql:"sid"`
	Source *string `sql:"source"`
	TTime  []*time `sql:"t_time"`
}

type TimeWithTimezoneOriginal struct {
	Sid               string `sql:"sid"`
	Source            string `sql:"source"`
	TTimeWithTimezone time   `sql:"t_time_with_timezone"`
}

type TimeWithTimezoneArrayOriginal struct {
	Sid               string `sql:"sid"`
	Source            string `sql:"source"`
	TTimeWithTimezone []time `sql:"t_time_with_timezone"`
}

type TimeWithTimezonePointer struct {
	Sid               *string `sql:"sid"`
	Source            *string `sql:"source"`
	TTimeWithTimezone *time   `sql:"t_time_with_timezone"`
}

type TimeWithTimezoneArrayPointer struct {
	Sid               *string `sql:"sid"`
	Source            *string `sql:"source"`
	TTimeWithTimezone []*time `sql:"t_time_with_timezone"`
}

type TimetzOriginal struct {
	Sid     string `sql:"sid"`
	Source  string `sql:"source"`
	TTimetz time   `sql:"t_timetz"`
}

type TimetzArrayOriginal struct {
	Sid     string `sql:"sid"`
	Source  string `sql:"source"`
	TTimetz []time `sql:"t_timetz"`
}

type TimetzPointer struct {
	Sid     *string `sql:"sid"`
	Source  *string `sql:"source"`
	TTimetz *time   `sql:"t_timetz"`
}

type TimetzArrayPointer struct {
	Sid     *string `sql:"sid"`
	Source  *string `sql:"source"`
	TTimetz []*time `sql:"t_timetz"`
}

type TimestampOriginal struct {
	Sid        string `sql:"sid"`
	Source     string `sql:"source"`
	TTimestamp time   `sql:"t_timestamp"`
}

type TimestampArrayOriginal struct {
	Sid        string `sql:"sid"`
	Source     string `sql:"source"`
	TTimestamp []time `sql:"t_timestamp"`
}

type TimestampPointer struct {
	Sid        *string `sql:"sid"`
	Source     *string `sql:"source"`
	TTimestamp *time   `sql:"t_timestamp"`
}

type TimestampArrayPointer struct {
	Sid        *string `sql:"sid"`
	Source     *string `sql:"source"`
	TTimestamp []*time `sql:"t_timestamp"`
}

type TimestampWithTimezoneOriginal struct {
	Sid                    string `sql:"sid"`
	Source                 string `sql:"source"`
	TTimestampWithTimezone time   `sql:"t_timestamp_with_timezone"`
}

type TimestampWithTimezoneArrayOriginal struct {
	Sid                    string `sql:"sid"`
	Source                 string `sql:"source"`
	TTimestampWithTimezone []time `sql:"t_timestamp_with_timezone"`
}

type TimestampWithTimezonePointer struct {
	Sid                    *string `sql:"sid"`
	Source                 *string `sql:"source"`
	TTimestampWithTimezone *time   `sql:"t_timestamp_with_timezone"`
}

type TimestampWithTimezoneArrayPointer struct {
	Sid                    *string `sql:"sid"`
	Source                 *string `sql:"source"`
	TTimestampWithTimezone []*time `sql:"t_timestamp_with_timezone"`
}

type TimestamptzOriginal struct {
	Sid          string `sql:"sid"`
	Source       string `sql:"source"`
	TTimestamptz time   `sql:"t_timestamptz"`
}

type TimestamptzArrayOriginal struct {
	Sid          string `sql:"sid"`
	Source       string `sql:"source"`
	TTimestamptz []time `sql:"t_timestamptz"`
}

type TimestamptzPointer struct {
	Sid          *string `sql:"sid"`
	Source       *string `sql:"source"`
	TTimestamptz *time   `sql:"t_timestamptz"`
}

type TimestamptzArrayPointer struct {
	Sid          *string `sql:"sid"`
	Source       *string `sql:"source"`
	TTimestamptz []*time `sql:"t_timestamptz"`
}
