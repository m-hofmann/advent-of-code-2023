package main

import (
	"fmt"
	"testing"
)

func TestHASH(t *testing.T) {
	var tests = []struct {
		s    string
		want int64
	}{
		{"HASH", 52},
		{"rn=1", 30},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("HASH(%s)==%d", tt.s, tt.want), func(t *testing.T) {
			hash := HASH(tt.s)
			if hash != tt.want {
				t.Errorf("HASH(%s) == %d but should be %d", tt.s, hash, tt.want)
			}
		})
	}
}
