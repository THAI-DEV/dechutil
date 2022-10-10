package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	dataInt := []int{1, 2, 3, 4, 5}
	dataInt = dechutil.RemoveSliceOnTop(dataInt, 2)

	dataStr := []string{"a", "b", "c"}
	dataStr = dechutil.RemoveSliceOnTop(dataStr, 2)

	dataFloat := []float64{11.5, 20.1, 13.4, 4.00, 5.15}
	dataFloat = dechutil.RemoveSliceOnTop(dataFloat, 2)

	fmt.Println(dataInt)
	fmt.Println(dataStr)
	fmt.Println(dataFloat)

}
