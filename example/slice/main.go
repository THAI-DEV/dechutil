package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func main() {
	fmt.Println("====================")
	s := []string{"a", "b", "c"}
	fmt.Println("Init :", s)

	s = dechutil.InsertSlice(s, "x", 1)
	fmt.Println("Insert x at index 1 :", s)

	s = dechutil.AppendSlice(s, "z")
	fmt.Println("Append z :", s)

	s = dechutil.RemoveSlice(s, 2)
	fmt.Println("Remove at index 2 :", s)

	s = dechutil.ReverseSlice(s)
	fmt.Println("Reverse :", s)

	fmt.Println("-----------------")
	i := []int{1, 2, 3}
	fmt.Println("Init :", i)

	i = dechutil.AppendSlice(i, 99)
	fmt.Println("Append 99 :", i)

	i = dechutil.InsertSlice(i, 0, 2)
	fmt.Println("Insert at index 2 :", i)

	i = dechutil.RemoveSlice(i, 3)
	fmt.Println("Remove at index 3 : ", i)

	i = dechutil.ReverseSlice(i)
	fmt.Println("Reverse :", i)
}
