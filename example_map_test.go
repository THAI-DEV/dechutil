package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleCloneMap() {
	data1 := map[string]string{
		"key1": "data1",
		"key2": "data2",
		"key3": "data3",
	}

	data2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	got1 := dechutil.CloneMap(data1)
	got1["key1"] = "new-data1"
	fmt.Println(data1["key1"], got1["key1"])

	got2 := dechutil.CloneMap(data2)
	got2["key1"] = 111
	fmt.Println(data2["key1"], got2["key1"])

	// Output:
	//data1 new-data1
	//1 111

}
