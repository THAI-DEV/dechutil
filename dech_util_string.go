package dechutil

import "unicode/utf8"

func ReplaceRuneAtIndex(input string, replaceRune rune, index int) string {
	runeSlice := []rune(input)
	runeSlice[index] = replaceRune
	return string(runeSlice)
}

func ReplaceStringAtIndex(input string, replaceString string, index int) string {
	replacementRune, _ := utf8.DecodeRuneInString(replaceString)
	return ReplaceRuneAtIndex(input, replacementRune, index)
}

func ReplaceByteAtIndex(input string, replaceByte byte, index int) string {
	sBytes := []byte(input)
	sBytes[index] = replaceByte

	return string(sBytes)
}
