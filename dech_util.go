package dechutil

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

func Version() string {
	return "DECH Util , Version : 1.0.0 , Last Build : 22/10/2022 16:26"
}

// * Convert a slice or array of a specific type to array of interface{}
func ConvertAnySliceToSliceInf(input interface{}) []interface{} {
	val := reflect.ValueOf(input)
	// There is no need to check, we want to panic if it's not slice or array
	inf := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		inf[i] = val.Index(i).Interface()
	}
	return inf
}

func PrintSliceInf(data []interface{}) {
	for _, v := range data {
		fmt.Printf("\n%#v\n", v)
	}
}

func PrintData(data interface{}) {
	cv := ConvertAnySliceToSliceInf(data)
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

func FormatCommaFloat(data float64, round int) string {
	str := strconv.FormatFloat(data, 'f', round, 64)
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

func ConvertNullString2String(data *string, nullStrVal string) string {
	result := nullStrVal
	if data != nil {
		result = *data
	}

	return result
}

// * input : "abc", "xxxx" output : abc false
// * input :"null val", "null val" output : <nil> true
func ConvertString2NullString(data string, defindStrVal string) (*string, bool) {
	var result *string
	var isNull bool

	if data == defindStrVal {
		result = nil
		isNull = true
	} else {
		result = &data
		isNull = false
	}

	return result, isNull
}

func IsNull(data *string) bool {
	result := false

	if data == nil {
		result = true
	}

	return result
}

func RoundFloat(data float64, numRound int) float64 {
	return math.Round(data*math.Pow(10, float64(numRound))) / math.Pow(10, float64(numRound))
}

// * input : []string{"a",""b},"'","'",","  output : 'a','b'
func ConvertSliceString2String(data []string, wrapStr1 string, wrapStr2 string, saperateStr string) string {
	result := ""
	for i, v := range data {
		result = result + wrapStr1 + v + wrapStr2
		if i < len(data)-1 {
			result = result + saperateStr
		}
	}
	// fmt.Println(result)
	return result
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomFloat(min, max float64, numRound int) float64 {
	result := min + rand.Float64()*(max-min)
	result = RoundFloat(result, numRound)
	return result
}

/*
func FormatComma2(number int) string {
	p := message.NewPrinter(language.English)
	withCommaThousandSep := p.Sprintf("%d", number)
	return withCommaThousandSep
}
*/
