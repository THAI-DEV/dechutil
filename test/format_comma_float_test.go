package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestFormatCommaFloat(t *testing.T) {
	type args struct {
		data  float64
		round int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{123, 0}, "123"},
		{"Case 2", args{1234.00, 2}, "1,234.00"},
		{"Case 3", args{1234.123, 2}, "1,234.12"},
		{"Case 4", args{1234567, 0}, "1,234,567"},
		{"Case 5", args{1234567.57, 2}, "1,234,567.57"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.FormatCommaFloat(tt.args.data, tt.args.round); got != tt.want {
				t.Errorf("FormatCommaFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_FormatCommaFloat(b *testing.B) {
	inputData1 := float64(23)
	inputData2 := 2

	for i := 0; i < b.N; i++ {
		dechutil.FormatCommaFloat(inputData1, inputData2)
	}
}
