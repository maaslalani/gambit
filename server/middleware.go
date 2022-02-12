package server

import (
	"fmt"
	"log"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/wish"
	"github.com/gliderlabs/ssh"
	"github.com/muesli/termenv"
)

// gambitMiddleware is a middleware that handles the Gambit ssh server. It
// creates rooms and assigns players to them.
func gambitMiddleware(srv *Server) wish.Middleware {
	return func(sh ssh.Handler) ssh.Handler {
		lipgloss.SetColorProfile(termenv.ANSI256)
		return func(s ssh.Session) {
			_, _, active := s.Pty()
			cmds := s.Command()
			if len(cmds) < 1 || !active {
				s.Write([]byte(help("No TTY")))
				s.Exit(1)
				return
			}
			password := ""
			id := cmds[0]
			if len(cmds) > 1 {
				password = cmds[1]
			}
			g := srv.FindRoom(id)
			if g == nil {
				log.Printf("room %s is created with password %q", id, password)
				g = srv.NewRoom(id, password)
			}
			if g.password == password {
				p, err := g.AddPlayer(s)
				if err != nil {
					s.Write([]byte(fmt.Sprintf("%s\n", err)))
					s.Exit(1)
					return
				}
				log.Printf("%s joined room %s [%s]", s.User(), id, s.RemoteAddr())
				p.StartGame()
				log.Printf("%s left room %s [%s]", s.User(), id, s.RemoteAddr())
			} else {
				s.Write([]byte(help("wrong password")))
				s.Exit(1)
				return
			}
			sh(s)
		}
	}
}

func help(s string) string {
	h := `Play chess in your terminal

Usage: ssh [<name>@]<host> -p <port> -t <id> [<password>]

`
	if s != "" {
		h += fmt.Sprintf("%s\n", s)
	}
	return h
}
