package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	currDir := dechutil.CurrentCallDir()
	fmt.Println("Curr Dir :", currDir)

	fmt.Println("--------------------------")

	sl1 := []string{"a", "b", "c"}
	sl2 := []int{1, 2, 3}

	arr1 := dechutil.ConvertAnySliceToSliceInf(sl1)
	arr2 := dechutil.ConvertAnySliceToSliceInf(sl2)

	fmt.Printf("%T %+v\n", sl1, sl1)
	fmt.Printf("%T %+v\n", arr1, arr1)

	fmt.Println("--------------------------")

	fmt.Printf("%T %+v\n", sl1, sl2)
	fmt.Printf("%T %+v\n", arr1, arr2)
}
