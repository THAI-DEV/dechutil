package dechutil

import (
	"fmt"
	"math/rand"
)

type sliceInf interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		string | bool
}

func RemoveSliceOnTop[T any](result []T, keepMax int) []T {
	for {
		if len(result) > keepMax {
			result, _ = removeElement(result, 0)
		} else {
			break
		}
	}

	return result
}

func RemoveSliceOnButtom[T any](result []T, keepMax int) []T {
	for {
		if len(result) > keepMax {
			result, _ = removeElement(result, len(result)-1)
		} else {
			break
		}
	}

	return result
}

func removeElement[T any](s []T, i int) ([]T, error) {
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

func AppendSlice[T any](slice []T, val T) []T {
	return append(slice, val)
}

func InsertSlice[T any](slice []T, val T, index int) []T {
	return append(slice[:index], append([]T{val}, slice[index:]...)...)
}

func RemoveSlice[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func ReverseSlice[T any](input []T) []T {
	inputLen := len(input)
	output := make([]T, inputLen)

	for i, n := range input {
		j := inputLen - i - 1

		output[j] = n
	}

	return output
}

func FindIndexSlice[T sliceInf](data []T, findVal T) int {
	for i, v := range data {
		if v == findVal {
			return i
		}
	}
	return -1
}

func IsDuplicateSlice[T sliceInf](data []T) bool {
	for i, v := range data {
		ind := findNextIndexInSliceByVal(data, v, i+1)
		if ind != -1 { //found
			return true
		}
	}

	return false
}

func ShuffleSlice[T any](data []T) []T {
	result := make([]T, len(data))

	perm := rand.Perm(len(data))
	for i, v := range perm {
		result[v] = data[i]
	}

	return result
}

func CloneSlice[T any](data []T) []T {
	result := make([]T, len(data))
	copy(result, data)

	return result
}

func CombineSameSlice[T sliceInf](sourceData []T, compareData []T) []T {
	result := []T{}

	for _, v := range sourceData {
		i := FindIndexSlice(compareData, v)
		if i != -1 {
			result = append(result, v)
		}
	}

	return result
}

func CombineDiffSlice[T sliceInf](sourceData []T, compareData []T) []T {
	result := []T{}

	for _, v := range sourceData {
		i := FindIndexSlice(compareData, v)
		if i == -1 {
			result = append(result, v)
		}
	}

	return result
}

func CombineUniqSlice[T sliceInf](sourceData []T) []T {
	result := []T{}

	for _, v := range sourceData {
		i := FindIndexSlice(result, v)
		if i == -1 {
			result = append(result, v)
		}
	}

	return result
}

func findNextIndexInSliceByVal[T sliceInf](data []T, findVal T, startInd int) int {
	for i := startInd; i < len(data); i++ {
		if data[i] == findVal {
			return i
		}
	}
	return -1
}
