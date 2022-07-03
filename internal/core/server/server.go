package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Server is a config of the server
type Server struct {
	Bind string
}

// New returns new server
func New(addr string) *Server {
	return &Server{
		Bind: addr,
	}
}

// ListenAndServe listens for incoming connections
func (s *Server) ListenAndServe(ctx context.Context, router http.Handler) {
	srv := http.Server{
		Addr:    s.Bind,
		Handler: router,
	}

	idleConnsClosed := make(chan bool)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		s := <-sig
		log.Printf("Received signal %s\n", s)

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
