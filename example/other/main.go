package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	a := dechutil.AnyValueToPointerValue("abc")
	fmt.Println(*a)
}
