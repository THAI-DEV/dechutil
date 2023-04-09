package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestToSnakeCase(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{"user login count"}, "user_login_count"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.ToSnakeCase(tt.args.inputStr); got != tt.want {
				t.Errorf("ToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ToSnakeCase(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ToSnakeCase(inputData)
	}
}
