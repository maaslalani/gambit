package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/gambit/cmd"
	"github.com/maaslalani/gambit/game"
	"github.com/muesli/coral"
)

var (
	rootCmd = &coral.Command{
		Use:                   "gambit",
		Short:                 "Play chess in your terminal",
		DisableFlagsInUseLine: true,
		RunE: func(cmd *coral.Command, args []string) error {
			if len(args) == 0 {
				startPos, _ := readStdin()

				debug := os.Getenv("DEBUG")
				if debug != "" {
					f, err := tea.LogToFile(debug, "")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()
				}

				p := tea.NewProgram(
					game.NewGameWithPosition(startPos),
					tea.WithAltScreen(),
					tea.WithMouseCellMotion(),
				)

				return p.Start()
			}

			return cmd.Help()
		},
	}
)

func init() {
	rootCmd.AddCommand(
		cmd.ServeCmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func readStdin() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if stat.Mode()&os.ModeNamedPipe == 0 && stat.Size() == 0 {
		return "", errors.New("no starting position provided")
	}

	reader := bufio.NewReader(os.Stdin)
	var b strings.Builder

	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		_, err = b.WriteRune(r)
		if err != nil {
			return "", err
		}
	}

	return b.String(), nil
}
