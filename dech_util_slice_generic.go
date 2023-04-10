package dechutil

import "fmt"

type sliceInf interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		string
}

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

func AppendSlice[T sliceInf](slice []T, val T) []T {
	return append(slice, val)
}

func InsertSlice[T sliceInf](slice []T, val T, index int) []T {
	return append(slice[:index], append([]T{val}, slice[index:]...)...)
}

func RemoveSlice[T sliceInf](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func ReverseSlice[T sliceInf](input []T) []T {
	inputLen := len(input)
	output := make([]T, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}
