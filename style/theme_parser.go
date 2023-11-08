package style

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	White          = "WHITE"
	Black          = "BLACK"
	LightSquare    = "LIGHTSQUARE"
	DarkSquare     = "DARKSQUARE"
	AvailableMoves = "AVAILABLEMOVES"
	SelectedPiece  = "SELECTEDPIECE"
	Check          = "CHECK"
)

// ParseThemeFile parses a file containing the theme information.
// File lines must be in the format `KEY=VALUE`.
func ParseThemeFile(filepath string) (Theme, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return Theme{}, fmt.Errorf("Cannot read file %s: %w", filepath, err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	theme := Theme{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splitLine := strings.Split(line, "=")
		if len(splitLine) != 2 {
			return theme, fmt.Errorf("Invalid line in theme file: %s", line)
		}
		key := splitLine[0]
		color := splitLine[1]
		switch key {
		case White:
			theme.WhitePieceColor = color
		case Black:
			theme.BlackPieceColor = color
		case LightSquare:
			theme.LightSquareColor = color
		case DarkSquare:
			theme.DarkSquareColor = color
		case AvailableMoves:
			theme.AvailableMovesColor = color
		case SelectedPiece:
			theme.SelectedPieceColor = color
		case Check:
			theme.CheckColor = color
		default:
			return Theme{}, fmt.Errorf("Invalid field in theme file: %s", key)
		}
	}
	return theme, nil
}
