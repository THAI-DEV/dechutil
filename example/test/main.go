package main

import (
	"fmt"
	"strconv"
)

func main() {
	var data float64 = 1234.006
	str := strconv.FormatFloat(data, 'f', 2, 64)
	// str := fmt.Sprint(data)
	fmt.Println(str)
}
