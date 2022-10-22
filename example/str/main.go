package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	str := "user login count"
	s := dechutil.ConvertToSnakeCaseAllCap(str)
	fmt.Println(s)

}
