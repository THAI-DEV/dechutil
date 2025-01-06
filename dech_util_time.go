package dechutil

import (
	"strconv"
	"time"
)

// * YYYY-MM-DD hh:mm:ss.ffffff
func TimeToStringDateTimeFull(timeVal time.Time) string {
	return timeVal.Format("2006-01-02 15:04:05.000000")
}

// * YYYY-MM-DD hh:mm:ss
func TimeToStringDateTime(timeVal time.Time) string {
	return timeVal.Format("2006-01-02 15:04:05")
}

// * YYYY-MM-DD
func TimeToStringDate(timeVal time.Time) string {
	return timeVal.Format("2006-01-02")
}

// * hh:mm:ss
func TimeToStringTime(timeVal time.Time) string {
	return timeVal.Format("15:04:05")
}

// * hh:mm:ss.ffffff
func TimeToStringTimeFull(timeVal time.Time) string {
	return timeVal.Format("15:04:05.000000")
}

// * DD/MM/YYYY hh:mm:ss
func TimeToShowStringDateTime(timVal time.Time) string {
	return timVal.Format("02/01/2006 15:04:05")
}

// * DD/MM/YYYY
func TimeToShowStringDate(curr time.Time) string {
	return curr.Format("02/01/2006")
}

// 2006-01-02 15:04:05.000000
func TimeToStringFormat(timeVal time.Time, format string) string {
	return timeVal.Format(format)
}

func UnitSecToMinute(sec int) float64 {
	min := float64(sec) / 60
	return min
}

func ElapseTime(milsDiff int64) string {
	if milsDiff < 1000 {
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
	timeElapsedText := [6]string{"second", "minute", "hour", "day", "week", "month"}

	timeElapsed[5] = int(toFix(milsDiff / monthInMillis)) // months
	milsDiff = milsDiff % monthInMillis
	timeElapsed[4] = int(toFix(milsDiff / weekInMillis)) // weeks
	milsDiff = milsDiff % weekInMillis
	timeElapsed[3] = int(toFix(milsDiff / dayInMillis)) // days
	milsDiff = milsDiff % dayInMillis
	timeElapsed[2] = int(toFix(milsDiff / hourInMillis)) // hours
	milsDiff = milsDiff % hourInMillis
	timeElapsed[1] = int(toFix(milsDiff / minuteInMillis)) // minutes
	milsDiff = milsDiff % minuteInMillis
	timeElapsed[0] = int(toFix(milsDiff / secondInMillis)) // seconds

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

func SecondTimeToUnitTime(secondTime uint64) (days uint64, hours uint64, minutes uint64, seconds uint64) {
	const (
		secondsPerMinute = 60
		secondsPerHour   = 60 * secondsPerMinute
		secondsPerDay    = 24 * secondsPerHour
	)

	days = secondTime / secondsPerDay
	secondTime %= secondsPerDay

	hours = secondTime / secondsPerHour
	secondTime %= secondsPerHour

	minutes = secondTime / secondsPerMinute
	seconds = secondTime % secondsPerMinute

	//fmt.Sprintf("%d day, %d hour, %d min, %d sec", days, hours, minutes, seconds)

	return
}

func CurrentBkkTime() time.Time {
	location := time.FixedZone("GMT+7", 7*3600)
	currentTime := time.Now().In(location)
	return currentTime
}

func toFix(value int64) int {
	f := float64(value)
	s := strconv.FormatFloat(f, 'f', 0, 64)
	i, _ := strconv.ParseInt(s, 10, 64)

	return int(i)
}
