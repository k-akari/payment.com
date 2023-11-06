package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/k-akari/payment.com/internal/handler"
)

func setCompanyIDToCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), handler.CompanyID, chi.URLParam(r, "companyID"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func setClientIDToCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), handler.ClientID, chi.URLParam(r, "clientID"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
