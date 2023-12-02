package main

import "testing"

func TestFindFirstLastDigit(t *testing.T) {
	var tests = []struct {
		input string
		a     int
		b     int
	}{
		{"abc4defg5hij6kl", 4, 6},
		{"treb7uchet", 7, 7},
		{"trebuchet", -1, -1},
		{"123", 1, 3},
		{"53", 5, 3},
		{"abc82", 8, 2},
		{"22611zfive", 2, 1},
		{"9eight8one9five", 9, 9},
		{"sevenone39sixsix41", 3, 1},
		{"7sqthfchpjklpn", 7, 7},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			a, b := findFirstLastDigit(tt.input)
			if !(a == tt.a && b == tt.b) {
				t.Errorf("findFirstLastDigit(%q) = %d, %d, want %d, %d", tt.input, a, b, tt.a, tt.b)
			}
		})
	}
}

func TestFindFirstLastNumber(t *testing.T) {
	var tests = []struct {
		input string
		a     int
		b     int
	}{
		{"sevenone39sixsix41", 7, 1},
		{"7sqthfchpjklpn", 7, 7},
		{"eighthree", 8, 3},
		{"four", 4, 4},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			a, b := findFirstLastNumber(tt.input)
			if !(a == tt.a && b == tt.b) {
				t.Errorf("findFirstLastNumber(%q) = %d, %d, want %d, %d", tt.input, a, b, tt.a, tt.b)
			}
		})
	}
}
