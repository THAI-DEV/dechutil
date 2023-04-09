package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	f := 150.257
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.AnyValueToString(f, 4))
	fmt.Println(dechutil.AnyValueToString(f, 3))
	fmt.Println(dechutil.AnyValueToString(f, 2))
	fmt.Println(dechutil.AnyValueToString(f, 1))
	fmt.Println(dechutil.AnyValueToString(f, 0))

	i := 72815
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.AnyValueToString(i, 0))
}
