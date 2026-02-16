// Package fingerprint identifies the type and version of discovered endpoints.
package fingerprint

import "time"

// AgentType represents a known AI coding agent.
type AgentType string

const (
	AgentCursor    AgentType = "cursor"
	AgentClaudeCode AgentType = "claude-code"
	AgentCopilot   AgentType = "github-copilot"
	AgentOpenCode  AgentType = "opencode"
	AgentWindsurf  AgentType = "windsurf"
	AgentUnknown   AgentType = "unknown"
)

// MCPServerInfo describes a discovered MCP server.
type MCPServerInfo struct {
	Name      string   `json:"name"`
	Transport string   `json:"transport"` // "stdio", "sse", "streamable-http"
	URI       string   `json:"uri"`
	Tools     []string `json:"tools"`
	Resources []string `json:"resources"`
}

// Result holds identification data for a discovered endpoint.
type Result struct {
	AgentType    AgentType       `json:"agent_type"`
	AgentVersion string          `json:"agent_version"`
	MCPServers   []MCPServerInfo `json:"mcp_servers"`
	Timestamp    time.Time       `json:"timestamp"`
}

// Fingerprinter identifies agent types from scan results.
type Fingerprinter struct{}

func New() *Fingerprinter { return &Fingerprinter{} }

// Identify attempts to fingerprint a target.
func (f *Fingerprinter) Identify(host string, port int, meta map[string]string) (*Result, error) {
	// TODO: Check MCP capabilities endpoint
	// TODO: Match process signatures against known agents
	// TODO: Parse config files for agent identification
	return &Result{AgentType: AgentUnknown, Timestamp: time.Now()}, nil
}
