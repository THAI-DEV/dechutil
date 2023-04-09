package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestToSnakeCaseAllCap(t *testing.T) {
	type args struct {
		inputStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{"user login count"}, "USER_LOGIN_COUNT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.ToSnakeCaseAllCap(tt.args.inputStr); got != tt.want {
				t.Errorf("ToSnakeCaseAllCap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ToSnakeCaseAllCap(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ToSnakeCaseAllCap(inputData)
	}
}
