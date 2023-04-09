package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	str1 := "test"
	str2 := ""
	var str3 *string

	var str4 *string
	s := "xxx"
	str4 = &s

	str5 := dechutil.AnyValueToPointerValue("xxx")

	str6 := dechutil.AnyValueToPointerValue(70)

	fmt.Println("--------------------------------------")
	fmt.Println("s1", dechutil.IsNullString(&str1), str1)
	fmt.Println("s2", dechutil.IsNullString(&str2), str2)
	fmt.Println("s3", dechutil.IsNullString(str3), str3)
	fmt.Println("s4", dechutil.IsNullString(str4), *str4)
	fmt.Println("s5", dechutil.IsNullString(str5), *str5)
	fmt.Println("s6", "_", *str6)

}
