// Package actions provides atomic response actions for playbooks.
package actions

import "context"

// Action is an executable response action.
type Action interface {
	Execute(ctx context.Context, params map[string]string) error
	Name() string
}

// RevokeAccess disconnects an agent from a specific MCP server.
type RevokeAccess struct{}

func (r *RevokeAccess) Name() string { return "revoke_access" }
func (r *RevokeAccess) Execute(ctx context.Context, params map[string]string) error {
	// TODO: Call gateway API to revoke MCP server access
	// TODO: Invalidate agent tokens
	// TODO: Log action to audit trail
	return nil
}

// Quarantine isolates an agent from all MCP connections.
type Quarantine struct{}

func (q *Quarantine) Name() string { return "quarantine" }
func (q *Quarantine) Execute(ctx context.Context, params map[string]string) error {
	// TODO: Revoke all MCP connections for the agent
	// TODO: Notify agent owner
	// TODO: Create incident record
	return nil
}

// GenerateReport creates a compliance/incident report.
type GenerateReport struct{}

func (g *GenerateReport) Name() string { return "generate_report" }
func (g *GenerateReport) Execute(ctx context.Context, params map[string]string) error {
	// TODO: Gather incident timeline
	// TODO: Map to compliance framework
	// TODO: Generate PDF/HTML report
	// TODO: Store in audit archive
	return nil
}
