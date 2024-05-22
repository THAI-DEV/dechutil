package dechutil_test

import (
	"fmt"
	"time"

	"github.com/THAI-DEV/dechutil"
)

const layoutDateTime = "2006-01-02 15:04:05.000000"

func ExampleTimeToShowStringDate() {
	t, _ := time.Parse(layoutDateTime, "2022-12-25 17:31:45.000000")

	got1 := dechutil.TimeToShowStringDate(t)

	fmt.Println(got1)

	// Output:
	// 25/12/2022
}

func ExampleTimeToShowStringDateTime() {
	t, _ := time.Parse(layoutDateTime, "2022-12-25 17:32:45.000000")

	got1 := dechutil.TimeToShowStringDateTime(t)

	fmt.Println(got1)

	// Output:
	// 25/12/2022 17:32:45
}

func ExampleTimeToStringDate() {
	t, _ := time.Parse(layoutDateTime, "2022-12-25 17:33:45.000000")

	got1 := dechutil.TimeToStringDate(t)

	fmt.Println(got1)

	// Output:
	// 2022-12-25
}

func ExampleTimeToStringDateTime() {
	t, _ := time.Parse(layoutDateTime, "2022-12-25 17:34:45.000000")

	got1 := dechutil.TimeToStringDateTime(t)

	fmt.Println(got1)

	// Output:
	// 2022-12-25 17:34:45
}

func ExampleTimeToStringDateTimeFull() {
	t, _ := time.Parse(layoutDateTime, "2022-12-25 17:30:45.123456")

	got1 := dechutil.TimeToStringDateTimeFull(t)

	fmt.Println(got1)

	// Output:
	// 2022-12-25 17:30:45.123456
}

func ExampleTimeToStringTime() {
	t, _ := time.Parse(layoutDateTime, "2022-12-25 17:30:45.000000")

	got1 := dechutil.TimeToStringTime(t)

	fmt.Println(got1)

	// Output:
	// 17:30:45
}

func ExampleTimeToStringTimeFull() {
	t, _ := time.Parse(layoutDateTime, "2022-12-25 17:30:45.123456")

	got1 := dechutil.TimeToStringTimeFull(t)

	fmt.Println(got1)

	// Output:
	// 17:30:45.123456
}

func ExampleTimeToStringFormat() {
	formatTime := "20060102 150405"
	t, _ := time.Parse(formatTime, "20221225 173045")

	got1 := dechutil.TimeToStringFormat(t, formatTime)

	fmt.Println(got1)

	// Output:
	// 20221225 173045
}
