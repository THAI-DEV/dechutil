package dechutil

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func Version() string {
	return "DECH Util , Version : 1.0.0 , Last Build : 22/05/2024 23:33"
}

func TypeAndKindOfObject(inf interface{}) (string, string) {
	s1 := reflect.TypeOf(inf).String()
	s2 := reflect.ValueOf(inf).Kind().String()
	return s1, s2
}

func PrintSliceOfAnyType(data []interface{}) {
	for _, v := range data {
		fmt.Printf("\n%#v\n", v)
	}
}

func PrintSliceOfSameType(data interface{}) {
	cv := AnySliceToSliceInf(data)
	for _, v := range cv {
		fmt.Printf("\n%#v\n", v)
	}
}

func FormatCommaInt(data int64) string {
	in := strconv.FormatInt(data, 10)
	numOfDigits := len(in)
	if data < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if data < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

func FormatCommaFloat(data float64, decimal int) string {
	str := strconv.FormatFloat(data, 'f', decimal, 64)
	strSlice := strings.SplitN(str, ".", -1)
	str1 := ""
	str2 := ""

	if len(strSlice) == 2 {
		str2 = strSlice[1]
	}

	str1 = strSlice[0]

	intVal, _ := strconv.ParseInt(str1, 10, 64)
	result := FormatCommaInt(intVal)

	if len(str2) > 1 {
		result = result + "." + str2
	}

	return result
}

func IsNullString(data *string) bool {
	result := false

	if data == nil {
		result = true
	}

	return result
}

func RoundFloat(data float64, decimal int) float64 {
	return math.Round(data*math.Pow(10, float64(decimal))) / math.Pow(10, float64(decimal))
}

/*
func FormatComma2(number int) string {
	p := message.NewPrinter(language.English)
	withCommaThousandSep := p.Sprintf("%d", number)
	return withCommaThousandSep
}
*/
