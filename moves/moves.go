package moves

import (
	"strings"

	dt "github.com/dylhunn/dragontoothmg"
)

// IsLegal determines whether it is legal to move to a destination square given
// all of the legal moves that can be made by a piece.
func IsLegal(legalMoves []dt.Move, destination string) bool {
	for _, move := range legalMoves {
		if strings.HasSuffix(move.String(), destination) {
			return true
		}

		if move.Promote() > 1 && strings.HasSuffix(move.String(), destination+"q") {
			return true
		}

	}
	return false
}

// LegalSelected returns the legal moves for a given piece based on an origin
// square given all the current legal moves for all pieces on the board.
func LegalSelected(moves []dt.Move, selected string) []dt.Move {
	var legalMoves []dt.Move

	// Return an empty slice if there is no square selected
	if selected == "" {
		return legalMoves
	}

	for _, move := range moves {
		if strings.HasPrefix(move.String(), selected) {
			legalMoves = append(legalMoves, move)
		}
	}

	return legalMoves
}
