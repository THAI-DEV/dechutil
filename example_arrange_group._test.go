package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleArrangGroupAsSeparate() {
	input := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09"}

	got1 := dechutil.ArrangGroupAsSeparate(input, 2, true)
	got2 := dechutil.ArrangGroupAsSeparate(input, 2, false)

	fmt.Println(got1[0])
	fmt.Println(got1[1])

	fmt.Println(got2[0])
	fmt.Println(got2[1])

	// Output:
	// 01 , 02 , 03 , 04 , 05
	// 06 , 07 , 08 , 09
	// 01 , 02 , 03 , 04 , 09
	// 05 , 06 , 07 , 08
}

func ExampleArrangGroupAsDistribute() {
	input := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09"}

	got1 := dechutil.ArrangGroupAsDistribute(input, 2)

	fmt.Println(got1[0])
	fmt.Println(got1[1])

	// Output:
	// 01 , 03 , 05 , 07 , 09
	// 02 , 04 , 06 , 08
}
