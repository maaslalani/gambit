package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/maaslalani/gambit/server"
	"github.com/muesli/coral"
)

var (
	host string
	port int
	key  string

	ServeCmd = &coral.Command{
		Use:     "serve",
		Aliases: []string{"server"},
		Short:   "Start a Gambit server",
		Args:    coral.NoArgs,
		RunE: func(cmd *coral.Command, args []string) error {
			k := os.Getenv("GAMBIT_SERVER_KEY_PATH")
			if k != "" {
				key = k
			}
			h := os.Getenv("GAMBIT_SERVER_HOST")
			if h != "" {
				host = h
			}
			p := os.Getenv("GAMBIT_SERVER_PORT")
			if p != "" {
				port, _ = strconv.Atoi(p)
			}
			s, err := server.NewServer(key, host, port)
			if err != nil {
				return err
			}

			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			log.Printf("Starting Gambit server on %s:%d", host, port)
			go func() {
				if err = s.Start(); err != nil {
					log.Fatalln(err)
				}
			}()

			<-done
			log.Print("Stopping Gambit server")
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer func() { cancel() }()
			if err := s.Shutdown(ctx); err != nil {
				log.Fatalln(err)
			}

			return nil
		},
	}
)

func init() {
	ServeCmd.Flags().StringVar(&key, "key", "gambit", "Server private key path")
	ServeCmd.Flags().StringVar(&host, "host", "", "Server host to bind to")
	ServeCmd.Flags().IntVar(&port, "port", 53531, "Server port to bind to")
}
