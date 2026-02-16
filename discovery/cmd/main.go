package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const version = "0.1.0-dev"

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		slog.Info("shutting down discovery scanner")
		cancel()
	}()

	slog.Info("starting OpenGov AI discovery scanner", "version", version)

	// TODO: Initialize scanner, fingerprinter, inventory, risk classifier
	// TODO: Start HTTP API for inventory queries
	// TODO: Run continuous discovery loop

	fmt.Println("discovery: scaffolding complete, implementation pending")
	<-ctx.Done()
}
