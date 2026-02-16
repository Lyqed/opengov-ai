// Package mintmcp implements the ingestion adapter for MintMCP gateway.
package mintmcp

import (
	"context"
	"log/slog"
)

type Adapter struct{ apiEndpoint, apiKey string }
type Config struct {
	APIEndpoint string `json:"api_endpoint"`
	APIKey      string `json:"api_key"`
}

func New(cfg Config) *Adapter {
	return &Adapter{apiEndpoint: cfg.APIEndpoint, apiKey: cfg.APIKey}
}

func (a *Adapter) Start(ctx context.Context) error {
	slog.Info("starting MintMCP adapter")
	// TODO: Connect to MintMCP Agent Monitor API
	// TODO: Poll or webhook for events
	return nil
}
