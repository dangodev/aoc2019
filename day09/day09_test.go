package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		intcode    Intcode
		pos        int
		input      []int
		inputCount int
		output     []int
		rb         int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "quine",
			args: args{
				intcode:    []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
				pos:        0,
				input:      []int{},
				inputCount: 0,
				output:     []int{},
				rb:         0,
			},
			want: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			name: "16dig",
			args: args{
				intcode:    []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
				pos:        0,
				input:      []int{},
				inputCount: 0,
				output:     []int{},
				rb:         0,
			},
			want: []int{12345678},
		},
		{
			name: "largeNum",
			args: args{
				intcode:    []int{104, 1125899906842624, 99},
				pos:        0,
				input:      []int{},
				inputCount: 0,
				output:     []int{},
				rb:         0,
			},
			want: []int{1125899906842624},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.intcode, tt.args.pos, tt.args.input, tt.args.inputCount, tt.args.output, tt.args.rb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
