package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleRoundFloat() {
	got1 := dechutil.RoundFloat(15.5075, 2)
	got2 := dechutil.RoundFloat(20.1542, 2)

	fmt.Println(got1)
	fmt.Println(got2)

	// Output:
	// 15.51
	// 20.15
}
