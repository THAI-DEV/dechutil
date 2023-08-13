package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleAnyToString() {
	got1 := dechutil.AnyToString("123", 0)
	got2 := dechutil.AnyToString(567, 0)
	got3 := dechutil.AnyToString(567.579, 2)
	got4 := dechutil.AnyToString(567.575, 2)
	got5 := dechutil.AnyToString(567.574, 2)
	got6 := dechutil.AnyToString(567.574, 0)

	fmt.Println(got1)
	fmt.Println(got2)
	fmt.Println(got3)
	fmt.Println(got4)
	fmt.Println(got5)
	fmt.Println(got6)

	// Output:
	// 123
	// 567
	// 567.58
	// 567.58
	// 567.57
	// 568
}

func ExampleAnySliceToSliceInf() {
	got1 := dechutil.AnySliceToSliceInf([]string{"a", "b"})

	fmt.Printf("%v %T\n", got1, got1)

	// Output:
	// [a b] []interface {}
}

func ExampleNullStringToString() {
	s1 := "test"
	var s2 *string

	got1 := dechutil.NullStringToString(&s1, "nullVal")
	got2 := dechutil.NullStringToString(s2, "nullVal")

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	// test
	// nullVal
}

func ExampleStringToNullString() {
	s1 := "abc"
	s2 := "xyz"

	got11, got12 := dechutil.StringToNullString(s1, s2)
	got21, got22 := dechutil.StringToNullString(s1, s1)

	if got12 {
		fmt.Println(got11, got12)
	} else {
		fmt.Println(*got11, got12)
	}

	if got22 {
		fmt.Println(got21, got22)
	} else {
		fmt.Println(*got21, got22)
	}

	// Output:
	// abc false
	// <nil> true
}

func ExampleStringToFloat() {
	got1 := dechutil.StringToFloat("123.453")
	got2 := dechutil.StringToFloat("123")

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	// 123.453
	// 123
}

func ExampleStringToInt() {
	got1 := dechutil.StringToInt("123")

	fmt.Println(got1)

	// Output:
	// 123
}

func ExampleStringToTime() {
	got1 := dechutil.StringToTime("2022-12-25 17:30:45.000000", "2006-01-02 15:04:05.000000")

	str1 := dechutil.TimeToStringDateTimeFull(got1)

	fmt.Println(str1)

	// Output:
	// 2022-12-25 17:30:45.000000
}
