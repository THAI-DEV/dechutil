package dechutil

import (
	"strconv"
	"time"
)

// * YYYY-MM-DD hh:mm:ss.ffffff
func ConverTime2StringDateTimeFull(timeVal time.Time) string {
	return timeVal.Format("2006-01-02 15:04:05.000000")
}

// * YYYY-MM-DD hh:mm:ss
func ConverTime2StringDateTime(timeVal time.Time) string {
	return timeVal.Format("2006-01-02 15:04:05")
}

// * YYYY-MM-DD
func ConverTime2StringDate(timeVal time.Time) string {
	return timeVal.Format("2006-01-02")
}

// * hh:mm:ss
func ConverTime2StringTime(timeVal time.Time) string {
	return timeVal.Format("15:04:05")
}

// * hh:mm:ss.ffffff
func ConverTime2StringTimeFull(timeVal time.Time) string {
	return timeVal.Format("15:04:05.000000")
}

// * DD/MM/YYYY hh:mm:ss
func ConverTime2ShowStringDateTime(timVal time.Time) string {
	return timVal.Format("02/01/2006 15:04:05")
}

// * DD/MM/YYYY
func ConverTime2ShowStringDate(curr time.Time) string {
	return curr.Format("02/01/2006")
}

func ConverttSecToMinute(sec int) float64 {
	min := float64(sec) / 60
	return min
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
