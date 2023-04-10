package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleAnyValueToString() {
	got1 := dechutil.AnyValueToString("123", 0)
	got2 := dechutil.AnyValueToString(567, 0)
	got3 := dechutil.AnyValueToString(567.579, 2)
	got4 := dechutil.AnyValueToString(567.575, 2)
	got5 := dechutil.AnyValueToString(567.574, 2)
	got6 := dechutil.AnyValueToString(567.574, 0)

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
