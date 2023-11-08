package style

import . "github.com/charmbracelet/lipgloss"

// Theme is a struct that contains the colors of the theme
type Theme struct {
	WhitePieceColor     string
	BlackPieceColor     string
	LightSquareColor    string
	DarkSquareColor     string
	AvailableMovesColor string
	CheckColor          string
	SelectedPieceColor  string
}

// IsValid returns a boolean indicating if all the colors of the theme
// are present
func (t *Theme) IsValid() bool {
	return t.WhitePieceColor != "" &&
		t.BlackPieceColor != "" &&
		t.LightSquareColor != "" &&
		t.DarkSquareColor != "" &&
		t.AvailableMovesColor != "" &&
		t.CheckColor != "" &&
		t.SelectedPieceColor != ""
}

// Bg sets the background of a string based on the color of the square
func (t *Theme) Bg(content string, isLightSquare bool) string {
	if isLightSquare {
		return NewStyle().Background(Color(t.LightSquareColor)).Render(content)
	}
	return NewStyle().Background(Color(t.DarkSquareColor)).Render(content)
}

// Fg sets the foreground of a string based on the color of the square
func (t *Theme) Fg(content string, isLightSquare bool) string {
	if isLightSquare {
		return NewStyle().Foreground(Color(t.LightSquareColor)).Render(content)
	}
	return NewStyle().Foreground(Color(t.DarkSquareColor)).Render(content)
}

func (t *Theme) borderLight(border string) string {
	return NewStyle().
		Background(Color(t.DarkSquareColor)).
		Foreground(Color(t.LightSquareColor)).
		Render(border)
}

func (t *Theme) borderDark(border string) string {
	return NewStyle().
		Background(Color(t.LightSquareColor)).
		Foreground(Color(t.DarkSquareColor)).
		Render(border)
}

// Border is used to draw the borders between ranks: it applies both
// a foreground and a background
func (t *Theme) Border(border string, isLightSquare bool) string {
	if isLightSquare {
		return NewStyle().
			Background(Color(t.DarkSquareColor)).
			Foreground(Color(t.LightSquareColor)).
			Render(border)
	}
	return NewStyle().
		Background(Color(t.LightSquareColor)).
		Foreground(Color(t.DarkSquareColor)).
		Render(border)
}

// SelectedSquare colors the content with the [SelectedPieceColor] color
// and applies the background based on the square color
func (t *Theme) SelectedSquare(content string, isLightSquare bool) string {
	if isLightSquare {
		return NewStyle().
			Foreground(Color(t.SelectedPieceColor)).
			Background(Color(t.LightSquareColor)).
			Render(content)
	}
	return NewStyle().
		Foreground(Color(t.SelectedPieceColor)).
		Background(Color(t.DarkSquareColor)).
		Render(content)
}

// CheckSquare colors the content with the [CheckColor] color
// and applies the background based on the square color
func (t *Theme) CheckSquare(content string, isLightSquare bool) string {
	if isLightSquare {
		return NewStyle().
			Foreground(Color(t.CheckColor)).
			Background(Color(t.LightSquareColor)).
			Render(content)
	}
	return NewStyle().
		Foreground(Color(t.CheckColor)).
		Background(Color(t.DarkSquareColor)).
		Render(content)
}

// AvailableMove colors the content with the [AvailableMove] color
// and applies the background based on the square color
func (t *Theme) AvailableMove(content string, isLightSquare bool) string {
	if isLightSquare {
		return NewStyle().
			Foreground(Color(t.AvailableMovesColor)).
			Background(Color(t.LightSquareColor)).
			Render(content)
	}
	return NewStyle().
		Foreground(Color(t.AvailableMovesColor)).
		Background(Color(t.DarkSquareColor)).
		Render(content)
}

// Piece colors the piece with white or black color
func (t *Theme) Piece(piece string, isWhite bool, isLightSquare bool) string {
	style := NewStyle()
	if isLightSquare {
		style = style.Background(Color(t.LightSquareColor))
	} else {
		style = style.Background(Color(t.DarkSquareColor))
	}
	if isWhite {
		style = style.Foreground(Color(t.WhitePieceColor))
	} else {
		style = style.Foreground(Color(t.BlackPieceColor))
	}
	return style.Render(piece)
}
