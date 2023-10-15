package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/pkg/validator"
)

type companyHandler struct {
	companyUsecase companyUsecase
}

func NewCompanyHandler(
	companyUsecase companyUsecase,
) *companyHandler {
	return &companyHandler{
		companyUsecase: companyUsecase,
	}
}

func (ch *companyHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var b struct {
		Name            string `json:"name" validate:"required"`
		Representative  string `json:"representative" validate:"required"`
		TelephoneNumber string `json:"telephone_number" validate:"required"`
		PostalCode      string `json:"postal_code" validate:"required"`
		Address         string `json:"address" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	err := validator.Struct(b)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	company := domain.Company{
		Name:            b.Name,
		Representative:  b.Representative,
		TelephoneNumber: b.TelephoneNumber,
		PostalCode:      b.PostalCode,
		Address:         b.Address,
	}

	cid, err := ch.companyUsecase.CreateCompany(ctx, &company)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID domain.CompanyID `json:"id"`
	}{ID: cid}

	respondJSON(ctx, w, resp, http.StatusOK)
}

func (ch *companyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	scid := chi.URLParam(r, "companyID")
	cid, err := strconv.Atoi(scid)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	c, err := ch.companyUsecase.GetCompanyByID(ctx, domain.CompanyID(cid))
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		Name            string `json:"name"`
		Representative  string `json:"representative"`
		TelephoneNumber string `json:"telephone_number"`
		PostalCode      string `json:"postal_code"`
		Address         string `json:"address"`
	}{
		Name:            c.Name,
		Representative:  c.Representative,
		TelephoneNumber: c.TelephoneNumber,
		PostalCode:      c.PostalCode,
		Address:         c.Address,
	}

	respondJSON(ctx, w, resp, http.StatusOK)
}
