// Package notifications routes alerts to the appropriate channels.
package notifications

import "context"

// Severity determines routing priority.
type Severity string

const (
	SeverityCritical Severity = "critical"
	SeverityHigh     Severity = "high"
	SeverityMedium   Severity = "medium"
	SeverityLow      Severity = "low"
)

// Alert is a notification to be routed.
type Alert struct {
	Title    string
	Message  string
	Severity Severity
	AssetID  string
	Channel  string // "slack", "pagerduty", "email"
}

// Router sends alerts to configured channels.
type Router struct{}

func NewRouter() *Router { return &Router{} }

func (r *Router) Route(ctx context.Context, alert Alert) error {
	// TODO: Look up routing rules for severity + asset
	// TODO: Send to configured channels
	// TODO: Track delivery status
	return nil
}
