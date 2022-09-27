package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestRoundFloat(t *testing.T) {
	type args struct {
		data     float64
		decimals int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"Case 1", args{123, 0}, 123},
		{"Case 2", args{0.12345, 1}, 0.1},
		{"Case 3", args{0.12345, 2}, 0.12},
		{"Case 4", args{0.12345, 3}, 0.123},
		{"Case 5", args{0.12345, 4}, 0.1235},
		{"Case 6", args{512.005, 2}, 512.01},
		{"Case 7", args{512.005, 1}, 512.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.RoundFloat(tt.args.data, tt.args.decimals); got != tt.want {
				t.Errorf("RoundFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_RoundFloat(b *testing.B) {
	inputData1 := 12.456
	inputData2 := 2

	for i := 0; i < b.N; i++ {
		dechutil.RoundFloat(inputData1, inputData2)
	}
}
