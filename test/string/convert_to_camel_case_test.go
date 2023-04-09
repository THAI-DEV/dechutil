package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestToCamelCase(t *testing.T) {
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
			if got := dechutil.ToCamelCase(tt.args.inputStr); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ToCamelCase(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ToCamelCase(inputData)
	}
}
