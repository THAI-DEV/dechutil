package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	str := "751250.579"
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.StringToFloat(str))
	fmt.Println(dechutil.StringToInt(str))

	str = "72815"
	fmt.Println("--------------------------------------")
	fmt.Println(dechutil.StringToFloat(str))
	fmt.Println(dechutil.StringToInt(str))
}
