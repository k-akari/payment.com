package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/k-akari/payment.com/internal/domain"
	"github.com/k-akari/payment.com/internal/pkg/validator"
)

type invoiceHandler struct {
	invoiceUsecase invoiceUsecase
}

func NewInvoiceHandler(
	invoiceUsecase invoiceUsecase,
) *invoiceHandler {
	return &invoiceHandler{
		invoiceUsecase: invoiceUsecase,
	}
}

func (ih *invoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
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
		IssuedDate:     &b.IssuedDate,
		PaidAmount:     b.PaidAmount,
		PaymentDueDate: &b.PaymentDueDate,
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

func (ih *invoiceHandler) ListByPaymentDueDateBetween(w http.ResponseWriter, r *http.Request) {
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
