package dechutil

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Version() string {
	return "DECH Util , Version : 1.0.0 , Last Build : 29/10/2022 11:30"
}

func TypeAndKindOfObject(inf interface{}) (string, string) {
	s1 := reflect.TypeOf(inf).String()
	s2 := reflect.ValueOf(inf).Kind().String()
	return s1, s2
}

func ConvertValueObjectToString(inf interface{}, decimal int) string {
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

func RoundFloat(data float64, decimal int) float64 {
	return math.Round(data*math.Pow(10, float64(decimal))) / math.Pow(10, float64(decimal))
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

// Random Min to Max-1
func RandomInt(min, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand2 := rand.New(seed)
	return rand2.Intn(max-min) + min
}

// Random Min to Max-1
func RandomFloat(min, max float64, decimal int) float64 {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)

	result := min + rand.Float64()*(max-min)
	result = RoundFloat(result, decimal)
	return result
}

func RandomString(num int, hasLower, hasUpper, hasNumber, isRemoveSameChar bool) string {
	charsetLower := "abcdefghijklmnopqrstuvwxyz"
	charsetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumber := "1234567890"
	charsetSame := "1lkKcCoO0pPquUxXyYzZ"

	var sb strings.Builder
	if hasLower {
		sb.WriteString(charsetLower)
	}

	if hasUpper {
		sb.WriteString(charsetUpper)
	}

	if hasNumber {
		sb.WriteString(charsetNumber)
	}

	charset := sb.String()

	if isRemoveSameChar {
		for _, v := range charsetSame {
			charset = strings.ReplaceAll(charset, string(v), "")
		}
	}

	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)

	result := ""
	ch := make(chan string)
	for {
		go randomChar(rand, charset, ch)
		result = result + string(<-ch)

		if len(result) >= num {
			break
		}
	}

	/*
		for {
			if len(result) >= num {
				return result
			}

			c := charset[rand.Intn(len(charset))]
			result = result + string(c)
		}
	*/
	return result
}

func randomChar(rand *rand.Rand, charset string, ch chan string) {
	c := charset[rand.Intn(len(charset))]
	ch <- string(c)
}

/*
func FormatComma2(number int) string {
	p := message.NewPrinter(language.English)
	withCommaThousandSep := p.Sprintf("%d", number)
	return withCommaThousandSep
}
*/
