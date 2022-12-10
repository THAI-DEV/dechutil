package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	a := dechutil.PointerOf("abc")
	fmt.Println(*a)
}
