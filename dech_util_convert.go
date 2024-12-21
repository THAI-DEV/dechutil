package dechutil

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func AnyToPointer[T any](value T) *T {
	return &value
}

func AnyToString(inf interface{}, decimal int) string {
	str, _ := TypeAndKindOfObject(inf)

	if str == "string" {
		val := inf.(string)
		return val
	}

	if str == "int" {
		val := inf.(int)
		return strconv.Itoa(val)
	}

	if str == "int8" {
		val := int(inf.(int8))
		return strconv.Itoa(val)
	}

	if str == "int16" {
		val := int(inf.(int16))
		return strconv.Itoa(val)
	}

	if str == "int32" {
		val := int(inf.(int32))
		return strconv.Itoa(val)
	}

	if str == "int64" {
		val := int(inf.(int64))
		return strconv.Itoa(val)
	}

	if str == "float32" {
		val := float64(inf.(float32))
		str := strconv.FormatFloat(val, 'f', decimal, 64)
		return str
	}

	if str == "float64" {
		val := float64(inf.(float64))
		str := strconv.FormatFloat(val, 'f', decimal, 64)
		return str
	}

	return str
}

// * Convert a slice or array of a specific type to array of interface{}
func AnySliceToSliceInf(input interface{}) []interface{} {
	val := reflect.ValueOf(input)
	// There is no need to check, we want to panic if it's not slice or array
	inf := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		inf[i] = val.Index(i).Interface()
	}
	return inf
}

func NullStringToString(data *string, nullStrVal string) string {
	result := nullStrVal
	if data != nil {
		result = *data
	}

	return result
}

// * input : "abc", "xxxx" output : abc false
// * input :"null val", "null val" output : <nil> true
func StringToNullString(data string, defineStrVal string) (*string, bool) {
	var result *string
	var isNull bool

	if data == defineStrVal {
		result = nil
		isNull = true
	} else {
		result = &data
		isNull = false
	}

	return result, isNull
}

func StringToFloat(strValue string) float64 {
	result, err := strconv.ParseFloat(strValue, 64)
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func StringToInt(strValue string) int {
	result, err := strconv.Atoi(strValue)
	if err != nil {
		fmt.Println(err)
	}

	return result
}

func StringToTime(strValue string, format string) time.Time {
	dt, err := time.Parse(format, strValue)
	if err != nil {
		fmt.Println(err)
	}

	return dt
}

func StringToLocalTime(strValue string, format string) time.Time {
	dt, err := time.ParseInLocation(format, strValue, time.Local)
	if err != nil {
		fmt.Println(err)
	}

	return dt
}
