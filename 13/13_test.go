package main

import (
	"fmt"
	"testing"
)

func TestGetChecksum(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want bool
	}{
		{0x10101, 0x10001, true},
		{0x10101, 0x10101, false},
		{0x11101, 0x10001, false},
		{0x1010001, 0x10001, true},
		{358, 102, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%x ~~ %x == %t", tt.a, tt.b, tt.want), func(t *testing.T) {
			equal := smudgeEqual(tt.a, tt.b)
			if equal != tt.want {
				t.Errorf("%x ~~ %x is %t but should be %t", tt.a, tt.b, equal, tt.want)
			}
		})
	}
}
