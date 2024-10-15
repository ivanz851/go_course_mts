package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task2/internal/server"
	"time"
)

func main() {
	_server := server.StartServer()

	go func() {
		log.Println("Starting server on :8080...")
		if err := _server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	performGracefulShutdown(_server)
}

func performGracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}

	log.Println("Server gracefully stopped")
}
