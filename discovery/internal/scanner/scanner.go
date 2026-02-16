// Package scanner discovers AI agents and MCP servers in the environment.
package scanner

import (
	"context"
	"log/slog"
	"time"
)

// Target represents a discovered endpoint.
type Target struct {
	Host      string            `json:"host"`
	Port      int               `json:"port"`
	Protocol  string            `json:"protocol"` // "mcp-sse", "mcp-stdio", "mcp-http", "http"
	FirstSeen time.Time         `json:"first_seen"`
	LastSeen  time.Time         `json:"last_seen"`
	Metadata  map[string]string `json:"metadata"`
}

// Scanner performs continuous discovery of AI agents and MCP servers.
type Scanner struct {
	interval time.Duration
	targets  []Target
}

// New creates a Scanner with the given interval in seconds.
func New(intervalSeconds int) *Scanner {
	return &Scanner{interval: time.Duration(intervalSeconds) * time.Second}
}

// Run starts the scan loop. Blocks until ctx is cancelled.
func (s *Scanner) Run(ctx context.Context) error {
	slog.Info("starting discovery scan loop", "interval", s.interval)
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	s.scan(ctx)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			s.scan(ctx)
		}
	}
}

func (s *Scanner) scan(ctx context.Context) {
	slog.Info("running discovery scan")
	// TODO: Scan for MCP servers on well-known ports (3000-3100, 8080-8090)
	// TODO: Query gateway APIs for registered servers
	// TODO: Check process tables for known agent binaries (cursor, claude, copilot)
	// TODO: Parse MCP config files (~/.cursor/mcp.json, ~/.claude/settings.json)
	// TODO: Scan for stdio-based MCP servers via process tree analysis
}

// Targets returns all discovered targets.
func (s *Scanner) Targets() []Target { return s.targets }
