package board

import "testing"

func TestPosition(t *testing.T) {
	tt := []struct {
		row int
		col int
		s   string
	}{
		{0, 7, "H1"},
		{1, 6, "G2"},
		{2, 5, "F3"},
		{3, 4, "E4"},
		{4, 3, "D5"},
		{5, 2, "C6"},
		{6, 1, "B7"},
		{7, 0, "A8"},
	}

	for i, tc := range tt {
		p := Position{tc.row, tc.col}
		if p.String() != tc.s {
			t.Errorf("Test %d: expected %s, got %s", i, tc.s, p.String())
		}
	}

	for i, tc := range tt {
		p := ToPosition(tc.s)
		if p.Col != tc.col || p.Row != tc.row {
			t.Errorf("Test %d: expected %s, got %s", i, tc.s, p.String())
		}
	}
}
