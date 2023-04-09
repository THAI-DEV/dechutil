package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestNullStringToString(t *testing.T) {
	val := "str"
	type args struct {
		data       *string
		nullStrVal string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{&val, "null val"}, "str"},
		{"Case 2", args{nil, "null val"}, "null val"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.NullStringToString(tt.args.data, tt.args.nullStrVal); got != tt.want {
				t.Errorf("NullStringToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ConvertNullStringToString(b *testing.B) {
	var inputData = "xxxx"
	for i := 0; i < b.N; i++ {
		dechutil.NullStringToString(&inputData, "null")
	}
}
