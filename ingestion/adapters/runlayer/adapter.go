// Package runlayer implements the ingestion adapter for RunLayer MCP gateway.
package runlayer

import (
	"context"
	"log/slog"
)

// Adapter connects to RunLayer and ingests events.
type Adapter struct {
	webhookURL string
	apiKey     string
}

// Config holds RunLayer adapter configuration.
type Config struct {
	WebhookURL string `json:"webhook_url"`
	APIKey     string `json:"api_key"`
}

// New creates a new RunLayer adapter.
func New(cfg Config) *Adapter {
	return &Adapter{webhookURL: cfg.WebhookURL, apiKey: cfg.APIKey}
}

// Start begins listening for RunLayer events.
func (a *Adapter) Start(ctx context.Context) error {
	slog.Info("starting RunLayer adapter", "webhook", a.webhookURL)
	// TODO: Register webhook with RunLayer API
	// TODO: Handle incoming events and normalize
	// TODO: Forward to event stream
	return nil
}
