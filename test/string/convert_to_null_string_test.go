package dechutil

import (
	"reflect"
	"testing"

	"github.com/THAI-DEV/dechutil"
)

func TestStringToNullString(t *testing.T) {
	var result *string
	s := "xxx"
	result = &s

	type args struct {
		data         string
		defindStrVal string
	}
	tests := []struct {
		name  string
		args  args
		want  *string
		want1 bool
	}{
		// TODO: Add test cases.
		{"Case 1", args{"null val", "null val"}, nil, true},
		{"Case 2", args{s, "null val"}, result, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := dechutil.StringToNullString(tt.args.data, tt.args.defindStrVal)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToNull() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StringToNull() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Benchmark_StringToNullString(b *testing.B) {
	inputData1 := "xxx"
	inputData2 := "xxx"

	for i := 0; i < b.N; i++ {
		dechutil.StringToNullString(inputData1, inputData2)
	}
}
