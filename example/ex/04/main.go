package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	f := 150.257
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.AnyToString(f, 4))
	fmt.Println(dechutil.AnyToString(f, 3))
	fmt.Println(dechutil.AnyToString(f, 2))
	fmt.Println(dechutil.AnyToString(f, 1))
	fmt.Println(dechutil.AnyToString(f, 0))

	i := 72815
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.AnyToString(i, 0))
}
