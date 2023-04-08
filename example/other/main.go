package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	a := dechutil.ConvertAnyValueToPointerValue("abc")
	fmt.Println(*a)
}
