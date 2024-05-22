package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleReplaceRuneAtIndex() {
	data1 := "Hello World"

	got1 := dechutil.ReplaceRuneAtIndex(data1, 'X', 0)
	got2 := dechutil.ReplaceRuneAtIndex(data1, 'w', 6)

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	//Xello World
	//Hello world
}

func ExampleReplaceStringAtIndex() {
	data1 := "Hello World"

	got1 := dechutil.ReplaceStringAtIndex(data1, "X", 0)
	got2 := dechutil.ReplaceStringAtIndex(data1, "w", 6)

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	//Xello World
	//Hello world
}
