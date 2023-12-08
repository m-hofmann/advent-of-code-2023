package main

import (
	"testing"
)

func TestHandTypePart1(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"Five of a kind AAAAA", "AAAAA", FIVE_OF_A_KIND},
		{"Four of a kind AA8AA", "AA8AA", FOUR_OF_A_KIND},
		{"Full house 23332", "23332", FULL_HOUSE},
		{"Three of a kind TTT98", "TTT98", THREE_OF_A_KIND},
		{"Two pair 23432", "23432", TWO_PAIR},
		{"One pair A23A4", "A23A4", ONE_PAIR},
		{"High card 23456", "23456", HIGH_CARD},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handType, err := handTypePart1(hand{cards: []byte(tt.input), bid: 1})
			if err != nil || handType != tt.want {
				t.Errorf("Hand type should be %d, is %d", handType, tt.want)
			}
		})
	}
}

func TestHandComparePart1(t *testing.T) {
	var tests = []struct {
		name    string
		left    string
		right   string
		smaller bool
	}{
		{"33332 is stronger than 2AAAA", "2AAAA", "33332", true},
		{"77888 is stronger than 77788", "77788", "77888", true},
		{"77888 is stronger than 23456", "23456", "A23A4", true},
		{"KK677 is stronger than KTJJT", "KTJJT", "KK677", true},
		{"KTJJT is not stronger than KK677", "KK677", "KTJJT", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.smaller != compareHandsPart1(hand{cards: []byte(tt.left)}, hand{cards: []byte(tt.right)}) {
				t.Errorf("Hand %s should be smaller than %s", tt.left, tt.right)
			}
		})
	}
}

func TestHandComparePart2(t *testing.T) {
	var tests = []struct {
		name    string
		left    string
		right   string
		smaller bool
	}{
		{"QQQJA is stronger than T55J5", "T55J5", "QQQJA", true},
		{"KTJJT is stronger than QQQJA", "QQQJA", "KTJJT", true},
		{"QQQQ2 is stronger than JKKK2", "JKKK2", "QQQQ2", true},
		{"2JJJJ is stronger than 2AAAA", "2AAAA", "2JJJJ", true},
		{"22222 is stronger than JJJJJ", "JJJJJ", "22222", true},
		{"8T342 is stronger than 782TQ", "782TQ", "8T342", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.smaller != compareHandsPart2(hand{cards: []byte(tt.left)}, hand{cards: []byte(tt.right)}) {
				t.Errorf("Hand %s should be smaller than %s", tt.left, tt.right)
			}
		})
	}
}

func TestHandTypePart2(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"Five of a kind AAAAA", "AAAAA", FIVE_OF_A_KIND},
		{"Four of a kind AA8AA", "AA8AA", FOUR_OF_A_KIND},
		{"Full house 23332", "23332", FULL_HOUSE},
		{"Three of a kind TTT98", "TTT98", THREE_OF_A_KIND},
		{"Four of a kind T55J5", "T55J5", FOUR_OF_A_KIND},
		{"Four of a kind KTJJT", "KTJJT", FOUR_OF_A_KIND},
		{"Four of a kind QQQJA", "QQQJA", FOUR_OF_A_KIND},
		{"Four of a kind KTJJT", "KTJJT", FOUR_OF_A_KIND},
		{"Five of a kind JJJJJ", "JJJJJ", FIVE_OF_A_KIND},
		{"Five of a kind KJKKJ", "KJKKJ", FIVE_OF_A_KIND},
		{"Five of a kind J33JJ", "J33JJ", FIVE_OF_A_KIND},
		{"Five of a kind 12QKA", "12QKA", HIGH_CARD},
		{"Full house 1122J", "1122J", FULL_HOUSE},
		{"One pair 123JA", "123JA", ONE_PAIR},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handType, err := handTypePart2(hand{cards: []byte(tt.input), bid: 1})
			if err != nil || handType != tt.want {
				t.Errorf("Hand type should be %d, is %d", handType, tt.want)
			}
		})
	}
}
