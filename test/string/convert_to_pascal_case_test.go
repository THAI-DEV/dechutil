package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestToPascalCase(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{"user login count"}, "UserLoginCount"},
		{"Case 2", args{"USER_login_count"}, "UserLoginCount"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.ToPascalCase(tt.args.inputStr); got != tt.want {
				t.Errorf("ToPascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ToPascallCase(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ToPascalCase(inputData)
	}
}
