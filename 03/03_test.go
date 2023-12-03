package main

import (
	"testing"
)

func TestIsDigit(t *testing.T) {
	var tests = []struct {
		name  string
		input byte
		digit bool
	}{
		{"9 is digit", '9', true},
		{"0 is digit", '0', true},
		{"5 is digit", '5', true},
		{". is not digit", '.', false},
		{"a is not digit", 'a', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := isDigit(tt.input)
			if a != tt.digit {
				t.Errorf("isDigit(%q) should be %t", tt.input, tt.digit)
			}
		})
	}
}
