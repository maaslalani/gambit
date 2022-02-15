package server

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/game"
	"github.com/maaslalani/gambit/style"
)

// NoteMsg is a message that is sent to the client when a message is added to
// the game.
type NoteMsg string

// SharedGame is a game that is shared between players. It wraps gambit bubble
// tea model and synchronizes messages among players and server.
type SharedGame struct {
	player      *Player
	game        *game.Game
	note        string
	turn        bool
	observer    bool
	whiteToMove *bool
	sync        chan tea.Msg
}

// NewSharedGame creates a new shared game for a player.
func NewSharedGame(p *Player, sync chan tea.Msg, roomTurn *bool, turn, observer bool, pos string) *SharedGame {
	g := game.NewGameWithPosition(pos)
	g.SetFlipped(!turn)
	r := &SharedGame{
		player:      p,
		game:        g,
		turn:        turn,
		observer:    observer,
		whiteToMove: roomTurn,
		sync:        sync,
	}
	return r
}

// Init implements bubble tea model.
func (r *SharedGame) Init() tea.Cmd {
	return nil
}

// SendMsg sends a message to the room.
func (r *SharedGame) SendMsg(msg tea.Msg) {
	go func() {
		r.sync <- msg
	}()
}

// Update implements bubble tea model.
func (r *SharedGame) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case game.NotifyMsg:
		if !r.observer && r.turn != msg.Turn {
			r.SendMsg(game.MoveMsg{From: msg.From, To: msg.To})
			if msg.Checkmate {
				r.SendMsg(NoteMsg(fmt.Sprintf("%s wins!", r.player)))
			}
		}
	case game.MoveMsg:
		g, cmd := r.game.Update(msg)
		r.game = g.(*game.Game)
		cmds = append(cmds, cmd)
	case NoteMsg:
		r.note = string(msg)
		return r, nil
	case tea.MouseMsg:
		if !r.observer && r.turn == *r.whiteToMove {
			if msg.Type != tea.MouseLeft {
				return r, nil
			}
			g, cmd := r.game.Update(msg)
			cmds = append(cmds, cmd)
			r.game = g.(*game.Game)
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			g, cmd := r.game.Update(msg)
			r.game = g.(*game.Game)
			cmds = append(cmds, cmd)
			cmds = append(cmds, tea.Quit)
		case "ctrl+f":
			g, cmd := r.game.Update(msg)
			cmds = append(cmds, cmd)
			r.game = g.(*game.Game)
		default:
			if !r.observer && r.turn == *r.whiteToMove {
				g, cmd := r.game.Update(msg)
				cmds = append(cmds, cmd)
				r.game = g.(*game.Game)
			}
		}
	default:
		if !r.observer && r.turn == *r.whiteToMove {
			g, cmd := r.game.Update(msg)
			cmds = append(cmds, cmd)
			r.game = g.(*game.Game)
		}
	}
	return r, tea.Batch(cmds...)
}

// View implements bubble tea model.
func (r *SharedGame) View() string {
	s := strings.Builder{}

	turn := "Black's move"
	if *r.whiteToMove {
		turn = "White's move"
	}

	s.WriteString(r.game.View())

	s.WriteRune('\n')
	s.WriteString(fmt.Sprintf(" %s %s", style.Title("Gambit"), turn))
	s.WriteRune('\n')
	s.WriteString(style.Faint(fmt.Sprintf(" Room %s as %s playing %s", r.player.room.id, r.player.session.User(), r.player.ptype)))
	s.WriteRune('\n')

	return s.String()
}
