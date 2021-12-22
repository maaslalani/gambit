package board

import (
	"fmt"
	"strings"
)

var (
	border = []string{
		"┌", "┬", "┐",
		"├", "┼", "┤",
		"└", "┴", "┘",
	}
)

const (
	vertical   = "│"
	horizontal = "─"
	marginLeft = "    "
)

var (
	files = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	ranks = []int{7, 6, 5, 4, 3, 2, 1, 0}
)

func (b Board) String() string {
	var s string
	if b.reversed {
		ranks = []int{0, 1, 2, 3, 4, 5, 6, 7}
	}
	for r, row := range ranks {
		if r == 0 {
			s += buildRow(border[0], border[1], border[2])
			s += "\n"
		}
		for c, cell := range b.grid[row] {
			if c == 0 {
				s += fmt.Sprintf(" %d  ", row+1)
			}
			s += fmt.Sprintf("%s %s ", vertical, cell)
			if c == len(ranks)-1 {
				s += vertical
			}
		}
		if r < len(b.grid)-1 {
			s += buildRow(border[3], border[4], border[5])
		} else {
			s += buildRow(border[6], border[7], border[8])
			s += "\n  " + marginLeft
			s += strings.Join(files, "   ")
		}
		s += "\n"
	}

	return s
}

func buildRow(left, middle, right string) string {
	var row []string
	row = append(row, left)
	for i := 0; i < 7; i++ {
		row = append(row, middle)
	}
	row = append(row, right)
	return fmt.Sprintf("\n%s%s", marginLeft, strings.Join(row, horizontal+horizontal+horizontal))
}
