package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	f := 150.257
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.ConvertAnyValueToString(f, 4))
	fmt.Println(dechutil.ConvertAnyValueToString(f, 3))
	fmt.Println(dechutil.ConvertAnyValueToString(f, 2))
	fmt.Println(dechutil.ConvertAnyValueToString(f, 1))
	fmt.Println(dechutil.ConvertAnyValueToString(f, 0))

	i := 72815
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.ConvertAnyValueToString(i, 0))
}
