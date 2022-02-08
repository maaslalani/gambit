package pieces

// Display maps pieces from their FEN representations to their ASCII
// representations for a more human readable experience.
var Display = map[string]string{
	"":  " ",
	"B": "♝",
	"K": "♚",
	"N": "♞",
	"P": "♟",
	"Q": "♛",
	"R": "♜",
	"b": "♗",
	"k": "♔",
	"n": "♘",
	"p": "♙",
	"q": "♕",
	"r": "♖",
}
