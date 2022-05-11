package dechutil

import (
	"fmt"
	"reflect"
)

func Version() string {
	return "DECH Util , Version : 1.0.0 , Last Build : 11/05/2022 17:55"
}

// Convert a slice or array of a specific type to array of interface{}
func ConvertSliceToArrInf(input interface{}) []interface{} {
	val := reflect.ValueOf(input)
	// There is no need to check, we want to panic if it's not slice or array
	inf := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		inf[i] = val.Index(i).Interface()
	}
	return inf
}

func PrintArrInf(data []interface{}) {
	for _, v := range data {
		fmt.Printf("\n%#v\n", v)
	}
}

func PrintData(data interface{}) {
	cv := ConvertSliceToArrInf(data)
	for _, v := range cv {
		fmt.Printf("\n%#v\n", v)
	}
}
