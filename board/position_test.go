package board

import "testing"

func TestPosition(t *testing.T) {
	tt := []struct {
		s   string
		row int
		col int
	}{
		{"A8", 7, 0},
		{"B7", 6, 1},
		{"C6", 5, 2},
		{"D5", 4, 3},
		{"E4", 3, 4},
		{"F3", 2, 5},
		{"G2", 1, 6},
		{"H1", 0, 7},
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
