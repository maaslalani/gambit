package main

import (
	"testing"
)

func TestPosition(t *testing.T) {
	tt := []struct {
		position position
		want     string
	}{
		{position{1, 1}, "A1"},
		{position{2, 2}, "B2"},
		{position{3, 3}, "C3"},
		{position{4, 4}, "D4"},
		{position{5, 5}, "E5"},
		{position{6, 6}, "F6"},
		{position{7, 7}, "G7"},
		{position{8, 8}, "H8"},
	}

	for _, tc := range tt {
		if tc.position.String() != tc.want {
			t.Fatalf("want: %s, got: %s", tc.want, tc.position)
		}
	}
}
