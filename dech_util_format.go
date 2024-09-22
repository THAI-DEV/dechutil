package dechutil

import (
	"strings"
)

func FormatStringComma(input string, commaPosition int) string {
	num1 := ""
	ind := strings.Index(input, ".")
	num2 := ""

	if ind > -1 {
		num1 = input[0:ind]
		num2 = input[ind:]
		input = num1
	}

	result := ""

	numLen := len(input)
	if numLen <= commaPosition {
		return input
	}

	numLoop := numLen / commaPosition
	if numLen%commaPosition > 0 {
		numLoop = numLoop + 1
	}

	list := []string{}
	index := numLen
	for i := numLoop; i >= 1; i-- {
		end := index
		start := end - commaPosition
		if start < 0 {
			start = 0
		}

		s := input[start:end]

		list = append(list, s)
		index = index - commaPosition
	}

	for i := len(list); i > 0; i-- {
		result = result + list[i-1]
		if i > 1 {
			result = result + ","
		}
	}

	if ind > 0 {
		result = result + num2
	}

	return result
}
