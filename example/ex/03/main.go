package main

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

const separator = "-------------------------------------------------"

var list = []string{}

func main() {
	fmt.Println(separator)
	case1()
	fmt.Println(separator)
	case2()
	fmt.Println(separator)
	case3()
	fmt.Println("-------------------------------------------------")
}

func case1() {
	fmt.Println(dechutil.RandomString(10, true, true, true, false))
	fmt.Println(dechutil.RandomString(10, false, false, true, false))
}

func case2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i+1, dechutil.RandomString(10, true, true, true, false))
	}
}

func case3() {
	for i := 0; i < 100; i++ {
		gen()
	}

	fmt.Println("List:", list)
	fmt.Println("Size:", len(list))
}

func gen() {
	s := dechutil.RandomString(8, true, true, true, true)

	// fmt.Println("String:", s)

	if isDuplicate(list, s) {
		fmt.Println("Duplicate:", s)
	} else {
		list = append(list, s)
	}
}

func isDuplicate(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
