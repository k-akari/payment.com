package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/pkg/validator"
)

type clientHandler struct {
	clientUsecase clientUsecase
}

func NewClientHandler(
	clientUsecase clientUsecase,
) *clientHandler {
	return &clientHandler{
		clientUsecase: clientUsecase,
	}
}

func (ch *clientHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	scoid, ok := ctx.Value("companyID").(string)
	if !ok {
		respondJSON(ctx, w, &errResponse{Message: "invalid company id"}, http.StatusInternalServerError)
		return
	}

	coid, err := strconv.Atoi(scoid)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

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

	err = validator.Struct(b)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusBadRequest)
		return
	}

	client := domain.Client{
		CompanyID:       domain.CompanyID(coid),
		Name:            b.Name,
		Representative:  b.Representative,
		TelephoneNumber: b.TelephoneNumber,
		PostalCode:      b.PostalCode,
		Address:         b.Address,
	}

	cid, err := ch.clientUsecase.Create(ctx, &client)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID domain.ClientID `json:"id"`
	}{ID: cid}

	respondJSON(ctx, w, resp, http.StatusOK)
}

func (ch *clientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sclid, ok := ctx.Value("clientID").(string)
	if !ok {
		respondJSON(ctx, w, &errResponse{Message: "invalid company id"}, http.StatusInternalServerError)
		return
	}

	clid, err := strconv.Atoi(sclid)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	c, err := ch.clientUsecase.GetByID(ctx, domain.ClientID(clid))
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		CompanyID       domain.CompanyID `json:"company_id"`
		Name            string           `json:"name"`
		Representative  string           `json:"representative"`
		TelephoneNumber string           `json:"telephone_number"`
		PostalCode      string           `json:"postal_code"`
		Address         string           `json:"address"`
	}{
		CompanyID:       c.CompanyID,
		Name:            c.Name,
		Representative:  c.Representative,
		TelephoneNumber: c.TelephoneNumber,
		PostalCode:      c.PostalCode,
		Address:         c.Address,
	}

	respondJSON(ctx, w, resp, http.StatusOK)
}
