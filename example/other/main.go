package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	a := dechutil.AnyToPointer("abc")
	fmt.Println(*a)
}
