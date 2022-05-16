package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	s1 := dechutil.SqlNullString(true, "")
	s2 := dechutil.SqlNullString(false, "xxx")

	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
}
