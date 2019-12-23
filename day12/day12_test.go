package main

import (
	"reflect"
	"testing"
)

func Test_step(t *testing.T) {
	type args struct {
		moonInput []moon
	}
	tests := []struct {
		name string
		args args
		want []moon
	}{
		{
			name: "step 1",
			args: args{
				moonInput: []moon{
					moon{
						pos:      coord{x: -1, y: 0, z: 2},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 2, y: -10, z: -7},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 4, y: -8, z: 8},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 3, y: 5, z: -1},
						velocity: coord{x: 0, y: 0, z: 0},
					},
				},
			},
			want: []moon{
				moon{
					pos:      coord{x: 2, y: -1, z: 1},
					velocity: coord{x: 3, y: -1, z: -1},
				},
				moon{
					pos:      coord{x: 3, y: -7, z: -4},
					velocity: coord{x: 1, y: 3, z: 3},
				},
				moon{
					pos:      coord{x: 1, y: -7, z: 5},
					velocity: coord{x: -3, y: 1, z: -3},
				},
				moon{
					pos:      coord{x: 2, y: 2, z: 0},
					velocity: coord{x: -1, y: -3, z: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := step(tt.args.moonInput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("step() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcRepeat(t *testing.T) {
	type args struct {
		moons []moon
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "long",
			args: args{
				moons: []moon{
					moon{
						pos:      coord{x: -1, y: -0, z: 2},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 2, y: -10, z: -7},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 4, y: -8, z: 8},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 3, y: 5, z: -1},
						velocity: coord{x: 0, y: 0, z: 0},
					},
				},
			},
			want: 2772,
		},
		{
			name: "long",
			args: args{
				moons: []moon{
					moon{
						pos:      coord{x: -8, y: -10, z: 0},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 5, y: 5, z: 10},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 2, y: -7, z: 3},
						velocity: coord{x: 0, y: 0, z: 0},
					},
					moon{
						pos:      coord{x: 9, y: -8, z: -3},
						velocity: coord{x: 0, y: 0, z: 0},
					},
				},
			},
			want: 4686774924,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcRepeat(tt.args.moons)
			if got != tt.want {
				t.Errorf("calcRepeat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
