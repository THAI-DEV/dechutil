package dechutil

import "testing"

func TestConvertNullString2String(t *testing.T) {
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
		{"Case 1", args{&val, "nill"}, "str"},
		{"Case 2", args{nil, "null val"}, "null val"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertNullString2String(tt.args.data, tt.args.nullStrVal); got != tt.want {
				t.Errorf("ConvertNullString2String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ConvertNullString2String(b *testing.B) {
	var inputData = "xxxx"
	for i := 0; i < b.N; i++ {
		ConvertNullString2String(&inputData, "null")
	}
}
