package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cleanandclean/cmd/boot"
)

func main() {
	project := boot.NewProject()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	serverErr := make(chan error, 1)

	go func() {
		log.Println("server starting on :8080")
		serverErr <- project.Run(":8080")
	}()

	select {
	case err := <-serverErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	case sig := <-shutdown:
		log.Printf("shutdown signal received: %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := project.Shutdown(ctx); err != nil {
			log.Printf("graceful shutdown failed: %v", err)
			if err := project.Close(); err != nil {
				log.Fatalf("forced shutdown failed: %v", err)
			}
		}

		log.Println("server stopped gracefully")
	}
}
