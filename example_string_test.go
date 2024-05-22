package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

const data = "Hello World"

func ExampleReplaceRuneAtIndex() {
	got1 := dechutil.ReplaceRuneAtIndex(data, 'X', 0)
	got2 := dechutil.ReplaceRuneAtIndex(data, 'w', 6)

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	//Xello World
	//Hello world
}

func ExampleReplaceStringAtIndex() {
	got1 := dechutil.ReplaceStringAtIndex(data, "X", 0)
	got2 := dechutil.ReplaceStringAtIndex(data, "w", 6)

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	//Xello World
	//Hello world
}

func ExampleReplaceByteAtIndex() {
	got1 := dechutil.ReplaceByteAtIndex(data, 'X', 0)
	got2 := dechutil.ReplaceByteAtIndex(data, 'w', 6)
	got3 := dechutil.ReplaceByteAtIndex(data, "z"[0], 1)

	fmt.Println(got1)
	fmt.Println(got2)
	fmt.Println(got3)

	// Output:
	//Xello World
	//Hello world
	//Hzllo World
}
