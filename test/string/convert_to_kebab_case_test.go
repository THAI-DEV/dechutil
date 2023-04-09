package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestToKebabCase(t *testing.T) {
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
			if got := dechutil.ToKebabCase(tt.args.inputStr); got != tt.want {
				t.Errorf("ToKebabCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ToKebabCase(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ToKebabCase(inputData)
	}
}
