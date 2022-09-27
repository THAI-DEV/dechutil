package dechutil

import (
	"database/sql"
	"time"
)

func SqlNullString(isSetNull bool, val string) sql.NullString {
	if isSetNull {
		return sql.NullString{}
	}

	return sql.NullString{
		String: val,
		Valid:  true,
	}
}

func SqlNullInt16(isSetNull bool, val int16) sql.NullInt16 {
	if isSetNull {
		return sql.NullInt16{}
	}

	return sql.NullInt16{
		Int16: val,
		Valid: true,
	}
}

func SqlNullInt32(isSetNull bool, val int32) sql.NullInt32 {
	if isSetNull {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: val,
		Valid: true,
	}
}

func SqlNullInt64(isSetNull bool, val int64) sql.NullInt64 {
	if isSetNull {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: val,
		Valid: true,
	}
}

func SqlNullFloat64(isSetNull bool, val float64) sql.NullFloat64 {
	if isSetNull {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{
		Float64: val,
		Valid:   true,
	}
}

func SqlNullTime(isSetNull bool, val time.Time) sql.NullTime {
	if isSetNull {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  val,
		Valid: true,
	}
}

func SqlNullBool(isSetNull bool, val bool) sql.NullBool {
	if isSetNull {
		return sql.NullBool{}
	}

	return sql.NullBool{
		Bool:  val,
		Valid: true,
	}
}

func SqlNullByte(isSetNull bool, val byte) sql.NullByte {
	if isSetNull {
		return sql.NullByte{}
	}

	return sql.NullByte{
		Byte:  val,
		Valid: true,
	}
}

//-------------------------------------------------------------------------------

func SqlNullStringRef(val *string) sql.NullString {
	if val == nil {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *val,
		Valid:  true,
	}
}

func SqlNullInt16Ref(val *int16) sql.NullInt16 {
	if val == nil {
		return sql.NullInt16{}
	}

	return sql.NullInt16{
		Int16: *val,
		Valid: true,
	}
}

func SqlNullInt32Ref(val *int32) sql.NullInt32 {
	if val == nil {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: *val,
		Valid: true,
	}
}

func SqlNullInt64Ref(val *int64) sql.NullInt64 {
	if val == nil {
		return sql.NullInt64{}
	}

	return sql.NullInt64{
		Int64: *val,
		Valid: true,
	}
}

func SqlNullFloat64Ref(val *float64) sql.NullFloat64 {
	if val == nil {
		return sql.NullFloat64{}
	}

	return sql.NullFloat64{
		Float64: *val,
		Valid:   true,
	}
}

func SqlNullTimeRef(val *time.Time) sql.NullTime {
	if val == nil {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  *val,
		Valid: true,
	}
}

func SqlNullBoolRef(val *bool) sql.NullBool {
	if val == nil {
		return sql.NullBool{}
	}

	return sql.NullBool{
		Bool:  *val,
		Valid: true,
	}
}

func SqlNullByteRef(val *byte) sql.NullByte {
	if val == nil {
		return sql.NullByte{}
	}

	return sql.NullByte{
		Byte:  *val,
		Valid: true,
	}
}

//-------------------------------------------------------------------------------

func StringSqlNullValue(val sql.NullString) *string {
	if val.Valid {
		return &val.String
	} else {
		return nil
	}
}

func Int16SqlNullValue(val sql.NullInt16) *int16 {
	if val.Valid {
		return &val.Int16
	} else {
		return nil
	}
}

func Int32SqlNullValue(val sql.NullInt32) *int32 {
	if val.Valid {
		return &val.Int32
	} else {
		return nil
	}
}

func Int64SqlNullValue(val sql.NullInt64) *int64 {
	if val.Valid {
		return &val.Int64
	} else {
		return nil
	}
}

func TimeSqlNullValue(val sql.NullTime) *time.Time {
	if val.Valid {
		return &val.Time
	} else {
		return nil
	}
}

func Float64SqlNullValue(val sql.NullFloat64) *float64 {
	if val.Valid {
		return &val.Float64
	} else {
		return nil
	}
}

func BoolSqlNullValue(val sql.NullBool) *bool {
	if val.Valid {
		return &val.Bool
	} else {
		return nil
	}
}

func ByteSqlNullValue(val sql.NullByte) *byte {
	if val.Valid {
		return &val.Byte
	} else {
		return nil
	}
}
