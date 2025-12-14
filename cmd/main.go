package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cleanandclean/internal"
	"cleanandclean/internal/infrastructure"
)

func main() {
	factory := infrastructure.NewDefaultFactory()
	project := internal.NewProject("Clean Project", factory)
	project.Initialize("Clean Application", "1.0.0")

	if err := project.Start(); err != nil {
		fmt.Printf("Failed to start: %v\n", err)
		os.Exit(1)
	}


	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := project.Stop(ctx); err != nil {
		fmt.Printf("Failed to stop: %v\n", err)
	}
}
