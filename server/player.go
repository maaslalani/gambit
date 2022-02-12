package server

import (
	"fmt"
	"log"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
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
	r    *Room
	s    ssh.Session
	p    *tea.Program
	g    *SharedGame
	t    PlayerType
	k    PublicKey
	once sync.Once
}

// String implements the Stringer interface.
func (p *Player) String() string {
	u := p.s.User()
	return fmt.Sprintf("%s (%s)", u, p.t)
}

// Position returns the player's board FEN position.
func (p *Player) Position() string {
	if p.g != nil && p.g.g != nil {
		return p.g.g.Position()
	}
	return ""
}

// Send sends a message to the bubble tea program.
func (p *Player) Send(m tea.Msg) {
	if p.p != nil {
		p.p.Send(m)
	} else {
		log.Printf("error sending message to player, program is nil")
	}
}

// Write writes data to the ssh session.
func (p *Player) Write(b []byte) (int, error) {
	return p.s.Write(b)
}

// WriteString writes a string to the ssh session.
func (p *Player) WriteString(s string) (int, error) {
	return p.s.Write([]byte(s))
}

// Close closes the the bubble tea program and deletes the player from the room.
func (p *Player) Close() error {
	p.once.Do(func() {
		defer delete(p.r.players, p.k.String())
		if p.p != nil {
			p.p.Kill()
		}
		p.s.Close()
	})
	return nil
}

// StartGame starts the bubble tea program.
func (p *Player) StartGame() {
	_, winch, _ := p.s.Pty()
	errc := make(chan error, 1)
	go func() {
		select {
		case err := <-errc:
			log.Printf("error starting program %s", err)
		case w := <-winch:
			if p.p != nil {
				p.p.Send(tea.WindowSizeMsg{Width: w.Width, Height: w.Height})
			}
		case <-p.s.Context().Done():
			p.Close()
		}
	}()
	defer p.r.SendMsg(NoteMsg(fmt.Sprintf("%s left the room", p)))
	m, err := p.p.StartReturningModel()
	if m != nil {
		p.g = m.(*SharedGame)
	}
	errc <- err
	p.Close()
}
