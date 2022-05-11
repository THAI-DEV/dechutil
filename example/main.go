package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	s := dechutil.CurrentCallDir()
	fmt.Println("Curr Dir :", s)
}
