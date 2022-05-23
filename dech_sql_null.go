package dechutil

import (
	"database/sql"
	"time"
)

func SqlNullString(isSetNull bool, s string) sql.NullString {
	if isSetNull {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
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

func StringSqlNullValue(v sql.NullString) *string {
	if v.Valid {
		return &v.String
	} else {
		return nil
	}
}

func Int16SqlNullValue(v sql.NullInt16) *int16 {
	if v.Valid {
		return &v.Int16
	} else {
		return nil
	}
}

func Int32SqlNullValue(v sql.NullInt32) *int32 {
	if v.Valid {
		return &v.Int32
	} else {
		return nil
	}
}

func Int64SqlNullValue(v sql.NullInt64) *int64 {
	if v.Valid {
		return &v.Int64
	} else {
		return nil
	}
}

func TimeSqlNullValue(v sql.NullTime) *time.Time {
	if v.Valid {
		return &v.Time
	} else {
		return nil
	}
}

func Float64SqlNullValue(v sql.NullFloat64) *float64 {
	if v.Valid {
		return &v.Float64
	} else {
		return nil
	}
}

func boolSqlNullValue(v sql.NullBool) *bool {
	if v.Valid {
		return &v.Bool
	} else {
		return nil
	}
}

func byteSqlNullValue(v sql.NullByte) *byte {
	if v.Valid {
		return &v.Byte
	} else {
		return nil
	}
}
