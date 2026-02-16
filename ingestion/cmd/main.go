package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() { <-sigCh; cancel() }()

	slog.Info("starting OpenGov AI ingestion service", "version", "0.1.0-dev")
	// TODO: Start HTTP server for webhook receivers
	// TODO: Initialize gateway adapters (RunLayer, Cloudflare, MintMCP)
	// TODO: Connect to NATS JetStream for event streaming
	fmt.Println("ingestion: scaffolding complete, implementation pending")
	<-ctx.Done()
}
