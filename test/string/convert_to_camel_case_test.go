package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestConvertString2CamelCase(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{"user login count"}, "userLoginCount"},
		{"Case 2", args{"USER_login_count"}, "userLoginCount"},
		{"Case 3", args{"USER"}, "user"},
		{"Case 4", args{"uSER"}, "user"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.ConvertToCamelCase(tt.args.inputStr); got != tt.want {
				t.Errorf("ConvertString2CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ConvertString2CamelCase(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ConvertToCamelCase(inputData)
	}
}
