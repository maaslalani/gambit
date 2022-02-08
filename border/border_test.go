package border_test

import (
	"testing"

	"github.com/maaslalani/gambit/border"
)

func TestTopBorder(t *testing.T) {
	tests := []struct {
		borderFunc func() string
		want       string
	}{

		{
			border.Top,
			"   ┌───┬───┬───┬───┬───┬───┬───┬───┐\n",
		},
		{
			border.Middle,
			"   ├───┼───┼───┼───┼───┼───┼───┼───┤\n",
		},
		{
			border.Bottom,
			"   └───┴───┴───┴───┴───┴───┴───┴───┘\n",
		},
	}

	for _, test := range tests {
		got := test.borderFunc()
		if got != test.want {
			t.Errorf("want %s, got %s", test.want, got)
		}
	}
}
