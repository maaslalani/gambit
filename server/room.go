package server

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/maaslalani/gambit/game"
)

var (
	idleTimeout = time.Minute * 3
)

// Room is a game room with a unique id, password, and a list of players.
type Room struct {
	id          string
	password    string
	players     map[string]*Player
	whiteToMove bool
	sync        chan tea.Msg
	done        chan struct{}
	finish      chan string
}

// String implements the Stringer interface.
func (r *Room) String() string {
	return r.id
}

// NewRoom creates a new room with a unique id and password.
func NewRoom(id, password string, finish chan string) *Room {
	s := make(chan tea.Msg)
	r := &Room{
		id:          id,
		password:    password,
		players:     make(map[string]*Player, 0),
		whiteToMove: true,
		sync:        s,
		done:        make(chan struct{}, 1),
		finish:      finish,
	}
	go func() {
		r.Listen()
	}()
	return r
}

// P1 returns the player with the first turn.
func (r *Room) P1() *Player {
	for _, p := range r.players {
		if p.ptype == whitePlayer {
			return p
		}
	}
	return nil
}

// P2 returns the player with the second turn.
func (r *Room) P2() *Player {
	for _, p := range r.players {
		if p.ptype == blackPlayer {
			return p
		}
	}
	return nil
}

// FindPlayer returns the player for the given public key.
func (r *Room) FindPlayer(pub PublicKey) *Player {
	p, ok := r.players[pub.String()]
	if ok {
		return p
	}
	return nil
}

// Write writes data to all players in the room.
func (r *Room) Write(b []byte) (n int, err error) {
	for _, p := range r.players {
		n, err = p.Write(b)
		if err != nil {
			return
		}
	}
	return
}

// WriteString writes a string to all players in the room.
func (r *Room) WriteString(s string) (int, error) {
	return r.Write([]byte(s))
}

// Close closes the room and deletes the room from the server memory.
func (r *Room) Close() {
	log.Printf("closing room %s", r)

	for _, p := range r.players {
		p.WriteString("Idle timeout.\n")
		p.Close()
	}

	r.done <- struct{}{}
	r.finish <- r.id
	close(r.sync)
	close(r.done)
}

// Listen listens for messages from players in the room and other events.
func (r *Room) Listen() {
	for {
		select {
		case <-r.done:
			return
		case <-time.After(idleTimeout):
			log.Printf("idle timeout for room %s", r)
			r.Close()
		case m := <-r.sync:
			color := whitePlayer
			if !r.whiteToMove {
				color = blackPlayer
			}
			switch msg := m.(type) {
			case NoteMsg:
				r.SendMsg(msg)
			case game.MoveMsg:
				note := fmt.Sprintf("%s moved %s to %s", color, msg.From, msg.To)
				r.whiteToMove = !r.whiteToMove
				r.SendMsg(m)
				r.SendMsg(NoteMsg(note))
			}
		}
	}
}

// SendMsg sends a bubble tea message to all players in the room.
func (r *Room) SendMsg(m tea.Msg) {
	go func() {
		for _, p := range r.players {
			p.Send(m)
		}
	}()
}

// Position returns the FEN position of the game in the room.
func (r *Room) Position() string {
	p1 := r.P1()
	p2 := r.P2()
	switch {
	case !r.whiteToMove && p1 != nil:
		fallthrough
	case r.whiteToMove && p2 == nil && p1 != nil:
		return p1.Position()
	case r.whiteToMove && p2 != nil:
		fallthrough
	case !r.whiteToMove && p1 == nil && p2 != nil:
		return p2.Position()
	default:
		for _, p := range r.players {
			if p.ptype == observerPlayer {
				return p.Position()
			}
		}
		return ""
	}
}

// MakePlayer creates a new player with the given type and session.
func (r *Room) MakePlayer(pt PlayerType, s ssh.Session) *Player {
	pos := r.Position()
	pl := &Player{
		room:    r,
		session: s,
		ptype:   pt,
		key:     PublicKey{key: s.PublicKey()},
	}
	m := NewSharedGame(pl, r.sync, &r.whiteToMove, pt == whitePlayer, pt == observerPlayer, pos)
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
		tea.WithInput(s),
		tea.WithOutput(s),
	)
	pl.program = p
	pl.game = m
	return pl
}

// AddPlayer adds a player to the room.
func (r *Room) AddPlayer(s ssh.Session) (*Player, error) {
	k := s.PublicKey()
	if k == nil {
		return nil, fmt.Errorf("no public key")
	}
	pub := PublicKey{key: k}
	p, ok := r.players[pub.String()]
	if ok {
		return nil, fmt.Errorf("Player %s is already in the room", p)
	}
	p1 := r.P1()
	p2 := r.P2()
	if p1 == nil {
		p = r.MakePlayer(whitePlayer, s)
	} else if p2 == nil {
		p = r.MakePlayer(blackPlayer, s)
	} else {
		p = r.MakePlayer(observerPlayer, s)
	}
	r.players[pub.String()] = p
	r.SendMsg(NoteMsg(fmt.Sprintf("%s joined the room", p)))
	return p, nil
}

// ObserversCount returns the number of observer players in the room.
func (r *Room) ObserversCount() int {
	n := 0
	for _, p := range r.players {
		if p.ptype == observerPlayer {
			n++
		}
	}
	return n
}
