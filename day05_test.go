package main

import (
	"reflect"
	"testing"
)

func Test_runv2(t *testing.T) {
	type args struct {
		intcode Intcode
		input   int
	}
	tests := []struct {
		name string
		args args
		want Intcode
	}{
		{
			name: "test1",
			args: args{
				intcode: Intcode{3, 9, 8, 9, 10, 9, 4, 9, 99, -1},
				input:   8,
			},
			want: Intcode{1002, 4, 3, 4, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runv2(tt.args.intcode, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runv2() = %v, want %v", got, tt.want)
			}
		})
	}
}
