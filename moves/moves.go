package moves

import (
	"strings"

	dt "github.com/dylhunn/dragontoothmg"
)

// IsLegal determines whether it is legal to move the the destination
// square given a piece's legal moves
func IsLegal(legalMoves []dt.Move, destination string) bool {
	for _, move := range legalMoves {
		if strings.HasSuffix(move.String(), destination) {
			return true
		}
	}
	return false
}

// LegalSelected returns the legal moves for a given piece this is usually
// for the selected piece so that we know to which we can move. If there is no
// selected piece we return an empty array of moves.
func LegalSelected(moves []dt.Move, selected string) []dt.Move {
	var legalMoves []dt.Move

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
