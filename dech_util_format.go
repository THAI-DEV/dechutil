package dechutil

func FormatStringComma(input string, commaPosition int) string {
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

	return result
}
