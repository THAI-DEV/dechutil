package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	str := "user login count"
	s := dechutil.ConvertString2SnakeCaseAllCap(str)
	fmt.Println(s)

}
