package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestFormatCommaInt(t *testing.T) {
	type args struct {
		data int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{123}, "123"},
		{"Case 2", args{1234}, "1,234"},
		{"Case 3", args{12345}, "12,345"},
		{"Case 4", args{1234567}, "1,234,567"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.FormatCommaInt(tt.args.data); got != tt.want {
				t.Errorf("FormatCommaInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzAdd(f *testing.F) {
	testcases := []int64{123, 1234567890}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, a int64) {
		dechutil.FormatCommaInt(a)
	})
}

func Benchmark_FormatCommaInt(b *testing.B) {
	inputData := int64(23)

	for i := 0; i < b.N; i++ {
		dechutil.FormatCommaInt(inputData)
	}
}
