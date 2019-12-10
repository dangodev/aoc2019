package main

import (
	"reflect"
	"testing"
)

func Test_thrusters(t *testing.T) {
	type args struct {
		phases  [5]int
		intcode Intcode
	}
	tests := []struct {
		name string
		args args
		want [5]int
	}{
		{
			name: "test1",
			args: args{
				phases:  [5]int{4, 3, 2, 1, 0},
				intcode: []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			},
			want: [5]int{4, 3, 2, 1, 0},
		},
		{
			name: "test2",
			args: args{
				phases:  [5]int{0, 1, 2, 3, 4},
				intcode: []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			},
			want: [5]int{5, 4, 3, 2, 1},
		},
		{
			name: "test3",
			args: args{
				phases:  [5]int{1, 0, 4, 3, 2},
				intcode: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			},
			want: [5]int{6, 5, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thrusters(tt.args.phases, tt.args.intcode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("thrusters() = %v, want %v", got, tt.want)
			}
		})
	}
}
