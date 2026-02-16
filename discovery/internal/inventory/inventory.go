// Package inventory maintains a deduplicated registry of discovered assets.
package inventory

import (
	"sync"
	"time"
)

// Asset represents a discovered AI agent or MCP server.
type Asset struct {
	ID          string            `json:"id"`
	Type        string            `json:"type"` // "agent", "mcp-server"
	Name        string            `json:"name"`
	Host        string            `json:"host"`
	Port        int               `json:"port"`
	Version     string            `json:"version"`
	Owner       string            `json:"owner"`
	Department  string            `json:"department"`
	RiskScore   float64           `json:"risk_score"`
	FirstSeen   time.Time         `json:"first_seen"`
	LastSeen    time.Time         `json:"last_seen"`
	Tags        map[string]string `json:"tags"`
	Connections []string          `json:"connections"` // IDs of connected assets
}

// Inventory is a thread-safe asset registry.
type Inventory struct {
	mu     sync.RWMutex
	assets map[string]*Asset
}

func New() *Inventory {
	return &Inventory{assets: make(map[string]*Asset)}
}

func (inv *Inventory) Upsert(a *Asset) {
	inv.mu.Lock()
	defer inv.mu.Unlock()
	if existing, ok := inv.assets[a.ID]; ok {
		existing.LastSeen = time.Now()
		existing.Version = a.Version
		existing.RiskScore = a.RiskScore
		return
	}
	a.FirstSeen = time.Now()
	a.LastSeen = a.FirstSeen
	inv.assets[a.ID] = a
}

func (inv *Inventory) Get(id string) (*Asset, bool) {
	inv.mu.RLock()
	defer inv.mu.RUnlock()
	a, ok := inv.assets[id]
	return a, ok
}

func (inv *Inventory) List() []*Asset {
	inv.mu.RLock()
	defer inv.mu.RUnlock()
	out := make([]*Asset, 0, len(inv.assets))
	for _, a := range inv.assets {
		out = append(out, a)
	}
	return out
}

func (inv *Inventory) Count() int {
	inv.mu.RLock()
	defer inv.mu.RUnlock()
	return len(inv.assets)
}
