package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestConvertString2SnakeCaseAllCap(t *testing.T) {
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
			if got := dechutil.ConvertToSnakeCaseAllCap(tt.args.inputStr); got != tt.want {
				t.Errorf("ConvertString2SnakeCaseAllCap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ConvertString2SnakeCaseAllCap(b *testing.B) {
	inputData := "user login count"

	for i := 0; i < b.N; i++ {
		dechutil.ConvertToSnakeCaseAllCap(inputData)
	}
}
