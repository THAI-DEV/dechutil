package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {

	s1 := dechutil.SqlNullString(true, "")     //? String: Valid:false}
	s2 := dechutil.SqlNullString(false, "xxx") //? {String:xxx Valid:true}
	s3 := dechutil.SqlNullStringRef(nil)       //? {String: Valid:false}

	s4 := dechutil.SqlNullValueString(s1) //? <nil>
	s5 := dechutil.SqlNullValueString(s2) //? xxx

	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
	fmt.Printf("%+v\n", s3)

	fmt.Printf("%+v\n", s4)
	fmt.Printf("%+v\n", *s5)

}
