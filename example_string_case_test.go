package dechutil_test

import (
	"fmt"

	"github.com/THAI-DEV/dechutil"
)

func ExampleToCamelCase() {
	str := "user login count"
	got1 := dechutil.ToCamelCase(str)

	fmt.Println(got1)

	// Output:
	// userLoginCount
}

func ExampleToKebabCase() {
	str := "user login count"
	got1 := dechutil.ToKebabCase(str)

	fmt.Println(got1)

	// Output:
	// user-login-count
}

func ExampleToLowerCap() {
	str := "user login count"
	got1 := dechutil.ToLowerCap(str)

	fmt.Println(got1)

	// Output:
	// user login count
}

func ExampleToUpperCap() {
	str := "user login count"
	got1 := dechutil.ToUpperCap(str)

	fmt.Println(got1)

	// Output:
	// User login count
}

func ExampleToPascalCase() {
	str := "user login count"
	got1 := dechutil.ToPascalCase(str)

	fmt.Println(got1)

	// Output:
	// UserLoginCount
}

func ExampleToSnakeCase() {
	str := "user login count"
	got1 := dechutil.ToSnakeCase(str)

	fmt.Println(got1)

	// Output:
	// user_login_count
}

func ExampleToSnakeCaseAllCap() {
	str := "user login count"
	got1 := dechutil.ToSnakeCaseAllCap(str)

	fmt.Println(got1)

	// Output:
	// USER_LOGIN_COUNT
}
