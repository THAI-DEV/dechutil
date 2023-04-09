package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	fmt.Println(dechutil.RandomString(10, true, true, true, false))
	fmt.Println(dechutil.RandomString(10, false, false, true, false))
}
