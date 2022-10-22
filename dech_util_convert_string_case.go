package dechutil

import (
	"strings"
)

//* Ref : https://betterprogramming.pub/string-case-styles-camel-pascal-snake-and-kebab-case-981407998841

// User
func ConvertToUpperCap(input string) string {
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
func ConvertToLowerCap(input string) string {
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
func ConvertToPascalCase(inputStr string) string {
	str := strings.ReplaceAll(inputStr, "_", " ")

	strSlice := strings.Split(str, " ")

	result := ""
	for _, v := range strSlice {
		result = result + ConvertToUpperCap(v)
	}

	return result
}

// userLoginCount
func ConvertToCamelCase(inputStr string) string {
	result := ConvertToPascalCase(inputStr)

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
func ConvertToSnakeCase(inputStr string) string {
	str := strings.ToLower(inputStr)
	result := strings.ReplaceAll(str, " ", "_")

	return result
}

// USER_LOGIN_COUNT
func ConvertToSnakeCaseAllCap(inputStr string) string {
	str := strings.ToUpper(inputStr)
	result := strings.ReplaceAll(str, " ", "_")

	return result
}

// user-login-count
func ConvertToKebabCase(inputStr string) string {
	str := strings.ToLower(inputStr)
	result := strings.ReplaceAll(str, " ", "-")

	return result
}
