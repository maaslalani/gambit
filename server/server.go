package server

import (
	"context"
	"fmt"
	"log"

	gossh "golang.org/x/crypto/ssh"

	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
)

// PublicKey wraps ssh.PublicKey.
type PublicKey struct {
	key ssh.PublicKey
}

// String implements the Stringer interface.
func (k PublicKey) String() string {
	return fmt.Sprintf("%s", gossh.MarshalAuthorizedKey(k.key))
}

// Server is a server that manages chess games.
type Server struct {
	host  string
	port  int
	srv   *ssh.Server
	games map[string]*Room
}

// NewServer creates a new server.
func NewServer(keyPath, host string, port int) (*Server, error) {
	s := &Server{
		host:  host,
		port:  port,
		games: make(map[string]*Room),
	}
	srv, err := wish.NewServer(
		ssh.PasswordAuth(passwordHandler),
		ssh.PublicKeyAuth(publicKeyHandler),
		wish.WithHostKeyPath(keyPath),
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithMiddleware(
			gambitMiddleware(s),
		),
	)
	if err != nil {
		return nil, err
	}
	s.srv = srv
	return s, nil
}

// Start starts the Gambit ssh server.
func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

// Shutdown shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	for _, g := range s.games {
		g.Close()
	}
	return s.srv.Shutdown(ctx)
}

func passwordHandler(ctx ssh.Context, password string) bool {
	return true
}

func publicKeyHandler(ctx ssh.Context, key ssh.PublicKey) bool {
	return true
}

// FindRoom finds a room with the given id.
func (s *Server) FindRoom(id string) *Room {
	g, ok := s.games[id]
	if !ok {
		return nil
	}
	return g
}

// NewRoom creates a new room with the given id and password.
func (s *Server) NewRoom(id, password string) *Room {
	finish := make(chan string, 1)
	go func() {
		id := <-finish
		log.Printf("deleting room %s", id)
		delete(s.games, id)
		close(finish)
	}()
	g := NewRoom(id, password, finish)
	s.games[id] = g
	return g
}
