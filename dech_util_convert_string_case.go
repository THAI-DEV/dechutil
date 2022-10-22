package dechutil

import (
	"strings"
)

//* Ref : https://betterprogramming.pub/string-case-styles-camel-pascal-snake-and-kebab-case-981407998841

// User
func ConvertToUpperCapString(input string) string {
	str := strings.ToLower(input)
	strSlice := strings.Split(str, "")

	result := ""
	for i, v := range strSlice {
		if i == 0 {
			result = result + strings.ToUpper(v)

		} else {
			result = result + v
		}
	}

	return result
}

// user
func ConvertToLowerCapString(input string) string {
	str := strings.ToLower(input)
	strSlice := strings.Split(str, "")

	result := ""
	for i, v := range strSlice {
		if i == 0 {
			result = result + strings.ToLower(v)

		} else {
			result = result + v
		}
	}

	return result
}

// UserLoginCount
func ConvertString2PascalCase(inputStr string) string {
	str := strings.ReplaceAll(inputStr, "_", " ")

	strSlice := strings.Split(str, " ")

	result := ""
	for _, v := range strSlice {
		result = result + ConvertToUpperCapString(v)
	}

	return result
}

// userLoginCount
func ConvertString2CamelCase(inputStr string) string {
	result := ConvertString2PascalCase(inputStr)

	strSlice2 := strings.Split(result, "")
	result = ""
	for i, v := range strSlice2 {
		if i == 0 {
			result = result + strings.ToLower(v)

		} else {
			result = result + v
		}
	}

	return result
}

// user_login_count
func ConvertString2SnakeCase(inputStr string) string {
	str := strings.ToLower(inputStr)
	result := strings.ReplaceAll(str, " ", "_")

	return result
}

// USER_LOGIN_COUNT
func ConvertString2SnakeCaseAllCap(inputStr string) string {
	str := strings.ToUpper(inputStr)
	result := strings.ReplaceAll(str, " ", "_")

	return result
}

// user-login-count
func ConvertString2KebabCase(inputStr string) string {
	str := strings.ToLower(inputStr)
	result := strings.ReplaceAll(str, " ", "-")

	return result
}
