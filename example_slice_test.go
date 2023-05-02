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

func ExampleFindIndexSlice() {
	dataStr := []string{"c", "a", "z", "x"}
	dataInt := []int{11, 20, 111, 2}
	dataFloat := []float64{11.50, 20.4, 111.43, 2.0}

	got1 := dechutil.FindIndexSlice(dataStr, "z")
	got2 := dechutil.FindIndexSlice(dataStr, "p")
	got3 := dechutil.FindIndexSlice(dataInt, 2)
	got4 := dechutil.FindIndexSlice(dataFloat, 111.43)

	fmt.Println(dechutil.AnyToString(got1, 0))
	fmt.Println(dechutil.AnyToString(got2, 0))
	fmt.Println(dechutil.AnyToString(got3, 0))
	fmt.Println(dechutil.AnyToString(got4, 0))

	// Output:
	// 2
	// -1
	// 3
	// 2
}

func ExampleIsDuplicateSlice() {
	dataStr1 := []string{"c", "a", "z", "x"}
	dataInt1 := []int{11, 20, 111, 2}
	dataFloat1 := []float64{11.50, 20.4, 111.43, 2.0}

	dataStr2 := []string{"c", "a", "c", "x"}
	dataInt2 := []int{11, 20, 111, 20}
	dataFloat2 := []float64{11.50, 20.4, 111.43, 11.50}

	got1 := dechutil.IsDuplicateSlice(dataStr1)
	got2 := dechutil.IsDuplicateSlice(dataInt1)
	got3 := dechutil.IsDuplicateSlice(dataFloat1)

	got4 := dechutil.IsDuplicateSlice(dataStr2)
	got5 := dechutil.IsDuplicateSlice(dataInt2)
	got6 := dechutil.IsDuplicateSlice(dataFloat2)

	fmt.Println(got1)
	fmt.Println(got2)
	fmt.Println(got3)
	fmt.Println(got4)
	fmt.Println(got5)
	fmt.Println(got6)

	// Output:
	// false
	// false
	// false
	// true
	// true
	// true

	for i, v := range []string{} {
		_ = v[i]

	}
}

func ExampleShuffleSlice() {
	data1 := []string{"a", "b", "c", "d", "e", "f"}
	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got1 := dechutil.ShuffleSlice(data1)
	fmt.Println(got1[0] != data1[0])

	got2 := dechutil.ShuffleSlice(data2)
	fmt.Println(got2[0] != data2[0])

	// Output:
	// true
	// true
}

func ExampleCloneSlice() {
	data1 := []string{"a", "b", "c", "d", "e", "f"}
	data2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got1 := dechutil.CloneSlice(data1)
	data1[0] = "aa"
	fmt.Println(got1[0], data1[0])

	got2 := dechutil.CloneSlice(data2)
	data2[0] = 10
	fmt.Println(got2[0], data2[0])

	// Output:
	// a aa
	// 1 10
}
