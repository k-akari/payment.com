package router

import (
	"net/http"
	"time"

	chimid "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/k-akari/payment.com/internal/handler"
	"github.com/k-akari/payment.com/internal/middleware"
)

func NewMux(coh *handler.CompanyHandler, clh *handler.ClientHandler, ih *handler.InvoiceHandler) http.Handler {
	mux := chi.NewRouter()

	mux.Use(chimid.RequestID)
	mux.Use(chimid.RealIP)
	mux.Use(chimid.Logger)
	mux.Use(chimid.Recoverer)
	mux.Use(chimid.Timeout(60 * time.Second))

	mux.Route("/companies", func(mux chi.Router) {
		mux.Post("/", coh.Create)
		mux.Route("/{companyID}", func(mux chi.Router) {
			mux.Use(middleware.SetCompanyIDToCtx)
			mux.Get("/", coh.GetByID)
			mux.Get("/invoices", ih.ListByPaymentDueDateBetween)
			mux.Route("/clients", func(mux chi.Router) {
				mux.Post("/", clh.Create)
				mux.Route("/{clientID}", func(mux chi.Router) {
					mux.Use(middleware.SetClientIDToCtx)
					mux.Get("/", clh.GetByID)
					mux.Route("/invoices", func(mux chi.Router) {
						mux.Post("/", ih.Create)
					})
				})
			})
		})
	})

	return mux
}
