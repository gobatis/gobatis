package postgresql

import (

)


type AllType struct{ 
        TBigint int64 `sql:"t_bigint"`
        TInt8 int8 `sql:"t_int8"`
        TBigserial int64 `sql:"t_bigserial"`
        TSerial8 int64 `sql:"t_serial8"`
        TBit byte `sql:"t_bit"`
        TBitVarying int64 `sql:"t_bit_varying"`
        TBoolean int64 `sql:"t_boolean"`
        TBool int64 `sql:"t_bool"`
        TBox int64 `sql:"t_box"`
        TBytea int64 `sql:"t_bytea"`
        TCharacter int64 `sql:"t_character"`
        TChar int64 `sql:"t_char"`
        TCharacterVarying int64 `sql:"t_character_varying"`
        TVarchar int64 `sql:"t_varchar"`
        TCidr int64 `sql:"t_cidr"`
        TCircle int64 `sql:"t_circle"`
        TDate int64 `sql:"t_date"`
        TDoublePrecision int64 `sql:"t_double_precision"`
        TFloat8 int64 `sql:"t_float8"`
        TInet int64 `sql:"t_inet"`
        TInteger int64 `sql:"t_integer"`
        TInt int64 `sql:"t_int"`
        TInt4 int64 `sql:"t_int4"`
        TInterval int64 `sql:"t_interval"`
        TJson int64 `sql:"t_json"`
        TJsonb int64 `sql:"t_jsonb"`
        TLine int64 `sql:"t_line"`
        TLseg int64 `sql:"t_lseg"`
        TMacaddr int64 `sql:"t_macaddr"`
        TMacaddr8 int64 `sql:"t_macaddr8"`
        TMoney int64 `sql:"t_money"`
        TNumeric int64 `sql:"t_numeric"`
        TDecimal int64 `sql:"t_decimal"`
        TPath int64 `sql:"t_path"`
        TPgLsn int64 `sql:"t_pg_lsn"`
        TPgSnapshot int64 `sql:"t_pg_snapshot"`
        TPoint int64 `sql:"t_point"`
        TPolygon int64 `sql:"t_polygon"`
        TReal int64 `sql:"t_real"`
        TFloat4 int64 `sql:"t_float4"`
        TSmallint int64 `sql:"t_smallint"`
        TInt2 int64 `sql:"t_int2"`
        TSmallserial int64 `sql:"t_smallserial"`
        TSerial2 int64 `sql:"t_serial2"`
        TSerial int64 `sql:"t_serial"`
        TSerial4 int64 `sql:"t_serial4"`
        TText int64 `sql:"t_text"`
        TTime int64 `sql:"t_time"`
        TTimeWithTimezone int64 `sql:"t_time_with_timezone"`
        TTimez int64 `sql:"t_timez"`
        TTimestamp int64 `sql:"t_timestamp"`
        TTimestampWithTimezone int64 `sql:"t_timestamp_with_timezone"`
        TTimestampz int64 `sql:"t_timestampz"`
        TTsquery int64 `sql:"t_tsquery"`
        TTsvector int64 `sql:"t_tsvector"`
        TTxidSnapshot int64 `sql:"t_txid_snapshot"`
        TUuid int64 `sql:"t_uuid"`
        TXml string `sql:"t_xml"`
}

type BigintOriginal struct{ 
        TBigint int64 `sql:"t_bigint"`
}

type BigintPointer struct{ 
        TBigint *int64 `sql:"t_bigint"`
}

type Int8Original struct{ 
        TInt8 int8 `sql:"t_int8"`
}

type Int8Pointer struct{ 
        TInt8 *int8 `sql:"t_int8"`
}

type BigserialOriginal struct{ 
        TBigserial int64 `sql:"t_bigserial"`
}

type BigserialPointer struct{ 
        TBigserial *int64 `sql:"t_bigserial"`
}

type Serial8Original struct{ 
        TSerial8 int64 `sql:"t_serial8"`
}

type Serial8Pointer struct{ 
        TSerial8 *int64 `sql:"t_serial8"`
}

type BitOriginal struct{ 
        TBit byte `sql:"t_bit"`
}

type BitPointer struct{ 
        TBit *byte `sql:"t_bit"`
}

type BitVaryingOriginal struct{ 
        TBitVarying int64 `sql:"t_bit_varying"`
}

type BitVaryingPointer struct{ 
        TBitVarying *int64 `sql:"t_bit_varying"`
}

type BooleanOriginal struct{ 
        TBoolean int64 `sql:"t_boolean"`
}

type BooleanPointer struct{ 
        TBoolean *int64 `sql:"t_boolean"`
}

type BoolOriginal struct{ 
        TBool int64 `sql:"t_bool"`
}

type BoolPointer struct{ 
        TBool *int64 `sql:"t_bool"`
}

type BoxOriginal struct{ 
        TBox int64 `sql:"t_box"`
}

type BoxPointer struct{ 
        TBox *int64 `sql:"t_box"`
}

type ByteaOriginal struct{ 
        TBytea int64 `sql:"t_bytea"`
}

type ByteaPointer struct{ 
        TBytea *int64 `sql:"t_bytea"`
}

type CharacterOriginal struct{ 
        TCharacter int64 `sql:"t_character"`
}

type CharacterPointer struct{ 
        TCharacter *int64 `sql:"t_character"`
}

type CharOriginal struct{ 
        TChar int64 `sql:"t_char"`
}

type CharPointer struct{ 
        TChar *int64 `sql:"t_char"`
}

type CharacterVaryingOriginal struct{ 
        TCharacterVarying int64 `sql:"t_character_varying"`
}

type CharacterVaryingPointer struct{ 
        TCharacterVarying *int64 `sql:"t_character_varying"`
}

type VarcharOriginal struct{ 
        TVarchar int64 `sql:"t_varchar"`
}

type VarcharPointer struct{ 
        TVarchar *int64 `sql:"t_varchar"`
}

type CidrOriginal struct{ 
        TCidr int64 `sql:"t_cidr"`
}

type CidrPointer struct{ 
        TCidr *int64 `sql:"t_cidr"`
}

type CircleOriginal struct{ 
        TCircle int64 `sql:"t_circle"`
}

type CirclePointer struct{ 
        TCircle *int64 `sql:"t_circle"`
}

type DateOriginal struct{ 
        TDate int64 `sql:"t_date"`
}

type DatePointer struct{ 
        TDate *int64 `sql:"t_date"`
}

type DoublePrecisionOriginal struct{ 
        TDoublePrecision int64 `sql:"t_double_precision"`
}

type DoublePrecisionPointer struct{ 
        TDoublePrecision *int64 `sql:"t_double_precision"`
}

type Float8Original struct{ 
        TFloat8 int64 `sql:"t_float8"`
}

type Float8Pointer struct{ 
        TFloat8 *int64 `sql:"t_float8"`
}

type InetOriginal struct{ 
        TInet int64 `sql:"t_inet"`
}

type InetPointer struct{ 
        TInet *int64 `sql:"t_inet"`
}

type IntegerOriginal struct{ 
        TInteger int64 `sql:"t_integer"`
}

type IntegerPointer struct{ 
        TInteger *int64 `sql:"t_integer"`
}

type IntOriginal struct{ 
        TInt int64 `sql:"t_int"`
}

type IntPointer struct{ 
        TInt *int64 `sql:"t_int"`
}

type Int4Original struct{ 
        TInt4 int64 `sql:"t_int4"`
}

type Int4Pointer struct{ 
        TInt4 *int64 `sql:"t_int4"`
}

type IntervalOriginal struct{ 
        TInterval int64 `sql:"t_interval"`
}

type IntervalPointer struct{ 
        TInterval *int64 `sql:"t_interval"`
}

type JsonOriginal struct{ 
        TJson int64 `sql:"t_json"`
}

type JsonPointer struct{ 
        TJson *int64 `sql:"t_json"`
}

type JsonbOriginal struct{ 
        TJsonb int64 `sql:"t_jsonb"`
}

type JsonbPointer struct{ 
        TJsonb *int64 `sql:"t_jsonb"`
}

type LineOriginal struct{ 
        TLine int64 `sql:"t_line"`
}

type LinePointer struct{ 
        TLine *int64 `sql:"t_line"`
}

type LsegOriginal struct{ 
        TLseg int64 `sql:"t_lseg"`
}

type LsegPointer struct{ 
        TLseg *int64 `sql:"t_lseg"`
}

type MacaddrOriginal struct{ 
        TMacaddr int64 `sql:"t_macaddr"`
}

type MacaddrPointer struct{ 
        TMacaddr *int64 `sql:"t_macaddr"`
}

type Macaddr8Original struct{ 
        TMacaddr8 int64 `sql:"t_macaddr8"`
}

type Macaddr8Pointer struct{ 
        TMacaddr8 *int64 `sql:"t_macaddr8"`
}

type MoneyOriginal struct{ 
        TMoney int64 `sql:"t_money"`
}

type MoneyPointer struct{ 
        TMoney *int64 `sql:"t_money"`
}

type NumericOriginal struct{ 
        TNumeric int64 `sql:"t_numeric"`
}

type NumericPointer struct{ 
        TNumeric *int64 `sql:"t_numeric"`
}

type DecimalOriginal struct{ 
        TDecimal int64 `sql:"t_decimal"`
}

type DecimalPointer struct{ 
        TDecimal *int64 `sql:"t_decimal"`
}

type PathOriginal struct{ 
        TPath int64 `sql:"t_path"`
}

type PathPointer struct{ 
        TPath *int64 `sql:"t_path"`
}

type PgLsnOriginal struct{ 
        TPgLsn int64 `sql:"t_pg_lsn"`
}

type PgLsnPointer struct{ 
        TPgLsn *int64 `sql:"t_pg_lsn"`
}

type PgSnapshotOriginal struct{ 
        TPgSnapshot int64 `sql:"t_pg_snapshot"`
}

type PgSnapshotPointer struct{ 
        TPgSnapshot *int64 `sql:"t_pg_snapshot"`
}

type PointOriginal struct{ 
        TPoint int64 `sql:"t_point"`
}

type PointPointer struct{ 
        TPoint *int64 `sql:"t_point"`
}

type PolygonOriginal struct{ 
        TPolygon int64 `sql:"t_polygon"`
}

type PolygonPointer struct{ 
        TPolygon *int64 `sql:"t_polygon"`
}

type RealOriginal struct{ 
        TReal int64 `sql:"t_real"`
}

type RealPointer struct{ 
        TReal *int64 `sql:"t_real"`
}

type Float4Original struct{ 
        TFloat4 int64 `sql:"t_float4"`
}

type Float4Pointer struct{ 
        TFloat4 *int64 `sql:"t_float4"`
}

type SmallintOriginal struct{ 
        TSmallint int64 `sql:"t_smallint"`
}

type SmallintPointer struct{ 
        TSmallint *int64 `sql:"t_smallint"`
}

type Int2Original struct{ 
        TInt2 int64 `sql:"t_int2"`
}

type Int2Pointer struct{ 
        TInt2 *int64 `sql:"t_int2"`
}

type SmallserialOriginal struct{ 
        TSmallserial int64 `sql:"t_smallserial"`
}

type SmallserialPointer struct{ 
        TSmallserial *int64 `sql:"t_smallserial"`
}

type Serial2Original struct{ 
        TSerial2 int64 `sql:"t_serial2"`
}

type Serial2Pointer struct{ 
        TSerial2 *int64 `sql:"t_serial2"`
}

type SerialOriginal struct{ 
        TSerial int64 `sql:"t_serial"`
}

type SerialPointer struct{ 
        TSerial *int64 `sql:"t_serial"`
}

type Serial4Original struct{ 
        TSerial4 int64 `sql:"t_serial4"`
}

type Serial4Pointer struct{ 
        TSerial4 *int64 `sql:"t_serial4"`
}

type TextOriginal struct{ 
        TText int64 `sql:"t_text"`
}

type TextPointer struct{ 
        TText *int64 `sql:"t_text"`
}

type TimeOriginal struct{ 
        TTime int64 `sql:"t_time"`
}

type TimePointer struct{ 
        TTime *int64 `sql:"t_time"`
}

type TimeWithTimezoneOriginal struct{ 
        TTimeWithTimezone int64 `sql:"t_time_with_timezone"`
}

type TimeWithTimezonePointer struct{ 
        TTimeWithTimezone *int64 `sql:"t_time_with_timezone"`
}

type TimezOriginal struct{ 
        TTimez int64 `sql:"t_timez"`
}

type TimezPointer struct{ 
        TTimez *int64 `sql:"t_timez"`
}

type TimestampOriginal struct{ 
        TTimestamp int64 `sql:"t_timestamp"`
}

type TimestampPointer struct{ 
        TTimestamp *int64 `sql:"t_timestamp"`
}

type TimestampWithTimezoneOriginal struct{ 
        TTimestampWithTimezone int64 `sql:"t_timestamp_with_timezone"`
}

type TimestampWithTimezonePointer struct{ 
        TTimestampWithTimezone *int64 `sql:"t_timestamp_with_timezone"`
}

type TimestampzOriginal struct{ 
        TTimestampz int64 `sql:"t_timestampz"`
}

type TimestampzPointer struct{ 
        TTimestampz *int64 `sql:"t_timestampz"`
}

type TsqueryOriginal struct{ 
        TTsquery int64 `sql:"t_tsquery"`
}

type TsqueryPointer struct{ 
        TTsquery *int64 `sql:"t_tsquery"`
}

type TsvectorOriginal struct{ 
        TTsvector int64 `sql:"t_tsvector"`
}

type TsvectorPointer struct{ 
        TTsvector *int64 `sql:"t_tsvector"`
}

type TxidSnapshotOriginal struct{ 
        TTxidSnapshot int64 `sql:"t_txid_snapshot"`
}

type TxidSnapshotPointer struct{ 
        TTxidSnapshot *int64 `sql:"t_txid_snapshot"`
}

type UuidOriginal struct{ 
        TUuid int64 `sql:"t_uuid"`
}

type UuidPointer struct{ 
        TUuid *int64 `sql:"t_uuid"`
}

type XmlOriginal struct{ 
        TXml string `sql:"t_xml"`
}

type XmlPointer struct{ 
        TXml *string `sql:"t_xml"`
}
