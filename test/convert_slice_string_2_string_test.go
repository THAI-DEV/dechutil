package dechutil

import (
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestConvertSliceString2String(t *testing.T) {
	type args struct {
		data        []string
		wrapStr1    string
		wrapStr2    string
		saperateStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Case 1", args{[]string{"a", "b", "c"}, "'", "'", ","}, "'a','b','c'"},
		{"Case 2", args{[]string{"a", "b", "c"}, "", "", "|"}, "a|b|c"},
		{"Case 3", args{[]string{"a"}, "'", "'", ","}, "'a'"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dechutil.ConvertSliceString2String(tt.args.data, tt.args.wrapStr1, tt.args.wrapStr2, tt.args.saperateStr); got != tt.want {
				t.Errorf("ConvertSliceString2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ConvertSliceString2String(b *testing.B) {
	inputData1 := []string{"a", "b", "c"}
	inputData2 := "'"
	inputData3 := "'"
	inputData4 := ","

	for i := 0; i < b.N; i++ {
		dechutil.ConvertSliceString2String(inputData1, inputData2, inputData3, inputData4)
	}
}
