package server

import (
	"fmt"
	"log"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
)

// PlayerType is the type of a player in a chess game.
type PlayerType int

const (
	whitePlayer PlayerType = iota
	blackPlayer
	observerPlayer
)

// String implements the Stringer interface.
func (pt PlayerType) String() string {
	switch pt {
	case whitePlayer:
		return "White"
	case blackPlayer:
		return "Black"
	case observerPlayer:
		return "Observer"
	default:
		return ""
	}
}

// Player is a player in a chess game who belongs to a room, has a ssh session
// and a bubble tea program.
type Player struct {
	room    *Room
	session ssh.Session
	program *tea.Program
	game    *SharedGame
	ptype   PlayerType
	key     PublicKey
	once    sync.Once
}

// String implements the Stringer interface.
func (p *Player) String() string {
	u := p.session.User()
	return fmt.Sprintf("%s (%s)", u, p.ptype)
}

// Position returns the player's board FEN position.
func (p *Player) Position() string {
	if p.game != nil && p.game.game != nil {
		return p.game.game.Position()
	}
	return ""
}

// Send sends a message to the bubble tea program.
func (p *Player) Send(m tea.Msg) {
	if p.program != nil {
		p.program.Send(m)
	} else {
		log.Printf("error sending message to player, program is nil")
	}
}

// Write writes data to the ssh session.
func (p *Player) Write(b []byte) (int, error) {
	return p.session.Write(b)
}

// WriteString writes a string to the ssh session.
func (p *Player) WriteString(s string) (int, error) {
	return p.session.Write([]byte(s))
}

// Close closes the the bubble tea program and deletes the player from the room.
func (p *Player) Close() error {
	p.once.Do(func() {
		defer delete(p.room.players, p.key.String())
		if p.program != nil {
			p.program.Kill()
		}
		p.session.Close()
	})
	return nil
}

// StartGame starts the bubble tea program.
func (p *Player) StartGame() {
	_, wchan, _ := p.session.Pty()
	errc := make(chan error, 1)
	go func() {
		select {
		case err := <-errc:
			log.Printf("error starting program %s", err)
		case w := <-wchan:
			if p.program != nil {
				p.program.Send(tea.WindowSizeMsg{Width: w.Width, Height: w.Height})
			}
		case <-p.session.Context().Done():
			p.Close()
		}
	}()
	defer p.room.SendMsg(NoteMsg(fmt.Sprintf("%s left the room", p)))
	m, err := p.program.StartReturningModel()
	if m != nil {
		p.game = m.(*SharedGame)
	}
	errc <- err
	p.Close()
}
