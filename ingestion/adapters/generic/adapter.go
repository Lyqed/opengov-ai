// Package generic implements a generic webhook/syslog adapter for custom gateways.
package generic

import (
	"context"
	"log/slog"
	"net/http"
)

type Adapter struct {
	listenAddr string
	handler    http.Handler
}

type Config struct {
	ListenAddr string `json:"listen_addr"` // e.g. ":8090"
}

func New(cfg Config) *Adapter {
	a := &Adapter{listenAddr: cfg.ListenAddr}
	mux := http.NewServeMux()
	mux.HandleFunc("/webhook", a.handleWebhook)
	a.handler = mux
	return a
}

func (a *Adapter) Start(ctx context.Context) error {
	slog.Info("starting generic webhook adapter", "addr", a.listenAddr)
	srv := &http.Server{Addr: a.listenAddr, Handler: a.handler}
	go func() { <-ctx.Done(); srv.Close() }()
	return srv.ListenAndServe()
}

func (a *Adapter) handleWebhook(w http.ResponseWriter, r *http.Request) {
	// TODO: Parse incoming webhook payload
	// TODO: Attempt auto-detection of source format
	// TODO: Normalize and forward to event stream
	w.WriteHeader(http.StatusAccepted)
}
