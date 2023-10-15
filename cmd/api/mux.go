package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/k-akari/payment.com/internal/handler"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
	"github.com/k-akari/payment.com/internal/infrastructure/repository"
	"github.com/k-akari/payment.com/internal/usecase"
)

func newMux(ctx context.Context, cfg *config) (http.Handler, func(), error) {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	db, cleanup, err := database.New(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName, cfg.DBPort)
	if err != nil {
		return nil, cleanup, err
	}

	dbc := database.NewClient(db)

	cr := repository.NewCompanyRepository(dbc)
	cu := usecase.NewCompanyUsecase(cr)
	ch := handler.NewCompanyHandler(cu)

	mux.Post("/companies", ch.Create)
	mux.Get("/companies/{companyID}", ch.GetByID)

	return mux, cleanup, nil
}
