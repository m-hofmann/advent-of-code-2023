package main

import (
	"fmt"
	"testing"
)

func TestDeterminant(t *testing.T) {
	var tests = []struct {
		a    coord
		b    coord
		want int
	}{
		{coord{8, 5}, coord{1, 6}, 43},
		{coord{y: 1, x: 3}, coord{x: 7, y: 2}, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("determinant(%v, %v) = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			det := determinant(tt.a, tt.b)
			if det != tt.want {
				t.Errorf("determinant(%v, %v) should be %d but is %d", tt.a, tt.b, tt.want, det)
			}
		})
	}
}

func TestShoelace(t *testing.T) {
	var tests = []struct {
		points []coord
		want   float64
	}{
		{[]coord{{1, 6}, {3, 1}, {7, 2}, {4, 4}, {8, 5}}, 16.5},
		{[]coord{{0, 0}, {2, 0}, {2, 2}, {0, 2}}, 16.5},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("shoelace(%v) = %f", tt.points, tt.want), func(t *testing.T) {
			area := shoelace(tt.points)
			if tt.want != area {
				t.Errorf("Area should be 16.5 but is %f", area)
			}
		})
	}

}
