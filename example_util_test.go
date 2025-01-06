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

func ExampleCalIndexByPageNo() {
	var got1 = [3]int{0, 0, 0}
	var got2 = [3]int{0, 0, 0}
	var got3 = [3]int{0, 0, 0}
	var got4 = [3]int{0, 0, 0}
	var got5 = [3]int{0, 0, 0}
	var got6 = [3]int{0, 0, 0}
	var got7 = [3]int{0, 0, 0}
	var got8 = [3]int{0, 0, 0}

	got1[0], got1[1], got1[2] = dechutil.CalIndexByPageNo(1, 10, 100)
	got2[0], got2[1], got2[2] = dechutil.CalIndexByPageNo(10, 10, 100)

	got3[0], got3[1], got3[2] = dechutil.CalIndexByPageNo(1, 4, 8)

	got4[0], got4[1], got4[2] = dechutil.CalIndexByPageNo(2, 5, 6)
	got5[0], got5[1], got5[2] = dechutil.CalIndexByPageNo(2, 5, 7)
	got6[0], got6[1], got6[2] = dechutil.CalIndexByPageNo(2, 5, 8)
	got7[0], got7[1], got7[2] = dechutil.CalIndexByPageNo(2, 5, 9)

	got8[0], got8[1], got8[2] = dechutil.CalIndexByPageNo(3, 5, 10)

	fmt.Println(got1[0], got1[1], got1[2])
	fmt.Println(got2[0], got2[1], got2[2])
	fmt.Println(got3[0], got3[1], got3[2])
	fmt.Println(got4[0], got4[1], got4[2])
	fmt.Println(got5[0], got5[1], got5[2])
	fmt.Println(got6[0], got6[1], got6[2])
	fmt.Println(got7[0], got7[1], got7[2])
	fmt.Println(got8[0], got8[1], got8[2])

	// Output:
	// 0 9 10
	// 90 99 10
	// 0 3 2
	// 5 5 2
	// 5 6 2
	// 5 7 2
	// 5 8 2
	// -1 -1 2
}
