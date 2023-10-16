package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/k-akari/payment.com/internal/domain"
)

type invoiceUsecase struct {
	invoiceRepository invoiceRepository
}

func NewInvoiceUsecase(
	invoiceRepository invoiceRepository,
) *invoiceUsecase {
	return &invoiceUsecase{
		invoiceRepository: invoiceRepository,
	}
}

func (u *invoiceUsecase) Create(
	ctx context.Context,
	invoice *domain.Invoice,
) (domain.InvoiceID, error) {
	invoice.Status = domain.InvoiceStatusUnpaid
	invoice.SetDefaultRate()
	invoice.CalcBilledAmount()

	iid, err := u.invoiceRepository.Create(ctx, invoice)
	if err != nil {
		return domain.InvoiceID(0), fmt.Errorf("failed to run u.invoiceRepository.Create: %w", err)
	}

	return iid, nil
}

func (u *invoiceUsecase) ListByPaymentDueDateBetween(
	ctx context.Context,
	coid domain.CompanyID,
	from *time.Time,
	to *time.Time,
) ([]*domain.Invoice, error) {
	invoices, err := u.invoiceRepository.ListByPaymentDueDateBetween(ctx, coid, from, to)
	if err != nil {
		return nil, fmt.Errorf("failed to run u.invoiceRepository.ListByPaymentDueDateBetween: %w", err)
	}

	return invoices, nil
}
