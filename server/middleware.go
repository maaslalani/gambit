package server

import (
	"log"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
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

			room := srv.FindRoom(id)
			if room == nil {
				log.Printf("room %s is created with password %q", id, password)
				room = srv.NewRoom(id, password)
			}

			if room.password != password {
				s.Write([]byte(help("Incorrect password")))
				s.Exit(1)
				return
			}

			p, err := room.AddPlayer(s)
			if err != nil {
				s.Write([]byte(err.Error() + "\n"))
				s.Exit(1)
				return
			}

			log.Printf("%s joined room %s [%s]", s.User(), id, s.RemoteAddr())
			p.StartGame()
			log.Printf("%s left room %s [%s]", s.User(), id, s.RemoteAddr())

			sh(s)
		}
	}
}

func help(s string) string {
	return strings.Join([]string{
		"Gambit: Play chess in your terminal",
		"Usage: ssh [<name>@]<host> -p <port> -t <room> [<password>]",
		s,
		"\n",
	}, "\n")
}
