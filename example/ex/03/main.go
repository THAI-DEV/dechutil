package main

import (
	"fmt"
	"sync"

	"github.com/THAI-DEV/dechutil"
)

const separator = "-------------------------------------------------"

var list = []string{}

func main() {
	fmt.Println(separator)
	// CaseNirmal()
	// fmt.Println(separator)

	// CaseLoop()
	// fmt.Println(separator)

	// CaseSync()
	// fmt.Println(separator)

	CaseASync()
	fmt.Println(separator)
}

func CaseNirmal() {
	fmt.Println(dechutil.RandomString(10, true, true, true, false))
	fmt.Println(dechutil.RandomString(10, false, false, true, false))
}

func CaseLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i+1, dechutil.RandomString(10, true, true, true, false))
	}
}

func CaseSync() {
	for i := 0; i < 100; i++ {
		gen()
	}

	fmt.Println("List:", list)
	fmt.Println("Size:", len(list))
}

func CaseASync() {
	n := 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < n; i++ {
		wg.Add(1)
		go genAsync(&wg, &mu)
	}

	wg.Wait()
	fmt.Println("List:", list)
	fmt.Println("Size:", len(list))
}

func gen() {
	s := dechutil.RandomString(8, true, true, true, true)

	if isDuplicate(list, s) {
		fmt.Println("Duplicate:", s)
	} else {
		list = append(list, s)
	}
}

func genAsync(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	s := dechutil.RandomString(8, true, true, true, true)

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
