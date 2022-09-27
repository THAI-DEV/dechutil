package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	s1, b1 := dechutil.ConvertString2Null("abc", "xxxx")          //? abc false
	s2, b2 := dechutil.ConvertString2Null("null val", "null val") //? <nil> true

	fmt.Println(*s1, b1)
	fmt.Println(s2, b2)
}
