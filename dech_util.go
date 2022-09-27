package dechutil

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func Version() string {
	return "DECH Util , Version : 1.0.0 , Last Build : 27/09/2022 14:43"
}

// Convert a slice or array of a specific type to array of interface{}
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

func FormatComma(data int64) string {
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

func ConvertNullString2String(data *string, nullStrVal string) string {
	result := nullStrVal
	if data != nil {
		result = *data
	}

	return result
}

func RoundFloat(data float64, numRound int) float64 {
	return math.Round(data*math.Pow(10, float64(numRound))) / math.Pow(10, float64(numRound))
}

/*
func FormatComma2(number int) string {
	p := message.NewPrinter(language.English)
	withCommaThousandSep := p.Sprintf("%d", number)
	return withCommaThousandSep
}
*/
