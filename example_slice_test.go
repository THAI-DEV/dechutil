package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleAppendSlice() {
	got1 := dechutil.AppendSlice([]string{"a"}, "b")

	fmt.Println(got1)

	// Output:
	// [a b]
}

func ExampleInsertSlice() {
	got1 := dechutil.InsertSlice([]string{"a", "c"}, "b", 1)

	fmt.Println(got1)

	// Output:
	// [a b c]
}

func ExampleRemoveSlice() {
	got1 := dechutil.RemoveSlice([]string{"a", "b", "c"}, 1)

	fmt.Println(got1)

	// Output:
	// [a c]
}

func ExampleReverseSlice() {
	got1 := dechutil.ReverseSlice([]string{"a", "b", "c"})

	fmt.Println(got1)

	// Output:
	// [c b a]
}

func ExampleRemoveSliceOnTop() {
	got1 := dechutil.RemoveSliceOnTop([]string{"a", "b", "c", "d", "e"}, 2)

	fmt.Println(got1)

	// Output:
	// [d e]
}

func ExampleRemoveSliceOnButtom() {
	got1 := dechutil.RemoveSliceOnButtom([]string{"a", "b", "c", "d", "e"}, 2)

	fmt.Println(got1)

	// Output:
	// [a b]
}

func ExampleSliceStringToString() {
	got1 := dechutil.SliceStringToString([]string{"a", "b", "c"}, "{", "}", "|")

	fmt.Println(got1)

	// Output:
	// {a}|{b}|{c}
}

func ExamplePrintSliceOfSameType() {
	dechutil.PrintSliceOfSameType([]string{"a", "b"})

	// Output:
	// "a"
	//
	// "b"
}

func ExamplePrintSliceOfAnyType() {
	var data []interface{}
	// var data []any

	data = append(data, 1)
	data = append(data, "a")

	dechutil.PrintSliceOfAnyType(data)

	// Output:
	// 1
	//
	// "a"
}