package main

import "testing"

func TestCardsStrength(t *testing.T) {
	tests := []struct {
		c        cards
		strength int
	}{
		{cards{{0}}, 11},
		{cards{{1}, {40}}, 21},
		{cards{{2}, {41}, {3}}, 12},
		{cards{{51}, {49}, {50}}, 0},
	}

	for i, tt := range tests {
		if tt.strength != tt.c.strength() {
			t.Errorf("case %d, expected: %d, actual: %d", i, tt.strength, tt.c.strength())
		}
	}
}
