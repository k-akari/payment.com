package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
)

func newMux(ctx context.Context, cfg *config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	_, cleanup, err := database.New(ctx, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName, cfg.DBPort)
	if err != nil {
		return nil, cleanup, err
	}

	return mux, cleanup, nil
}
