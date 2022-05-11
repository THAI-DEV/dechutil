package dechutil

type sliceInf interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		string
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
