package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetChecksum(t *testing.T) {
	var tests = []struct {
		input string
		want  []int
	}{
		{".###....##.#", []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s is %v", tt.input, tt.want), func(t *testing.T) {
			checksum := getChecksum([]byte(tt.input))
			if !reflect.DeepEqual(checksum, tt.want) {
				t.Errorf("%s should have checksum %v but has %v", tt.input, tt.want, checksum)
			}
		})
	}
}
