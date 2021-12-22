package piece

import "strings"

func (p Piece) ToFen() string {
	if p.Color == NoColor {
		return ""
	}

	if p.Color == Black {
		return strings.ToLower(string(p.Type))
	}

	return string(p.Type)
}

func FromFen(fen string) Piece {
	t := Type(strings.ToUpper(fen))
	var c Color
	if strings.ToUpper(fen) == fen {
		c = White
	} else {
		c = Black
	}
	return Piece{Type: t, Color: c}
}
