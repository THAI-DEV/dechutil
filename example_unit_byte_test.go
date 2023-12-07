package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleUnitByteToByte() {
	got1 := dechutil.UnitByteToByte(1, dechutil.GB)
	got2 := dechutil.UnitByteToByte(1, dechutil.KB)

	fmt.Println(dechutil.FormatCommaFloat(float64(got1), 0))
	fmt.Println(dechutil.FormatCommaFloat(float64(got2), 0))

	// Output:
	// 1,073,741,824
	// 1,024
}

func ExampleByteToUnitByte() {
	got1 := dechutil.ByteToUnitByte(1_073_741_824, dechutil.GB)
	got2 := dechutil.ByteToUnitByte(1_024, dechutil.KB)

	fmt.Println(dechutil.FormatCommaFloat(float64(got1), 0))
	fmt.Println(dechutil.FormatCommaFloat(float64(got2), 0))

	// Output:
	// 1
	// 1
}
