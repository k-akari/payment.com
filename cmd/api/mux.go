package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func newMux(ctx context.Context, cfg *config) (http.Handler, error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	return mux, nil
}
