package dechutil

import "fmt"

func RemoveSliceOnTop[T sliceInf](result []T, keepMax int) []T {
	for {
		if len(result) > keepMax {
			result, _ = removeElement(result, 0)
		} else {
			break
		}
	}

	return result
}

func removeElement[T sliceInf](s []T, i int) ([]T, error) {
	// perform bounds checking first to prevent a panic!
	if i >= len(s) || i < 0 {
		return nil, fmt.Errorf("index is out of range. index is %d with slice length %d", i, len(s))
	}

	// This creates a new slice by creating 2 slices from the original:
	// s[:i] -> [1, 2]
	// s[i+1:] -> [4, 5, 6]
	// and joining them together using `append`
	return append(s[:i], s[i+1:]...), nil
}

// * input : []string{"a",""b},"'","'",","  output : 'a','b'
func ConvertSliceStringToString(data []string, wrapStr1 string, wrapStr2 string, saperateStr string) string {
	result := ""
	for i, v := range data {
		result = result + wrapStr1 + v + wrapStr2
		if i < len(data)-1 {
			result = result + saperateStr
		}
	}
	// fmt.Println(result)
	return result
}
