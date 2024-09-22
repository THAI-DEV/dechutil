package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleFormatStringComma() {
	commaPosition := 3
	data := []string{
		"12",
		"123",
		"1234",
		"12345",
		"123456",
		"1234567",
		"12345678",
		"123456789",
		"1234567890",
	}

	fmt.Println()
	for _, v := range data {
		got := dechutil.FormatStringComma(v, commaPosition)
		fmt.Println(got)
	}

	// Output:
	//12
	//123
	//1,234
	//12,345
	//123,456
	//1,234,567
	//12,345,678
	//123,456,789
	//1,234,567,890

}
