package main

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/k-akari/payment.com/internal/handler"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
	"github.com/k-akari/payment.com/internal/infrastructure/repository"
	"github.com/k-akari/payment.com/internal/usecase"
)

func newMux(db *sqlx.DB) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	dbc := database.NewClient(db)
	cor := repository.NewCompanyRepository(dbc)
	clr := repository.NewClientRepository(dbc)
	ir := repository.NewInvoiceRepository(dbc)
	cou := usecase.NewCompanyUsecase(cor)
	clu := usecase.NewClientUsecase(clr)
	iu := usecase.NewInvoiceUsecase(ir)
	coh := handler.NewCompanyHandler(cou)
	clh := handler.NewClientHandler(clu)
	ih := handler.NewInvoiceHandler(iu)

	mux.Route("/companies", func(mux chi.Router) {
		mux.Post("/", coh.Create)
		mux.Route("/{companyID}", func(mux chi.Router) {
			mux.Use(companyCtx)
			mux.Get("/", coh.GetByID)
			mux.Get("/invoices", ih.ListByPaymentDueDateBetween)
			mux.Route("/clients", func(mux chi.Router) {
				mux.Post("/", clh.Create)
				mux.Route("/{clientID}", func(mux chi.Router) {
					mux.Use(clientCtx)
					mux.Get("/", clh.GetByID)
					mux.Route("/invoices", func(mux chi.Router) {
						mux.Post("/", ih.Create)
					})
				})
			})
		})
	})

	return mux, nil
}

func companyCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), handler.CompanyID, chi.URLParam(r, "companyID"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func clientCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), handler.ClientID, chi.URLParam(r, "clientID"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
