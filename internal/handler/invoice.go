package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/k-akari/golang-rest-api-sample/internal/domain"
	"github.com/k-akari/golang-rest-api-sample/internal/pkg/validator"
)

type InvoiceHandler struct {
	invoiceUsecase invoiceUsecase
}

func NewInvoiceHandler(
	invoiceUsecase invoiceUsecase,
) *InvoiceHandler {
	return &InvoiceHandler{
		invoiceUsecase: invoiceUsecase,
	}
}

func (ih *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	coid, err := getCompanyIDFromCtx(ctx)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	clid, err := getClientIDFromCtx(ctx)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	var b struct {
		IssuedDate     time.Time `json:"issued_date" validate:"required"`
		PaidAmount     int64     `json:"paid_amount" validate:"required"`
		PaymentDueDate time.Time `json:"payment_due_date" validate:"required"`
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

	invoice := domain.Invoice{
		CompanyID:      domain.CompanyID(coid),
		ClientID:       domain.ClientID(clid),
		IssuedDate:     b.IssuedDate,
		PaidAmount:     b.PaidAmount,
		PaymentDueDate: b.PaymentDueDate,
	}

	iid, err := ih.invoiceUsecase.Create(ctx, &invoice)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	resp := struct {
		ID domain.InvoiceID `json:"id"`
	}{ID: iid}

	respondJSON(ctx, w, resp, http.StatusOK)
}

func (ih *InvoiceHandler) ListByPaymentDueDateBetween(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	coid, err := getCompanyIDFromCtx(ctx)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	var b struct {
		From *time.Time `json:"from" validate:"required"`
		To   *time.Time `json:"to" validate:"required"`
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

	is, err := ih.invoiceUsecase.ListByPaymentDueDateBetween(ctx, domain.CompanyID(coid), b.From, b.To)
	if err != nil {
		respondJSON(ctx, w, &errResponse{Message: err.Error()}, http.StatusInternalServerError)
		return
	}

	respondJSON(ctx, w, is, http.StatusOK)
}
