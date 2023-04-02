package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	f := 150.257
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.ConvertValueObjectToString(f, 4))
	fmt.Println(dechutil.ConvertValueObjectToString(f, 3))
	fmt.Println(dechutil.ConvertValueObjectToString(f, 2))
	fmt.Println(dechutil.ConvertValueObjectToString(f, 1))
	fmt.Println(dechutil.ConvertValueObjectToString(f, 0))

	i := 72815
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.ConvertValueObjectToString(i, 0))
}
