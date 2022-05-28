package dechutil

import (
	"fmt"
	"reflect"
	"strconv"
)

func Version() string {
	return "DECH Util , Version : 1.0.0 , Last Build : 28/05/2022 18:50"
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

func ElapseTime(milisDiff int64) string {
	if milisDiff < 1000 {
		return "0 second"
	}

	formattedTime := ""
	const secondInMillis = 1000
	const minuteInMillis = secondInMillis * 60
	const hourInMillis = minuteInMillis * 60
	const dayInMillis = hourInMillis * 24
	const weekInMillis = dayInMillis * 7
	const monthInMillis = dayInMillis * 30

	timeElapsed := [6]int{0, 0, 0, 0, 0, 0}
	timeElapsedText := [6]string{"second", "minute", "hour", "'day", "week", "month"}

	timeElapsed[5] = int(toFix(milisDiff / monthInMillis)) // months
	milisDiff = milisDiff % monthInMillis
	timeElapsed[4] = int(toFix(milisDiff / weekInMillis)) // weeks
	milisDiff = milisDiff % weekInMillis
	timeElapsed[3] = int(toFix(milisDiff / dayInMillis)) // days
	milisDiff = milisDiff % dayInMillis
	timeElapsed[2] = int(toFix(milisDiff / hourInMillis)) // hours
	milisDiff = milisDiff % hourInMillis
	timeElapsed[1] = int(toFix(milisDiff / minuteInMillis)) // minutes
	milisDiff = milisDiff % minuteInMillis
	timeElapsed[0] = int(toFix(milisDiff / secondInMillis)) // seconds

	// Only adds 3 significant high valued units
	j := 0
	for i := len(timeElapsed) - 1; i >= 0 && j < 3; i-- {
		// loop from high to low time unit
		if timeElapsed[i] > 0 {
			if j > 0 {
				formattedTime += ", "
			} else {
				formattedTime += ""
			}

			formattedTime += strconv.Itoa(timeElapsed[i]) + " " + timeElapsedText[i]

			if j > 1 {
				formattedTime += "s"
			} else {
				formattedTime += ""
			}

			j++
		}
	}

	return formattedTime
}

func toFix(value int64) int {
	f := float64(value)
	s := strconv.FormatFloat(f, 'f', 0, 64)
	i, _ := strconv.ParseInt(s, 10, 64)

	return int(i)
}
