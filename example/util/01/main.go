package main

import (
	"fmt"
	"time"

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

	fmt.Println("--------------------------")

	date1 := time.Date(2022, time.Month(5), 16, 16, 2, 1, 0, time.UTC)
	date2 := time.Date(2022, time.Month(5), 16, 20, 5, 17, 0, time.UTC)
	diff := date2.Sub(date1)

	s := dechutil.ElapseTime(diff.Milliseconds())
	fmt.Println(s)
}
