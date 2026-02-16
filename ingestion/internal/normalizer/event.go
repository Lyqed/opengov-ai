// Package normalizer converts gateway-specific events into a unified schema.
package normalizer

import "time"

// Event is the normalized event schema for all gateway events.
type Event struct {
	ID          string            `json:"id"`
	Timestamp   time.Time         `json:"timestamp"`
	Source      string            `json:"source"`       // "runlayer", "cloudflare", "mintmcp", "generic"
	AgentID     string            `json:"agent_id"`
	UserID      string            `json:"user_id"`
	Action      Action            `json:"action"`
	MCPServer   string            `json:"mcp_server"`
	Tool        string            `json:"tool"`
	Status      string            `json:"status"`       // "allowed", "blocked", "pending"
	RiskSignals []RiskSignal      `json:"risk_signals"`
	Metadata    map[string]string `json:"metadata"`
	RawPayload  []byte            `json:"raw_payload,omitempty"`
}

// Action represents the type of agent action observed.
type Action string

const (
	ActionToolCall     Action = "tool_call"
	ActionFileRead     Action = "file_read"
	ActionFileWrite    Action = "file_write"
	ActionCLIExec      Action = "cli_exec"
	ActionNetworkReq   Action = "network_request"
	ActionMCPConnect   Action = "mcp_connect"
	ActionMCPDisconnect Action = "mcp_disconnect"
)

// RiskSignal is a risk indicator detected during event processing.
type RiskSignal struct {
	Type     string  `json:"type"`     // "secret_detected", "pii_detected", "anomaly", "policy_violation"
	Severity string  `json:"severity"` // "low", "medium", "high", "critical"
	Detail   string  `json:"detail"`
	Score    float64 `json:"score"`
}

// Normalizer converts raw gateway events into the unified Event schema.
type Normalizer struct{}

func New() *Normalizer { return &Normalizer{} }

// Normalize converts a raw payload from a specific source into a unified Event.
func (n *Normalizer) Normalize(source string, payload []byte) (*Event, error) {
	// TODO: Dispatch to source-specific parser
	// TODO: Extract identity information
	// TODO: Map gateway-specific action types to unified Action enum
	return nil, nil
}
