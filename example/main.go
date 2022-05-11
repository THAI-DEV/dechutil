package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	currDir := dechutil.CurrentCallDir()
	fmt.Println("Curr Dir :", currDir)
}
