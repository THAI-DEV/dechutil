package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	currDir := dechutil.CurrentCallDir()
	fmt.Println("Curr Dir :", currDir)

	fmt.Println("--------------------------")

	slice1 := []string{"a", "b", "c"}
	slice2 := []int{1, 2, 3}

	arr1 := dechutil.ConvertAnySliceToSliceInf(slice1)
	arr2 := dechutil.ConvertAnySliceToSliceInf(slice2)

	fmt.Printf("%T %+v\n", slice1, slice1)
	fmt.Printf("%T %+v\n", arr1, arr1)

	fmt.Println("--------------------------")

	fmt.Printf("%T %+v\n", slice1, slice2)
	fmt.Printf("%T %+v\n", arr1, arr2)
}
