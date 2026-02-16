// Package stream publishes normalized events to NATS JetStream.
package stream

import "log/slog"

// Publisher sends normalized events to the event bus.
type Publisher struct {
	natsURL string
	subject string
}

type Config struct {
	NatsURL string `json:"nats_url"`
	Subject string `json:"subject"` // e.g. "opengov.events.>"
}

func New(cfg Config) *Publisher {
	return &Publisher{natsURL: cfg.NatsURL, subject: cfg.Subject}
}

func (p *Publisher) Connect() error {
	slog.Info("connecting to NATS JetStream", "url", p.natsURL)
	// TODO: Connect to NATS
	// TODO: Ensure JetStream stream exists
	return nil
}

func (p *Publisher) Publish(eventJSON []byte) error {
	// TODO: Publish to NATS JetStream subject
	return nil
}

func (p *Publisher) Close() error {
	// TODO: Drain and close NATS connection
	return nil
}
