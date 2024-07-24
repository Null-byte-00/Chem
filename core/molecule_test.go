package core

import (
	"testing"
	_ "github.com/Null-byte-00/Chem"
)

func TestElementToNumber(t *testing.T) {
	tests := []struct {
		element string
		want int
	}{
		{"H", 1},
		{"C", 6},
		{"N", 7},
		{"O", 8},
		{"F", 9},
		{"P", 15},
		{"S", 16},
		{"Cl", 17},
		{"Br", 35},
		{"I", 53},
		{"X", 0},
	}
	for _, tt := range tests {
		t.Run(tt.element, func(t *testing.T) {
			if got := ElementToNumber(tt.element); got != tt.want {
				t.Errorf("ElementToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}