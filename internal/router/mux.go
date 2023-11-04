package router

import (
	"net/http"
	"time"

	chimid "github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/k-akari/payment.com/internal/handler"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
	"github.com/k-akari/payment.com/internal/infrastructure/repository"
	"github.com/k-akari/payment.com/internal/middleware"
	"github.com/k-akari/payment.com/internal/usecase"
)

func NewMux(db *sqlx.DB) (http.Handler, error) {
	mux := chi.NewRouter()

	mux.Use(chimid.RequestID)
	mux.Use(chimid.RealIP)
	mux.Use(chimid.Logger)
	mux.Use(chimid.Recoverer)
	mux.Use(chimid.Timeout(60 * time.Second))

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

	return mux, nil
}
