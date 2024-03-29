package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleRoundFloat() {
	got1 := dechutil.RoundFloat(15.5075, 2)
	got2 := dechutil.RoundFloat(20.1542, 2)
	got3 := dechutil.RoundFloat(45.1978, 0)

	fmt.Println(got1)
	fmt.Println(got2)
	fmt.Println(got3)

	// Output:
	// 15.51
	// 20.15
	// 45
}

func ExampleFormatCommaInt() {
	got1 := dechutil.FormatCommaInt(123)
	got2 := dechutil.FormatCommaInt(12345)

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	// 123
	// 12,345
}

func ExampleFormatCommaFloat() {
	got1 := dechutil.FormatCommaFloat(123.0, 2)
	got2 := dechutil.FormatCommaFloat(12345.0, 2)
	got3 := dechutil.FormatCommaFloat(12345.678, 2)
	got4 := dechutil.FormatCommaFloat(12345.123, 2)

	fmt.Println(got1)
	fmt.Println(got2)
	fmt.Println(got3)
	fmt.Println(got4)

	// Output:
	// 123.00
	// 12,345.00
	// 12,345.68
	// 12,345.12
}

func ExampleIsNullString() {
	str1 := "test"
	str2 := ""
	var str3 *string

	var str4 *string
	s := "xxx"
	str4 = &s

	str5 := dechutil.AnyToPointer("xxx")

	got1 := dechutil.IsNullString(&str1)
	got2 := dechutil.IsNullString(&str2)
	got3 := dechutil.IsNullString(str3)
	got4 := dechutil.IsNullString(str4)
	got5 := dechutil.IsNullString(str5)

	fmt.Println(got1)
	fmt.Println(got2)
	fmt.Println(got3)
	fmt.Println(got4)
	fmt.Println(got5)

	// Output:
	// false
	// false
	// true
	// false
	// false
}

func ExampleAnyToPointer() {
	got1 := dechutil.AnyToPointer("xxx")
	got2 := dechutil.AnyToPointer("")

	var nulStr *string
	got3 := dechutil.AnyToPointer(nulStr)

	got4 := dechutil.AnyToPointer(7)
	got5 := dechutil.AnyToPointer(20.5)
	got6 := dechutil.AnyToPointer(true)

	fmt.Println(*got1)
	fmt.Println(*got2)
	fmt.Println(*got3)
	fmt.Println(*got4)
	fmt.Println(*got5)
	fmt.Println(*got6)

	// Output:
	// xxx
	//
	// <nil>
	// 7
	// 20.5
	// true
}
func ExampleTypeAndKindOfObject() {
	got11, got12 := dechutil.TypeAndKindOfObject("test")
	got21, got22 := dechutil.TypeAndKindOfObject(169)
	got31, got32 := dechutil.TypeAndKindOfObject(5.13)
	got41, got42 := dechutil.TypeAndKindOfObject(true)
	got51, got52 := dechutil.TypeAndKindOfObject([]string{})

	fmt.Println(got11, got12)
	fmt.Println(got21, got22)
	fmt.Println(got31, got32)
	fmt.Println(got41, got42)
	fmt.Println(got51, got52)

	// Output:
	// string string
	// int int
	// float64 float64
	// bool bool
	// []string slice
}
