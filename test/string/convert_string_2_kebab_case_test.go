package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestConvertString2KebabCase(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{"user login count"}, "user-login-count"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.ConvertString2KebabCase(tt.args.inputStr); got != tt.want {
				t.Errorf("ConvertString2KebabCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ConvertString2KebabCase(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ConvertString2KebabCase(inputData)
	}
}
